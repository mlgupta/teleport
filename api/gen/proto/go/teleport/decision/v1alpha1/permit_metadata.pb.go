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
// source: teleport/decision/v1alpha1/permit_metadata.proto

package decisionpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Metadata for access permits.
type PermitMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FeatureAssertions is a list of EnforcementFeature that the PEP (Policy
	// Enforcement Point) *must* implement in order to correctly enforce the
	// decision. Note that where possible new features should be structured to
	// "fail safe" rather than relying on feature assertions.
	FeatureAssertions []EnforcementFeature `protobuf:"varint,1,rep,packed,name=feature_assertions,json=featureAssertions,proto3,enum=teleport.decision.v1alpha1.EnforcementFeature" json:"feature_assertions,omitempty"`
	// PdpVersion is the version of the PDP (Policy Decision Point) that evaluated
	// the decision request.
	PdpVersion string `protobuf:"bytes,2,opt,name=pdp_version,json=pdpVersion,proto3" json:"pdp_version,omitempty"`
}

func (x *PermitMetadata) Reset() {
	*x = PermitMetadata{}
	mi := &file_teleport_decision_v1alpha1_permit_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermitMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermitMetadata) ProtoMessage() {}

func (x *PermitMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_permit_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermitMetadata.ProtoReflect.Descriptor instead.
func (*PermitMetadata) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *PermitMetadata) GetFeatureAssertions() []EnforcementFeature {
	if x != nil {
		return x.FeatureAssertions
	}
	return nil
}

func (x *PermitMetadata) GetPdpVersion() string {
	if x != nil {
		return x.PdpVersion
	}
	return ""
}

var File_teleport_decision_v1alpha1_permit_metadata_proto protoreflect.FileDescriptor

var file_teleport_decision_v1alpha1_permit_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x34,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x12, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x11, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x64, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x64, 0x70,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescOnce sync.Once
	file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescData = file_teleport_decision_v1alpha1_permit_metadata_proto_rawDesc
)

func file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescGZIP() []byte {
	file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescOnce.Do(func() {
		file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescData)
	})
	return file_teleport_decision_v1alpha1_permit_metadata_proto_rawDescData
}

var file_teleport_decision_v1alpha1_permit_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_decision_v1alpha1_permit_metadata_proto_goTypes = []any{
	(*PermitMetadata)(nil),  // 0: teleport.decision.v1alpha1.PermitMetadata
	(EnforcementFeature)(0), // 1: teleport.decision.v1alpha1.EnforcementFeature
}
var file_teleport_decision_v1alpha1_permit_metadata_proto_depIdxs = []int32{
	1, // 0: teleport.decision.v1alpha1.PermitMetadata.feature_assertions:type_name -> teleport.decision.v1alpha1.EnforcementFeature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_decision_v1alpha1_permit_metadata_proto_init() }
func file_teleport_decision_v1alpha1_permit_metadata_proto_init() {
	if File_teleport_decision_v1alpha1_permit_metadata_proto != nil {
		return
	}
	file_teleport_decision_v1alpha1_enforcement_feature_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_decision_v1alpha1_permit_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_decision_v1alpha1_permit_metadata_proto_goTypes,
		DependencyIndexes: file_teleport_decision_v1alpha1_permit_metadata_proto_depIdxs,
		MessageInfos:      file_teleport_decision_v1alpha1_permit_metadata_proto_msgTypes,
	}.Build()
	File_teleport_decision_v1alpha1_permit_metadata_proto = out.File
	file_teleport_decision_v1alpha1_permit_metadata_proto_rawDesc = nil
	file_teleport_decision_v1alpha1_permit_metadata_proto_goTypes = nil
	file_teleport_decision_v1alpha1_permit_metadata_proto_depIdxs = nil
}
