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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: ttn/lorawan/v3/events.proto

package ttnpb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the event. This can be used to find the (localized) event description.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Time at which the event was triggered.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// Identifiers of the entity (or entities) involved.
	Identifiers []*EntityIdentifiers `protobuf:"bytes,3,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// Optional data attached to the event.
	Data *anypb.Any `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Correlation IDs can be used to find related events and actions such as API calls.
	CorrelationIds []string `protobuf:"bytes,5,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// The origin of the event. Typically the hostname of the server that created it.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// Event context, internal use only.
	Context map[string][]byte `protobuf:"bytes,7,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The event will be visible to a caller that has any of these rights.
	Visibility *Rights `protobuf:"bytes,8,opt,name=visibility,proto3" json:"visibility,omitempty"`
	// Details on the authentication provided by the caller that triggered this event.
	Authentication *Event_Authentication `protobuf:"bytes,9,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// The IP address of the caller that triggered this event.
	RemoteIp string `protobuf:"bytes,10,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The IP address of the caller that triggered this event.
	UserAgent string `protobuf:"bytes,11,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The unique identifier of the event, assigned on creation.
	UniqueId string `protobuf:"bytes,12,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetIdentifiers() []*EntityIdentifiers {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *Event) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetCorrelationIds() []string {
	if x != nil {
		return x.CorrelationIds
	}
	return nil
}

func (x *Event) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Event) GetContext() map[string][]byte {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Event) GetVisibility() *Rights {
	if x != nil {
		return x.Visibility
	}
	return nil
}

func (x *Event) GetAuthentication() *Event_Authentication {
	if x != nil {
		return x.Authentication
	}
	return nil
}

func (x *Event) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

func (x *Event) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Event) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

type StreamEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiers []*EntityIdentifiers `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// If greater than zero, this will return historical events, up to this maximum when the stream starts.
	// If used in combination with "after", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	Tail uint32 `protobuf:"varint,2,opt,name=tail,proto3" json:"tail,omitempty"`
	// If not empty, this will return historical events after the given time when the stream starts.
	// If used in combination with "tail", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	After *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	// If provided, this will filter events, so that only events with the given names are returned.
	// Names can be provided as either exact event names (e.g. 'gs.up.receive'),
	// or as regular expressions (e.g. '/^gs\..+/').
	Names []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *StreamEventsRequest) Reset() {
	*x = StreamEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsRequest) ProtoMessage() {}

func (x *StreamEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_events_proto_rawDescGZIP(), []int{1}
}

func (x *StreamEventsRequest) GetIdentifiers() []*EntityIdentifiers {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *StreamEventsRequest) GetTail() uint32 {
	if x != nil {
		return x.Tail
	}
	return 0
}

func (x *StreamEventsRequest) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *StreamEventsRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type FindRelatedEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
}

func (x *FindRelatedEventsRequest) Reset() {
	*x = FindRelatedEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRelatedEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRelatedEventsRequest) ProtoMessage() {}

func (x *FindRelatedEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRelatedEventsRequest.ProtoReflect.Descriptor instead.
func (*FindRelatedEventsRequest) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_events_proto_rawDescGZIP(), []int{2}
}

func (x *FindRelatedEventsRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

type FindRelatedEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *FindRelatedEventsResponse) Reset() {
	*x = FindRelatedEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRelatedEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRelatedEventsResponse) ProtoMessage() {}

func (x *FindRelatedEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRelatedEventsResponse.ProtoReflect.Descriptor instead.
func (*FindRelatedEventsResponse) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_events_proto_rawDescGZIP(), []int{3}
}

func (x *FindRelatedEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event_Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of authentication that was used. This is typically a bearer token.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The type of token that was used. Common types are APIKey, AccessToken and SessionToken.
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// The ID of the token that was used.
	TokenId string `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *Event_Authentication) Reset() {
	*x = Event_Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Authentication) ProtoMessage() {}

func (x *Event_Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Authentication.ProtoReflect.Descriptor instead.
func (*Event_Authentication) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_events_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Event_Authentication) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event_Authentication) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Event_Authentication) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

var File_ttn_lorawan_v3_events_proto protoreflect.FileDescriptor

var file_ttn_lorawan_v3_events_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x74, 0x6e, 0x2f, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x05, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52,
	0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x0e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5e,
	0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0xb6,
	0x01, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x0b,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x30, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0xe1, 0x01, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x5a, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x30, 0x01, 0x12, 0x7b, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x28, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ttn_lorawan_v3_events_proto_rawDescOnce sync.Once
	file_ttn_lorawan_v3_events_proto_rawDescData = file_ttn_lorawan_v3_events_proto_rawDesc
)

func file_ttn_lorawan_v3_events_proto_rawDescGZIP() []byte {
	file_ttn_lorawan_v3_events_proto_rawDescOnce.Do(func() {
		file_ttn_lorawan_v3_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_ttn_lorawan_v3_events_proto_rawDescData)
	})
	return file_ttn_lorawan_v3_events_proto_rawDescData
}

var file_ttn_lorawan_v3_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ttn_lorawan_v3_events_proto_goTypes = []interface{}{
	(*Event)(nil),                     // 0: ttn.lorawan.v3.Event
	(*StreamEventsRequest)(nil),       // 1: ttn.lorawan.v3.StreamEventsRequest
	(*FindRelatedEventsRequest)(nil),  // 2: ttn.lorawan.v3.FindRelatedEventsRequest
	(*FindRelatedEventsResponse)(nil), // 3: ttn.lorawan.v3.FindRelatedEventsResponse
	nil,                               // 4: ttn.lorawan.v3.Event.ContextEntry
	(*Event_Authentication)(nil),      // 5: ttn.lorawan.v3.Event.Authentication
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
	(*EntityIdentifiers)(nil),         // 7: ttn.lorawan.v3.EntityIdentifiers
	(*anypb.Any)(nil),                 // 8: google.protobuf.Any
	(*Rights)(nil),                    // 9: ttn.lorawan.v3.Rights
}
var file_ttn_lorawan_v3_events_proto_depIdxs = []int32{
	6,  // 0: ttn.lorawan.v3.Event.time:type_name -> google.protobuf.Timestamp
	7,  // 1: ttn.lorawan.v3.Event.identifiers:type_name -> ttn.lorawan.v3.EntityIdentifiers
	8,  // 2: ttn.lorawan.v3.Event.data:type_name -> google.protobuf.Any
	4,  // 3: ttn.lorawan.v3.Event.context:type_name -> ttn.lorawan.v3.Event.ContextEntry
	9,  // 4: ttn.lorawan.v3.Event.visibility:type_name -> ttn.lorawan.v3.Rights
	5,  // 5: ttn.lorawan.v3.Event.authentication:type_name -> ttn.lorawan.v3.Event.Authentication
	7,  // 6: ttn.lorawan.v3.StreamEventsRequest.identifiers:type_name -> ttn.lorawan.v3.EntityIdentifiers
	6,  // 7: ttn.lorawan.v3.StreamEventsRequest.after:type_name -> google.protobuf.Timestamp
	0,  // 8: ttn.lorawan.v3.FindRelatedEventsResponse.events:type_name -> ttn.lorawan.v3.Event
	1,  // 9: ttn.lorawan.v3.Events.Stream:input_type -> ttn.lorawan.v3.StreamEventsRequest
	2,  // 10: ttn.lorawan.v3.Events.FindRelated:input_type -> ttn.lorawan.v3.FindRelatedEventsRequest
	0,  // 11: ttn.lorawan.v3.Events.Stream:output_type -> ttn.lorawan.v3.Event
	3,  // 12: ttn.lorawan.v3.Events.FindRelated:output_type -> ttn.lorawan.v3.FindRelatedEventsResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_ttn_lorawan_v3_events_proto_init() }
func file_ttn_lorawan_v3_events_proto_init() {
	if File_ttn_lorawan_v3_events_proto != nil {
		return
	}
	file_ttn_lorawan_v3_identifiers_proto_init()
	file_ttn_lorawan_v3_rights_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ttn_lorawan_v3_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ttn_lorawan_v3_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ttn_lorawan_v3_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRelatedEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ttn_lorawan_v3_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRelatedEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ttn_lorawan_v3_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Authentication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ttn_lorawan_v3_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ttn_lorawan_v3_events_proto_goTypes,
		DependencyIndexes: file_ttn_lorawan_v3_events_proto_depIdxs,
		MessageInfos:      file_ttn_lorawan_v3_events_proto_msgTypes,
	}.Build()
	File_ttn_lorawan_v3_events_proto = out.File
	file_ttn_lorawan_v3_events_proto_rawDesc = nil
	file_ttn_lorawan_v3_events_proto_goTypes = nil
	file_ttn_lorawan_v3_events_proto_depIdxs = nil
}
