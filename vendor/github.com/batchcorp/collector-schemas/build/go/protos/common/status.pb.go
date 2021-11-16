// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package common

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

type Status struct {
	// A simple error code that can be easily handled by the client. The
	// actual error code is defined by `google.rpc.Code`.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing human-readable error message in English. It should
	// both explain the error and offer an actionable resolution to it.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
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

func init() {
	proto.RegisterType((*Status)(nil), "events.Status")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0xe9, 0xcf, 0x6f, 0xc5, 0x20, 0x0a, 0x99, 0x3a, 0x16, 0xa7, 0x2e, 0x26, 0x83, 0xe0,
	0xe6, 0xd2, 0x47, 0xa8, 0x9b, 0x93, 0xc9, 0xed, 0xa5, 0x29, 0x34, 0xbd, 0xa5, 0xf7, 0xc6, 0xe7,
	0x17, 0x52, 0xdc, 0xce, 0x77, 0x38, 0x7c, 0x1c, 0x75, 0x64, 0x71, 0x92, 0xd8, 0x2c, 0x2b, 0x09,
	0xe9, 0x12, 0x3f, 0x38, 0x0b, 0x5f, 0xee, 0xaa, 0x7c, 0xe6, 0x5e, 0x6b, 0xf5, 0x0f, 0xd4, 0x63,
	0x55, 0xd4, 0x45, 0xb3, 0xeb, 0x72, 0xd6, 0x95, 0xda, 0x47, 0x64, 0x76, 0x03, 0x56, 0x7f, 0x75,
	0xd1, 0x1c, 0xba, 0x1f, 0xb6, 0x6f, 0x75, 0xe6, 0x60, 0xbc, 0x13, 0x08, 0x66, 0x53, 0xb5, 0xa7,
	0x0e, 0x79, 0xa1, 0x99, 0x71, 0x13, 0xbe, 0x1e, 0xc3, 0x28, 0x21, 0x79, 0x03, 0x14, 0x6d, 0x1e,
	0x02, 0xad, 0x8b, 0x05, 0x9a, 0x26, 0x04, 0xa1, 0xf5, 0xca, 0x10, 0x30, 0x3a, 0xb6, 0x3e, 0x8d,
	0x53, 0x6f, 0x07, 0xb2, 0xf9, 0x19, 0x5b, 0xa0, 0x18, 0x69, 0xf6, 0x65, 0xc6, 0xdb, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x51, 0xf8, 0x88, 0x8e, 0xb8, 0x00, 0x00, 0x00,
}
