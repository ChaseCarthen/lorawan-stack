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
// source: ttn/lorawan/v3/client_services.proto

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
	ClientRegistry_Create_FullMethodName  = "/ttn.lorawan.v3.ClientRegistry/Create"
	ClientRegistry_Get_FullMethodName     = "/ttn.lorawan.v3.ClientRegistry/Get"
	ClientRegistry_List_FullMethodName    = "/ttn.lorawan.v3.ClientRegistry/List"
	ClientRegistry_Update_FullMethodName  = "/ttn.lorawan.v3.ClientRegistry/Update"
	ClientRegistry_Delete_FullMethodName  = "/ttn.lorawan.v3.ClientRegistry/Delete"
	ClientRegistry_Restore_FullMethodName = "/ttn.lorawan.v3.ClientRegistry/Restore"
	ClientRegistry_Purge_FullMethodName   = "/ttn.lorawan.v3.ClientRegistry/Purge"
)

// ClientRegistryClient is the client API for ClientRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientRegistryClient interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error)
	// List OAuth clients where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the OAuth clients the caller
	// has access to.
	// Similar to Get, this selects the fields specified in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error)
	// Update the OAuth client, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error)
	// Delete the OAuth client. This may not release the client ID for reuse.
	Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Restore a recently deleted client.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Purge the client. This will release the client ID for reuse.
	Purge(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewClientRegistryClient(cc grpc.ClientConnInterface) ClientRegistryClient {
	return &clientRegistryClient{cc}
}

func (c *clientRegistryClient) Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, ClientRegistry_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, ClientRegistry_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error) {
	out := new(Clients)
	err := c.cc.Invoke(ctx, ClientRegistry_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, ClientRegistry_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientRegistry_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Restore(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientRegistry_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Purge(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientRegistry_Purge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientRegistryServer is the server API for ClientRegistry service.
// All implementations must embed UnimplementedClientRegistryServer
// for forward compatibility
type ClientRegistryServer interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateClientRequest) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetClientRequest) (*Client, error)
	// List OAuth clients where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the OAuth clients the caller
	// has access to.
	// Similar to Get, this selects the fields specified in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListClientsRequest) (*Clients, error)
	// Update the OAuth client, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateClientRequest) (*Client, error)
	// Delete the OAuth client. This may not release the client ID for reuse.
	Delete(context.Context, *ClientIdentifiers) (*emptypb.Empty, error)
	// Restore a recently deleted client.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *ClientIdentifiers) (*emptypb.Empty, error)
	// Purge the client. This will release the client ID for reuse.
	Purge(context.Context, *ClientIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientRegistryServer()
}

// UnimplementedClientRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedClientRegistryServer struct {
}

func (UnimplementedClientRegistryServer) Create(context.Context, *CreateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClientRegistryServer) Get(context.Context, *GetClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedClientRegistryServer) List(context.Context, *ListClientsRequest) (*Clients, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClientRegistryServer) Update(context.Context, *UpdateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClientRegistryServer) Delete(context.Context, *ClientIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientRegistryServer) Restore(context.Context, *ClientIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedClientRegistryServer) Purge(context.Context, *ClientIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedClientRegistryServer) mustEmbedUnimplementedClientRegistryServer() {}

// UnsafeClientRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientRegistryServer will
// result in compilation errors.
type UnsafeClientRegistryServer interface {
	mustEmbedUnimplementedClientRegistryServer()
}

func RegisterClientRegistryServer(s grpc.ServiceRegistrar, srv ClientRegistryServer) {
	s.RegisterService(&ClientRegistry_ServiceDesc, srv)
}

func _ClientRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Create(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Get(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).List(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Update(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Delete(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Restore(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientRegistry_Purge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Purge(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientRegistry_ServiceDesc is the grpc.ServiceDesc for ClientRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientRegistry",
	HandlerType: (*ClientRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClientRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClientRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _ClientRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _ClientRegistry_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/client_services.proto",
}

const (
	ClientAccess_ListRights_FullMethodName         = "/ttn.lorawan.v3.ClientAccess/ListRights"
	ClientAccess_GetCollaborator_FullMethodName    = "/ttn.lorawan.v3.ClientAccess/GetCollaborator"
	ClientAccess_SetCollaborator_FullMethodName    = "/ttn.lorawan.v3.ClientAccess/SetCollaborator"
	ClientAccess_ListCollaborators_FullMethodName  = "/ttn.lorawan.v3.ClientAccess/ListCollaborators"
	ClientAccess_DeleteCollaborator_FullMethodName = "/ttn.lorawan.v3.ClientAccess/DeleteCollaborator"
)

// ClientAccessClient is the client API for ClientAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientAccessClient interface {
	// List the rights the caller has on this application.
	ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the OAuth client.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List the collaborators on this OAuth client.
	ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
	// DeleteCollaborator removes a collaborator from a client.
	DeleteCollaborator(ctx context.Context, in *DeleteClientCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewClientAccessClient(cc grpc.ClientConnInterface) ClientAccessClient {
	return &clientAccessClient{cc}
}

func (c *clientAccessClient) ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, ClientAccess_ListRights_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, ClientAccess_GetCollaborator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientAccess_SetCollaborator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, ClientAccess_ListCollaborators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) DeleteCollaborator(ctx context.Context, in *DeleteClientCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientAccess_DeleteCollaborator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAccessServer is the server API for ClientAccess service.
// All implementations must embed UnimplementedClientAccessServer
// for forward compatibility
type ClientAccessServer interface {
	// List the rights the caller has on this application.
	ListRights(context.Context, *ClientIdentifiers) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the OAuth client.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetClientCollaboratorRequest) (*emptypb.Empty, error)
	// List the collaborators on this OAuth client.
	ListCollaborators(context.Context, *ListClientCollaboratorsRequest) (*Collaborators, error)
	// DeleteCollaborator removes a collaborator from a client.
	DeleteCollaborator(context.Context, *DeleteClientCollaboratorRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientAccessServer()
}

// UnimplementedClientAccessServer must be embedded to have forward compatible implementations.
type UnimplementedClientAccessServer struct {
}

func (UnimplementedClientAccessServer) ListRights(context.Context, *ClientIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (UnimplementedClientAccessServer) GetCollaborator(context.Context, *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (UnimplementedClientAccessServer) SetCollaborator(context.Context, *SetClientCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (UnimplementedClientAccessServer) ListCollaborators(context.Context, *ListClientCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}
func (UnimplementedClientAccessServer) DeleteCollaborator(context.Context, *DeleteClientCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollaborator not implemented")
}
func (UnimplementedClientAccessServer) mustEmbedUnimplementedClientAccessServer() {}

// UnsafeClientAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientAccessServer will
// result in compilation errors.
type UnsafeClientAccessServer interface {
	mustEmbedUnimplementedClientAccessServer()
}

func RegisterClientAccessServer(s grpc.ServiceRegistrar, srv ClientAccessServer) {
	s.RegisterService(&ClientAccess_ServiceDesc, srv)
}

func _ClientAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAccess_ListRights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListRights(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAccess_GetCollaborator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).GetCollaborator(ctx, req.(*GetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAccess_SetCollaborator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).SetCollaborator(ctx, req.(*SetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAccess_ListCollaborators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListCollaborators(ctx, req.(*ListClientCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_DeleteCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).DeleteCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAccess_DeleteCollaborator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).DeleteCollaborator(ctx, req.(*DeleteClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientAccess_ServiceDesc is the grpc.ServiceDesc for ClientAccess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientAccess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientAccess",
	HandlerType: (*ClientAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _ClientAccess_ListRights_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _ClientAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _ClientAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _ClientAccess_ListCollaborators_Handler,
		},
		{
			MethodName: "DeleteCollaborator",
			Handler:    _ClientAccess_DeleteCollaborator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/client_services.proto",
}
