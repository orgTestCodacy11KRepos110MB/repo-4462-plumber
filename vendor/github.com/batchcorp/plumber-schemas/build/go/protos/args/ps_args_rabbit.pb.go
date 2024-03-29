// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_rabbit.proto

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

type RabbitConn struct {
	// @gotags: kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"`
	// @gotags: kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"
	UseTls bool `protobuf:"varint,2,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"`
	// @gotags: kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"
	TlsSkipVerify        bool     `protobuf:"varint,3,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitConn) Reset()         { *m = RabbitConn{} }
func (m *RabbitConn) String() string { return proto.CompactTextString(m) }
func (*RabbitConn) ProtoMessage()    {}
func (*RabbitConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{0}
}

func (m *RabbitConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitConn.Unmarshal(m, b)
}
func (m *RabbitConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitConn.Marshal(b, m, deterministic)
}
func (m *RabbitConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitConn.Merge(m, src)
}
func (m *RabbitConn) XXX_Size() int {
	return xxx_messageInfo_RabbitConn.Size(m)
}
func (m *RabbitConn) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitConn.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitConn proto.InternalMessageInfo

func (m *RabbitConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RabbitConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *RabbitConn) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

type RabbitReadArgs struct {
	// @gotags: kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"`
	// @gotags: kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"
	QueueName string `protobuf:"bytes,2,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"`
	// @gotags: kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"
	BindingKey string `protobuf:"bytes,3,opt,name=binding_key,json=bindingKey,proto3" json:"binding_key,omitempty" kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"`
	// @gotags: kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"
	QueueExclusive bool `protobuf:"varint,4,opt,name=queue_exclusive,json=queueExclusive,proto3" json:"queue_exclusive,omitempty" kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"`
	// @gotags: kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"
	QueueDeclare bool `protobuf:"varint,5,opt,name=queue_declare,json=queueDeclare,proto3" json:"queue_declare,omitempty" kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"`
	// @gotags: kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"
	QueueDurable bool `protobuf:"varint,6,opt,name=queue_durable,json=queueDurable,proto3" json:"queue_durable,omitempty" kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"`
	// @gotags: kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"
	AutoAck bool `protobuf:"varint,7,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty" kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"`
	// @gotags: kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"
	ConsumerTag string `protobuf:"bytes,8,opt,name=consumer_tag,json=consumerTag,proto3" json:"consumer_tag,omitempty" kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"`
	// @gotags: kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"
	QueueDelete bool `protobuf:"varint,9,opt,name=queue_delete,json=queueDelete,proto3" json:"queue_delete,omitempty" kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"`
	// @gotags: kong:"help='Key=Value pair for sending additional queue argument to RabbitMQ. Example: --queue-arg x-dead-letter-exchange=mydlexchange --queue-arg x-queue-mode=lazy',env='PLUMBER_RELAY_RABBIT_QUEUE_ARGS'"
	QueueArg map[string]string `protobuf:"bytes,10,rep,name=queue_arg,json=queueArg,proto3" json:"queue_arg,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" kong:"help='Key=Value pair for sending additional queue argument to RabbitMQ. Example: --queue-arg x-dead-letter-exchange=mydlexchange --queue-arg x-queue-mode=lazy',env='PLUMBER_RELAY_RABBIT_QUEUE_ARGS'"`
	// @gotags: kong:"help='Exclude messages with routing key matching regex',env='PLUMBER_RELAY_RABBIT_BINDING_KEY_EXCLUDE_REGEX'"
	ExcludeBindingKeyRegex string `protobuf:"bytes,11,opt,name=exclude_binding_key_regex,json=excludeBindingKeyRegex,proto3" json:"exclude_binding_key_regex,omitempty" kong:"help='Exclude messages with routing key matching regex',env='PLUMBER_RELAY_RABBIT_BINDING_KEY_EXCLUDE_REGEX'"`
	// @gotags: kong:"help='Consumed messages are intended to be dead-lettered',env='PLUMBER_RELAY_RABBIT_DEAD_LETTER'"
	DeadLetter           bool     `protobuf:"varint,12,opt,name=dead_letter,json=deadLetter,proto3" json:"dead_letter,omitempty" kong:"help='Consumed messages are intended to be dead-lettered',env='PLUMBER_RELAY_RABBIT_DEAD_LETTER'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitReadArgs) Reset()         { *m = RabbitReadArgs{} }
func (m *RabbitReadArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitReadArgs) ProtoMessage()    {}
func (*RabbitReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{1}
}

func (m *RabbitReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitReadArgs.Unmarshal(m, b)
}
func (m *RabbitReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitReadArgs.Marshal(b, m, deterministic)
}
func (m *RabbitReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitReadArgs.Merge(m, src)
}
func (m *RabbitReadArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitReadArgs.Size(m)
}
func (m *RabbitReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitReadArgs proto.InternalMessageInfo

func (m *RabbitReadArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *RabbitReadArgs) GetBindingKey() string {
	if m != nil {
		return m.BindingKey
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueExclusive() bool {
	if m != nil {
		return m.QueueExclusive
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDeclare() bool {
	if m != nil {
		return m.QueueDeclare
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDurable() bool {
	if m != nil {
		return m.QueueDurable
	}
	return false
}

func (m *RabbitReadArgs) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *RabbitReadArgs) GetConsumerTag() string {
	if m != nil {
		return m.ConsumerTag
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueDelete() bool {
	if m != nil {
		return m.QueueDelete
	}
	return false
}

func (m *RabbitReadArgs) GetQueueArg() map[string]string {
	if m != nil {
		return m.QueueArg
	}
	return nil
}

func (m *RabbitReadArgs) GetExcludeBindingKeyRegex() string {
	if m != nil {
		return m.ExcludeBindingKeyRegex
	}
	return ""
}

func (m *RabbitReadArgs) GetDeadLetter() bool {
	if m != nil {
		return m.DeadLetter
	}
	return false
}

type RabbitWriteArgs struct {
	// @gotags: kong:"help='Exchange to write message(s) to',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Exchange to write message(s) to',required"`
	// @gotags: kong:"help='Routing key to write message(s) to',required"
	RoutingKey string `protobuf:"bytes,2,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty" kong:"help='Routing key to write message(s) to',required"`
	// @gotags: kong:"help='Fills message properties $app_id with this value',default=plumber"
	AppId string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" kong:"help='Fills message properties  with this value',default=plumber"`
	// @gotags: kong:"help='The type of exchange we are working with',enum='direct,topic,headers,fanout',default=topic,group=exchange"
	ExchangeType string `protobuf:"bytes,4,opt,name=exchange_type,json=exchangeType,proto3" json:"exchange_type,omitempty" kong:"help='The type of exchange we are working with',enum='direct,topic,headers,fanout',default=topic,group=exchange"`
	// @gotags: kong:"help='Whether to declare an exchange (if it does not exist)',group=exchange"
	ExchangeDeclare bool `protobuf:"varint,5,opt,name=exchange_declare,json=exchangeDeclare,proto3" json:"exchange_declare,omitempty" kong:"help='Whether to declare an exchange (if it does not exist)',group=exchange"`
	// @gotags: kong:"help='Whether to make a declared exchange durable',group=exchange"
	ExchangeDurable bool `protobuf:"varint,6,opt,name=exchange_durable,json=exchangeDurable,proto3" json:"exchange_durable,omitempty" kong:"help='Whether to make a declared exchange durable',group=exchange"`
	// @gotags: kong:"help='Whether to auto-delete the exchange (after writes)',group=exchange"
	ExchangeAutoDelete   bool     `protobuf:"varint,7,opt,name=exchange_auto_delete,json=exchangeAutoDelete,proto3" json:"exchange_auto_delete,omitempty" kong:"help='Whether to auto-delete the exchange (after writes)',group=exchange"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitWriteArgs) Reset()         { *m = RabbitWriteArgs{} }
func (m *RabbitWriteArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitWriteArgs) ProtoMessage()    {}
func (*RabbitWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{2}
}

func (m *RabbitWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitWriteArgs.Unmarshal(m, b)
}
func (m *RabbitWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitWriteArgs.Marshal(b, m, deterministic)
}
func (m *RabbitWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitWriteArgs.Merge(m, src)
}
func (m *RabbitWriteArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitWriteArgs.Size(m)
}
func (m *RabbitWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitWriteArgs proto.InternalMessageInfo

func (m *RabbitWriteArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitWriteArgs) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *RabbitWriteArgs) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RabbitWriteArgs) GetExchangeType() string {
	if m != nil {
		return m.ExchangeType
	}
	return ""
}

func (m *RabbitWriteArgs) GetExchangeDeclare() bool {
	if m != nil {
		return m.ExchangeDeclare
	}
	return false
}

func (m *RabbitWriteArgs) GetExchangeDurable() bool {
	if m != nil {
		return m.ExchangeDurable
	}
	return false
}

func (m *RabbitWriteArgs) GetExchangeAutoDelete() bool {
	if m != nil {
		return m.ExchangeAutoDelete
	}
	return false
}

func init() {
	proto.RegisterType((*RabbitConn)(nil), "protos.args.RabbitConn")
	proto.RegisterType((*RabbitReadArgs)(nil), "protos.args.RabbitReadArgs")
	proto.RegisterMapType((map[string]string)(nil), "protos.args.RabbitReadArgs.QueueArgEntry")
	proto.RegisterType((*RabbitWriteArgs)(nil), "protos.args.RabbitWriteArgs")
}

func init() { proto.RegisterFile("ps_args_rabbit.proto", fileDescriptor_01d1eee3dc8ebf97) }

var fileDescriptor_01d1eee3dc8ebf97 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x95, 0xb4, 0xcd, 0xcb, 0x38, 0x6d, 0xaa, 0x55, 0x9f, 0x87, 0x2d, 0x12, 0x6a, 0x29,
	0x12, 0xb4, 0x07, 0x12, 0x04, 0x17, 0x4a, 0x4f, 0x29, 0x14, 0x09, 0x81, 0x90, 0x30, 0x15, 0x48,
	0x5c, 0x56, 0x6b, 0x7b, 0x70, 0xac, 0xac, 0xed, 0xed, 0xbe, 0x54, 0xcd, 0x67, 0xe1, 0x1b, 0xf0,
	0x29, 0xd1, 0xee, 0xda, 0x25, 0xbe, 0x71, 0xb2, 0xe7, 0x37, 0x7f, 0xcf, 0x8c, 0x67, 0xfe, 0x70,
	0x20, 0x35, 0xe3, 0x2a, 0xd7, 0x4c, 0xf1, 0x24, 0x29, 0xcc, 0x4c, 0xaa, 0xda, 0xd4, 0x24, 0xf2,
	0x0f, 0x3d, 0x73, 0x99, 0x93, 0x1c, 0x20, 0xf6, 0xc9, 0xb7, 0x75, 0x55, 0x11, 0x0a, 0x43, 0x9e,
	0x65, 0x0a, 0xb5, 0xa6, 0xbd, 0xe3, 0xde, 0xe9, 0x38, 0x6e, 0x43, 0xf2, 0x00, 0x86, 0x56, 0x23,
	0x33, 0x42, 0xd3, 0xfe, 0x71, 0xef, 0x74, 0x14, 0x0f, 0xac, 0xc6, 0x6b, 0xa1, 0xc9, 0x53, 0x98,
	0x1a, 0xa1, 0x99, 0x5e, 0x15, 0x92, 0xdd, 0xa2, 0x2a, 0x7e, 0xae, 0xe9, 0x96, 0x17, 0xec, 0x1a,
	0xa1, 0xbf, 0xae, 0x0a, 0xf9, 0xcd, 0xc3, 0x93, 0xdf, 0xdb, 0xb0, 0x17, 0x3a, 0xc5, 0xc8, 0xb3,
	0x85, 0xca, 0x35, 0x79, 0x02, 0xbb, 0x78, 0x97, 0x2e, 0x79, 0x95, 0x23, 0xab, 0x78, 0x89, 0x4d,
	0xcf, 0x49, 0x0b, 0x3f, 0xf3, 0x12, 0xc9, 0x23, 0x80, 0x1b, 0x8b, 0xb6, 0x51, 0xf4, 0xbd, 0x62,
	0xec, 0x89, 0x4f, 0x1f, 0x41, 0x94, 0x14, 0x55, 0x56, 0x54, 0x39, 0x5b, 0x61, 0x68, 0x3d, 0x8e,
	0xa1, 0x41, 0x1f, 0x71, 0x4d, 0x9e, 0xc1, 0x34, 0x7c, 0x8f, 0x77, 0xa9, 0xb0, 0xba, 0xb8, 0x45,
	0xba, 0xed, 0xe7, 0xdb, 0xf3, 0xf8, 0xaa, 0xa5, 0x6e, 0x9a, 0x20, 0xcc, 0x30, 0x15, 0x5c, 0x21,
	0xdd, 0xf1, 0xb2, 0x89, 0x87, 0xef, 0x02, 0xdb, 0x10, 0x59, 0xc5, 0x13, 0x81, 0x74, 0xb0, 0x29,
	0x0a, 0x8c, 0x1c, 0xc2, 0x88, 0x5b, 0x53, 0x33, 0x9e, 0xae, 0xe8, 0xd0, 0xe7, 0x87, 0x2e, 0x5e,
	0xa4, 0x2b, 0xf2, 0x18, 0x26, 0x69, 0x5d, 0x69, 0x5b, 0xa2, 0x62, 0x86, 0xe7, 0x74, 0xe4, 0xe7,
	0x8d, 0x5a, 0x76, 0xcd, 0x73, 0x27, 0x69, 0xe7, 0x10, 0x68, 0x90, 0x8e, 0x7d, 0x85, 0xa8, 0x19,
	0xc3, 0x21, 0xf2, 0x1e, 0xc2, 0x06, 0xdc, 0x71, 0x29, 0x1c, 0x6f, 0x9d, 0x46, 0x2f, 0xcf, 0x66,
	0x1b, 0x57, 0x9d, 0x75, 0x17, 0x3d, 0xfb, 0xe2, 0xc4, 0x0b, 0x95, 0x5f, 0x55, 0x46, 0xad, 0xe3,
	0xd1, 0x4d, 0x13, 0x92, 0x73, 0x38, 0xf4, 0x5b, 0xc9, 0x90, 0x6d, 0x2c, 0x91, 0x29, 0xcc, 0xf1,
	0x8e, 0x46, 0x7e, 0xb4, 0xff, 0x1b, 0xc1, 0xe5, 0xfd, 0x46, 0x63, 0x97, 0x75, 0x7b, 0xcf, 0x90,
	0x67, 0x4c, 0xa0, 0x31, 0xa8, 0xe8, 0xc4, 0x0f, 0x09, 0x0e, 0x7d, 0xf2, 0xe4, 0xe1, 0x05, 0xec,
	0x76, 0xda, 0x92, 0x7d, 0xd8, 0x72, 0x17, 0x0a, 0x37, 0x76, 0xaf, 0xe4, 0x00, 0x76, 0x6e, 0xb9,
	0xb0, 0xed, 0x55, 0x43, 0xf0, 0xa6, 0xff, 0xba, 0x77, 0xf2, 0xab, 0x0f, 0xd3, 0xf0, 0x0f, 0xdf,
	0x55, 0x61, 0xf0, 0xdf, 0xdd, 0x72, 0x04, 0x91, 0xaa, 0xad, 0x69, 0xed, 0x10, 0x0a, 0x43, 0x83,
	0x9c, 0x1d, 0xfe, 0x83, 0x01, 0x97, 0x92, 0x15, 0x59, 0x63, 0x95, 0x1d, 0x2e, 0xe5, 0x87, 0xac,
	0x53, 0xdc, 0xac, 0x65, 0xf0, 0xc8, 0x46, 0xf1, 0xeb, 0xb5, 0x44, 0x72, 0x06, 0xfb, 0xf7, 0xa2,
	0xae, 0x49, 0xa6, 0x2d, 0x6f, 0x7d, 0xd2, 0x91, 0x76, 0xac, 0xf2, 0x57, 0xda, 0xb8, 0xe5, 0x05,
	0x1c, 0xdc, 0x4b, 0xbd, 0x6d, 0x9a, 0xbb, 0x07, 0xe7, 0x90, 0x36, 0xb7, 0xb0, 0xa6, 0x0e, 0xe7,
	0xbf, 0xbc, 0xf8, 0x71, 0x9e, 0x17, 0x66, 0x69, 0x93, 0x59, 0x5a, 0x97, 0xf3, 0x84, 0x9b, 0x74,
	0x99, 0xd6, 0x4a, 0xce, 0xa5, 0xb0, 0x65, 0x82, 0xea, 0xb9, 0x4e, 0x97, 0x58, 0x72, 0x3d, 0x4f,
	0x6c, 0x21, 0xb2, 0x79, 0x5e, 0xcf, 0x83, 0x35, 0xe6, 0xce, 0x1a, 0xc9, 0xc0, 0x07, 0xaf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa2, 0x1b, 0x00, 0x1c, 0x04, 0x00, 0x00,
}
