/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitational/trace"
	"github.com/julienschmidt/httprouter"

	"github.com/gravitational/teleport/lib/httplib"
	"github.com/gravitational/teleport/lib/integrations/samlidp"
	"github.com/gravitational/teleport/lib/integrations/samlidp/samlidpconfig"
	libutils "github.com/gravitational/teleport/lib/utils"
	"github.com/gravitational/teleport/lib/web/scripts/oneoff"
)

func (h *Handler) gcpWorkforceConfigScript(w http.ResponseWriter, r *http.Request, p httprouter.Params) (any, error) {
	queryParams := r.URL.Query()

	// We aren't going to run any service in this method but the NewGCPWorkforceService performs
	// GCP resource name validation which I think is handy before handing off values to script
	// generation func below.
	gcpWorkforce, err := samlidp.NewGCPWorkforceService(samlidp.GCPWorkforceService{
		APIParams: samlidpconfig.GCPWorkforceAPIParams{
			OrganizationID:     queryParams.Get("orgId"),
			PoolName:           queryParams.Get("poolName"),
			PoolProviderName:   queryParams.Get("poolProviderName"),
			SAMLIdPMetadataURL: fmt.Sprintf("https://%s/enterprise/saml-idp/metadata", h.PublicProxyAddr()),
		},
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// The script must execute the following command:
	// teleport integration configure samlidp gcp-workforce
	argsList := []string{
		"integration", "configure", "samlidp", "gcp-workforce",
		fmt.Sprintf("--org-id=%s", libutils.UnixShellQuote(gcpWorkforce.APIParams.OrganizationID)),
		fmt.Sprintf("--pool-name=%s", libutils.UnixShellQuote(gcpWorkforce.APIParams.PoolName)),
		fmt.Sprintf("--pool-provider-name=%s", libutils.UnixShellQuote(gcpWorkforce.APIParams.PoolProviderName)),
		fmt.Sprintf("--idp-metadata-url=%s", libutils.UnixShellQuote(gcpWorkforce.APIParams.SAMLIdPMetadataURL)),
	}
	script, err := oneoff.BuildScript(oneoff.OneOffScriptParams{
		TeleportArgs:   strings.Join(argsList, " "),
		SuccessMessage: "Success! You can now go back to the browser to complete the database enrollment.",
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	httplib.SetScriptHeaders(w.Header())
	_, err = fmt.Fprint(w, script)

	return nil, trace.Wrap(err)
}
