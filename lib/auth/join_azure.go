/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package auth

import (
	"cmp"
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/coreos/go-oidc"
	"github.com/digitorus/pkcs7"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"

	"github.com/gravitational/teleport/api/client"
	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/cloud/azure"
	"github.com/gravitational/teleport/lib/utils"
)

const azureAccessTokenAudience = "https://management.azure.com/"

// Structs for unmarshaling attested data. Schema can be found at
// https://learn.microsoft.com/en-us/azure/virtual-machines/linux/instance-metadata-service?tabs=linux#response-2

type signedAttestedData struct {
	Encoding  string `json:"encoding"`
	Signature string `json:"signature"`
}

type plan struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
}

type timestamp struct {
	CreatedOn string `json:"createdOn"`
	ExpiresOn string `json:"expiresOn"`
}

type attestedData struct {
	LicenseType    string    `json:"licenseType"`
	Nonce          string    `json:"nonce"`
	Plan           plan      `json:"plan"`
	Timestamp      timestamp `json:"timestamp"`
	ID             string    `json:"vmId"`
	SubscriptionID string    `json:"subscriptionId"`
	SKU            string    `json:"sku"`
}

type accessTokenClaims struct {
	jwt.Claims
	TenantID string `json:"tid"`
	Version  string `json:"ver"`

	// Azure JWT tokens include two optional claims that can be used to validate
	// the subscription and resource group of a joining node. These claims hold
	// different values depending on the assigned Managed Identity of the Azure VM:
	// - xms_mirid:
	//   - For System-Assigned Identity it represents the resource id of the VM.
	//   - For User-Assigned Identity it represents the resource id of the user-assigned identity.
	// - xms_az_rid:
	//   - For System-Assigned Identity this claim is omitted.
	//   - For User-Assigned Identity it represents the resource id of the VM.
	//
	// More details at: https://learn.microsoft.com/en-us/answers/questions/1282788/existence-of-xms-az-rid-field-in-activity-logs-of

	ManangedIdentityResourceID string `json:"xms_mirid"`
	AzureResourceID            string `json:"xms_az_rid"`
}

type azureVerifyTokenFunc func(ctx context.Context, rawIDToken string) (*accessTokenClaims, error)

type vmClientGetter func(subscriptionID string, token *azure.StaticCredential) (azure.VirtualMachinesClient, error)

type azureRegisterConfig struct {
	clock                  clockwork.Clock
	certificateAuthorities []*x509.Certificate
	verify                 azureVerifyTokenFunc
	getVMClient            vmClientGetter
}

func azureVerifyFuncFromOIDCVerifier(cfg *oidc.Config) azureVerifyTokenFunc {
	return func(ctx context.Context, rawIDToken string) (*accessTokenClaims, error) {
		token, err := jwt.ParseSigned(rawIDToken)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		// Need to get the tenant ID before we verify so we can construct the issuer URL.
		var unverifiedClaims accessTokenClaims
		if err := token.UnsafeClaimsWithoutVerification(&unverifiedClaims); err != nil {
			return nil, trace.Wrap(err)
		}
		issuer, err := url.JoinPath("https://sts.windows.net", unverifiedClaims.TenantID, "/")
		if err != nil {
			return nil, trace.Wrap(err)
		}
		provider, err := oidc.NewProvider(ctx, issuer)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		verifiedToken, err := provider.Verifier(cfg).Verify(ctx, rawIDToken)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		var tokenClaims accessTokenClaims
		if err := verifiedToken.Claims(&tokenClaims); err != nil {
			return nil, trace.Wrap(err)
		}
		return &tokenClaims, nil
	}
}

func (cfg *azureRegisterConfig) CheckAndSetDefaults(ctx context.Context) error {
	if cfg.clock == nil {
		cfg.clock = clockwork.NewRealClock()
	}
	if cfg.verify == nil {
		oidcConfig := &oidc.Config{
			ClientID: azureAccessTokenAudience,
			Now:      cfg.clock.Now,
		}
		cfg.verify = azureVerifyFuncFromOIDCVerifier(oidcConfig)
	}

	if cfg.certificateAuthorities == nil {
		certs, err := getAzureRootCerts()
		if err != nil {
			return trace.Wrap(err)
		}
		cfg.certificateAuthorities = certs
	}
	if cfg.getVMClient == nil {
		cfg.getVMClient = func(subscriptionID string, token *azure.StaticCredential) (azure.VirtualMachinesClient, error) {
			opts := &armpolicy.ClientOptions{
				ClientOptions: policy.ClientOptions{
					Telemetry: policy.TelemetryOptions{
						ApplicationID: "teleport",
					},
				},
			}
			client, err := azure.NewVirtualMachinesClient(subscriptionID, token, opts)
			return client, trace.Wrap(err)
		}
	}
	return nil
}

type azureRegisterOption func(cfg *azureRegisterConfig)

// parseAndVeryAttestedData verifies that an attested data document was signed
// by Azure. If verification is successful, it returns the ID of the VM that
// produced the document.
func parseAndVerifyAttestedData(ctx context.Context, adBytes []byte, challenge string, certs []*x509.Certificate) (subscriptionID, vmID string, err error) {
	var signedAD signedAttestedData
	if err := utils.FastUnmarshal(adBytes, &signedAD); err != nil {
		return "", "", trace.Wrap(err)
	}
	if signedAD.Encoding != "pkcs7" {
		return "", "", trace.AccessDenied("unsupported signature type: %v", signedAD.Encoding)
	}

	sigPEM := "-----BEGIN PKCS7-----\n" + signedAD.Signature + "\n-----END PKCS7-----"
	sigBER, _ := pem.Decode([]byte(sigPEM))
	if sigBER == nil {
		return "", "", trace.AccessDenied("unable to decode attested data document")
	}

	p7, err := pkcs7.Parse(sigBER.Bytes)
	if err != nil {
		return "", "", trace.Wrap(err)
	}
	var ad attestedData
	if err := utils.FastUnmarshal(p7.Content, &ad); err != nil {
		return "", "", trace.Wrap(err)
	}
	if ad.Nonce != challenge {
		return "", "", trace.AccessDenied("challenge is missing or does not match")
	}

	if len(p7.Certificates) == 0 {
		return "", "", trace.AccessDenied("no certificates for signature")
	}
	fixAzureSigningAlgorithm(p7)

	// Azure only sends the leaf cert, so we have to fetch the intermediate.
	intermediate, err := getAzureIssuerCert(ctx, p7.Certificates[0])
	if err != nil {
		return "", "", trace.Wrap(err)
	}
	if intermediate != nil {
		p7.Certificates = append(p7.Certificates, intermediate)
	}

	pool := x509.NewCertPool()
	for _, cert := range certs {
		pool.AddCert(cert)
	}

	if err := p7.VerifyWithChain(pool); err != nil {
		return "", "", trace.Wrap(err)
	}

	return ad.SubscriptionID, ad.ID, nil
}

// verifyToken verifies the token and validates the expected claims.
func verifyToken(ctx context.Context, cfg *azureRegisterConfig, accessToken string, requestStart time.Time) (*accessTokenClaims, error) {
	tokenClaims, err := cfg.verify(ctx, accessToken)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	expectedIssuer, err := url.JoinPath("https://sts.windows.net", tokenClaims.TenantID, "/")
	if err != nil {
		return nil, trace.Wrap(err)
	}
	// v2 tokens have the version appended to the issuer.
	if tokenClaims.Version == "2.0" {
		expectedIssuer, err = url.JoinPath(expectedIssuer, "2.0")
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}

	expectedClaims := jwt.Expected{
		Issuer:   expectedIssuer,
		Audience: jwt.Audience{azureAccessTokenAudience},
		Time:     requestStart,
	}

	if err := tokenClaims.Validate(expectedClaims); err != nil {
		return nil, trace.Wrap(err)
	}

	return tokenClaims, nil
}

// verifyVMIdentity verifies that the provided access token came from the
// correct Azure VM.
func verifyVMIdentity(ctx context.Context, cfg *azureRegisterConfig, tokenClaims *accessTokenClaims, accessToken, subscriptionID, vmID string) (*azure.VirtualMachine, error) {
	tokenCredential := azure.NewStaticCredential(azcore.AccessToken{
		Token:     accessToken,
		ExpiresOn: tokenClaims.Expiry.Time(),
	})
	vmClient, err := cfg.getVMClient(subscriptionID, tokenCredential)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	resourceID, err := arm.ParseResourceID(tokenClaims.ManangedIdentityResourceID)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	var vm *azure.VirtualMachine

	// If the token is from the system-assigned managed identity, the resource ID
	// is for the VM itself and we can use it to look up the VM.
	if slices.Contains(resourceID.ResourceType.Types, "virtualMachines") {
		vm, err = vmClient.Get(ctx, tokenClaims.ManangedIdentityResourceID)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		if vm.VMID != vmID {
			return nil, trace.AccessDenied("vm ID does not match")
		}

		// If the token is from a user-assigned managed identity, the resource ID is
		// for the identity and we need to look the VM up by VM ID.
	} else {
		vm, err = vmClient.GetByVMID(ctx, vmID)
		if err != nil {
			if trace.IsNotFound(err) {
				return nil, trace.AccessDenied("no VM found with matching VM ID")
			}
			return nil, trace.Wrap(err)
		}
	}

	return vm, nil
}

func checkAzureAllowRulesWithClaims(claims *accessTokenClaims, token string, allowRules []*types.ProvisionTokenSpecV2Azure_Rule) error {
	// xms_az_rid claim is omitted when the VM is assigned a System-Assigned Identity.
	// The xms_mirid claim should be used instead.
	rid := cmp.Or(claims.AzureResourceID, claims.ManangedIdentityResourceID)

	resourceID, err := arm.ParseResourceID(rid)
	if err != nil {
		return trace.Wrap(err, "failed to parse resource id from claims")
	}

	if !slices.Contains(resourceID.ResourceType.Types, "virtualMachines") {
		return trace.BadParameter("unexpected resource type: %q", resourceID.ResourceType.Type)
	}

	if err := checkAzureAllowRules(resourceID.SubscriptionID, resourceID.ResourceGroupName, allowRules); err != nil {
		return trace.AccessDenied("instance %v did not match any allow rules in token %v", resourceID.Name, token)
	}
	return nil
}

func checkAzureAllowRulesWithVMs(vm *azure.VirtualMachine, token string, allowRules []*types.ProvisionTokenSpecV2Azure_Rule) error {
	if err := checkAzureAllowRules(vm.Subscription, vm.ResourceGroup, allowRules); err != nil {
		return trace.AccessDenied("instance %v did not match any allow rules in token %v", vm.Name, token)
	}
	return nil
}

func checkAzureAllowRules(subscription, resourceGroup string, allowRules []*types.ProvisionTokenSpecV2Azure_Rule) error {
	for _, rule := range allowRules {
		if rule.Subscription != subscription {
			continue
		}
		if !azureResourceGroupIsAllowed(rule.ResourceGroups, resourceGroup) {
			continue
		}
		return nil
	}
	return trace.AccessDenied("matching allow rule not found")
}

func azureResourceGroupIsAllowed(allowedResourceGroups []string, vmResourceGroup string) bool {
	if len(allowedResourceGroups) == 0 {
		return true
	}

	// ResourceGroups are case insensitive.
	// https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/frequently-asked-questions#are-resource-group-names-case-sensitive
	// The API returns them using capital case, but docs don't mention a specific case.
	// Converting everything to the same case will ensure a proper comparison.
	resourceGroup := strings.ToUpper(vmResourceGroup)
	for _, allowedResourceGroup := range allowedResourceGroups {
		if strings.EqualFold(resourceGroup, allowedResourceGroup) {
			return true
		}
	}

	return false
}

func (a *Server) checkAzureRequest(ctx context.Context, challenge string, req *proto.RegisterUsingAzureMethodRequest, cfg *azureRegisterConfig) error {
	requestStart := a.clock.Now()
	tokenName := req.RegisterUsingTokenRequest.Token
	provisionToken, err := a.GetToken(ctx, tokenName)
	if err != nil {
		return trace.Wrap(err)
	}
	if provisionToken.GetJoinMethod() != types.JoinMethodAzure {
		return trace.AccessDenied("this token does not support the Azure join method")
	}

	subID, vmID, err := parseAndVerifyAttestedData(ctx, req.AttestedData, challenge, cfg.certificateAuthorities)
	if err != nil {
		return trace.Wrap(err)
	}

	claims, err := verifyToken(ctx, cfg, req.AccessToken, requestStart)
	if err != nil {
		return trace.Wrap(err)
	}

	token, ok := provisionToken.(*types.ProvisionTokenV2)
	if !ok {
		return trace.BadParameter("azure join method only supports ProvisionTokenV2, '%T' was provided", provisionToken)
	}

	if err := checkAzureAllowRulesWithClaims(claims, token.GetName(), token.Spec.Azure.Allow); err == nil {
		return nil
	}
	a.logger.WarnContext(ctx, "Failed to validate Azure allow rules with claims. Attempting to validate with VMs.",
		"error", err)

	// Required claims for validation are only present for source resource types
	// that have onboarded to SNI auth. Fallback to validation with VMs if
	// unable to validate with claims.
	vm, err := verifyVMIdentity(ctx, cfg, claims, req.AccessToken, subID, vmID)
	if err != nil {
		return trace.Wrap(err)
	}
	return trace.Wrap(checkAzureAllowRulesWithVMs(vm, token.GetName(), token.Spec.Azure.Allow))
}

func generateAzureChallenge() (string, error) {
	challenge, err := generateChallenge(base64.RawURLEncoding, 24)
	return challenge, trace.Wrap(err)
}

// RegisterUsingAzureMethod registers the caller using the Azure join method
// and returns signed certs to join the cluster.
//
// The caller must provide a ChallengeResponseFunc which returns a
// *proto.RegisterUsingAzureMethodRequest with a signed attested data document
// including the challenge as a nonce.
func (a *Server) RegisterUsingAzureMethod(
	ctx context.Context,
	challengeResponse client.RegisterAzureChallengeResponseFunc,
	opts ...azureRegisterOption,
) (certs *proto.Certs, err error) {
	var provisionToken types.ProvisionToken
	var joinRequest *types.RegisterUsingTokenRequest
	defer func() {
		// Emit a log message and audit event on join failure.
		if err != nil {
			a.handleJoinFailure(
				err, provisionToken, nil, joinRequest,
			)
		}
	}()

	cfg := &azureRegisterConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	if err := cfg.CheckAndSetDefaults(ctx); err != nil {
		return nil, trace.Wrap(err)
	}

	challenge, err := generateAzureChallenge()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	req, err := challengeResponse(challenge)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	joinRequest = req.RegisterUsingTokenRequest

	if err := req.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	provisionToken, err = a.checkTokenJoinRequestCommon(ctx, req.RegisterUsingTokenRequest)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := a.checkAzureRequest(ctx, challenge, req, cfg); err != nil {
		return nil, trace.Wrap(err)
	}

	if req.RegisterUsingTokenRequest.Role == types.RoleBot {
		certs, err := a.generateCertsBot(
			ctx,
			provisionToken,
			req.RegisterUsingTokenRequest,
			nil,
		)
		return certs, trace.Wrap(err)
	}
	certs, err = a.generateCerts(
		ctx,
		provisionToken,
		req.RegisterUsingTokenRequest,
		nil,
	)
	return certs, trace.Wrap(err)
}

// fixAzureSigningAlgorithm fixes a mismatch between the object IDs of the
// hashing algorithm sent by Azure vs the ones expected by the pkcs7 library.
// Specifically, Azure (incorrectly?) sends a [digest encryption algorithm]
// where the pkcs7 structure's [signerInfo] expects a [digest algorithm].
//
// [signerInfo]: https://www.rfc-editor.org/rfc/rfc2315#section-6.4
// [digest algorithm]: https://www.rfc-editor.org/rfc/rfc2315#section-6.3
// [digest encryption algorithm]: https://www.rfc-editor.org/rfc/rfc2315#section-6.4
func fixAzureSigningAlgorithm(p7 *pkcs7.PKCS7) {
	for i, signer := range p7.Signers {
		if signer.DigestAlgorithm.Algorithm.Equal(pkcs7.OIDEncryptionAlgorithmRSASHA256) {
			p7.Signers[i].DigestAlgorithm.Algorithm = pkcs7.OIDDigestAlgorithmSHA256
		}
	}
}
