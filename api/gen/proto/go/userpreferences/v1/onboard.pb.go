// Copyright 2023 Gravitational, Inc
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
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/onboard.proto

package userpreferencesv1

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

// Resources are the Resource options in the onboarding questionnaire
type Resource int32

const (
	Resource_RESOURCE_UNSPECIFIED      Resource = 0
	Resource_RESOURCE_WINDOWS_DESKTOPS Resource = 1
	Resource_RESOURCE_SERVER_SSH       Resource = 2
	Resource_RESOURCE_DATABASES        Resource = 3
	Resource_RESOURCE_KUBERNETES       Resource = 4
	Resource_RESOURCE_WEB_APPLICATIONS Resource = 5
)

// Enum value maps for Resource.
var (
	Resource_name = map[int32]string{
		0: "RESOURCE_UNSPECIFIED",
		1: "RESOURCE_WINDOWS_DESKTOPS",
		2: "RESOURCE_SERVER_SSH",
		3: "RESOURCE_DATABASES",
		4: "RESOURCE_KUBERNETES",
		5: "RESOURCE_WEB_APPLICATIONS",
	}
	Resource_value = map[string]int32{
		"RESOURCE_UNSPECIFIED":      0,
		"RESOURCE_WINDOWS_DESKTOPS": 1,
		"RESOURCE_SERVER_SSH":       2,
		"RESOURCE_DATABASES":        3,
		"RESOURCE_KUBERNETES":       4,
		"RESOURCE_WEB_APPLICATIONS": 5,
	}
)

func (x Resource) Enum() *Resource {
	p := new(Resource)
	*p = x
	return p
}

func (x Resource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resource) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_userpreferences_v1_onboard_proto_enumTypes[0].Descriptor()
}

func (Resource) Type() protoreflect.EnumType {
	return &file_teleport_userpreferences_v1_onboard_proto_enumTypes[0]
}

func (x Resource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resource.Descriptor instead.
func (Resource) EnumDescriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_onboard_proto_rawDescGZIP(), []int{0}
}

// MarketingParams are the parameters associated with a user via marketing campaign at the time of sign up.
// They contain both traditional Urchin Tracking Module (UTM) parameters as well as custom parameters.
type MarketingParams struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// campaign is the UTM campaign parameter which identifies a specific product promotion
	Campaign string `protobuf:"bytes,1,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// source is the UTM source parameter which identifies which site sent the traffic
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// medium is the UTM medium parameter which identifies what type of link was used
	Medium string `protobuf:"bytes,3,opt,name=medium,proto3" json:"medium,omitempty"`
	// intent is the internal query param, which identifies any additional marketing intentions
	// via internally set and directed parameters.
	Intent        string `protobuf:"bytes,4,opt,name=intent,proto3" json:"intent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarketingParams) Reset() {
	*x = MarketingParams{}
	mi := &file_teleport_userpreferences_v1_onboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketingParams) ProtoMessage() {}

func (x *MarketingParams) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_onboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketingParams.ProtoReflect.Descriptor instead.
func (*MarketingParams) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_onboard_proto_rawDescGZIP(), []int{0}
}

func (x *MarketingParams) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *MarketingParams) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MarketingParams) GetMedium() string {
	if x != nil {
		return x.Medium
	}
	return ""
}

func (x *MarketingParams) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

// OnboardUserPreferences is the user preferences selected during onboarding.
type OnboardUserPreferences struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// preferredResources is an array of the resources a user selected during their onboarding questionnaire.
	PreferredResources []Resource `protobuf:"varint,1,rep,packed,name=preferred_resources,json=preferredResources,proto3,enum=teleport.userpreferences.v1.Resource" json:"preferred_resources,omitempty"`
	// marketingParams are the parameters associated with a user via marketing campaign at the time of sign up
	MarketingParams *MarketingParams `protobuf:"bytes,2,opt,name=marketing_params,json=marketingParams,proto3" json:"marketing_params,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OnboardUserPreferences) Reset() {
	*x = OnboardUserPreferences{}
	mi := &file_teleport_userpreferences_v1_onboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnboardUserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnboardUserPreferences) ProtoMessage() {}

func (x *OnboardUserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_onboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnboardUserPreferences.ProtoReflect.Descriptor instead.
func (*OnboardUserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_onboard_proto_rawDescGZIP(), []int{1}
}

func (x *OnboardUserPreferences) GetPreferredResources() []Resource {
	if x != nil {
		return x.PreferredResources
	}
	return nil
}

func (x *OnboardUserPreferences) GetMarketingParams() *MarketingParams {
	if x != nil {
		return x.MarketingParams
	}
	return nil
}

var File_teleport_userpreferences_v1_onboard_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_onboard_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x75, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0xc9, 0x01, 0x0a, 0x16, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x12,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x57, 0x0a, 0x10, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a, 0xac, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x53, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x53, 0x53, 0x48, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x53,
	0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4b,
	0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x41, 0x50, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x05, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_userpreferences_v1_onboard_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_onboard_proto_rawDescData = file_teleport_userpreferences_v1_onboard_proto_rawDesc
)

func file_teleport_userpreferences_v1_onboard_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_onboard_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_onboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userpreferences_v1_onboard_proto_rawDescData)
	})
	return file_teleport_userpreferences_v1_onboard_proto_rawDescData
}

var file_teleport_userpreferences_v1_onboard_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_userpreferences_v1_onboard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_userpreferences_v1_onboard_proto_goTypes = []any{
	(Resource)(0),                  // 0: teleport.userpreferences.v1.Resource
	(*MarketingParams)(nil),        // 1: teleport.userpreferences.v1.MarketingParams
	(*OnboardUserPreferences)(nil), // 2: teleport.userpreferences.v1.OnboardUserPreferences
}
var file_teleport_userpreferences_v1_onboard_proto_depIdxs = []int32{
	0, // 0: teleport.userpreferences.v1.OnboardUserPreferences.preferred_resources:type_name -> teleport.userpreferences.v1.Resource
	1, // 1: teleport.userpreferences.v1.OnboardUserPreferences.marketing_params:type_name -> teleport.userpreferences.v1.MarketingParams
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_onboard_proto_init() }
func file_teleport_userpreferences_v1_onboard_proto_init() {
	if File_teleport_userpreferences_v1_onboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_userpreferences_v1_onboard_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userpreferences_v1_onboard_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_onboard_proto_depIdxs,
		EnumInfos:         file_teleport_userpreferences_v1_onboard_proto_enumTypes,
		MessageInfos:      file_teleport_userpreferences_v1_onboard_proto_msgTypes,
	}.Build()
	File_teleport_userpreferences_v1_onboard_proto = out.File
	file_teleport_userpreferences_v1_onboard_proto_rawDesc = nil
	file_teleport_userpreferences_v1_onboard_proto_goTypes = nil
	file_teleport_userpreferences_v1_onboard_proto_depIdxs = nil
}
