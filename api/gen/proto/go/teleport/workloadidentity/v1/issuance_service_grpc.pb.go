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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/workloadidentity/v1/issuance_service.proto

package workloadidentityv1

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
	WorkloadIdentityIssuanceService_IssueWorkloadIdentity_FullMethodName   = "/teleport.workloadidentity.v1.WorkloadIdentityIssuanceService/IssueWorkloadIdentity"
	WorkloadIdentityIssuanceService_IssueWorkloadIdentities_FullMethodName = "/teleport.workloadidentity.v1.WorkloadIdentityIssuanceService/IssueWorkloadIdentities"
)

// WorkloadIdentityIssuanceServiceClient is the client API for WorkloadIdentityIssuanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkloadIdentityIssuanceServiceClient interface {
	// IssueWorkloadIdentity issues a workload identity credential for the named
	// WorkloadIdentity resource. If it is unable to issue a credential,
	// an error will be returned.
	IssueWorkloadIdentity(ctx context.Context, in *IssueWorkloadIdentityRequest, opts ...grpc.CallOption) (*IssueWorkloadIdentityResponse, error)
	// IssueWorkloadIdentities can issue multiple workload identity credentials
	// based on label selectors for the WorkloadIdentity resources.
	IssueWorkloadIdentities(ctx context.Context, in *IssueWorkloadIdentitiesRequest, opts ...grpc.CallOption) (*IssueWorkloadIdentitiesResponse, error)
}

type workloadIdentityIssuanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkloadIdentityIssuanceServiceClient(cc grpc.ClientConnInterface) WorkloadIdentityIssuanceServiceClient {
	return &workloadIdentityIssuanceServiceClient{cc}
}

func (c *workloadIdentityIssuanceServiceClient) IssueWorkloadIdentity(ctx context.Context, in *IssueWorkloadIdentityRequest, opts ...grpc.CallOption) (*IssueWorkloadIdentityResponse, error) {
	out := new(IssueWorkloadIdentityResponse)
	err := c.cc.Invoke(ctx, WorkloadIdentityIssuanceService_IssueWorkloadIdentity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadIdentityIssuanceServiceClient) IssueWorkloadIdentities(ctx context.Context, in *IssueWorkloadIdentitiesRequest, opts ...grpc.CallOption) (*IssueWorkloadIdentitiesResponse, error) {
	out := new(IssueWorkloadIdentitiesResponse)
	err := c.cc.Invoke(ctx, WorkloadIdentityIssuanceService_IssueWorkloadIdentities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkloadIdentityIssuanceServiceServer is the server API for WorkloadIdentityIssuanceService service.
// All implementations must embed UnimplementedWorkloadIdentityIssuanceServiceServer
// for forward compatibility
type WorkloadIdentityIssuanceServiceServer interface {
	// IssueWorkloadIdentity issues a workload identity credential for the named
	// WorkloadIdentity resource. If it is unable to issue a credential,
	// an error will be returned.
	IssueWorkloadIdentity(context.Context, *IssueWorkloadIdentityRequest) (*IssueWorkloadIdentityResponse, error)
	// IssueWorkloadIdentities can issue multiple workload identity credentials
	// based on label selectors for the WorkloadIdentity resources.
	IssueWorkloadIdentities(context.Context, *IssueWorkloadIdentitiesRequest) (*IssueWorkloadIdentitiesResponse, error)
	mustEmbedUnimplementedWorkloadIdentityIssuanceServiceServer()
}

// UnimplementedWorkloadIdentityIssuanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkloadIdentityIssuanceServiceServer struct {
}

func (UnimplementedWorkloadIdentityIssuanceServiceServer) IssueWorkloadIdentity(context.Context, *IssueWorkloadIdentityRequest) (*IssueWorkloadIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueWorkloadIdentity not implemented")
}
func (UnimplementedWorkloadIdentityIssuanceServiceServer) IssueWorkloadIdentities(context.Context, *IssueWorkloadIdentitiesRequest) (*IssueWorkloadIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueWorkloadIdentities not implemented")
}
func (UnimplementedWorkloadIdentityIssuanceServiceServer) mustEmbedUnimplementedWorkloadIdentityIssuanceServiceServer() {
}

// UnsafeWorkloadIdentityIssuanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkloadIdentityIssuanceServiceServer will
// result in compilation errors.
type UnsafeWorkloadIdentityIssuanceServiceServer interface {
	mustEmbedUnimplementedWorkloadIdentityIssuanceServiceServer()
}

func RegisterWorkloadIdentityIssuanceServiceServer(s grpc.ServiceRegistrar, srv WorkloadIdentityIssuanceServiceServer) {
	s.RegisterService(&WorkloadIdentityIssuanceService_ServiceDesc, srv)
}

func _WorkloadIdentityIssuanceService_IssueWorkloadIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueWorkloadIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadIdentityIssuanceServiceServer).IssueWorkloadIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkloadIdentityIssuanceService_IssueWorkloadIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadIdentityIssuanceServiceServer).IssueWorkloadIdentity(ctx, req.(*IssueWorkloadIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadIdentityIssuanceService_IssueWorkloadIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueWorkloadIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadIdentityIssuanceServiceServer).IssueWorkloadIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkloadIdentityIssuanceService_IssueWorkloadIdentities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadIdentityIssuanceServiceServer).IssueWorkloadIdentities(ctx, req.(*IssueWorkloadIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkloadIdentityIssuanceService_ServiceDesc is the grpc.ServiceDesc for WorkloadIdentityIssuanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkloadIdentityIssuanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.workloadidentity.v1.WorkloadIdentityIssuanceService",
	HandlerType: (*WorkloadIdentityIssuanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueWorkloadIdentity",
			Handler:    _WorkloadIdentityIssuanceService_IssueWorkloadIdentity_Handler,
		},
		{
			MethodName: "IssueWorkloadIdentities",
			Handler:    _WorkloadIdentityIssuanceService_IssueWorkloadIdentities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/workloadidentity/v1/issuance_service.proto",
}
