// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_encoding_options.proto

package encoding

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

type EncodeType int32

const (
	EncodeType_ENCODE_TYPE_UNSET  EncodeType = 0
	EncodeType_ENCODE_TYPE_JSONPB EncodeType = 1
	EncodeType_ENCODE_TYPE_AVRO   EncodeType = 2
)

var EncodeType_name = map[int32]string{
	0: "ENCODE_TYPE_UNSET",
	1: "ENCODE_TYPE_JSONPB",
	2: "ENCODE_TYPE_AVRO",
}

var EncodeType_value = map[string]int32{
	"ENCODE_TYPE_UNSET":  0,
	"ENCODE_TYPE_JSONPB": 1,
	"ENCODE_TYPE_AVRO":   2,
}

func (x EncodeType) String() string {
	return proto.EnumName(EncodeType_name, int32(x))
}

func (EncodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{0}
}

type DecodeType int32

const (
	DecodeType_DECODE_TYPE_UNSET      DecodeType = 0
	DecodeType_DECODE_TYPE_PROTOBUF   DecodeType = 1
	DecodeType_DECODE_TYPE_AVRO       DecodeType = 2
	DecodeType_DECODE_TYPE_THRIFT     DecodeType = 3
	DecodeType_DECODE_TYPE_FLATBUFFER DecodeType = 4
)

var DecodeType_name = map[int32]string{
	0: "DECODE_TYPE_UNSET",
	1: "DECODE_TYPE_PROTOBUF",
	2: "DECODE_TYPE_AVRO",
	3: "DECODE_TYPE_THRIFT",
	4: "DECODE_TYPE_FLATBUFFER",
}

var DecodeType_value = map[string]int32{
	"DECODE_TYPE_UNSET":      0,
	"DECODE_TYPE_PROTOBUF":   1,
	"DECODE_TYPE_AVRO":       2,
	"DECODE_TYPE_THRIFT":     3,
	"DECODE_TYPE_FLATBUFFER": 4,
}

func (x DecodeType) String() string {
	return proto.EnumName(DecodeType_name, int32(x))
}

func (DecodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{1}
}

type EnvelopeType int32

const (
	EnvelopeType_ENVELOPE_TYPE_UNSET   EnvelopeType = 0
	EnvelopeType_ENVELOPE_TYPE_DEEP    EnvelopeType = 1
	EnvelopeType_ENVELOPE_TYPE_SHALLOW EnvelopeType = 2
)

var EnvelopeType_name = map[int32]string{
	0: "ENVELOPE_TYPE_UNSET",
	1: "ENVELOPE_TYPE_DEEP",
	2: "ENVELOPE_TYPE_SHALLOW",
}

var EnvelopeType_value = map[string]int32{
	"ENVELOPE_TYPE_UNSET":   0,
	"ENVELOPE_TYPE_DEEP":    1,
	"ENVELOPE_TYPE_SHALLOW": 2,
}

func (x EnvelopeType) String() string {
	return proto.EnumName(EnvelopeType_name, int32(x))
}

func (EnvelopeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{2}
}

type ProtobufSettings struct {
	// @gotags: kong:"help='Input message(s) should be encoded with this message envelope'"
	ProtobufRootMessage string `protobuf:"bytes,1,opt,name=protobuf_root_message,json=protobufRootMessage,proto3" json:"protobuf_root_message,omitempty" kong:"help='Input message(s) should be encoded with this message envelope'"`
	// Desktop/server should not use this.
	// @gotags: kong:"help='One or more directories which contains protobuf schemas',existingdir"
	ProtobufDirs []string `protobuf:"bytes,2,rep,name=protobuf_dirs,json=protobufDirs,proto3" json:"protobuf_dirs,omitempty" kong:"help='One or more directories which contains protobuf schemas',existingdir"`
	// Directory where protos are stored (used for github import)
	// @gotags: kong:"-"
	XProtobufRootDir string `protobuf:"bytes,3,opt,name=_protobuf_root_dir,json=ProtobufRootDir,proto3" json:"_protobuf_root_dir,omitempty" kong:"-"`
	// Used by server/desktop when creating a read without an existing schema
	// @gotags: kong:"-"
	Archive []byte `protobuf:"bytes,4,opt,name=archive,proto3" json:"archive,omitempty" kong:"-"`
	// Used internally by the server
	// @gotags: kong:"-"
	XMessageDescriptor []byte `protobuf:"bytes,5,opt,name=_message_descriptor,json=MessageDescriptor,proto3" json:"_message_descriptor,omitempty" kong:"-"`
	// @gotags: kong:"help='Envelope type (options: deep, shallow)',type=pbenum,pbenum_strip_prefix=ENVELOPE_TYPE_,pbenum_lowercase,default=deep"
	ProtobufEnvelopeType EnvelopeType `protobuf:"varint,6,opt,name=protobuf_envelope_type,json=protobufEnvelopeType,proto3,enum=protos.encoding.EnvelopeType" json:"protobuf_envelope_type,omitempty" kong:"help='Envelope type (options: deep, shallow)',type=pbenum,pbenum_strip_prefix=ENVELOPE_TYPE_,pbenum_lowercase,default=deep"`
	// @gotags: kong:"help='For shallow envelope messages, the payload field should be encoded with this message name'"
	ShallowEnvelopeMessage string `protobuf:"bytes,7,opt,name=shallow_envelope_message,json=shallowEnvelopeMessage,proto3" json:"shallow_envelope_message,omitempty" kong:"help='For shallow envelope messages, the payload field should be encoded with this message name'"`
	// @gotags: kong:"help='For shallow envelope messages, the field number of the root message that contains the shallow envelope payload'"
	ShallowEnvelopeFieldNumber int32 `protobuf:"varint,8,opt,name=shallow_envelope_field_number,json=shallowEnvelopeFieldNumber,proto3" json:"shallow_envelope_field_number,omitempty" kong:"help='For shallow envelope messages, the field number of the root message that contains the shallow envelope payload'"`
	// Used internally by the server
	// @gotags: kong:"-"
	XShallowEnvelopeMessageDescriptor []byte `protobuf:"bytes,9,opt,name=_shallow_envelope_message_descriptor,json=ShallowEnvelopeMessageDescriptor,proto3" json:"_shallow_envelope_message_descriptor,omitempty" kong:"-"`
	// @gotags: kong:"help='Protobuf descriptor set(.protoset or .fds file)'"
	ProtobufDescriptorSet string   `protobuf:"bytes,10,opt,name=protobuf_descriptor_set,json=protobufDescriptorSet,proto3" json:"protobuf_descriptor_set,omitempty" kong:"help='Protobuf descriptor set(.protoset or .fds file)'"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ProtobufSettings) Reset()         { *m = ProtobufSettings{} }
func (m *ProtobufSettings) String() string { return proto.CompactTextString(m) }
func (*ProtobufSettings) ProtoMessage()    {}
func (*ProtobufSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{0}
}

func (m *ProtobufSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtobufSettings.Unmarshal(m, b)
}
func (m *ProtobufSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtobufSettings.Marshal(b, m, deterministic)
}
func (m *ProtobufSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtobufSettings.Merge(m, src)
}
func (m *ProtobufSettings) XXX_Size() int {
	return xxx_messageInfo_ProtobufSettings.Size(m)
}
func (m *ProtobufSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtobufSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ProtobufSettings proto.InternalMessageInfo

func (m *ProtobufSettings) GetProtobufRootMessage() string {
	if m != nil {
		return m.ProtobufRootMessage
	}
	return ""
}

func (m *ProtobufSettings) GetProtobufDirs() []string {
	if m != nil {
		return m.ProtobufDirs
	}
	return nil
}

func (m *ProtobufSettings) GetXProtobufRootDir() string {
	if m != nil {
		return m.XProtobufRootDir
	}
	return ""
}

func (m *ProtobufSettings) GetArchive() []byte {
	if m != nil {
		return m.Archive
	}
	return nil
}

func (m *ProtobufSettings) GetXMessageDescriptor() []byte {
	if m != nil {
		return m.XMessageDescriptor
	}
	return nil
}

func (m *ProtobufSettings) GetProtobufEnvelopeType() EnvelopeType {
	if m != nil {
		return m.ProtobufEnvelopeType
	}
	return EnvelopeType_ENVELOPE_TYPE_UNSET
}

func (m *ProtobufSettings) GetShallowEnvelopeMessage() string {
	if m != nil {
		return m.ShallowEnvelopeMessage
	}
	return ""
}

func (m *ProtobufSettings) GetShallowEnvelopeFieldNumber() int32 {
	if m != nil {
		return m.ShallowEnvelopeFieldNumber
	}
	return 0
}

func (m *ProtobufSettings) GetXShallowEnvelopeMessageDescriptor() []byte {
	if m != nil {
		return m.XShallowEnvelopeMessageDescriptor
	}
	return nil
}

func (m *ProtobufSettings) GetProtobufDescriptorSet() string {
	if m != nil {
		return m.ProtobufDescriptorSet
	}
	return ""
}

type AvroSettings struct {
	// Used by CLI; desktop should not set/use this.
	// @gotags: kong:"help='If encode-type is set to avro, must specify avro schema file',existingfile"
	AvroSchemaFile string `protobuf:"bytes,1,opt,name=avro_schema_file,json=avroSchemaFile,proto3" json:"avro_schema_file,omitempty" kong:"help='If encode-type is set to avro, must specify avro schema file',existingfile"`
	// Used by desktop; CLI should not set/use this.
	// @gotags: kong:"-"
	Schema               []byte   `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty" kong:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvroSettings) Reset()         { *m = AvroSettings{} }
func (m *AvroSettings) String() string { return proto.CompactTextString(m) }
func (*AvroSettings) ProtoMessage()    {}
func (*AvroSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{1}
}

func (m *AvroSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvroSettings.Unmarshal(m, b)
}
func (m *AvroSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvroSettings.Marshal(b, m, deterministic)
}
func (m *AvroSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvroSettings.Merge(m, src)
}
func (m *AvroSettings) XXX_Size() int {
	return xxx_messageInfo_AvroSettings.Size(m)
}
func (m *AvroSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AvroSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AvroSettings proto.InternalMessageInfo

func (m *AvroSettings) GetAvroSchemaFile() string {
	if m != nil {
		return m.AvroSchemaFile
	}
	return ""
}

func (m *AvroSettings) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type ThriftSettings struct {
	// @gotags: kong:"help='One or more directories containing Thrift IDL files'"
	ThriftDirs []string `protobuf:"bytes,1,rep,name=thrift_dirs,json=thriftDirs,proto3" json:"thrift_dirs,omitempty" kong:"help='One or more directories containing Thrift IDL files'"`
	// @gotags: kong:"help='Namespace and struct name to decode the message with. Ex: com.mycompany.Account'"
	ThriftStruct         string   `protobuf:"bytes,2,opt,name=thrift_struct,json=thriftStruct,proto3" json:"thrift_struct,omitempty" kong:"help='Namespace and struct name to decode the message with. Ex: com.mycompany.Account'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThriftSettings) Reset()         { *m = ThriftSettings{} }
func (m *ThriftSettings) String() string { return proto.CompactTextString(m) }
func (*ThriftSettings) ProtoMessage()    {}
func (*ThriftSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{2}
}

func (m *ThriftSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftSettings.Unmarshal(m, b)
}
func (m *ThriftSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftSettings.Marshal(b, m, deterministic)
}
func (m *ThriftSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftSettings.Merge(m, src)
}
func (m *ThriftSettings) XXX_Size() int {
	return xxx_messageInfo_ThriftSettings.Size(m)
}
func (m *ThriftSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftSettings proto.InternalMessageInfo

func (m *ThriftSettings) GetThriftDirs() []string {
	if m != nil {
		return m.ThriftDirs
	}
	return nil
}

func (m *ThriftSettings) GetThriftStruct() string {
	if m != nil {
		return m.ThriftStruct
	}
	return ""
}

type JSONSchemaSettings struct {
	// Used by desktop; CLI should not set/use this.
	// @gotags: kong:"-"
	Schema               []byte   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty" kong:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONSchemaSettings) Reset()         { *m = JSONSchemaSettings{} }
func (m *JSONSchemaSettings) String() string { return proto.CompactTextString(m) }
func (*JSONSchemaSettings) ProtoMessage()    {}
func (*JSONSchemaSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{3}
}

func (m *JSONSchemaSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONSchemaSettings.Unmarshal(m, b)
}
func (m *JSONSchemaSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONSchemaSettings.Marshal(b, m, deterministic)
}
func (m *JSONSchemaSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONSchemaSettings.Merge(m, src)
}
func (m *JSONSchemaSettings) XXX_Size() int {
	return xxx_messageInfo_JSONSchemaSettings.Size(m)
}
func (m *JSONSchemaSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONSchemaSettings.DiscardUnknown(m)
}

var xxx_messageInfo_JSONSchemaSettings proto.InternalMessageInfo

func (m *JSONSchemaSettings) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type EncodeOptions struct {
	// Use an existing schema for encoding (and ignore all other encode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// @gotags: kong:"help='Encode type (options: unset, jsonpb, avro)',default=unset,type=pbenum,pbenum_strip_prefix=ENCODE_TYPE_,pbenum_lowercase"
	EncodeType EncodeType `protobuf:"varint,2,opt,name=encode_type,json=encodeType,proto3,enum=protos.encoding.EncodeType" json:"encode_type,omitempty" kong:"help='Encode type (options: unset, jsonpb, avro)',default=unset,type=pbenum,pbenum_strip_prefix=ENCODE_TYPE_,pbenum_lowercase"`
	// @gotags: kong:"embed,group=protobuf"
	ProtobufSettings *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty" kong:"embed,group=protobuf"`
	// @gotags: kong:"embed,group=avro"
	AvroSettings         *AvroSettings `protobuf:"bytes,4,opt,name=avro_settings,json=avroSettings,proto3" json:"avro_settings,omitempty" kong:"embed,group=avro"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EncodeOptions) Reset()         { *m = EncodeOptions{} }
func (m *EncodeOptions) String() string { return proto.CompactTextString(m) }
func (*EncodeOptions) ProtoMessage()    {}
func (*EncodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{4}
}

func (m *EncodeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodeOptions.Unmarshal(m, b)
}
func (m *EncodeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodeOptions.Marshal(b, m, deterministic)
}
func (m *EncodeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodeOptions.Merge(m, src)
}
func (m *EncodeOptions) XXX_Size() int {
	return xxx_messageInfo_EncodeOptions.Size(m)
}
func (m *EncodeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_EncodeOptions proto.InternalMessageInfo

func (m *EncodeOptions) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *EncodeOptions) GetEncodeType() EncodeType {
	if m != nil {
		return m.EncodeType
	}
	return EncodeType_ENCODE_TYPE_UNSET
}

func (m *EncodeOptions) GetProtobufSettings() *ProtobufSettings {
	if m != nil {
		return m.ProtobufSettings
	}
	return nil
}

func (m *EncodeOptions) GetAvroSettings() *AvroSettings {
	if m != nil {
		return m.AvroSettings
	}
	return nil
}

type DecodeOptions struct {
	// Use an existing schema for decoding (and ignore all other decode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// @gotags: kong:"help='Decode type (options: unset, protobuf, avro, thrift, flatbuffer)',type=pbenum,pbenum_strip_prefix=DECODE_TYPE_,pbenum_lowercase,default=unset"
	DecodeType DecodeType `protobuf:"varint,2,opt,name=decode_type,json=decodeType,proto3,enum=protos.encoding.DecodeType" json:"decode_type,omitempty" kong:"help='Decode type (options: unset, protobuf, avro, thrift, flatbuffer)',type=pbenum,pbenum_strip_prefix=DECODE_TYPE_,pbenum_lowercase,default=unset"`
	// @gotags: kong:"embed,group=protobuf"
	ProtobufSettings *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty" kong:"embed,group=protobuf"`
	// @gotags: kong:"embed,group=avro"
	AvroSettings *AvroSettings `protobuf:"bytes,4,opt,name=avro_settings,json=avroSettings,proto3" json:"avro_settings,omitempty" kong:"embed,group=avro"`
	// @gotags: kong:"embed,group=thrift"
	ThriftSettings       *ThriftSettings `protobuf:"bytes,5,opt,name=thrift_settings,json=thriftSettings,proto3" json:"thrift_settings,omitempty" kong:"embed,group=thrift"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DecodeOptions) Reset()         { *m = DecodeOptions{} }
func (m *DecodeOptions) String() string { return proto.CompactTextString(m) }
func (*DecodeOptions) ProtoMessage()    {}
func (*DecodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{5}
}

func (m *DecodeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeOptions.Unmarshal(m, b)
}
func (m *DecodeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeOptions.Marshal(b, m, deterministic)
}
func (m *DecodeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeOptions.Merge(m, src)
}
func (m *DecodeOptions) XXX_Size() int {
	return xxx_messageInfo_DecodeOptions.Size(m)
}
func (m *DecodeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeOptions proto.InternalMessageInfo

func (m *DecodeOptions) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *DecodeOptions) GetDecodeType() DecodeType {
	if m != nil {
		return m.DecodeType
	}
	return DecodeType_DECODE_TYPE_UNSET
}

func (m *DecodeOptions) GetProtobufSettings() *ProtobufSettings {
	if m != nil {
		return m.ProtobufSettings
	}
	return nil
}

func (m *DecodeOptions) GetAvroSettings() *AvroSettings {
	if m != nil {
		return m.AvroSettings
	}
	return nil
}

func (m *DecodeOptions) GetThriftSettings() *ThriftSettings {
	if m != nil {
		return m.ThriftSettings
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.encoding.EncodeType", EncodeType_name, EncodeType_value)
	proto.RegisterEnum("protos.encoding.DecodeType", DecodeType_name, DecodeType_value)
	proto.RegisterEnum("protos.encoding.EnvelopeType", EnvelopeType_name, EnvelopeType_value)
	proto.RegisterType((*ProtobufSettings)(nil), "protos.encoding.ProtobufSettings")
	proto.RegisterType((*AvroSettings)(nil), "protos.encoding.AvroSettings")
	proto.RegisterType((*ThriftSettings)(nil), "protos.encoding.ThriftSettings")
	proto.RegisterType((*JSONSchemaSettings)(nil), "protos.encoding.JSONSchemaSettings")
	proto.RegisterType((*EncodeOptions)(nil), "protos.encoding.EncodeOptions")
	proto.RegisterType((*DecodeOptions)(nil), "protos.encoding.DecodeOptions")
}

func init() { proto.RegisterFile("ps_encoding_options.proto", fileDescriptor_b8f6b0c5cf15b9fa) }

var fileDescriptor_b8f6b0c5cf15b9fa = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0x5e, 0x13, 0xf2, 0xc3, 0x09, 0x10, 0x67, 0x92, 0x10, 0x27, 0x51, 0x14, 0x96, 0xec, 0x05,
	0xca, 0xee, 0x82, 0x94, 0x95, 0x56, 0x7b, 0xb1, 0xab, 0x15, 0x14, 0xa3, 0xa4, 0xa2, 0xd8, 0xb5,
	0x9d, 0x54, 0xcd, 0x8d, 0x05, 0xf6, 0x00, 0x23, 0x19, 0xc6, 0xb2, 0x07, 0xaa, 0x3c, 0x41, 0x5f,
	0xaa, 0x0f, 0xd0, 0xa7, 0xaa, 0x2a, 0xcf, 0xd8, 0xc6, 0x40, 0xaa, 0xaa, 0x77, 0xbd, 0x82, 0xf9,
	0xce, 0x39, 0x9f, 0xcf, 0xf9, 0xce, 0x37, 0x03, 0x67, 0x7e, 0x68, 0xe3, 0x99, 0x43, 0x5d, 0x32,
	0x1b, 0xdb, 0xd4, 0x67, 0x84, 0xce, 0xc2, 0x86, 0x1f, 0x50, 0x46, 0xd1, 0x01, 0xff, 0x09, 0x1b,
	0x49, 0xb8, 0xf6, 0x29, 0x0f, 0xb2, 0x1e, 0x61, 0xc3, 0xf9, 0xc8, 0xc4, 0x8c, 0x91, 0xd9, 0x38,
	0x44, 0xb7, 0x70, 0xe2, 0xc7, 0x98, 0x1d, 0x50, 0xca, 0xec, 0x29, 0x0e, 0xc3, 0xc1, 0x18, 0x2b,
	0x52, 0x55, 0xaa, 0x17, 0x8c, 0xa3, 0x24, 0x68, 0x50, 0xca, 0xde, 0x88, 0x10, 0xba, 0x86, 0x52,
	0x5a, 0xe3, 0x92, 0x20, 0x54, 0x72, 0xd5, 0xad, 0x7a, 0xc1, 0x28, 0x26, 0x60, 0x87, 0x04, 0x21,
	0xfa, 0x1d, 0x90, 0xbd, 0xca, 0xec, 0x92, 0x40, 0xd9, 0xe2, 0xac, 0x07, 0x7a, 0x86, 0xb5, 0x43,
	0x02, 0xa4, 0xc0, 0xee, 0x20, 0x70, 0x26, 0x64, 0x81, 0x95, 0x7c, 0x55, 0xaa, 0x17, 0x8d, 0xe4,
	0x88, 0x1a, 0x70, 0x94, 0xb4, 0x64, 0xbb, 0x38, 0x74, 0x02, 0xe2, 0x33, 0x1a, 0x28, 0xdb, 0x3c,
	0xeb, 0x30, 0xee, 0xa8, 0x93, 0x06, 0x90, 0x09, 0x95, 0xf4, 0xab, 0x78, 0xb6, 0xc0, 0x1e, 0xf5,
	0xb1, 0xcd, 0x9e, 0x7d, 0xac, 0xec, 0x54, 0xa5, 0x7a, 0xf9, 0xf6, 0xb2, 0xb1, 0x26, 0x4b, 0x43,
	0x8d, 0xb3, 0xac, 0x67, 0x1f, 0x1b, 0xc7, 0x49, 0x71, 0x16, 0x45, 0xff, 0x80, 0x12, 0x4e, 0x06,
	0x9e, 0x47, 0x3f, 0x2c, 0x39, 0x13, 0x9d, 0x76, 0xf9, 0x44, 0x95, 0x38, 0x9e, 0x94, 0x25, 0x52,
	0xb5, 0xe0, 0x72, 0xa3, 0x72, 0x44, 0xb0, 0xe7, 0xda, 0xb3, 0xf9, 0x74, 0x88, 0x03, 0x65, 0xaf,
	0x2a, 0xd5, 0xb7, 0x8d, 0xf3, 0xb5, 0xf2, 0x6e, 0x94, 0xd2, 0xe7, 0x19, 0xa8, 0x0f, 0xbf, 0xd9,
	0xdf, 0xfa, 0x7a, 0x56, 0x92, 0x02, 0x97, 0xa4, 0x6a, 0xbe, 0xd8, 0x48, 0x46, 0xa1, 0xbf, 0xe1,
	0x74, 0xb9, 0xbd, 0x14, 0xb6, 0x43, 0xcc, 0x14, 0xe0, 0xb3, 0xa4, 0x86, 0x58, 0x16, 0x99, 0x98,
	0xd5, 0x74, 0x28, 0xb6, 0x16, 0x01, 0x4d, 0x9d, 0x53, 0x07, 0x79, 0xb0, 0x08, 0xa8, 0x1d, 0x3a,
	0x13, 0x3c, 0x1d, 0xd8, 0x23, 0xe2, 0x25, 0xa6, 0x29, 0x47, 0xb8, 0xc9, 0xe1, 0x2e, 0xf1, 0x30,
	0xaa, 0xc0, 0x8e, 0x48, 0x52, 0x72, 0xbc, 0xc7, 0xf8, 0x54, 0x7b, 0x84, 0xb2, 0x35, 0x09, 0xc8,
	0x88, 0xa5, 0x9c, 0x57, 0xb0, 0xcf, 0x38, 0x22, 0x7c, 0x25, 0x71, 0x5f, 0x81, 0x80, 0xb8, 0xab,
	0xae, 0xa1, 0x14, 0x27, 0x84, 0x2c, 0x98, 0x3b, 0x8c, 0x33, 0x16, 0x8c, 0xa2, 0x00, 0x4d, 0x8e,
	0xd5, 0xfe, 0x00, 0xf4, 0xda, 0xd4, 0xfa, 0xa2, 0x83, 0x94, 0x7b, 0xd9, 0x85, 0xb4, 0xd2, 0xc5,
	0x17, 0x09, 0x4a, 0x6a, 0x64, 0x06, 0xac, 0x89, 0xfb, 0x83, 0x2e, 0xa0, 0x10, 0x0f, 0x45, 0xdc,
	0x78, 0xa4, 0x3d, 0x01, 0xdc, 0xbb, 0xe8, 0x5f, 0xd8, 0xe7, 0xd6, 0x89, 0x5d, 0x95, 0xe3, 0xae,
	0xba, 0x78, 0xc1, 0x55, 0x51, 0x0e, 0xf7, 0x14, 0xe0, 0xf4, 0x3f, 0xea, 0xc3, 0x61, 0x2a, 0x7e,
	0x18, 0x77, 0xc6, 0x2f, 0xc5, 0xfe, 0xed, 0xaf, 0x1b, 0x1c, 0xeb, 0x97, 0xd5, 0x90, 0xfd, 0xf5,
	0xeb, 0xdb, 0x86, 0x92, 0x58, 0x42, 0xc2, 0x95, 0xe7, 0x5c, 0x9b, 0x2e, 0xcf, 0xae, 0xce, 0x28,
	0x0e, 0x32, 0xa7, 0xda, 0xe7, 0x1c, 0x94, 0x3a, 0xf8, 0x47, 0x04, 0x70, 0xf1, 0xf7, 0x05, 0x10,
	0x8c, 0x42, 0x00, 0x17, 0xff, 0xcc, 0x02, 0xa0, 0x3b, 0x38, 0x48, 0x4c, 0x95, 0xb0, 0x6c, 0x73,
	0x96, 0xab, 0x0d, 0x96, 0x55, 0xbf, 0x1a, 0x65, 0xb6, 0x72, 0xbe, 0x79, 0x0b, 0xb0, 0x5c, 0x3c,
	0x3a, 0x81, 0x43, 0xb5, 0xff, 0x4a, 0xeb, 0xa8, 0xb6, 0xf5, 0x5e, 0x57, 0xed, 0x87, 0xbe, 0xa9,
	0x5a, 0xf2, 0x2f, 0xa8, 0x02, 0x28, 0x0b, 0x47, 0x56, 0xd5, 0xdb, 0xb2, 0x84, 0x8e, 0x41, 0xce,
	0xe2, 0xad, 0x47, 0x43, 0x93, 0x73, 0x37, 0x1f, 0x25, 0x80, 0xa5, 0x96, 0x11, 0x67, 0x47, 0xdd,
	0xe4, 0x54, 0xe0, 0x38, 0x0b, 0xeb, 0x86, 0x66, 0x69, 0xed, 0x87, 0xae, 0x60, 0xcd, 0x46, 0x04,
	0x6b, 0xd4, 0x43, 0x16, 0xb5, 0xee, 0x8c, 0xfb, 0xae, 0x25, 0x6f, 0xa1, 0x73, 0xa8, 0x64, 0xf1,
	0x6e, 0xaf, 0x65, 0xb5, 0x1f, 0xba, 0x5d, 0xd5, 0x90, 0xf3, 0x37, 0x4f, 0x50, 0x5c, 0x79, 0x15,
	0x4f, 0xe1, 0x48, 0xed, 0x3f, 0xaa, 0x3d, 0x4d, 0x7f, 0x69, 0xc0, 0x6c, 0xa0, 0xa3, 0xaa, 0xba,
	0x2c, 0xa1, 0x33, 0x38, 0x59, 0xc5, 0xcd, 0xbb, 0x56, 0xaf, 0xa7, 0xbd, 0x93, 0x73, 0xed, 0xff,
	0x9f, 0xfe, 0x1b, 0x13, 0x36, 0x99, 0x0f, 0x1b, 0x0e, 0x9d, 0x36, 0x87, 0x03, 0xe6, 0x4c, 0x1c,
	0x1a, 0xf8, 0x4d, 0xdf, 0xe3, 0x8f, 0xe0, 0x9f, 0xc2, 0x7d, 0x61, 0x73, 0x38, 0x27, 0x9e, 0xdb,
	0x1c, 0xd3, 0xa6, 0x58, 0x4c, 0x33, 0x59, 0xcc, 0x70, 0x87, 0x03, 0x7f, 0x7d, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x74, 0xbb, 0x60, 0xa5, 0x11, 0x07, 0x00, 0x00,
}
