// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: ttn/lorawan/v3/gateway_configuration.proto

package ttnpb

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
	GatewayConfigurationService_GetGatewayConfiguration_FullMethodName = "/ttn.lorawan.v3.GatewayConfigurationService/GetGatewayConfiguration"
)

// GatewayConfigurationServiceClient is the client API for GatewayConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayConfigurationServiceClient interface {
	GetGatewayConfiguration(ctx context.Context, in *GetGatewayConfigurationRequest, opts ...grpc.CallOption) (*GetGatewayConfigurationResponse, error)
}

type gatewayConfigurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayConfigurationServiceClient(cc grpc.ClientConnInterface) GatewayConfigurationServiceClient {
	return &gatewayConfigurationServiceClient{cc}
}

func (c *gatewayConfigurationServiceClient) GetGatewayConfiguration(ctx context.Context, in *GetGatewayConfigurationRequest, opts ...grpc.CallOption) (*GetGatewayConfigurationResponse, error) {
	out := new(GetGatewayConfigurationResponse)
	err := c.cc.Invoke(ctx, GatewayConfigurationService_GetGatewayConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayConfigurationServiceServer is the server API for GatewayConfigurationService service.
// All implementations must embed UnimplementedGatewayConfigurationServiceServer
// for forward compatibility
type GatewayConfigurationServiceServer interface {
	GetGatewayConfiguration(context.Context, *GetGatewayConfigurationRequest) (*GetGatewayConfigurationResponse, error)
	mustEmbedUnimplementedGatewayConfigurationServiceServer()
}

// UnimplementedGatewayConfigurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayConfigurationServiceServer struct {
}

func (UnimplementedGatewayConfigurationServiceServer) GetGatewayConfiguration(context.Context, *GetGatewayConfigurationRequest) (*GetGatewayConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayConfiguration not implemented")
}
func (UnimplementedGatewayConfigurationServiceServer) mustEmbedUnimplementedGatewayConfigurationServiceServer() {
}

// UnsafeGatewayConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayConfigurationServiceServer will
// result in compilation errors.
type UnsafeGatewayConfigurationServiceServer interface {
	mustEmbedUnimplementedGatewayConfigurationServiceServer()
}

func RegisterGatewayConfigurationServiceServer(s grpc.ServiceRegistrar, srv GatewayConfigurationServiceServer) {
	s.RegisterService(&GatewayConfigurationService_ServiceDesc, srv)
}

func _GatewayConfigurationService_GetGatewayConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigurationServiceServer).GetGatewayConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfigurationService_GetGatewayConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigurationServiceServer).GetGatewayConfiguration(ctx, req.(*GetGatewayConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayConfigurationService_ServiceDesc is the grpc.ServiceDesc for GatewayConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayConfigurationService",
	HandlerType: (*GatewayConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayConfiguration",
			Handler:    _GatewayConfigurationService_GetGatewayConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/gateway_configuration.proto",
}
