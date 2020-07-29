// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/types/status.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PermissionDeniedDetails_Reason int32

const (
	// Reason unknown.
	PermissionDeniedDetails_UNKNOWN PermissionDeniedDetails_Reason = 0
	// Agent identity has expired.
	PermissionDeniedDetails_AGENT_EXPIRED PermissionDeniedDetails_Reason = 1
	// Identity is not an attested agent.
	PermissionDeniedDetails_AGENT_NOT_ATTESTED PermissionDeniedDetails_Reason = 2
	// Identity is not the active agent identity.
	PermissionDeniedDetails_AGENT_NOT_ACTIVE PermissionDeniedDetails_Reason = 3
	// Agent has been banned.
	PermissionDeniedDetails_AGENT_BANNED PermissionDeniedDetails_Reason = 4
)

var PermissionDeniedDetails_Reason_name = map[int32]string{
	0: "UNKNOWN",
	1: "AGENT_EXPIRED",
	2: "AGENT_NOT_ATTESTED",
	3: "AGENT_NOT_ACTIVE",
	4: "AGENT_BANNED",
}

var PermissionDeniedDetails_Reason_value = map[string]int32{
	"UNKNOWN":            0,
	"AGENT_EXPIRED":      1,
	"AGENT_NOT_ATTESTED": 2,
	"AGENT_NOT_ACTIVE":   3,
	"AGENT_BANNED":       4,
}

func (x PermissionDeniedDetails_Reason) String() string {
	return proto.EnumName(PermissionDeniedDetails_Reason_name, int32(x))
}

func (PermissionDeniedDetails_Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6d85e11b7a283fe, []int{1, 0}
}

type Status struct {
	// A status code, which should be an enum value of google.rpc.Code.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6d85e11b7a283fe, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PermissionDeniedDetails struct {
	Reason               PermissionDeniedDetails_Reason `protobuf:"varint,1,opt,name=reason,proto3,enum=spire.types.PermissionDeniedDetails_Reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PermissionDeniedDetails) Reset()         { *m = PermissionDeniedDetails{} }
func (m *PermissionDeniedDetails) String() string { return proto.CompactTextString(m) }
func (*PermissionDeniedDetails) ProtoMessage()    {}
func (*PermissionDeniedDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6d85e11b7a283fe, []int{1}
}

func (m *PermissionDeniedDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionDeniedDetails.Unmarshal(m, b)
}
func (m *PermissionDeniedDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionDeniedDetails.Marshal(b, m, deterministic)
}
func (m *PermissionDeniedDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionDeniedDetails.Merge(m, src)
}
func (m *PermissionDeniedDetails) XXX_Size() int {
	return xxx_messageInfo_PermissionDeniedDetails.Size(m)
}
func (m *PermissionDeniedDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionDeniedDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionDeniedDetails proto.InternalMessageInfo

func (m *PermissionDeniedDetails) GetReason() PermissionDeniedDetails_Reason {
	if m != nil {
		return m.Reason
	}
	return PermissionDeniedDetails_UNKNOWN
}

func init() {
	proto.RegisterEnum("spire.types.PermissionDeniedDetails_Reason", PermissionDeniedDetails_Reason_name, PermissionDeniedDetails_Reason_value)
	proto.RegisterType((*Status)(nil), "spire.types.Status")
	proto.RegisterType((*PermissionDeniedDetails)(nil), "spire.types.PermissionDeniedDetails")
}

func init() {
	proto.RegisterFile("spire/types/status.proto", fileDescriptor_f6d85e11b7a283fe)
}

var fileDescriptor_f6d85e11b7a283fe = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0xf4, 0xad, 0x29, 0x4e, 0x55, 0xd6, 0x45, 0x34, 0xc7, 0xd2, 0x53, 0xa5, 0xb0,
	0x01, 0x05, 0xef, 0x69, 0xb3, 0x48, 0x11, 0xb6, 0x25, 0x5d, 0xff, 0xe0, 0xa5, 0xa4, 0xed, 0xb4,
	0x5d, 0x30, 0xd9, 0x90, 0xd9, 0x1e, 0xfc, 0x86, 0x7e, 0x2c, 0x61, 0xa3, 0x90, 0x8b, 0xb7, 0x99,
	0xdf, 0xc3, 0x33, 0x03, 0x3f, 0x88, 0xa8, 0x32, 0x35, 0xc6, 0xee, 0xb3, 0x42, 0x8a, 0xc9, 0xe5,
	0xee, 0x48, 0xa2, 0xaa, 0xad, 0xb3, 0xbc, 0xef, 0x13, 0xe1, 0x93, 0xe1, 0x03, 0x84, 0x4b, 0x1f,
	0x72, 0x0e, 0xdd, 0x8d, 0xdd, 0x62, 0x14, 0x0c, 0x82, 0xd1, 0x49, 0xe6, 0x67, 0x1e, 0x41, 0xaf,
	0x40, 0xa2, 0x7c, 0x8f, 0x51, 0x67, 0x10, 0x8c, 0x4e, 0xb3, 0xdf, 0x75, 0xf8, 0x15, 0xc0, 0xcd,
	0x02, 0xeb, 0xc2, 0x10, 0x19, 0x5b, 0xa6, 0x58, 0x1a, 0xdc, 0xa6, 0xe8, 0x72, 0xf3, 0x41, 0x7c,
	0x0a, 0x61, 0x8d, 0x39, 0xd9, 0xd2, 0xdf, 0xba, 0xb8, 0x1b, 0x8b, 0xd6, 0x47, 0xf1, 0x47, 0x4b,
	0x64, 0xbe, 0x92, 0xfd, 0x54, 0x87, 0x07, 0x08, 0x1b, 0xc2, 0xfb, 0xd0, 0x7b, 0x56, 0x4f, 0x6a,
	0xfe, 0xaa, 0xd8, 0x3f, 0x7e, 0x09, 0xe7, 0xc9, 0xa3, 0x54, 0x7a, 0x25, 0xdf, 0x16, 0xb3, 0x4c,
	0xa6, 0x2c, 0xe0, 0xd7, 0xc0, 0x1b, 0xa4, 0xe6, 0x7a, 0x95, 0x68, 0x2d, 0x97, 0x5a, 0xa6, 0xac,
	0xc3, 0xaf, 0x80, 0xb5, 0xf8, 0x54, 0xcf, 0x5e, 0x24, 0xfb, 0xcf, 0x19, 0x9c, 0x35, 0x74, 0x92,
	0x28, 0x25, 0x53, 0xd6, 0x9d, 0x8c, 0xdf, 0x6f, 0xf7, 0xc6, 0x1d, 0x8e, 0x6b, 0xb1, 0xb1, 0x45,
	0x4c, 0x95, 0xd9, 0xed, 0x30, 0x6e, 0xec, 0x79, 0x61, 0x71, 0xcb, 0xe4, 0x3a, 0xf4, 0xe8, 0xfe,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x79, 0x2a, 0xbd, 0x89, 0x5f, 0x01, 0x00, 0x00,
}