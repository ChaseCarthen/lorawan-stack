// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
// source: ttn/lorawan/v3/deviceclaimingserver.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EndDeviceClaimingServer_Claim_FullMethodName                  = "/ttn.lorawan.v3.EndDeviceClaimingServer/Claim"
	EndDeviceClaimingServer_Unclaim_FullMethodName                = "/ttn.lorawan.v3.EndDeviceClaimingServer/Unclaim"
	EndDeviceClaimingServer_GetInfoByJoinEUI_FullMethodName       = "/ttn.lorawan.v3.EndDeviceClaimingServer/GetInfoByJoinEUI"
	EndDeviceClaimingServer_GetClaimStatus_FullMethodName         = "/ttn.lorawan.v3.EndDeviceClaimingServer/GetClaimStatus"
	EndDeviceClaimingServer_AuthorizeApplication_FullMethodName   = "/ttn.lorawan.v3.EndDeviceClaimingServer/AuthorizeApplication"
	EndDeviceClaimingServer_UnauthorizeApplication_FullMethodName = "/ttn.lorawan.v3.EndDeviceClaimingServer/UnauthorizeApplication"
)

// EndDeviceClaimingServerClient is the client API for EndDeviceClaimingServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceClaimingServerClient interface {
	// Claims the end device on a Join Server by claim authentication code or QR code.
	Claim(ctx context.Context, in *ClaimEndDeviceRequest, opts ...grpc.CallOption) (*EndDeviceIdentifiers, error)
	// Unclaims the end device on a Join Server.
	// EUIs provided in the request are ignored and the end device is looked up by the given identifiers.
	Unclaim(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Return whether claiming is available for a given JoinEUI.
	GetInfoByJoinEUI(ctx context.Context, in *GetInfoByJoinEUIRequest, opts ...grpc.CallOption) (*GetInfoByJoinEUIResponse, error)
	// Gets the claim status of an end device.
	// EUIs provided in the request are ignored and the end device is looked up by the given identifiers.
	GetClaimStatus(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*GetClaimStatusResponse, error)
	// Authorize the End Device Claiming Server to claim devices registered in the given application. The application
	// identifiers are the source application, where the devices are registered before they are claimed.
	// The API key is used to access the application, find the device, verify the claim request and delete the end device
	// from the source application.
	// DEPRECATED: Device claiming that transfers devices between applications is no longer supported and will be removed
	// in a future version of The Things Stack.
	AuthorizeApplication(ctx context.Context, in *AuthorizeApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Unauthorize the End Device Claiming Server to claim devices in the given application.
	// This reverts the authorization given with rpc AuthorizeApplication.
	// DEPRECATED: Device claiming that transfers devices between applications is no longer supported and will be removed
	// in a future version of The Things Stack.
	UnauthorizeApplication(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type endDeviceClaimingServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceClaimingServerClient(cc grpc.ClientConnInterface) EndDeviceClaimingServerClient {
	return &endDeviceClaimingServerClient{cc}
}

func (c *endDeviceClaimingServerClient) Claim(ctx context.Context, in *ClaimEndDeviceRequest, opts ...grpc.CallOption) (*EndDeviceIdentifiers, error) {
	out := new(EndDeviceIdentifiers)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_Claim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceClaimingServerClient) Unclaim(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_Unclaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceClaimingServerClient) GetInfoByJoinEUI(ctx context.Context, in *GetInfoByJoinEUIRequest, opts ...grpc.CallOption) (*GetInfoByJoinEUIResponse, error) {
	out := new(GetInfoByJoinEUIResponse)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_GetInfoByJoinEUI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceClaimingServerClient) GetClaimStatus(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*GetClaimStatusResponse, error) {
	out := new(GetClaimStatusResponse)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_GetClaimStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceClaimingServerClient) AuthorizeApplication(ctx context.Context, in *AuthorizeApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_AuthorizeApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceClaimingServerClient) UnauthorizeApplication(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceClaimingServer_UnauthorizeApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceClaimingServerServer is the server API for EndDeviceClaimingServer service.
// All implementations must embed UnimplementedEndDeviceClaimingServerServer
// for forward compatibility
type EndDeviceClaimingServerServer interface {
	// Claims the end device on a Join Server by claim authentication code or QR code.
	Claim(context.Context, *ClaimEndDeviceRequest) (*EndDeviceIdentifiers, error)
	// Unclaims the end device on a Join Server.
	// EUIs provided in the request are ignored and the end device is looked up by the given identifiers.
	Unclaim(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error)
	// Return whether claiming is available for a given JoinEUI.
	GetInfoByJoinEUI(context.Context, *GetInfoByJoinEUIRequest) (*GetInfoByJoinEUIResponse, error)
	// Gets the claim status of an end device.
	// EUIs provided in the request are ignored and the end device is looked up by the given identifiers.
	GetClaimStatus(context.Context, *EndDeviceIdentifiers) (*GetClaimStatusResponse, error)
	// Authorize the End Device Claiming Server to claim devices registered in the given application. The application
	// identifiers are the source application, where the devices are registered before they are claimed.
	// The API key is used to access the application, find the device, verify the claim request and delete the end device
	// from the source application.
	// DEPRECATED: Device claiming that transfers devices between applications is no longer supported and will be removed
	// in a future version of The Things Stack.
	AuthorizeApplication(context.Context, *AuthorizeApplicationRequest) (*emptypb.Empty, error)
	// Unauthorize the End Device Claiming Server to claim devices in the given application.
	// This reverts the authorization given with rpc AuthorizeApplication.
	// DEPRECATED: Device claiming that transfers devices between applications is no longer supported and will be removed
	// in a future version of The Things Stack.
	UnauthorizeApplication(context.Context, *ApplicationIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedEndDeviceClaimingServerServer()
}

// UnimplementedEndDeviceClaimingServerServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceClaimingServerServer struct {
}

func (UnimplementedEndDeviceClaimingServerServer) Claim(context.Context, *ClaimEndDeviceRequest) (*EndDeviceIdentifiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) Unclaim(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unclaim not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) GetInfoByJoinEUI(context.Context, *GetInfoByJoinEUIRequest) (*GetInfoByJoinEUIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByJoinEUI not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) GetClaimStatus(context.Context, *EndDeviceIdentifiers) (*GetClaimStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimStatus not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) AuthorizeApplication(context.Context, *AuthorizeApplicationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeApplication not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) UnauthorizeApplication(context.Context, *ApplicationIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnauthorizeApplication not implemented")
}
func (UnimplementedEndDeviceClaimingServerServer) mustEmbedUnimplementedEndDeviceClaimingServerServer() {
}

// UnsafeEndDeviceClaimingServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceClaimingServerServer will
// result in compilation errors.
type UnsafeEndDeviceClaimingServerServer interface {
	mustEmbedUnimplementedEndDeviceClaimingServerServer()
}

func RegisterEndDeviceClaimingServerServer(s grpc.ServiceRegistrar, srv EndDeviceClaimingServerServer) {
	s.RegisterService(&EndDeviceClaimingServer_ServiceDesc, srv)
}

func _EndDeviceClaimingServer_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_Claim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).Claim(ctx, req.(*ClaimEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceClaimingServer_Unclaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).Unclaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_Unclaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).Unclaim(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceClaimingServer_GetInfoByJoinEUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoByJoinEUIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).GetInfoByJoinEUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_GetInfoByJoinEUI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).GetInfoByJoinEUI(ctx, req.(*GetInfoByJoinEUIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceClaimingServer_GetClaimStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).GetClaimStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_GetClaimStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).GetClaimStatus(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceClaimingServer_AuthorizeApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).AuthorizeApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_AuthorizeApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).AuthorizeApplication(ctx, req.(*AuthorizeApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceClaimingServer_UnauthorizeApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceClaimingServerServer).UnauthorizeApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceClaimingServer_UnauthorizeApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceClaimingServerServer).UnauthorizeApplication(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceClaimingServer_ServiceDesc is the grpc.ServiceDesc for EndDeviceClaimingServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceClaimingServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceClaimingServer",
	HandlerType: (*EndDeviceClaimingServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claim",
			Handler:    _EndDeviceClaimingServer_Claim_Handler,
		},
		{
			MethodName: "Unclaim",
			Handler:    _EndDeviceClaimingServer_Unclaim_Handler,
		},
		{
			MethodName: "GetInfoByJoinEUI",
			Handler:    _EndDeviceClaimingServer_GetInfoByJoinEUI_Handler,
		},
		{
			MethodName: "GetClaimStatus",
			Handler:    _EndDeviceClaimingServer_GetClaimStatus_Handler,
		},
		{
			MethodName: "AuthorizeApplication",
			Handler:    _EndDeviceClaimingServer_AuthorizeApplication_Handler,
		},
		{
			MethodName: "UnauthorizeApplication",
			Handler:    _EndDeviceClaimingServer_UnauthorizeApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/deviceclaimingserver.proto",
}

const (
	EndDeviceBatchClaimingServer_Unclaim_FullMethodName           = "/ttn.lorawan.v3.EndDeviceBatchClaimingServer/Unclaim"
	EndDeviceBatchClaimingServer_GetInfoByJoinEUIs_FullMethodName = "/ttn.lorawan.v3.EndDeviceBatchClaimingServer/GetInfoByJoinEUIs"
)

// EndDeviceBatchClaimingServerClient is the client API for EndDeviceBatchClaimingServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceBatchClaimingServerClient interface {
	// Unclaims multiple end devices on an external Join Server.
	// All devices must have the same application ID.
	// Check the response for devices that could not be unclaimed.
	Unclaim(ctx context.Context, in *BatchUnclaimEndDevicesRequest, opts ...grpc.CallOption) (*BatchUnclaimEndDevicesResponse, error)
	// Return whether claiming is supported for each Join EUI in a given list.
	GetInfoByJoinEUIs(ctx context.Context, in *GetInfoByJoinEUIsRequest, opts ...grpc.CallOption) (*GetInfoByJoinEUIsResponse, error)
}

type endDeviceBatchClaimingServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceBatchClaimingServerClient(cc grpc.ClientConnInterface) EndDeviceBatchClaimingServerClient {
	return &endDeviceBatchClaimingServerClient{cc}
}

func (c *endDeviceBatchClaimingServerClient) Unclaim(ctx context.Context, in *BatchUnclaimEndDevicesRequest, opts ...grpc.CallOption) (*BatchUnclaimEndDevicesResponse, error) {
	out := new(BatchUnclaimEndDevicesResponse)
	err := c.cc.Invoke(ctx, EndDeviceBatchClaimingServer_Unclaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceBatchClaimingServerClient) GetInfoByJoinEUIs(ctx context.Context, in *GetInfoByJoinEUIsRequest, opts ...grpc.CallOption) (*GetInfoByJoinEUIsResponse, error) {
	out := new(GetInfoByJoinEUIsResponse)
	err := c.cc.Invoke(ctx, EndDeviceBatchClaimingServer_GetInfoByJoinEUIs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceBatchClaimingServerServer is the server API for EndDeviceBatchClaimingServer service.
// All implementations must embed UnimplementedEndDeviceBatchClaimingServerServer
// for forward compatibility
type EndDeviceBatchClaimingServerServer interface {
	// Unclaims multiple end devices on an external Join Server.
	// All devices must have the same application ID.
	// Check the response for devices that could not be unclaimed.
	Unclaim(context.Context, *BatchUnclaimEndDevicesRequest) (*BatchUnclaimEndDevicesResponse, error)
	// Return whether claiming is supported for each Join EUI in a given list.
	GetInfoByJoinEUIs(context.Context, *GetInfoByJoinEUIsRequest) (*GetInfoByJoinEUIsResponse, error)
	mustEmbedUnimplementedEndDeviceBatchClaimingServerServer()
}

// UnimplementedEndDeviceBatchClaimingServerServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceBatchClaimingServerServer struct {
}

func (UnimplementedEndDeviceBatchClaimingServerServer) Unclaim(context.Context, *BatchUnclaimEndDevicesRequest) (*BatchUnclaimEndDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unclaim not implemented")
}
func (UnimplementedEndDeviceBatchClaimingServerServer) GetInfoByJoinEUIs(context.Context, *GetInfoByJoinEUIsRequest) (*GetInfoByJoinEUIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByJoinEUIs not implemented")
}
func (UnimplementedEndDeviceBatchClaimingServerServer) mustEmbedUnimplementedEndDeviceBatchClaimingServerServer() {
}

// UnsafeEndDeviceBatchClaimingServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceBatchClaimingServerServer will
// result in compilation errors.
type UnsafeEndDeviceBatchClaimingServerServer interface {
	mustEmbedUnimplementedEndDeviceBatchClaimingServerServer()
}

func RegisterEndDeviceBatchClaimingServerServer(s grpc.ServiceRegistrar, srv EndDeviceBatchClaimingServerServer) {
	s.RegisterService(&EndDeviceBatchClaimingServer_ServiceDesc, srv)
}

func _EndDeviceBatchClaimingServer_Unclaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUnclaimEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceBatchClaimingServerServer).Unclaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceBatchClaimingServer_Unclaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceBatchClaimingServerServer).Unclaim(ctx, req.(*BatchUnclaimEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceBatchClaimingServer_GetInfoByJoinEUIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoByJoinEUIsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceBatchClaimingServerServer).GetInfoByJoinEUIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceBatchClaimingServer_GetInfoByJoinEUIs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceBatchClaimingServerServer).GetInfoByJoinEUIs(ctx, req.(*GetInfoByJoinEUIsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceBatchClaimingServer_ServiceDesc is the grpc.ServiceDesc for EndDeviceBatchClaimingServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceBatchClaimingServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceBatchClaimingServer",
	HandlerType: (*EndDeviceBatchClaimingServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unclaim",
			Handler:    _EndDeviceBatchClaimingServer_Unclaim_Handler,
		},
		{
			MethodName: "GetInfoByJoinEUIs",
			Handler:    _EndDeviceBatchClaimingServer_GetInfoByJoinEUIs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/deviceclaimingserver.proto",
}

const (
	GatewayClaimingServer_Claim_FullMethodName               = "/ttn.lorawan.v3.GatewayClaimingServer/Claim"
	GatewayClaimingServer_AuthorizeGateway_FullMethodName    = "/ttn.lorawan.v3.GatewayClaimingServer/AuthorizeGateway"
	GatewayClaimingServer_UnauthorizeGateway_FullMethodName  = "/ttn.lorawan.v3.GatewayClaimingServer/UnauthorizeGateway"
	GatewayClaimingServer_GetInfoByGatewayEUI_FullMethodName = "/ttn.lorawan.v3.GatewayClaimingServer/GetInfoByGatewayEUI"
)

// GatewayClaimingServerClient is the client API for GatewayClaimingServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClaimingServerClient interface {
	// Claims a gateway by claim authentication code or QR code and transfers the gateway to the target user.
	Claim(ctx context.Context, in *ClaimGatewayRequest, opts ...grpc.CallOption) (*GatewayIdentifiers, error)
	// AuthorizeGateway allows a gateway to be claimed.
	AuthorizeGateway(ctx context.Context, in *AuthorizeGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UnauthorizeGateway prevents a gateway from being claimed.
	UnauthorizeGateway(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Return whether claiming is available for a given gateway EUI.
	GetInfoByGatewayEUI(ctx context.Context, in *GetInfoByGatewayEUIRequest, opts ...grpc.CallOption) (*GetInfoByGatewayEUIResponse, error)
}

type gatewayClaimingServerClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClaimingServerClient(cc grpc.ClientConnInterface) GatewayClaimingServerClient {
	return &gatewayClaimingServerClient{cc}
}

func (c *gatewayClaimingServerClient) Claim(ctx context.Context, in *ClaimGatewayRequest, opts ...grpc.CallOption) (*GatewayIdentifiers, error) {
	out := new(GatewayIdentifiers)
	err := c.cc.Invoke(ctx, GatewayClaimingServer_Claim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClaimingServerClient) AuthorizeGateway(ctx context.Context, in *AuthorizeGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayClaimingServer_AuthorizeGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClaimingServerClient) UnauthorizeGateway(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayClaimingServer_UnauthorizeGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClaimingServerClient) GetInfoByGatewayEUI(ctx context.Context, in *GetInfoByGatewayEUIRequest, opts ...grpc.CallOption) (*GetInfoByGatewayEUIResponse, error) {
	out := new(GetInfoByGatewayEUIResponse)
	err := c.cc.Invoke(ctx, GatewayClaimingServer_GetInfoByGatewayEUI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayClaimingServerServer is the server API for GatewayClaimingServer service.
// All implementations must embed UnimplementedGatewayClaimingServerServer
// for forward compatibility
type GatewayClaimingServerServer interface {
	// Claims a gateway by claim authentication code or QR code and transfers the gateway to the target user.
	Claim(context.Context, *ClaimGatewayRequest) (*GatewayIdentifiers, error)
	// AuthorizeGateway allows a gateway to be claimed.
	AuthorizeGateway(context.Context, *AuthorizeGatewayRequest) (*emptypb.Empty, error)
	// UnauthorizeGateway prevents a gateway from being claimed.
	UnauthorizeGateway(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error)
	// Return whether claiming is available for a given gateway EUI.
	GetInfoByGatewayEUI(context.Context, *GetInfoByGatewayEUIRequest) (*GetInfoByGatewayEUIResponse, error)
	mustEmbedUnimplementedGatewayClaimingServerServer()
}

// UnimplementedGatewayClaimingServerServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayClaimingServerServer struct {
}

func (UnimplementedGatewayClaimingServerServer) Claim(context.Context, *ClaimGatewayRequest) (*GatewayIdentifiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}
func (UnimplementedGatewayClaimingServerServer) AuthorizeGateway(context.Context, *AuthorizeGatewayRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeGateway not implemented")
}
func (UnimplementedGatewayClaimingServerServer) UnauthorizeGateway(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnauthorizeGateway not implemented")
}
func (UnimplementedGatewayClaimingServerServer) GetInfoByGatewayEUI(context.Context, *GetInfoByGatewayEUIRequest) (*GetInfoByGatewayEUIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByGatewayEUI not implemented")
}
func (UnimplementedGatewayClaimingServerServer) mustEmbedUnimplementedGatewayClaimingServerServer() {}

// UnsafeGatewayClaimingServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayClaimingServerServer will
// result in compilation errors.
type UnsafeGatewayClaimingServerServer interface {
	mustEmbedUnimplementedGatewayClaimingServerServer()
}

func RegisterGatewayClaimingServerServer(s grpc.ServiceRegistrar, srv GatewayClaimingServerServer) {
	s.RegisterService(&GatewayClaimingServer_ServiceDesc, srv)
}

func _GatewayClaimingServer_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayClaimingServerServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayClaimingServer_Claim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayClaimingServerServer).Claim(ctx, req.(*ClaimGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayClaimingServer_AuthorizeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayClaimingServerServer).AuthorizeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayClaimingServer_AuthorizeGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayClaimingServerServer).AuthorizeGateway(ctx, req.(*AuthorizeGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayClaimingServer_UnauthorizeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayClaimingServerServer).UnauthorizeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayClaimingServer_UnauthorizeGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayClaimingServerServer).UnauthorizeGateway(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayClaimingServer_GetInfoByGatewayEUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoByGatewayEUIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayClaimingServerServer).GetInfoByGatewayEUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayClaimingServer_GetInfoByGatewayEUI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayClaimingServerServer).GetInfoByGatewayEUI(ctx, req.(*GetInfoByGatewayEUIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayClaimingServer_ServiceDesc is the grpc.ServiceDesc for GatewayClaimingServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayClaimingServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayClaimingServer",
	HandlerType: (*GatewayClaimingServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claim",
			Handler:    _GatewayClaimingServer_Claim_Handler,
		},
		{
			MethodName: "AuthorizeGateway",
			Handler:    _GatewayClaimingServer_AuthorizeGateway_Handler,
		},
		{
			MethodName: "UnauthorizeGateway",
			Handler:    _GatewayClaimingServer_UnauthorizeGateway_Handler,
		},
		{
			MethodName: "GetInfoByGatewayEUI",
			Handler:    _GatewayClaimingServer_GetInfoByGatewayEUI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/deviceclaimingserver.proto",
}
