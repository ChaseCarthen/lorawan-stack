// Copyright © 2023 The Things Network Foundation, The Things Industries B.V.
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
// source: ttn/lorawan/v3/networkserver_relay.proto

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
	NsRelayConfigurationService_CreateRelay_FullMethodName                     = "/ttn.lorawan.v3.NsRelayConfigurationService/CreateRelay"
	NsRelayConfigurationService_GetRelay_FullMethodName                        = "/ttn.lorawan.v3.NsRelayConfigurationService/GetRelay"
	NsRelayConfigurationService_UpdateRelay_FullMethodName                     = "/ttn.lorawan.v3.NsRelayConfigurationService/UpdateRelay"
	NsRelayConfigurationService_DeleteRelay_FullMethodName                     = "/ttn.lorawan.v3.NsRelayConfigurationService/DeleteRelay"
	NsRelayConfigurationService_CreateRelayUplinkForwardingRule_FullMethodName = "/ttn.lorawan.v3.NsRelayConfigurationService/CreateRelayUplinkForwardingRule"
	NsRelayConfigurationService_GetRelayUplinkForwardingRule_FullMethodName    = "/ttn.lorawan.v3.NsRelayConfigurationService/GetRelayUplinkForwardingRule"
	NsRelayConfigurationService_ListRelayUplinkForwardingRules_FullMethodName  = "/ttn.lorawan.v3.NsRelayConfigurationService/ListRelayUplinkForwardingRules"
	NsRelayConfigurationService_UpdateRelayUplinkForwardingRule_FullMethodName = "/ttn.lorawan.v3.NsRelayConfigurationService/UpdateRelayUplinkForwardingRule"
	NsRelayConfigurationService_DeleteRelayUplinkForwardingRule_FullMethodName = "/ttn.lorawan.v3.NsRelayConfigurationService/DeleteRelayUplinkForwardingRule"
)

// NsRelayConfigurationServiceClient is the client API for NsRelayConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NsRelayConfigurationServiceClient interface {
	// Create a relay.
	CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error)
	// Get a relay.
	GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error)
	// Update a relay.
	UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error)
	// Delete a relay.
	DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error)
	// Create an uplink forwarding rule.
	CreateRelayUplinkForwardingRule(ctx context.Context, in *CreateRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*CreateRelayUplinkForwardingRuleResponse, error)
	// Get an uplink forwarding rule.
	GetRelayUplinkForwardingRule(ctx context.Context, in *GetRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*GetRelayUplinkForwardingRuleResponse, error)
	// List uplink forwarding rules.
	ListRelayUplinkForwardingRules(ctx context.Context, in *ListRelayUplinkForwardingRulesRequest, opts ...grpc.CallOption) (*ListRelayUplinkForwardingRulesResponse, error)
	// Update an uplink forwarding rule.
	UpdateRelayUplinkForwardingRule(ctx context.Context, in *UpdateRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*UpdateRelayUplinkForwardingRuleResponse, error)
	// Delete an uplink forwarding rule.
	DeleteRelayUplinkForwardingRule(ctx context.Context, in *DeleteRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*DeleteRelayUplinkForwardingRuleResponse, error)
}

type nsRelayConfigurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNsRelayConfigurationServiceClient(cc grpc.ClientConnInterface) NsRelayConfigurationServiceClient {
	return &nsRelayConfigurationServiceClient{cc}
}

func (c *nsRelayConfigurationServiceClient) CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error) {
	out := new(CreateRelayResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_CreateRelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error) {
	out := new(GetRelayResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_GetRelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error) {
	out := new(UpdateRelayResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_UpdateRelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error) {
	out := new(DeleteRelayResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_DeleteRelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) CreateRelayUplinkForwardingRule(ctx context.Context, in *CreateRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*CreateRelayUplinkForwardingRuleResponse, error) {
	out := new(CreateRelayUplinkForwardingRuleResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_CreateRelayUplinkForwardingRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) GetRelayUplinkForwardingRule(ctx context.Context, in *GetRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*GetRelayUplinkForwardingRuleResponse, error) {
	out := new(GetRelayUplinkForwardingRuleResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_GetRelayUplinkForwardingRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) ListRelayUplinkForwardingRules(ctx context.Context, in *ListRelayUplinkForwardingRulesRequest, opts ...grpc.CallOption) (*ListRelayUplinkForwardingRulesResponse, error) {
	out := new(ListRelayUplinkForwardingRulesResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_ListRelayUplinkForwardingRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) UpdateRelayUplinkForwardingRule(ctx context.Context, in *UpdateRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*UpdateRelayUplinkForwardingRuleResponse, error) {
	out := new(UpdateRelayUplinkForwardingRuleResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_UpdateRelayUplinkForwardingRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsRelayConfigurationServiceClient) DeleteRelayUplinkForwardingRule(ctx context.Context, in *DeleteRelayUplinkForwardingRuleRequest, opts ...grpc.CallOption) (*DeleteRelayUplinkForwardingRuleResponse, error) {
	out := new(DeleteRelayUplinkForwardingRuleResponse)
	err := c.cc.Invoke(ctx, NsRelayConfigurationService_DeleteRelayUplinkForwardingRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsRelayConfigurationServiceServer is the server API for NsRelayConfigurationService service.
// All implementations must embed UnimplementedNsRelayConfigurationServiceServer
// for forward compatibility
type NsRelayConfigurationServiceServer interface {
	// Create a relay.
	CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error)
	// Get a relay.
	GetRelay(context.Context, *GetRelayRequest) (*GetRelayResponse, error)
	// Update a relay.
	UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error)
	// Delete a relay.
	DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error)
	// Create an uplink forwarding rule.
	CreateRelayUplinkForwardingRule(context.Context, *CreateRelayUplinkForwardingRuleRequest) (*CreateRelayUplinkForwardingRuleResponse, error)
	// Get an uplink forwarding rule.
	GetRelayUplinkForwardingRule(context.Context, *GetRelayUplinkForwardingRuleRequest) (*GetRelayUplinkForwardingRuleResponse, error)
	// List uplink forwarding rules.
	ListRelayUplinkForwardingRules(context.Context, *ListRelayUplinkForwardingRulesRequest) (*ListRelayUplinkForwardingRulesResponse, error)
	// Update an uplink forwarding rule.
	UpdateRelayUplinkForwardingRule(context.Context, *UpdateRelayUplinkForwardingRuleRequest) (*UpdateRelayUplinkForwardingRuleResponse, error)
	// Delete an uplink forwarding rule.
	DeleteRelayUplinkForwardingRule(context.Context, *DeleteRelayUplinkForwardingRuleRequest) (*DeleteRelayUplinkForwardingRuleResponse, error)
	mustEmbedUnimplementedNsRelayConfigurationServiceServer()
}

// UnimplementedNsRelayConfigurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNsRelayConfigurationServiceServer struct {
}

func (UnimplementedNsRelayConfigurationServiceServer) CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelay not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) GetRelay(context.Context, *GetRelayRequest) (*GetRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelay not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelay not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelay not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) CreateRelayUplinkForwardingRule(context.Context, *CreateRelayUplinkForwardingRuleRequest) (*CreateRelayUplinkForwardingRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelayUplinkForwardingRule not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) GetRelayUplinkForwardingRule(context.Context, *GetRelayUplinkForwardingRuleRequest) (*GetRelayUplinkForwardingRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelayUplinkForwardingRule not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) ListRelayUplinkForwardingRules(context.Context, *ListRelayUplinkForwardingRulesRequest) (*ListRelayUplinkForwardingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRelayUplinkForwardingRules not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) UpdateRelayUplinkForwardingRule(context.Context, *UpdateRelayUplinkForwardingRuleRequest) (*UpdateRelayUplinkForwardingRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelayUplinkForwardingRule not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) DeleteRelayUplinkForwardingRule(context.Context, *DeleteRelayUplinkForwardingRuleRequest) (*DeleteRelayUplinkForwardingRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelayUplinkForwardingRule not implemented")
}
func (UnimplementedNsRelayConfigurationServiceServer) mustEmbedUnimplementedNsRelayConfigurationServiceServer() {
}

// UnsafeNsRelayConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NsRelayConfigurationServiceServer will
// result in compilation errors.
type UnsafeNsRelayConfigurationServiceServer interface {
	mustEmbedUnimplementedNsRelayConfigurationServiceServer()
}

func RegisterNsRelayConfigurationServiceServer(s grpc.ServiceRegistrar, srv NsRelayConfigurationServiceServer) {
	s.RegisterService(&NsRelayConfigurationService_ServiceDesc, srv)
}

func _NsRelayConfigurationService_CreateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).CreateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_CreateRelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).CreateRelay(ctx, req.(*CreateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_GetRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).GetRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_GetRelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).GetRelay(ctx, req.(*GetRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_UpdateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).UpdateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_UpdateRelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).UpdateRelay(ctx, req.(*UpdateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_DeleteRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).DeleteRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_DeleteRelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).DeleteRelay(ctx, req.(*DeleteRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_CreateRelayUplinkForwardingRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayUplinkForwardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).CreateRelayUplinkForwardingRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_CreateRelayUplinkForwardingRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).CreateRelayUplinkForwardingRule(ctx, req.(*CreateRelayUplinkForwardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_GetRelayUplinkForwardingRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelayUplinkForwardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).GetRelayUplinkForwardingRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_GetRelayUplinkForwardingRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).GetRelayUplinkForwardingRule(ctx, req.(*GetRelayUplinkForwardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_ListRelayUplinkForwardingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRelayUplinkForwardingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).ListRelayUplinkForwardingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_ListRelayUplinkForwardingRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).ListRelayUplinkForwardingRules(ctx, req.(*ListRelayUplinkForwardingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_UpdateRelayUplinkForwardingRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayUplinkForwardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).UpdateRelayUplinkForwardingRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_UpdateRelayUplinkForwardingRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).UpdateRelayUplinkForwardingRule(ctx, req.(*UpdateRelayUplinkForwardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsRelayConfigurationService_DeleteRelayUplinkForwardingRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayUplinkForwardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsRelayConfigurationServiceServer).DeleteRelayUplinkForwardingRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsRelayConfigurationService_DeleteRelayUplinkForwardingRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsRelayConfigurationServiceServer).DeleteRelayUplinkForwardingRule(ctx, req.(*DeleteRelayUplinkForwardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NsRelayConfigurationService_ServiceDesc is the grpc.ServiceDesc for NsRelayConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NsRelayConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsRelayConfigurationService",
	HandlerType: (*NsRelayConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRelay",
			Handler:    _NsRelayConfigurationService_CreateRelay_Handler,
		},
		{
			MethodName: "GetRelay",
			Handler:    _NsRelayConfigurationService_GetRelay_Handler,
		},
		{
			MethodName: "UpdateRelay",
			Handler:    _NsRelayConfigurationService_UpdateRelay_Handler,
		},
		{
			MethodName: "DeleteRelay",
			Handler:    _NsRelayConfigurationService_DeleteRelay_Handler,
		},
		{
			MethodName: "CreateRelayUplinkForwardingRule",
			Handler:    _NsRelayConfigurationService_CreateRelayUplinkForwardingRule_Handler,
		},
		{
			MethodName: "GetRelayUplinkForwardingRule",
			Handler:    _NsRelayConfigurationService_GetRelayUplinkForwardingRule_Handler,
		},
		{
			MethodName: "ListRelayUplinkForwardingRules",
			Handler:    _NsRelayConfigurationService_ListRelayUplinkForwardingRules_Handler,
		},
		{
			MethodName: "UpdateRelayUplinkForwardingRule",
			Handler:    _NsRelayConfigurationService_UpdateRelayUplinkForwardingRule_Handler,
		},
		{
			MethodName: "DeleteRelayUplinkForwardingRule",
			Handler:    _NsRelayConfigurationService_DeleteRelayUplinkForwardingRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver_relay.proto",
}
