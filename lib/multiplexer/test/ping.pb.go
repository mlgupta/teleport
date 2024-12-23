//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: teleport/lib/multiplexer/test/ping.proto

package test

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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payload       string                 `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_teleport_lib_multiplexer_test_ping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_multiplexer_test_ping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_teleport_lib_multiplexer_test_ping_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payload       string                 `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_teleport_lib_multiplexer_test_ping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_multiplexer_test_ping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_teleport_lib_multiplexer_test_ping_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_teleport_lib_multiplexer_test_ping_proto protoreflect.FileDescriptor

var file_teleport_lib_multiplexer_test_ping_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x78, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x24,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x32, 0x61, 0x0a, 0x06, 0x50, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x57,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_lib_multiplexer_test_ping_proto_rawDescOnce sync.Once
	file_teleport_lib_multiplexer_test_ping_proto_rawDescData = file_teleport_lib_multiplexer_test_ping_proto_rawDesc
)

func file_teleport_lib_multiplexer_test_ping_proto_rawDescGZIP() []byte {
	file_teleport_lib_multiplexer_test_ping_proto_rawDescOnce.Do(func() {
		file_teleport_lib_multiplexer_test_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_lib_multiplexer_test_ping_proto_rawDescData)
	})
	return file_teleport_lib_multiplexer_test_ping_proto_rawDescData
}

var file_teleport_lib_multiplexer_test_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_lib_multiplexer_test_ping_proto_goTypes = []any{
	(*Request)(nil),  // 0: teleport.lib.multiplexer.test.Request
	(*Response)(nil), // 1: teleport.lib.multiplexer.test.Response
}
var file_teleport_lib_multiplexer_test_ping_proto_depIdxs = []int32{
	0, // 0: teleport.lib.multiplexer.test.Pinger.Ping:input_type -> teleport.lib.multiplexer.test.Request
	1, // 1: teleport.lib.multiplexer.test.Pinger.Ping:output_type -> teleport.lib.multiplexer.test.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_lib_multiplexer_test_ping_proto_init() }
func file_teleport_lib_multiplexer_test_ping_proto_init() {
	if File_teleport_lib_multiplexer_test_ping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_lib_multiplexer_test_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_lib_multiplexer_test_ping_proto_goTypes,
		DependencyIndexes: file_teleport_lib_multiplexer_test_ping_proto_depIdxs,
		MessageInfos:      file_teleport_lib_multiplexer_test_ping_proto_msgTypes,
	}.Build()
	File_teleport_lib_multiplexer_test_ping_proto = out.File
	file_teleport_lib_multiplexer_test_ping_proto_rawDesc = nil
	file_teleport_lib_multiplexer_test_ping_proto_goTypes = nil
	file_teleport_lib_multiplexer_test_ping_proto_depIdxs = nil
}
