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
// source: ttn/lorawan/v3/user_services.proto

package ttnpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ttn_lorawan_v3_user_services_proto protoreflect.FileDescriptor

var file_ttn_lorawan_v3_user_services_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76,
	0x33, 0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2, 0x07, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x5e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x4f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x67, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a,
	0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2e, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x2c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x82, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x27, 0x3a, 0x01, 0x2a, 0x1a, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x1f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x18, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x60, 0x0a, 0x05,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x32, 0x9d,
	0x07, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x66, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x7e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01,
	0x2a, 0x22, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x7a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79,
	0x73, 0x12, 0x7e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x24,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x33, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x1a, 0x2f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65,
	0x79, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x69, 0x64, 0x7d, 0x12,
	0x84, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79,
	0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x2a, 0x2b, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x95, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x26, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x32, 0xc0,
	0x02, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x62, 0x0a, 0x04, 0x53, 0x65, 0x6e,
	0x64, 0x12, 0x25, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x5f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x32, 0x94, 0x02, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x79, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24,
	0x12, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x2a, 0x2f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76,
	0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_ttn_lorawan_v3_user_services_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),              // 0: ttn.lorawan.v3.CreateUserRequest
	(*GetUserRequest)(nil),                 // 1: ttn.lorawan.v3.GetUserRequest
	(*ListUsersRequest)(nil),               // 2: ttn.lorawan.v3.ListUsersRequest
	(*UpdateUserRequest)(nil),              // 3: ttn.lorawan.v3.UpdateUserRequest
	(*CreateTemporaryPasswordRequest)(nil), // 4: ttn.lorawan.v3.CreateTemporaryPasswordRequest
	(*UpdateUserPasswordRequest)(nil),      // 5: ttn.lorawan.v3.UpdateUserPasswordRequest
	(*UserIdentifiers)(nil),                // 6: ttn.lorawan.v3.UserIdentifiers
	(*CreateUserAPIKeyRequest)(nil),        // 7: ttn.lorawan.v3.CreateUserAPIKeyRequest
	(*ListUserAPIKeysRequest)(nil),         // 8: ttn.lorawan.v3.ListUserAPIKeysRequest
	(*GetUserAPIKeyRequest)(nil),           // 9: ttn.lorawan.v3.GetUserAPIKeyRequest
	(*UpdateUserAPIKeyRequest)(nil),        // 10: ttn.lorawan.v3.UpdateUserAPIKeyRequest
	(*DeleteUserAPIKeyRequest)(nil),        // 11: ttn.lorawan.v3.DeleteUserAPIKeyRequest
	(*CreateLoginTokenRequest)(nil),        // 12: ttn.lorawan.v3.CreateLoginTokenRequest
	(*SendInvitationRequest)(nil),          // 13: ttn.lorawan.v3.SendInvitationRequest
	(*ListInvitationsRequest)(nil),         // 14: ttn.lorawan.v3.ListInvitationsRequest
	(*DeleteInvitationRequest)(nil),        // 15: ttn.lorawan.v3.DeleteInvitationRequest
	(*ListUserSessionsRequest)(nil),        // 16: ttn.lorawan.v3.ListUserSessionsRequest
	(*UserSessionIdentifiers)(nil),         // 17: ttn.lorawan.v3.UserSessionIdentifiers
	(*User)(nil),                           // 18: ttn.lorawan.v3.User
	(*Users)(nil),                          // 19: ttn.lorawan.v3.Users
	(*emptypb.Empty)(nil),                  // 20: google.protobuf.Empty
	(*Rights)(nil),                         // 21: ttn.lorawan.v3.Rights
	(*APIKey)(nil),                         // 22: ttn.lorawan.v3.APIKey
	(*APIKeys)(nil),                        // 23: ttn.lorawan.v3.APIKeys
	(*CreateLoginTokenResponse)(nil),       // 24: ttn.lorawan.v3.CreateLoginTokenResponse
	(*Invitation)(nil),                     // 25: ttn.lorawan.v3.Invitation
	(*Invitations)(nil),                    // 26: ttn.lorawan.v3.Invitations
	(*UserSessions)(nil),                   // 27: ttn.lorawan.v3.UserSessions
}
var file_ttn_lorawan_v3_user_services_proto_depIdxs = []int32{
	0,  // 0: ttn.lorawan.v3.UserRegistry.Create:input_type -> ttn.lorawan.v3.CreateUserRequest
	1,  // 1: ttn.lorawan.v3.UserRegistry.Get:input_type -> ttn.lorawan.v3.GetUserRequest
	2,  // 2: ttn.lorawan.v3.UserRegistry.List:input_type -> ttn.lorawan.v3.ListUsersRequest
	3,  // 3: ttn.lorawan.v3.UserRegistry.Update:input_type -> ttn.lorawan.v3.UpdateUserRequest
	4,  // 4: ttn.lorawan.v3.UserRegistry.CreateTemporaryPassword:input_type -> ttn.lorawan.v3.CreateTemporaryPasswordRequest
	5,  // 5: ttn.lorawan.v3.UserRegistry.UpdatePassword:input_type -> ttn.lorawan.v3.UpdateUserPasswordRequest
	6,  // 6: ttn.lorawan.v3.UserRegistry.Delete:input_type -> ttn.lorawan.v3.UserIdentifiers
	6,  // 7: ttn.lorawan.v3.UserRegistry.Restore:input_type -> ttn.lorawan.v3.UserIdentifiers
	6,  // 8: ttn.lorawan.v3.UserRegistry.Purge:input_type -> ttn.lorawan.v3.UserIdentifiers
	6,  // 9: ttn.lorawan.v3.UserAccess.ListRights:input_type -> ttn.lorawan.v3.UserIdentifiers
	7,  // 10: ttn.lorawan.v3.UserAccess.CreateAPIKey:input_type -> ttn.lorawan.v3.CreateUserAPIKeyRequest
	8,  // 11: ttn.lorawan.v3.UserAccess.ListAPIKeys:input_type -> ttn.lorawan.v3.ListUserAPIKeysRequest
	9,  // 12: ttn.lorawan.v3.UserAccess.GetAPIKey:input_type -> ttn.lorawan.v3.GetUserAPIKeyRequest
	10, // 13: ttn.lorawan.v3.UserAccess.UpdateAPIKey:input_type -> ttn.lorawan.v3.UpdateUserAPIKeyRequest
	11, // 14: ttn.lorawan.v3.UserAccess.DeleteAPIKey:input_type -> ttn.lorawan.v3.DeleteUserAPIKeyRequest
	12, // 15: ttn.lorawan.v3.UserAccess.CreateLoginToken:input_type -> ttn.lorawan.v3.CreateLoginTokenRequest
	13, // 16: ttn.lorawan.v3.UserInvitationRegistry.Send:input_type -> ttn.lorawan.v3.SendInvitationRequest
	14, // 17: ttn.lorawan.v3.UserInvitationRegistry.List:input_type -> ttn.lorawan.v3.ListInvitationsRequest
	15, // 18: ttn.lorawan.v3.UserInvitationRegistry.Delete:input_type -> ttn.lorawan.v3.DeleteInvitationRequest
	16, // 19: ttn.lorawan.v3.UserSessionRegistry.List:input_type -> ttn.lorawan.v3.ListUserSessionsRequest
	17, // 20: ttn.lorawan.v3.UserSessionRegistry.Delete:input_type -> ttn.lorawan.v3.UserSessionIdentifiers
	18, // 21: ttn.lorawan.v3.UserRegistry.Create:output_type -> ttn.lorawan.v3.User
	18, // 22: ttn.lorawan.v3.UserRegistry.Get:output_type -> ttn.lorawan.v3.User
	19, // 23: ttn.lorawan.v3.UserRegistry.List:output_type -> ttn.lorawan.v3.Users
	18, // 24: ttn.lorawan.v3.UserRegistry.Update:output_type -> ttn.lorawan.v3.User
	20, // 25: ttn.lorawan.v3.UserRegistry.CreateTemporaryPassword:output_type -> google.protobuf.Empty
	20, // 26: ttn.lorawan.v3.UserRegistry.UpdatePassword:output_type -> google.protobuf.Empty
	20, // 27: ttn.lorawan.v3.UserRegistry.Delete:output_type -> google.protobuf.Empty
	20, // 28: ttn.lorawan.v3.UserRegistry.Restore:output_type -> google.protobuf.Empty
	20, // 29: ttn.lorawan.v3.UserRegistry.Purge:output_type -> google.protobuf.Empty
	21, // 30: ttn.lorawan.v3.UserAccess.ListRights:output_type -> ttn.lorawan.v3.Rights
	22, // 31: ttn.lorawan.v3.UserAccess.CreateAPIKey:output_type -> ttn.lorawan.v3.APIKey
	23, // 32: ttn.lorawan.v3.UserAccess.ListAPIKeys:output_type -> ttn.lorawan.v3.APIKeys
	22, // 33: ttn.lorawan.v3.UserAccess.GetAPIKey:output_type -> ttn.lorawan.v3.APIKey
	22, // 34: ttn.lorawan.v3.UserAccess.UpdateAPIKey:output_type -> ttn.lorawan.v3.APIKey
	20, // 35: ttn.lorawan.v3.UserAccess.DeleteAPIKey:output_type -> google.protobuf.Empty
	24, // 36: ttn.lorawan.v3.UserAccess.CreateLoginToken:output_type -> ttn.lorawan.v3.CreateLoginTokenResponse
	25, // 37: ttn.lorawan.v3.UserInvitationRegistry.Send:output_type -> ttn.lorawan.v3.Invitation
	26, // 38: ttn.lorawan.v3.UserInvitationRegistry.List:output_type -> ttn.lorawan.v3.Invitations
	20, // 39: ttn.lorawan.v3.UserInvitationRegistry.Delete:output_type -> google.protobuf.Empty
	27, // 40: ttn.lorawan.v3.UserSessionRegistry.List:output_type -> ttn.lorawan.v3.UserSessions
	20, // 41: ttn.lorawan.v3.UserSessionRegistry.Delete:output_type -> google.protobuf.Empty
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ttn_lorawan_v3_user_services_proto_init() }
func file_ttn_lorawan_v3_user_services_proto_init() {
	if File_ttn_lorawan_v3_user_services_proto != nil {
		return
	}
	file_ttn_lorawan_v3_identifiers_proto_init()
	file_ttn_lorawan_v3_rights_proto_init()
	file_ttn_lorawan_v3_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ttn_lorawan_v3_user_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_ttn_lorawan_v3_user_services_proto_goTypes,
		DependencyIndexes: file_ttn_lorawan_v3_user_services_proto_depIdxs,
	}.Build()
	File_ttn_lorawan_v3_user_services_proto = out.File
	file_ttn_lorawan_v3_user_services_proto_rawDesc = nil
	file_ttn_lorawan_v3_user_services_proto_goTypes = nil
	file_ttn_lorawan_v3_user_services_proto_depIdxs = nil
}
