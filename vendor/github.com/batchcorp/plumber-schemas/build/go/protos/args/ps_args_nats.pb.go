// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_nats.proto

package args

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

type NatsConn struct {
	// @gotags: kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_DSN'"
	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_DSN'"`
	// @gotags: kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_CREDENTIALS'"
	UserCredentials []byte `protobuf:"bytes,2,opt,name=user_credentials,json=userCredentials,proto3" json:"user_credentials,omitempty" kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_CREDENTIALS'"`
	// @gotags: kong:"embed"
	TlsOptions           *NatsTLSOptions `protobuf:"bytes,3,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty" kong:"embed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NatsConn) Reset()         { *m = NatsConn{} }
func (m *NatsConn) String() string { return proto.CompactTextString(m) }
func (*NatsConn) ProtoMessage()    {}
func (*NatsConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5ddeb96e248a57, []int{0}
}

func (m *NatsConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsConn.Unmarshal(m, b)
}
func (m *NatsConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsConn.Marshal(b, m, deterministic)
}
func (m *NatsConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsConn.Merge(m, src)
}
func (m *NatsConn) XXX_Size() int {
	return xxx_messageInfo_NatsConn.Size(m)
}
func (m *NatsConn) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsConn.DiscardUnknown(m)
}

var xxx_messageInfo_NatsConn proto.InternalMessageInfo

func (m *NatsConn) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

func (m *NatsConn) GetUserCredentials() []byte {
	if m != nil {
		return m.UserCredentials
	}
	return nil
}

func (m *NatsConn) GetTlsOptions() *NatsTLSOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

type NatsTLSOptions struct {
	// @gotags: kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_VERIFY_CERT'"
	TlsSkipVerify bool `protobuf:"varint,1,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_VERIFY_CERT'"`
	// @gotags: kong:"help='CA file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CA_CERT'"
	TlsCaCert []byte `protobuf:"bytes,2,opt,name=tls_ca_cert,json=tlsCaCert,proto3" json:"tls_ca_cert,omitempty" kong:"help='CA file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CA_CERT'"`
	// @gotags: kong:"help='Client cert file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_CERT'"
	TlsClientCert []byte `protobuf:"bytes,3,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='Client cert file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_CERT'"`
	// @gotags: kong:"help='client key file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_KEY'"
	TlsClientKey         []byte   `protobuf:"bytes,4,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty" kong:"help='client key file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_KEY'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsTLSOptions) Reset()         { *m = NatsTLSOptions{} }
func (m *NatsTLSOptions) String() string { return proto.CompactTextString(m) }
func (*NatsTLSOptions) ProtoMessage()    {}
func (*NatsTLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5ddeb96e248a57, []int{1}
}

func (m *NatsTLSOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsTLSOptions.Unmarshal(m, b)
}
func (m *NatsTLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsTLSOptions.Marshal(b, m, deterministic)
}
func (m *NatsTLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsTLSOptions.Merge(m, src)
}
func (m *NatsTLSOptions) XXX_Size() int {
	return xxx_messageInfo_NatsTLSOptions.Size(m)
}
func (m *NatsTLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsTLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_NatsTLSOptions proto.InternalMessageInfo

func (m *NatsTLSOptions) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

func (m *NatsTLSOptions) GetTlsCaCert() []byte {
	if m != nil {
		return m.TlsCaCert
	}
	return nil
}

func (m *NatsTLSOptions) GetTlsClientCert() []byte {
	if m != nil {
		return m.TlsClientCert
	}
	return nil
}

func (m *NatsTLSOptions) GetTlsClientKey() []byte {
	if m != nil {
		return m.TlsClientKey
	}
	return nil
}

type NatsReadArgs struct {
	// @gotags: kong:"help='NATS Subject. Ex: foo.bar.*',env='PLUMBER_RELAY_NATS_SUBJECT'"
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" kong:"help='NATS Subject. Ex: foo.bar.*',env='PLUMBER_RELAY_NATS_SUBJECT'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsReadArgs) Reset()         { *m = NatsReadArgs{} }
func (m *NatsReadArgs) String() string { return proto.CompactTextString(m) }
func (*NatsReadArgs) ProtoMessage()    {}
func (*NatsReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5ddeb96e248a57, []int{2}
}

func (m *NatsReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsReadArgs.Unmarshal(m, b)
}
func (m *NatsReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsReadArgs.Marshal(b, m, deterministic)
}
func (m *NatsReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsReadArgs.Merge(m, src)
}
func (m *NatsReadArgs) XXX_Size() int {
	return xxx_messageInfo_NatsReadArgs.Size(m)
}
func (m *NatsReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_NatsReadArgs proto.InternalMessageInfo

func (m *NatsReadArgs) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type NatsWriteArgs struct {
	// @gotags: kong:"help='NATS Subject. Ex: foo.bar.*'"
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" kong:"help='NATS Subject. Ex: foo.bar.*'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsWriteArgs) Reset()         { *m = NatsWriteArgs{} }
func (m *NatsWriteArgs) String() string { return proto.CompactTextString(m) }
func (*NatsWriteArgs) ProtoMessage()    {}
func (*NatsWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5ddeb96e248a57, []int{3}
}

func (m *NatsWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsWriteArgs.Unmarshal(m, b)
}
func (m *NatsWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsWriteArgs.Marshal(b, m, deterministic)
}
func (m *NatsWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsWriteArgs.Merge(m, src)
}
func (m *NatsWriteArgs) XXX_Size() int {
	return xxx_messageInfo_NatsWriteArgs.Size(m)
}
func (m *NatsWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_NatsWriteArgs proto.InternalMessageInfo

func (m *NatsWriteArgs) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func init() {
	proto.RegisterType((*NatsConn)(nil), "protos.args.NatsConn")
	proto.RegisterType((*NatsTLSOptions)(nil), "protos.args.NatsTLSOptions")
	proto.RegisterType((*NatsReadArgs)(nil), "protos.args.NatsReadArgs")
	proto.RegisterType((*NatsWriteArgs)(nil), "protos.args.NatsWriteArgs")
}

func init() { proto.RegisterFile("ps_args_nats.proto", fileDescriptor_ee5ddeb96e248a57) }

var fileDescriptor_ee5ddeb96e248a57 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0x13, 0xdd, 0xb2, 0xbf, 0xe4, 0xa9, 0x20, 0xc8, 0x28, 0x22, 0xdd, 0x83, 0x2d,
	0xe8, 0x93, 0xe8, 0x8b, 0xf6, 0x51, 0x51, 0xe8, 0x44, 0xc1, 0x97, 0x92, 0xa6, 0xb1, 0x8b, 0xcb,
	0x9a, 0x92, 0x7b, 0x2b, 0xec, 0x03, 0xf8, 0x55, 0xfc, 0x9c, 0x92, 0x74, 0xd3, 0xf9, 0xe2, 0x53,
	0xee, 0x3d, 0xf9, 0x71, 0xcf, 0xe1, 0x10, 0x5a, 0x43, 0xc6, 0x4c, 0x09, 0x59, 0xc5, 0x10, 0xa2,
	0xda, 0x68, 0xd4, 0xb4, 0xef, 0x1e, 0x88, 0xac, 0x1e, 0x7c, 0x7a, 0xa4, 0xfb, 0xc0, 0x10, 0x12,
	0x5d, 0x55, 0x74, 0x42, 0x3a, 0x05, 0x54, 0xbe, 0x37, 0xf5, 0xc2, 0x5e, 0x6a, 0x47, 0x3a, 0x23,
	0x93, 0x06, 0x84, 0xc9, 0xb8, 0x11, 0x85, 0xa8, 0x50, 0x32, 0x05, 0xfe, 0xde, 0xd4, 0x0b, 0x07,
	0xe9, 0xd8, 0xea, 0xc9, 0xaf, 0x4c, 0xaf, 0x49, 0x1f, 0x15, 0x64, 0xba, 0x46, 0xa9, 0x2b, 0xf0,
	0x3b, 0x53, 0x2f, 0xec, 0x9f, 0x1f, 0x45, 0x3b, 0x66, 0x91, 0x35, 0x7a, 0xba, 0x9f, 0x3f, 0xb6,
	0x48, 0x4a, 0x50, 0xc1, 0x66, 0x0e, 0xbe, 0x3c, 0x32, 0xfa, 0xfb, 0x4d, 0x4f, 0xc9, 0xd8, 0x1e,
	0x84, 0xa5, 0xac, 0xb3, 0x0f, 0x61, 0xe4, 0xdb, 0xda, 0x25, 0xeb, 0xa6, 0x43, 0x54, 0x30, 0x5f,
	0xca, 0xfa, 0xd9, 0x89, 0xf4, 0xb8, 0x35, 0xe6, 0x2c, 0xe3, 0xc2, 0xe0, 0x26, 0x5e, 0x0f, 0x15,
	0x24, 0x2c, 0x11, 0x06, 0xb7, 0x77, 0xb8, 0x92, 0xa2, 0xc2, 0x96, 0xe9, 0x38, 0xc6, 0xde, 0x49,
	0x9c, 0xea, 0xb8, 0x13, 0x32, 0xda, 0xe1, 0x96, 0x62, 0xed, 0xef, 0x3b, 0x6c, 0xf0, 0x83, 0xdd,
	0x89, 0x75, 0x10, 0x92, 0x81, 0xcd, 0x99, 0x0a, 0x56, 0xdc, 0x98, 0x12, 0xa8, 0x4f, 0x0e, 0xa1,
	0xc9, 0xdf, 0x05, 0xc7, 0x4d, 0x6f, 0xdb, 0x35, 0x98, 0x91, 0xa1, 0x25, 0x5f, 0x8c, 0x44, 0xf1,
	0x3f, 0x7a, 0x7b, 0xf5, 0x7a, 0x59, 0x4a, 0x5c, 0x34, 0x79, 0xc4, 0xf5, 0x2a, 0xce, 0x19, 0xf2,
	0x05, 0xd7, 0xa6, 0x8e, 0x6b, 0xd5, 0xac, 0x72, 0x61, 0xce, 0x80, 0x2f, 0xc4, 0x8a, 0x41, 0x9c,
	0x37, 0x52, 0x15, 0x71, 0xa9, 0xe3, 0xb6, 0xd5, 0xd8, 0xb6, 0x9a, 0x1f, 0xb8, 0xe5, 0xe2, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0xfa, 0x2c, 0x03, 0xec, 0x01, 0x00, 0x00,
}
