// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/rights.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Right is the enum that defines all the different rights to do something in the network.
type Right int32

const (
	Right_right_invalid Right = 0
	// The right to view user information.
	Right_RIGHT_USER_INFO Right = 1
	// The right to edit basic user settings.
	Right_RIGHT_USER_SETTINGS_BASIC Right = 2
	// The right to view and edit user API keys.
	Right_RIGHT_USER_SETTINGS_API_KEYS Right = 3
	// The right to delete user account.
	Right_RIGHT_USER_DELETE Right = 4
	// The right to view and edit authorized OAuth clients of the user.
	Right_RIGHT_USER_AUTHORIZED_CLIENTS Right = 5
	// The right to list applications the user is a collaborator of.
	Right_RIGHT_USER_APPLICATIONS_LIST Right = 6
	// The right to create an application under the user account.
	Right_RIGHT_USER_APPLICATIONS_CREATE Right = 7
	// The right to list gateways the user is a collaborator of.
	Right_RIGHT_USER_GATEWAYS_LIST Right = 8
	// The right to create a gateway under the account of the user.
	Right_RIGHT_USER_GATEWAYS_CREATE Right = 9
	// The right to list OAuth clients the user is a collaborator of.
	Right_RIGHT_USER_CLIENTS_LIST Right = 10
	// The right to create an OAuth client under the account of the user.
	Right_RIGHT_USER_CLIENTS_CREATE Right = 11
	// The right to list organizations the user is a member of.
	Right_RIGHT_USER_ORGANIZATIONS_LIST Right = 12
	// The right to create an organization under the user account.
	Right_RIGHT_USER_ORGANIZATIONS_CREATE Right = 13
	// The right to read notifications sent to the user.
	Right_RIGHT_USER_NOTIFICATIONS_READ Right = 59
	// The pseudo-right for all (current and future) user rights.
	Right_RIGHT_USER_ALL Right = 14
	// The right to view application information.
	Right_RIGHT_APPLICATION_INFO Right = 15
	// The right to edit basic application settings.
	Right_RIGHT_APPLICATION_SETTINGS_BASIC Right = 16
	// The right to view and edit application API keys.
	Right_RIGHT_APPLICATION_SETTINGS_API_KEYS Right = 17
	// The right to view and edit application collaborators.
	Right_RIGHT_APPLICATION_SETTINGS_COLLABORATORS Right = 18
	// The right to view and edit application packages and associations.
	Right_RIGHT_APPLICATION_SETTINGS_PACKAGES Right = 56
	// The right to delete application.
	Right_RIGHT_APPLICATION_DELETE Right = 19
	// The right to view devices in application.
	Right_RIGHT_APPLICATION_DEVICES_READ Right = 20
	// The right to create devices in application.
	Right_RIGHT_APPLICATION_DEVICES_WRITE Right = 21
	// The right to view device keys in application.
	// Note that keys may not be stored in a way that supports viewing them.
	Right_RIGHT_APPLICATION_DEVICES_READ_KEYS Right = 22
	// The right to edit device keys in application.
	Right_RIGHT_APPLICATION_DEVICES_WRITE_KEYS Right = 23
	// The right to read application traffic (uplink and downlink).
	Right_RIGHT_APPLICATION_TRAFFIC_READ Right = 24
	// The right to write uplink application traffic.
	Right_RIGHT_APPLICATION_TRAFFIC_UP_WRITE Right = 25
	// The right to write downlink application traffic.
	Right_RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE Right = 26
	// The right to link as Application to a Network Server for traffic exchange,
	// i.e. read uplink and write downlink (API keys only).
	// This right is typically only given to an Application Server.
	// This right implies RIGHT_APPLICATION_INFO, RIGHT_APPLICATION_TRAFFIC_READ,
	// and RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE.
	Right_RIGHT_APPLICATION_LINK Right = 27
	// The pseudo-right for all (current and future) application rights.
	Right_RIGHT_APPLICATION_ALL Right = 28
	// The pseudo-right for all (current and future) OAuth client rights.
	Right_RIGHT_CLIENT_ALL Right = 29
	// The right to view gateway information.
	Right_RIGHT_GATEWAY_INFO Right = 30
	// The right to edit basic gateway settings.
	Right_RIGHT_GATEWAY_SETTINGS_BASIC Right = 31
	// The right to view and edit gateway API keys.
	Right_RIGHT_GATEWAY_SETTINGS_API_KEYS Right = 32
	// The right to view and edit gateway collaborators.
	Right_RIGHT_GATEWAY_SETTINGS_COLLABORATORS Right = 33
	// The right to delete gateway.
	Right_RIGHT_GATEWAY_DELETE Right = 34
	// The right to read gateway traffic.
	Right_RIGHT_GATEWAY_TRAFFIC_READ Right = 35
	// The right to write downlink gateway traffic.
	Right_RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE Right = 36
	// The right to link as Gateway to a Gateway Server for traffic exchange,
	// i.e. write uplink and read downlink (API keys only)
	// This right is typically only given to a gateway.
	// This right implies RIGHT_GATEWAY_INFO.
	Right_RIGHT_GATEWAY_LINK Right = 37
	// The right to view gateway status.
	Right_RIGHT_GATEWAY_STATUS_READ Right = 38
	// The right to view view gateway location.
	Right_RIGHT_GATEWAY_LOCATION_READ Right = 39
	// The right to store secrets associated with this gateway.
	Right_RIGHT_GATEWAY_WRITE_SECRETS Right = 57
	// The right to retrieve secrets associated with this gateway.
	Right_RIGHT_GATEWAY_READ_SECRETS Right = 58
	// The pseudo-right for all (current and future) gateway rights.
	Right_RIGHT_GATEWAY_ALL Right = 40
	// The right to view organization information.
	Right_RIGHT_ORGANIZATION_INFO Right = 41
	// The right to edit basic organization settings.
	Right_RIGHT_ORGANIZATION_SETTINGS_BASIC Right = 42
	// The right to view and edit organization API keys.
	Right_RIGHT_ORGANIZATION_SETTINGS_API_KEYS Right = 43
	// The right to view and edit organization members.
	Right_RIGHT_ORGANIZATION_SETTINGS_MEMBERS Right = 44
	// The right to delete organization.
	Right_RIGHT_ORGANIZATION_DELETE Right = 45
	// The right to list the applications the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST Right = 46
	// The right to create an application under the organization.
	Right_RIGHT_ORGANIZATION_APPLICATIONS_CREATE Right = 47
	// The right to list the gateways the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_GATEWAYS_LIST Right = 48
	// The right to create a gateway under the organization.
	Right_RIGHT_ORGANIZATION_GATEWAYS_CREATE Right = 49
	// The right to list the OAuth clients the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_CLIENTS_LIST Right = 50
	// The right to create an OAuth client under the organization.
	Right_RIGHT_ORGANIZATION_CLIENTS_CREATE Right = 51
	// The right to add the organization as a collaborator on an existing entity.
	Right_RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR Right = 52
	// The pseudo-right for all (current and future) organization rights.
	Right_RIGHT_ORGANIZATION_ALL Right = 53
	// The right to send invites to new users.
	// Note that this is not prefixed with "USER_"; it is not a right on the user entity.
	Right_RIGHT_SEND_INVITES Right = 54
	// The pseudo-right for all (current and future) possible rights.
	Right_RIGHT_ALL Right = 55
)

var Right_name = map[int32]string{
	0:  "right_invalid",
	1:  "RIGHT_USER_INFO",
	2:  "RIGHT_USER_SETTINGS_BASIC",
	3:  "RIGHT_USER_SETTINGS_API_KEYS",
	4:  "RIGHT_USER_DELETE",
	5:  "RIGHT_USER_AUTHORIZED_CLIENTS",
	6:  "RIGHT_USER_APPLICATIONS_LIST",
	7:  "RIGHT_USER_APPLICATIONS_CREATE",
	8:  "RIGHT_USER_GATEWAYS_LIST",
	9:  "RIGHT_USER_GATEWAYS_CREATE",
	10: "RIGHT_USER_CLIENTS_LIST",
	11: "RIGHT_USER_CLIENTS_CREATE",
	12: "RIGHT_USER_ORGANIZATIONS_LIST",
	13: "RIGHT_USER_ORGANIZATIONS_CREATE",
	59: "RIGHT_USER_NOTIFICATIONS_READ",
	14: "RIGHT_USER_ALL",
	15: "RIGHT_APPLICATION_INFO",
	16: "RIGHT_APPLICATION_SETTINGS_BASIC",
	17: "RIGHT_APPLICATION_SETTINGS_API_KEYS",
	18: "RIGHT_APPLICATION_SETTINGS_COLLABORATORS",
	56: "RIGHT_APPLICATION_SETTINGS_PACKAGES",
	19: "RIGHT_APPLICATION_DELETE",
	20: "RIGHT_APPLICATION_DEVICES_READ",
	21: "RIGHT_APPLICATION_DEVICES_WRITE",
	22: "RIGHT_APPLICATION_DEVICES_READ_KEYS",
	23: "RIGHT_APPLICATION_DEVICES_WRITE_KEYS",
	24: "RIGHT_APPLICATION_TRAFFIC_READ",
	25: "RIGHT_APPLICATION_TRAFFIC_UP_WRITE",
	26: "RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE",
	27: "RIGHT_APPLICATION_LINK",
	28: "RIGHT_APPLICATION_ALL",
	29: "RIGHT_CLIENT_ALL",
	30: "RIGHT_GATEWAY_INFO",
	31: "RIGHT_GATEWAY_SETTINGS_BASIC",
	32: "RIGHT_GATEWAY_SETTINGS_API_KEYS",
	33: "RIGHT_GATEWAY_SETTINGS_COLLABORATORS",
	34: "RIGHT_GATEWAY_DELETE",
	35: "RIGHT_GATEWAY_TRAFFIC_READ",
	36: "RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE",
	37: "RIGHT_GATEWAY_LINK",
	38: "RIGHT_GATEWAY_STATUS_READ",
	39: "RIGHT_GATEWAY_LOCATION_READ",
	57: "RIGHT_GATEWAY_WRITE_SECRETS",
	58: "RIGHT_GATEWAY_READ_SECRETS",
	40: "RIGHT_GATEWAY_ALL",
	41: "RIGHT_ORGANIZATION_INFO",
	42: "RIGHT_ORGANIZATION_SETTINGS_BASIC",
	43: "RIGHT_ORGANIZATION_SETTINGS_API_KEYS",
	44: "RIGHT_ORGANIZATION_SETTINGS_MEMBERS",
	45: "RIGHT_ORGANIZATION_DELETE",
	46: "RIGHT_ORGANIZATION_APPLICATIONS_LIST",
	47: "RIGHT_ORGANIZATION_APPLICATIONS_CREATE",
	48: "RIGHT_ORGANIZATION_GATEWAYS_LIST",
	49: "RIGHT_ORGANIZATION_GATEWAYS_CREATE",
	50: "RIGHT_ORGANIZATION_CLIENTS_LIST",
	51: "RIGHT_ORGANIZATION_CLIENTS_CREATE",
	52: "RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR",
	53: "RIGHT_ORGANIZATION_ALL",
	54: "RIGHT_SEND_INVITES",
	55: "RIGHT_ALL",
}

var Right_value = map[string]int32{
	"right_invalid":                            0,
	"RIGHT_USER_INFO":                          1,
	"RIGHT_USER_SETTINGS_BASIC":                2,
	"RIGHT_USER_SETTINGS_API_KEYS":             3,
	"RIGHT_USER_DELETE":                        4,
	"RIGHT_USER_AUTHORIZED_CLIENTS":            5,
	"RIGHT_USER_APPLICATIONS_LIST":             6,
	"RIGHT_USER_APPLICATIONS_CREATE":           7,
	"RIGHT_USER_GATEWAYS_LIST":                 8,
	"RIGHT_USER_GATEWAYS_CREATE":               9,
	"RIGHT_USER_CLIENTS_LIST":                  10,
	"RIGHT_USER_CLIENTS_CREATE":                11,
	"RIGHT_USER_ORGANIZATIONS_LIST":            12,
	"RIGHT_USER_ORGANIZATIONS_CREATE":          13,
	"RIGHT_USER_NOTIFICATIONS_READ":            59,
	"RIGHT_USER_ALL":                           14,
	"RIGHT_APPLICATION_INFO":                   15,
	"RIGHT_APPLICATION_SETTINGS_BASIC":         16,
	"RIGHT_APPLICATION_SETTINGS_API_KEYS":      17,
	"RIGHT_APPLICATION_SETTINGS_COLLABORATORS": 18,
	"RIGHT_APPLICATION_SETTINGS_PACKAGES":      56,
	"RIGHT_APPLICATION_DELETE":                 19,
	"RIGHT_APPLICATION_DEVICES_READ":           20,
	"RIGHT_APPLICATION_DEVICES_WRITE":          21,
	"RIGHT_APPLICATION_DEVICES_READ_KEYS":      22,
	"RIGHT_APPLICATION_DEVICES_WRITE_KEYS":     23,
	"RIGHT_APPLICATION_TRAFFIC_READ":           24,
	"RIGHT_APPLICATION_TRAFFIC_UP_WRITE":       25,
	"RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE":     26,
	"RIGHT_APPLICATION_LINK":                   27,
	"RIGHT_APPLICATION_ALL":                    28,
	"RIGHT_CLIENT_ALL":                         29,
	"RIGHT_GATEWAY_INFO":                       30,
	"RIGHT_GATEWAY_SETTINGS_BASIC":             31,
	"RIGHT_GATEWAY_SETTINGS_API_KEYS":          32,
	"RIGHT_GATEWAY_SETTINGS_COLLABORATORS":     33,
	"RIGHT_GATEWAY_DELETE":                     34,
	"RIGHT_GATEWAY_TRAFFIC_READ":               35,
	"RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE":         36,
	"RIGHT_GATEWAY_LINK":                       37,
	"RIGHT_GATEWAY_STATUS_READ":                38,
	"RIGHT_GATEWAY_LOCATION_READ":              39,
	"RIGHT_GATEWAY_WRITE_SECRETS":              57,
	"RIGHT_GATEWAY_READ_SECRETS":               58,
	"RIGHT_GATEWAY_ALL":                        40,
	"RIGHT_ORGANIZATION_INFO":                  41,
	"RIGHT_ORGANIZATION_SETTINGS_BASIC":        42,
	"RIGHT_ORGANIZATION_SETTINGS_API_KEYS":     43,
	"RIGHT_ORGANIZATION_SETTINGS_MEMBERS":      44,
	"RIGHT_ORGANIZATION_DELETE":                45,
	"RIGHT_ORGANIZATION_APPLICATIONS_LIST":     46,
	"RIGHT_ORGANIZATION_APPLICATIONS_CREATE":   47,
	"RIGHT_ORGANIZATION_GATEWAYS_LIST":         48,
	"RIGHT_ORGANIZATION_GATEWAYS_CREATE":       49,
	"RIGHT_ORGANIZATION_CLIENTS_LIST":          50,
	"RIGHT_ORGANIZATION_CLIENTS_CREATE":        51,
	"RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR":   52,
	"RIGHT_ORGANIZATION_ALL":                   53,
	"RIGHT_SEND_INVITES":                       54,
	"RIGHT_ALL":                                55,
}

func (x Right) String() string {
	return proto.EnumName(Right_name, int32(x))
}

func (Right) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{0}
}

type Rights struct {
	Rights               []Right  `protobuf:"varint,1,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rights) Reset()         { *m = Rights{} }
func (m *Rights) String() string { return proto.CompactTextString(m) }
func (*Rights) ProtoMessage()    {}
func (*Rights) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{0}
}
func (m *Rights) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rights.Unmarshal(m, b)
}
func (m *Rights) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rights.Marshal(b, m, deterministic)
}
func (m *Rights) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rights.Merge(m, src)
}
func (m *Rights) XXX_Size() int {
	return xxx_messageInfo_Rights.Size(m)
}
func (m *Rights) XXX_DiscardUnknown() {
	xxx_messageInfo_Rights.DiscardUnknown(m)
}

var xxx_messageInfo_Rights proto.InternalMessageInfo

func (m *Rights) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type APIKey struct {
	// Immutable and unique public identifier for the API key.
	// Generated by the Access Server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable and unique secret value of the API key.
	// Generated by the Access Server.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// User-defined (friendly) name for the API key.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Rights that are granted to this API key.
	Rights               []Right          `protobuf:"varint,4,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	CreatedAt            *types.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ExpiresAt            *types.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *APIKey) Reset()         { *m = APIKey{} }
func (m *APIKey) String() string { return proto.CompactTextString(m) }
func (*APIKey) ProtoMessage()    {}
func (*APIKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{1}
}
func (m *APIKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKey.Unmarshal(m, b)
}
func (m *APIKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKey.Marshal(b, m, deterministic)
}
func (m *APIKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKey.Merge(m, src)
}
func (m *APIKey) XXX_Size() int {
	return xxx_messageInfo_APIKey.Size(m)
}
func (m *APIKey) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKey.DiscardUnknown(m)
}

var xxx_messageInfo_APIKey proto.InternalMessageInfo

func (m *APIKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *APIKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *APIKey) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *APIKey) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *APIKey) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *APIKey) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type APIKeys struct {
	ApiKeys              []*APIKey `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *APIKeys) Reset()         { *m = APIKeys{} }
func (m *APIKeys) String() string { return proto.CompactTextString(m) }
func (*APIKeys) ProtoMessage()    {}
func (*APIKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{2}
}
func (m *APIKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKeys.Unmarshal(m, b)
}
func (m *APIKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKeys.Marshal(b, m, deterministic)
}
func (m *APIKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKeys.Merge(m, src)
}
func (m *APIKeys) XXX_Size() int {
	return xxx_messageInfo_APIKeys.Size(m)
}
func (m *APIKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKeys.DiscardUnknown(m)
}

var xxx_messageInfo_APIKeys proto.InternalMessageInfo

func (m *APIKeys) GetApiKeys() []*APIKey {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

type Collaborator struct {
	Ids                  *OrganizationOrUserIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	Rights               []Right                        `protobuf:"varint,2,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Collaborator) Reset()         { *m = Collaborator{} }
func (m *Collaborator) String() string { return proto.CompactTextString(m) }
func (*Collaborator) ProtoMessage()    {}
func (*Collaborator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{3}
}
func (m *Collaborator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collaborator.Unmarshal(m, b)
}
func (m *Collaborator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collaborator.Marshal(b, m, deterministic)
}
func (m *Collaborator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collaborator.Merge(m, src)
}
func (m *Collaborator) XXX_Size() int {
	return xxx_messageInfo_Collaborator.Size(m)
}
func (m *Collaborator) XXX_DiscardUnknown() {
	xxx_messageInfo_Collaborator.DiscardUnknown(m)
}

var xxx_messageInfo_Collaborator proto.InternalMessageInfo

func (m *Collaborator) GetIds() *OrganizationOrUserIdentifiers {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Collaborator) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type GetCollaboratorResponse struct {
	Ids                  *OrganizationOrUserIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	Rights               []Right                        `protobuf:"varint,2,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetCollaboratorResponse) Reset()         { *m = GetCollaboratorResponse{} }
func (m *GetCollaboratorResponse) String() string { return proto.CompactTextString(m) }
func (*GetCollaboratorResponse) ProtoMessage()    {}
func (*GetCollaboratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{4}
}
func (m *GetCollaboratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollaboratorResponse.Unmarshal(m, b)
}
func (m *GetCollaboratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollaboratorResponse.Marshal(b, m, deterministic)
}
func (m *GetCollaboratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollaboratorResponse.Merge(m, src)
}
func (m *GetCollaboratorResponse) XXX_Size() int {
	return xxx_messageInfo_GetCollaboratorResponse.Size(m)
}
func (m *GetCollaboratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollaboratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollaboratorResponse proto.InternalMessageInfo

func (m *GetCollaboratorResponse) GetIds() *OrganizationOrUserIdentifiers {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GetCollaboratorResponse) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type Collaborators struct {
	Collaborators        []*Collaborator `protobuf:"bytes,1,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Collaborators) Reset()         { *m = Collaborators{} }
func (m *Collaborators) String() string { return proto.CompactTextString(m) }
func (*Collaborators) ProtoMessage()    {}
func (*Collaborators) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{5}
}
func (m *Collaborators) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collaborators.Unmarshal(m, b)
}
func (m *Collaborators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collaborators.Marshal(b, m, deterministic)
}
func (m *Collaborators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collaborators.Merge(m, src)
}
func (m *Collaborators) XXX_Size() int {
	return xxx_messageInfo_Collaborators.Size(m)
}
func (m *Collaborators) XXX_DiscardUnknown() {
	xxx_messageInfo_Collaborators.DiscardUnknown(m)
}

var xxx_messageInfo_Collaborators proto.InternalMessageInfo

func (m *Collaborators) GetCollaborators() []*Collaborator {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.Right", Right_name, Right_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.Right", Right_name, Right_value)
	proto.RegisterType((*Rights)(nil), "ttn.lorawan.v3.Rights")
	golang_proto.RegisterType((*Rights)(nil), "ttn.lorawan.v3.Rights")
	proto.RegisterType((*APIKey)(nil), "ttn.lorawan.v3.APIKey")
	golang_proto.RegisterType((*APIKey)(nil), "ttn.lorawan.v3.APIKey")
	proto.RegisterType((*APIKeys)(nil), "ttn.lorawan.v3.APIKeys")
	golang_proto.RegisterType((*APIKeys)(nil), "ttn.lorawan.v3.APIKeys")
	proto.RegisterType((*Collaborator)(nil), "ttn.lorawan.v3.Collaborator")
	golang_proto.RegisterType((*Collaborator)(nil), "ttn.lorawan.v3.Collaborator")
	proto.RegisterType((*GetCollaboratorResponse)(nil), "ttn.lorawan.v3.GetCollaboratorResponse")
	golang_proto.RegisterType((*GetCollaboratorResponse)(nil), "ttn.lorawan.v3.GetCollaboratorResponse")
	proto.RegisterType((*Collaborators)(nil), "ttn.lorawan.v3.Collaborators")
	golang_proto.RegisterType((*Collaborators)(nil), "ttn.lorawan.v3.Collaborators")
}

func init() { proto.RegisterFile("lorawan-stack/api/rights.proto", fileDescriptor_9bb69af2cf8904c5) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/rights.proto", fileDescriptor_9bb69af2cf8904c5)
}

var fileDescriptor_9bb69af2cf8904c5 = []byte{
	// 1251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4d, 0x53, 0xdb, 0x56,
	0x14, 0xad, 0xfc, 0xed, 0x47, 0x4c, 0x5e, 0x5e, 0x02, 0x31, 0x86, 0x80, 0x63, 0xf2, 0xe1, 0xd2,
	0x58, 0x6e, 0x4c, 0xd3, 0x36, 0x69, 0x66, 0x5a, 0x59, 0x16, 0x8e, 0x40, 0xb1, 0x3c, 0x92, 0x08,
	0x13, 0x36, 0x1e, 0x81, 0x85, 0x50, 0x31, 0x92, 0x47, 0x7a, 0x90, 0xd0, 0x65, 0x97, 0x5d, 0xf6,
	0x1f, 0x74, 0x9b, 0x4d, 0x67, 0xfa, 0x0b, 0xba, 0xee, 0xdf, 0xe8, 0x74, 0xd3, 0x65, 0x97, 0xac,
	0x3a, 0x96, 0x9e, 0xd1, 0x87, 0xe5, 0xb8, 0x74, 0x27, 0xdd, 0x7b, 0xee, 0xd5, 0x79, 0xe7, 0x9e,
	0x77, 0x47, 0x60, 0x75, 0x60, 0xd9, 0xea, 0x3b, 0xd5, 0xac, 0x39, 0x58, 0x3d, 0x3c, 0xa9, 0xab,
	0x43, 0xa3, 0x6e, 0x1b, 0xfa, 0x31, 0x76, 0xe8, 0xa1, 0x6d, 0x61, 0x0b, 0xcd, 0x63, 0x6c, 0xd2,
	0x04, 0x43, 0x9f, 0x6f, 0x96, 0x5a, 0xba, 0x81, 0x8f, 0xcf, 0x0e, 0xe8, 0x43, 0xeb, 0xb4, 0xae,
	0x1c, 0x6b, 0xca, 0xb1, 0x61, 0xea, 0x0e, 0x6f, 0xf6, 0xcf, 0x1c, 0x6c, 0x1b, 0x9a, 0x53, 0x77,
	0xab, 0x0e, 0x6b, 0xba, 0x66, 0xd6, 0x74, 0xab, 0x76, 0x34, 0x50, 0x75, 0xa7, 0xae, 0x9a, 0xa6,
	0x85, 0x55, 0x6c, 0x58, 0x26, 0xe9, 0x5a, 0x62, 0x02, 0x5d, 0x34, 0xf3, 0xdc, 0xba, 0x18, 0xda,
	0xd6, 0xfb, 0x8b, 0x60, 0xf1, 0xb9, 0x3a, 0x30, 0xfa, 0x2a, 0xd6, 0xea, 0x13, 0x0f, 0xa4, 0x45,
	0x2d, 0xd0, 0x42, 0xb7, 0x74, 0xcb, 0x2b, 0x3e, 0x38, 0x3b, 0x72, 0xdf, 0xdc, 0x17, 0xf7, 0x89,
	0xc0, 0xd9, 0x6b, 0xf1, 0xfe, 0xde, 0xb1, 0xcc, 0x18, 0xda, 0x6b, 0xba, 0x65, 0xe9, 0x03, 0xcd,
	0xff, 0x14, 0x36, 0x4e, 0x35, 0x07, 0xab, 0xa7, 0x43, 0x02, 0x58, 0x9f, 0x54, 0xd3, 0xe8, 0x6b,
	0x26, 0x36, 0x8e, 0x0c, 0xcd, 0x26, 0x5d, 0x2a, 0x5b, 0x20, 0x23, 0xb9, 0x12, 0xa3, 0x97, 0x20,
	0xe3, 0x89, 0x5d, 0xa4, 0xca, 0xc9, 0xea, 0x7c, 0x63, 0x81, 0x0e, 0xab, 0x4d, 0xbb, 0xb8, 0x66,
	0xe1, 0xb2, 0x09, 0x7e, 0xa6, 0xb2, 0x95, 0xf4, 0x8f, 0x54, 0x02, 0x52, 0x12, 0xa9, 0xa9, 0xfc,
	0x91, 0x00, 0x19, 0xa6, 0xcb, 0xef, 0x68, 0x17, 0x68, 0x1e, 0x24, 0x8c, 0x7e, 0x91, 0x2a, 0x53,
	0xd5, 0xbc, 0x94, 0x30, 0xfa, 0x08, 0x82, 0xe4, 0x89, 0x76, 0x51, 0x4c, 0xb8, 0x81, 0xd1, 0x23,
	0x5a, 0x06, 0x29, 0x53, 0x3d, 0xd5, 0x8a, 0xc9, 0x51, 0xa8, 0x99, 0xbd, 0x6c, 0xa6, 0xec, 0x44,
	0xb1, 0x21, 0xb9, 0xc1, 0x00, 0x8f, 0xd4, 0xf5, 0x79, 0xa0, 0xe7, 0x00, 0x1c, 0xda, 0x9a, 0x8a,
	0xb5, 0x7e, 0x4f, 0xc5, 0xc5, 0x74, 0x99, 0xaa, 0xce, 0x35, 0x4a, 0xb4, 0x27, 0x15, 0x3d, 0x96,
	0x8a, 0x56, 0xc6, 0x52, 0x49, 0x79, 0x82, 0x66, 0xf0, 0xa8, 0xf4, 0x6c, 0xd8, 0x1f, 0x97, 0x66,
	0x66, 0x97, 0x12, 0x34, 0x83, 0x11, 0x0b, 0x80, 0xf6, 0x7e, 0x68, 0xd8, 0x9a, 0x33, 0x2a, 0xcd,
	0xce, 0x2a, 0x6d, 0xe6, 0x2e, 0x9b, 0xe9, 0xdf, 0xa8, 0xc4, 0x77, 0x94, 0x94, 0x27, 0x75, 0x0c,
	0x7e, 0x91, 0xfb, 0xe7, 0xc3, 0x52, 0x2a, 0x47, 0x41, 0xaa, 0xf2, 0x12, 0x64, 0x3d, 0x2d, 0x1d,
	0xf4, 0x14, 0xe4, 0xd4, 0xa1, 0xd1, 0x3b, 0xd1, 0x2e, 0xbc, 0xb9, 0xcc, 0x35, 0x16, 0xa3, 0x7a,
	0x78, 0x50, 0x29, 0xab, 0x0e, 0x8d, 0x51, 0x49, 0xe5, 0x57, 0x0a, 0xdc, 0x60, 0xad, 0xc1, 0x40,
	0x3d, 0xb0, 0x6c, 0x15, 0x5b, 0x36, 0xe2, 0x41, 0xd2, 0xe8, 0x3b, 0xee, 0x44, 0xe6, 0x1a, 0xb5,
	0x68, 0xb9, 0x68, 0xeb, 0xaa, 0x69, 0xfc, 0xe0, 0x7a, 0x4b, 0xb4, 0x77, 0x1d, 0xcd, 0xe6, 0x7d,
	0x97, 0xb8, 0x4c, 0x7f, 0x72, 0x15, 0x1e, 0xf5, 0x08, 0x0c, 0x27, 0x71, 0xfd, 0xe1, 0x6c, 0xa7,
	0x72, 0x49, 0x98, 0xda, 0x4e, 0xe5, 0x52, 0x30, 0xbd, 0x9d, 0xca, 0xa5, 0x61, 0x66, 0x3b, 0x95,
	0xcb, 0xc0, 0x6c, 0xe5, 0x17, 0x0a, 0xdc, 0x6d, 0x6b, 0x38, 0x48, 0x5a, 0xd2, 0x9c, 0xa1, 0x65,
	0x3a, 0x1a, 0xfa, 0xf6, 0xff, 0x93, 0xf7, 0x28, 0xd7, 0xfe, 0x13, 0xe5, 0x99, 0x1c, 0x65, 0x50,
	0x08, 0xf2, 0x73, 0x50, 0x13, 0x14, 0x0e, 0x83, 0x01, 0x32, 0x9e, 0x95, 0x68, 0xfb, 0xd0, 0xa9,
	0xc2, 0x25, 0x1b, 0x7f, 0xdd, 0x04, 0x69, 0xf7, 0xf3, 0xe8, 0x16, 0x28, 0xb8, 0x04, 0x7a, 0x86,
	0xe9, 0xee, 0x16, 0xf8, 0x09, 0xba, 0x0d, 0x6e, 0x4a, 0x7c, 0xfb, 0x95, 0xd2, 0xdb, 0x95, 0x39,
	0xa9, 0xc7, 0x77, 0xb6, 0x44, 0x48, 0xa1, 0x7b, 0x60, 0x29, 0x10, 0x94, 0x39, 0x45, 0xe1, 0x3b,
	0x6d, 0xb9, 0xd7, 0x64, 0x64, 0x9e, 0x85, 0x09, 0x54, 0x06, 0x2b, 0x71, 0x69, 0xa6, 0xcb, 0xf7,
	0x76, 0xb8, 0xb7, 0x32, 0x4c, 0xa2, 0x05, 0x70, 0x2b, 0x80, 0x68, 0x71, 0x02, 0xa7, 0x70, 0x30,
	0x85, 0xee, 0x83, 0x7b, 0x81, 0x30, 0xb3, 0xab, 0xbc, 0x12, 0x25, 0x7e, 0x9f, 0x6b, 0xf5, 0x58,
	0x81, 0xe7, 0x3a, 0x8a, 0x0c, 0xd3, 0x91, 0xde, 0x4c, 0xb7, 0x2b, 0xf0, 0x2c, 0xa3, 0xf0, 0x62,
	0x47, 0xee, 0x09, 0xbc, 0xac, 0xc0, 0x0c, 0xaa, 0x80, 0xd5, 0x69, 0x08, 0x56, 0xe2, 0x18, 0x85,
	0x83, 0x59, 0xb4, 0x02, 0x8a, 0x01, 0x4c, 0x9b, 0x51, 0xb8, 0x3d, 0xe6, 0x2d, 0xe9, 0x90, 0x43,
	0xab, 0xa0, 0x14, 0x97, 0x25, 0xd5, 0x79, 0xb4, 0x0c, 0xee, 0x06, 0xf2, 0x84, 0x9b, 0x57, 0x0c,
	0x22, 0xda, 0x8c, 0x93, 0xa4, 0x76, 0x2e, 0x72, 0x44, 0x51, 0x6a, 0x33, 0x1d, 0x7e, 0x3f, 0x78,
	0x80, 0x1b, 0x68, 0x1d, 0xac, 0x4d, 0x85, 0x90, 0x3e, 0x85, 0x48, 0x9f, 0x8e, 0xa8, 0xf0, 0x5b,
	0x57, 0xc7, 0x94, 0x38, 0xa6, 0x05, 0xbf, 0x41, 0x08, 0xcc, 0x07, 0x85, 0x10, 0x04, 0x38, 0x8f,
	0x4a, 0x60, 0xd1, 0x8b, 0x05, 0x74, 0xf1, 0xa6, 0x7a, 0x13, 0x3d, 0x00, 0xe5, 0xc9, 0x5c, 0x64,
	0xb8, 0x10, 0x3d, 0x06, 0xeb, 0x1f, 0x41, 0x5d, 0xcd, 0xf8, 0x16, 0x7a, 0x02, 0xaa, 0x1f, 0x01,
	0xb2, 0xa2, 0x20, 0x30, 0x4d, 0x51, 0x62, 0x14, 0x51, 0x92, 0x21, 0x9a, 0xd1, 0xb6, 0xcb, 0xb0,
	0x3b, 0x4c, 0x9b, 0x93, 0xe1, 0xd7, 0xfe, 0xe8, 0x82, 0x40, 0xe2, 0xa0, 0xdb, 0xfe, 0xf0, 0xc3,
	0xd9, 0x37, 0x3c, 0xcb, 0x11, 0x5d, 0xee, 0xf8, 0xfa, 0xc6, 0x61, 0xf6, 0x24, 0x5e, 0xe1, 0xe0,
	0x42, 0x3c, 0x9f, 0x60, 0x23, 0xef, 0x98, 0x8b, 0xa8, 0x0a, 0x1e, 0xcc, 0xe8, 0xe6, 0x21, 0xef,
	0xc6, 0x73, 0x53, 0x24, 0x66, 0x6b, 0x8b, 0x67, 0x3d, 0x6e, 0x45, 0xf4, 0x08, 0x54, 0xa6, 0x63,
	0x76, 0xbb, 0x84, 0xde, 0x52, 0xfc, 0x57, 0xc7, 0xb8, 0x96, 0xb8, 0xd7, 0x21, 0xc8, 0x52, 0xfc,
	0xc4, 0x05, 0xbe, 0xb3, 0x03, 0x97, 0xd1, 0x12, 0x58, 0x98, 0xcc, 0x8d, 0x8c, 0xb2, 0x82, 0xee,
	0x00, 0xe8, 0xa5, 0x3c, 0x07, 0xbb, 0xd1, 0x7b, 0x68, 0x11, 0x20, 0x2f, 0x4a, 0x2e, 0x85, 0x67,
	0x9d, 0x55, 0xff, 0x56, 0x8e, 0xe3, 0x11, 0xdb, 0xac, 0xf9, 0xa2, 0x4f, 0x20, 0xae, 0x2c, 0x53,
	0xf6, 0x4f, 0x35, 0x01, 0x0a, 0xdb, 0xe5, 0x3e, 0x2a, 0x82, 0x3b, 0x61, 0x24, 0x71, 0x40, 0xc5,
	0xbf, 0xbc, 0xe3, 0x4c, 0x48, 0xe1, 0x75, 0xdf, 0xe5, 0xd1, 0x7c, 0x40, 0xb5, 0x07, 0x93, 0x07,
	0x75, 0x15, 0x7b, 0xe8, 0xdf, 0xee, 0x2b, 0x86, 0x0a, 0xa3, 0xec, 0x12, 0x6b, 0x3d, 0x42, 0x6b,
	0x60, 0x39, 0x52, 0x26, 0x12, 0x55, 0x5d, 0xc0, 0xe3, 0x49, 0x80, 0xe7, 0x10, 0x99, 0x63, 0x25,
	0x4e, 0x91, 0xe1, 0xf3, 0x49, 0xfa, 0xae, 0xd7, 0xc6, 0xf9, 0x17, 0xfe, 0xe6, 0x1c, 0xe7, 0x47,
	0x83, 0xa9, 0xfa, 0x2b, 0x29, 0xb8, 0x2e, 0xbc, 0xe9, 0x7c, 0x8a, 0x1e, 0x82, 0xfb, 0x31, 0xc9,
	0xc8, 0x88, 0x36, 0x7c, 0xf5, 0xe3, 0x61, 0x57, 0x73, 0xfa, 0xcc, 0xbf, 0x1c, 0xf1, 0xc8, 0xd7,
	0xdc, 0xeb, 0x26, 0x27, 0xc9, 0xf0, 0x89, 0x2f, 0x57, 0x08, 0x48, 0x66, 0x55, 0x9b, 0xf2, 0xc5,
	0xc9, 0xa5, 0x4e, 0xa3, 0x0d, 0xf0, 0x68, 0x16, 0x92, 0xac, 0xc6, 0xba, 0x3f, 0xe1, 0x10, 0x36,
	0xbc, 0xe4, 0x3f, 0xf7, 0x6f, 0x5a, 0x3c, 0x8a, 0x74, 0x7b, 0xea, 0x1b, 0x37, 0x84, 0x0b, 0x2d,
	0xfd, 0xc6, 0x14, 0x85, 0x23, 0xcb, 0x7f, 0x73, 0xda, 0x29, 0x5a, 0xad, 0x1e, 0x13, 0xb6, 0x38,
	0xfc, 0xc2, 0xbf, 0xb7, 0x61, 0xac, 0x20, 0xc0, 0x67, 0xbe, 0x3b, 0x65, 0xae, 0xd3, 0xea, 0xf1,
	0x9d, 0x37, 0xbc, 0xc2, 0xc9, 0xf0, 0x4b, 0x54, 0x00, 0x79, 0x72, 0x9f, 0x05, 0x01, 0x7e, 0x55,
	0x2a, 0xfc, 0xfd, 0x61, 0x29, 0x5f, 0xa4, 0x36, 0xd2, 0x6e, 0xb0, 0xf9, 0xec, 0xf7, 0x3f, 0x57,
	0xa9, 0xfd, 0xba, 0x6e, 0xd1, 0xf8, 0x58, 0xc3, 0xee, 0xef, 0x3e, 0x6d, 0x6a, 0xf8, 0x9d, 0x65,
	0x9f, 0xd4, 0xc3, 0x7f, 0xe9, 0xe7, 0x9b, 0xf5, 0xe1, 0x89, 0x5e, 0xc7, 0xd8, 0x1c, 0x1e, 0x1c,
	0x64, 0xdc, 0x5f, 0xc7, 0xcd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc8, 0x95, 0x93, 0x18,
	0x0d, 0x00, 0x00,
}
