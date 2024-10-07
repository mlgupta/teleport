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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/machineid/v1/workload_identity_service.proto

package machineidv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WorkloadIdentityService_SignX509SVIDs_FullMethodName = "/teleport.machineid.v1.WorkloadIdentityService/SignX509SVIDs"
	WorkloadIdentityService_SignJWTSVIDs_FullMethodName  = "/teleport.machineid.v1.WorkloadIdentityService/SignJWTSVIDs"
)

// WorkloadIdentityServiceClient is the client API for WorkloadIdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkloadIdentityServiceClient interface {
	// SignX509SVIDs generates signed x509 SVIDs based on the SVIDs provided in
	// the request.
	SignX509SVIDs(ctx context.Context, in *SignX509SVIDsRequest, opts ...grpc.CallOption) (*SignX509SVIDsResponse, error)
	// SignJWTSVIDs generates signed JWT SVIDs based on the requested SVIDs.
	SignJWTSVIDs(ctx context.Context, in *SignJWTSVIDsRequest, opts ...grpc.CallOption) (*SignJWTSVIDsResponse, error)
}

type workloadIdentityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkloadIdentityServiceClient(cc grpc.ClientConnInterface) WorkloadIdentityServiceClient {
	return &workloadIdentityServiceClient{cc}
}

func (c *workloadIdentityServiceClient) SignX509SVIDs(ctx context.Context, in *SignX509SVIDsRequest, opts ...grpc.CallOption) (*SignX509SVIDsResponse, error) {
	out := new(SignX509SVIDsResponse)
	err := c.cc.Invoke(ctx, WorkloadIdentityService_SignX509SVIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadIdentityServiceClient) SignJWTSVIDs(ctx context.Context, in *SignJWTSVIDsRequest, opts ...grpc.CallOption) (*SignJWTSVIDsResponse, error) {
	out := new(SignJWTSVIDsResponse)
	err := c.cc.Invoke(ctx, WorkloadIdentityService_SignJWTSVIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkloadIdentityServiceServer is the server API for WorkloadIdentityService service.
// All implementations must embed UnimplementedWorkloadIdentityServiceServer
// for forward compatibility
type WorkloadIdentityServiceServer interface {
	// SignX509SVIDs generates signed x509 SVIDs based on the SVIDs provided in
	// the request.
	SignX509SVIDs(context.Context, *SignX509SVIDsRequest) (*SignX509SVIDsResponse, error)
	// SignJWTSVIDs generates signed JWT SVIDs based on the requested SVIDs.
	SignJWTSVIDs(context.Context, *SignJWTSVIDsRequest) (*SignJWTSVIDsResponse, error)
	mustEmbedUnimplementedWorkloadIdentityServiceServer()
}

// UnimplementedWorkloadIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkloadIdentityServiceServer struct {
}

func (UnimplementedWorkloadIdentityServiceServer) SignX509SVIDs(context.Context, *SignX509SVIDsRequest) (*SignX509SVIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignX509SVIDs not implemented")
}
func (UnimplementedWorkloadIdentityServiceServer) SignJWTSVIDs(context.Context, *SignJWTSVIDsRequest) (*SignJWTSVIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignJWTSVIDs not implemented")
}
func (UnimplementedWorkloadIdentityServiceServer) mustEmbedUnimplementedWorkloadIdentityServiceServer() {
}

// UnsafeWorkloadIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkloadIdentityServiceServer will
// result in compilation errors.
type UnsafeWorkloadIdentityServiceServer interface {
	mustEmbedUnimplementedWorkloadIdentityServiceServer()
}

func RegisterWorkloadIdentityServiceServer(s grpc.ServiceRegistrar, srv WorkloadIdentityServiceServer) {
	s.RegisterService(&WorkloadIdentityService_ServiceDesc, srv)
}

func _WorkloadIdentityService_SignX509SVIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignX509SVIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadIdentityServiceServer).SignX509SVIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkloadIdentityService_SignX509SVIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadIdentityServiceServer).SignX509SVIDs(ctx, req.(*SignX509SVIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadIdentityService_SignJWTSVIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJWTSVIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadIdentityServiceServer).SignJWTSVIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkloadIdentityService_SignJWTSVIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadIdentityServiceServer).SignJWTSVIDs(ctx, req.(*SignJWTSVIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkloadIdentityService_ServiceDesc is the grpc.ServiceDesc for WorkloadIdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkloadIdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.machineid.v1.WorkloadIdentityService",
	HandlerType: (*WorkloadIdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignX509SVIDs",
			Handler:    _WorkloadIdentityService_SignX509SVIDs_Handler,
		},
		{
			MethodName: "SignJWTSVIDs",
			Handler:    _WorkloadIdentityService_SignJWTSVIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/machineid/v1/workload_identity_service.proto",
}
