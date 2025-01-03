// Copyright 2024 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: teleport/machineid/v1/bot_instance.proto

package machineidv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	v11 "github.com/gravitational/teleport/api/gen/proto/go/teleport/workloadidentity/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A BotInstance
type BotInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of resource represented.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Differentiates variations of the same kind. All resources should
	// contain one, even if it is never populated.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// The version of the resource being represented.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Common metadata that all resources share.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The configured properties of a BotInstance.
	Spec *BotInstanceSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	// Fields that are set by the server as results of operations. These should
	// not be modified by users.
	Status *BotInstanceStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BotInstance) Reset() {
	*x = BotInstance{}
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BotInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInstance) ProtoMessage() {}

func (x *BotInstance) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInstance.ProtoReflect.Descriptor instead.
func (*BotInstance) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP(), []int{0}
}

func (x *BotInstance) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *BotInstance) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *BotInstance) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BotInstance) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BotInstance) GetSpec() *BotInstanceSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *BotInstance) GetStatus() *BotInstanceStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// BotInstanceSpec contains fields
type BotInstanceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the bot associated with this instance.
	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	// The unique identifier for this instance.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *BotInstanceSpec) Reset() {
	*x = BotInstanceSpec{}
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BotInstanceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInstanceSpec) ProtoMessage() {}

func (x *BotInstanceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInstanceSpec.ProtoReflect.Descriptor instead.
func (*BotInstanceSpec) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP(), []int{1}
}

func (x *BotInstanceSpec) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *BotInstanceSpec) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

// BotInstanceStatusHeartbeat contains information self-reported by an instance
// of a Bot. This information is not verified by the server and should not be
// trusted.
type BotInstanceStatusHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp that the heartbeat was recorded by the Auth Server. Any
	// value submitted by `tbot` for this field will be ignored.
	RecordedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
	// Indicates whether this is the heartbeat submitted by `tbot` on startup.
	IsStartup bool `protobuf:"varint,2,opt,name=is_startup,json=isStartup,proto3" json:"is_startup,omitempty"`
	// The version of `tbot` that submitted this heartbeat.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The hostname of the host that `tbot` is running on.
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The duration that `tbot` has been running for when it submitted this
	// heartbeat.
	Uptime *durationpb.Duration `protobuf:"bytes,5,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// The currently configured join_method.
	JoinMethod string `protobuf:"bytes,6,opt,name=join_method,json=joinMethod,proto3" json:"join_method,omitempty"`
	// Indicates whether `tbot` is running in one-shot mode.
	OneShot bool `protobuf:"varint,7,opt,name=one_shot,json=oneShot,proto3" json:"one_shot,omitempty"`
	// The architecture of the host that `tbot` is running on, determined by
	// runtime.GOARCH.
	Architecture string `protobuf:"bytes,8,opt,name=architecture,proto3" json:"architecture,omitempty"`
	// The OS of the host that `tbot` is running on, determined by runtime.GOOS.
	Os string `protobuf:"bytes,9,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *BotInstanceStatusHeartbeat) Reset() {
	*x = BotInstanceStatusHeartbeat{}
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BotInstanceStatusHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInstanceStatusHeartbeat) ProtoMessage() {}

func (x *BotInstanceStatusHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInstanceStatusHeartbeat.ProtoReflect.Descriptor instead.
func (*BotInstanceStatusHeartbeat) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP(), []int{2}
}

func (x *BotInstanceStatusHeartbeat) GetRecordedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordedAt
	}
	return nil
}

func (x *BotInstanceStatusHeartbeat) GetIsStartup() bool {
	if x != nil {
		return x.IsStartup
	}
	return false
}

func (x *BotInstanceStatusHeartbeat) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BotInstanceStatusHeartbeat) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *BotInstanceStatusHeartbeat) GetUptime() *durationpb.Duration {
	if x != nil {
		return x.Uptime
	}
	return nil
}

func (x *BotInstanceStatusHeartbeat) GetJoinMethod() string {
	if x != nil {
		return x.JoinMethod
	}
	return ""
}

func (x *BotInstanceStatusHeartbeat) GetOneShot() bool {
	if x != nil {
		return x.OneShot
	}
	return false
}

func (x *BotInstanceStatusHeartbeat) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *BotInstanceStatusHeartbeat) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

// BotInstanceStatusAuthentication contains information about a join or renewal.
// Ths information is entirely sourced by the Auth Server and can be trusted.
type BotInstanceStatusAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp that the join or renewal was authenticated by the Auth
	// Server.
	AuthenticatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=authenticated_at,json=authenticatedAt,proto3" json:"authenticated_at,omitempty"`
	// The join method used for this join or renewal.
	JoinMethod string `protobuf:"bytes,2,opt,name=join_method,json=joinMethod,proto3" json:"join_method,omitempty"`
	// The join token used for this join or renewal. This is only populated for
	// delegated join methods as the value for `token` join methods is sensitive.
	JoinToken string `protobuf:"bytes,3,opt,name=join_token,json=joinToken,proto3" json:"join_token,omitempty"`
	// The metadata sourced from the join method.
	Metadata *structpb.Struct `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// On each renewal, this generation is incremented. For delegated join
	// methods, this counter is not checked during renewal. For the `token` join
	// method, this counter is checked during renewal and the Bot is locked out if
	// the counter in the certificate does not match the counter of the last
	// authentication.
	Generation int32 `protobuf:"varint,5,opt,name=generation,proto3" json:"generation,omitempty"`
	// The public key of the Bot instance. This must be a PEM wrapped, PKIX DER
	// encoded public key. This provides consistency and supports multiple types
	// of public key algorithm.
	PublicKey []byte `protobuf:"bytes,6,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The attributes generated during the join process. Typically, this is
	// information from the join attestation process itself. This field will
	// eventually replace the `metadata` field, which is structureless.
	JoinAttrs *v11.JoinAttrs `protobuf:"bytes,8,opt,name=join_attrs,json=joinAttrs,proto3" json:"join_attrs,omitempty"`
}

func (x *BotInstanceStatusAuthentication) Reset() {
	*x = BotInstanceStatusAuthentication{}
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BotInstanceStatusAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInstanceStatusAuthentication) ProtoMessage() {}

func (x *BotInstanceStatusAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInstanceStatusAuthentication.ProtoReflect.Descriptor instead.
func (*BotInstanceStatusAuthentication) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP(), []int{3}
}

func (x *BotInstanceStatusAuthentication) GetAuthenticatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AuthenticatedAt
	}
	return nil
}

func (x *BotInstanceStatusAuthentication) GetJoinMethod() string {
	if x != nil {
		return x.JoinMethod
	}
	return ""
}

func (x *BotInstanceStatusAuthentication) GetJoinToken() string {
	if x != nil {
		return x.JoinToken
	}
	return ""
}

func (x *BotInstanceStatusAuthentication) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BotInstanceStatusAuthentication) GetGeneration() int32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *BotInstanceStatusAuthentication) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *BotInstanceStatusAuthentication) GetJoinAttrs() *v11.JoinAttrs {
	if x != nil {
		return x.JoinAttrs
	}
	return nil
}

// BotInstanceStatus holds the status of a BotInstance.
type BotInstanceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The initial authentication status for this bot instance.
	InitialAuthentication *BotInstanceStatusAuthentication `protobuf:"bytes,1,opt,name=initial_authentication,json=initialAuthentication,proto3" json:"initial_authentication,omitempty"`
	// The N most recent authentication status records for this bot instance.
	LatestAuthentications []*BotInstanceStatusAuthentication `protobuf:"bytes,2,rep,name=latest_authentications,json=latestAuthentications,proto3" json:"latest_authentications,omitempty"`
	// The initial heartbeat status for this bot instance.
	InitialHeartbeat *BotInstanceStatusHeartbeat `protobuf:"bytes,3,opt,name=initial_heartbeat,json=initialHeartbeat,proto3" json:"initial_heartbeat,omitempty"`
	// The N most recent heartbeats for this bot instance.
	LatestHeartbeats []*BotInstanceStatusHeartbeat `protobuf:"bytes,4,rep,name=latest_heartbeats,json=latestHeartbeats,proto3" json:"latest_heartbeats,omitempty"`
}

func (x *BotInstanceStatus) Reset() {
	*x = BotInstanceStatus{}
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BotInstanceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInstanceStatus) ProtoMessage() {}

func (x *BotInstanceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_instance_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInstanceStatus.ProtoReflect.Descriptor instead.
func (*BotInstanceStatus) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP(), []int{4}
}

func (x *BotInstanceStatus) GetInitialAuthentication() *BotInstanceStatusAuthentication {
	if x != nil {
		return x.InitialAuthentication
	}
	return nil
}

func (x *BotInstanceStatus) GetLatestAuthentications() []*BotInstanceStatusAuthentication {
	if x != nil {
		return x.LatestAuthentications
	}
	return nil
}

func (x *BotInstanceStatus) GetInitialHeartbeat() *BotInstanceStatusHeartbeat {
	if x != nil {
		return x.InitialHeartbeat
	}
	return nil
}

func (x *BotInstanceStatus) GetLatestHeartbeats() []*BotInstanceStatusHeartbeat {
	if x != nil {
		return x.LatestHeartbeats
	}
	return nil
}

var File_teleport_machineid_v1_bot_instance_proto protoreflect.FileDescriptor

var file_teleport_machineid_v1_bot_instance_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x0f, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0xd1, 0x02,
	0x0a, 0x1a, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x65, 0x53, 0x68, 0x6f, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f,
	0x73, 0x22, 0xf7, 0x02, 0x0a, 0x1f, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x46, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x09, 0x6a,
	0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x52, 0x0b,
	0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0xb1, 0x03, 0x0a, 0x11,
	0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x6d, 0x0a, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x6d, 0x0a, 0x16, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x5e, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x10, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x5e, 0x0a, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x10, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x42,
	0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72,
	0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x69, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_machineid_v1_bot_instance_proto_rawDescOnce sync.Once
	file_teleport_machineid_v1_bot_instance_proto_rawDescData = file_teleport_machineid_v1_bot_instance_proto_rawDesc
)

func file_teleport_machineid_v1_bot_instance_proto_rawDescGZIP() []byte {
	file_teleport_machineid_v1_bot_instance_proto_rawDescOnce.Do(func() {
		file_teleport_machineid_v1_bot_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_machineid_v1_bot_instance_proto_rawDescData)
	})
	return file_teleport_machineid_v1_bot_instance_proto_rawDescData
}

var file_teleport_machineid_v1_bot_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_machineid_v1_bot_instance_proto_goTypes = []any{
	(*BotInstance)(nil),                     // 0: teleport.machineid.v1.BotInstance
	(*BotInstanceSpec)(nil),                 // 1: teleport.machineid.v1.BotInstanceSpec
	(*BotInstanceStatusHeartbeat)(nil),      // 2: teleport.machineid.v1.BotInstanceStatusHeartbeat
	(*BotInstanceStatusAuthentication)(nil), // 3: teleport.machineid.v1.BotInstanceStatusAuthentication
	(*BotInstanceStatus)(nil),               // 4: teleport.machineid.v1.BotInstanceStatus
	(*v1.Metadata)(nil),                     // 5: teleport.header.v1.Metadata
	(*timestamppb.Timestamp)(nil),           // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),             // 7: google.protobuf.Duration
	(*structpb.Struct)(nil),                 // 8: google.protobuf.Struct
	(*v11.JoinAttrs)(nil),                   // 9: teleport.workloadidentity.v1.JoinAttrs
}
var file_teleport_machineid_v1_bot_instance_proto_depIdxs = []int32{
	5,  // 0: teleport.machineid.v1.BotInstance.metadata:type_name -> teleport.header.v1.Metadata
	1,  // 1: teleport.machineid.v1.BotInstance.spec:type_name -> teleport.machineid.v1.BotInstanceSpec
	4,  // 2: teleport.machineid.v1.BotInstance.status:type_name -> teleport.machineid.v1.BotInstanceStatus
	6,  // 3: teleport.machineid.v1.BotInstanceStatusHeartbeat.recorded_at:type_name -> google.protobuf.Timestamp
	7,  // 4: teleport.machineid.v1.BotInstanceStatusHeartbeat.uptime:type_name -> google.protobuf.Duration
	6,  // 5: teleport.machineid.v1.BotInstanceStatusAuthentication.authenticated_at:type_name -> google.protobuf.Timestamp
	8,  // 6: teleport.machineid.v1.BotInstanceStatusAuthentication.metadata:type_name -> google.protobuf.Struct
	9,  // 7: teleport.machineid.v1.BotInstanceStatusAuthentication.join_attrs:type_name -> teleport.workloadidentity.v1.JoinAttrs
	3,  // 8: teleport.machineid.v1.BotInstanceStatus.initial_authentication:type_name -> teleport.machineid.v1.BotInstanceStatusAuthentication
	3,  // 9: teleport.machineid.v1.BotInstanceStatus.latest_authentications:type_name -> teleport.machineid.v1.BotInstanceStatusAuthentication
	2,  // 10: teleport.machineid.v1.BotInstanceStatus.initial_heartbeat:type_name -> teleport.machineid.v1.BotInstanceStatusHeartbeat
	2,  // 11: teleport.machineid.v1.BotInstanceStatus.latest_heartbeats:type_name -> teleport.machineid.v1.BotInstanceStatusHeartbeat
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_teleport_machineid_v1_bot_instance_proto_init() }
func file_teleport_machineid_v1_bot_instance_proto_init() {
	if File_teleport_machineid_v1_bot_instance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_machineid_v1_bot_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_machineid_v1_bot_instance_proto_goTypes,
		DependencyIndexes: file_teleport_machineid_v1_bot_instance_proto_depIdxs,
		MessageInfos:      file_teleport_machineid_v1_bot_instance_proto_msgTypes,
	}.Build()
	File_teleport_machineid_v1_bot_instance_proto = out.File
	file_teleport_machineid_v1_bot_instance_proto_rawDesc = nil
	file_teleport_machineid_v1_bot_instance_proto_goTypes = nil
	file_teleport_machineid_v1_bot_instance_proto_depIdxs = nil
}
