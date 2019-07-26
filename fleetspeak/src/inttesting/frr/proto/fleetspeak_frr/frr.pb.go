// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/inttesting/frr/proto/fleetspeak_frr/frr.proto

package fleetspeak_frr

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	grpc "google.golang.org/grpc"
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

// Contains the information needed to configure a frr server component.
type Config struct {
	// The address to reach the master frr server over grpc.
	MasterServer         string   `protobuf:"bytes,1,opt,name=master_server,json=masterServer,proto3" json:"master_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetMasterServer() string {
	if m != nil {
		return m.MasterServer
	}
	return ""
}

// A TrafficRequest message is sent from the server to the client which tells the
// client to send random data back.
type TrafficRequestData struct {
	// An identifier used to identify the frr master instance responsible for this.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// An identifier used to link responses to requests.
	RequestId int64 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// How many messages to send back for this request. Defaults to 1.
	NumMessages int64 `protobuf:"varint,3,opt,name=num_messages,json=numMessages,proto3" json:"num_messages,omitempty"`
	// How large should each message be, in bytes. Defaults to 1024.
	MessageSize int64 `protobuf:"varint,4,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	// How long to wait between messages. Defaults to 0.
	MessageDelayMs int64 `protobuf:"varint,5,opt,name=message_delay_ms,json=messageDelayMs,proto3" json:"message_delay_ms,omitempty"`
	// How much to jitter the previous parameters - all parameters will be
	// multiplied by a number between 1.0 and 1.0 + jitter.
	Jitter               float32  `protobuf:"fixed32,6,opt,name=jitter,proto3" json:"jitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficRequestData) Reset()         { *m = TrafficRequestData{} }
func (m *TrafficRequestData) String() string { return proto.CompactTextString(m) }
func (*TrafficRequestData) ProtoMessage()    {}
func (*TrafficRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{1}
}

func (m *TrafficRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficRequestData.Unmarshal(m, b)
}
func (m *TrafficRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficRequestData.Marshal(b, m, deterministic)
}
func (m *TrafficRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficRequestData.Merge(m, src)
}
func (m *TrafficRequestData) XXX_Size() int {
	return xxx_messageInfo_TrafficRequestData.Size(m)
}
func (m *TrafficRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficRequestData proto.InternalMessageInfo

func (m *TrafficRequestData) GetMasterId() int64 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *TrafficRequestData) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *TrafficRequestData) GetNumMessages() int64 {
	if m != nil {
		return m.NumMessages
	}
	return 0
}

func (m *TrafficRequestData) GetMessageSize() int64 {
	if m != nil {
		return m.MessageSize
	}
	return 0
}

func (m *TrafficRequestData) GetMessageDelayMs() int64 {
	if m != nil {
		return m.MessageDelayMs
	}
	return 0
}

func (m *TrafficRequestData) GetJitter() float32 {
	if m != nil {
		return m.Jitter
	}
	return 0
}

type TrafficResponseData struct {
	MasterId      int64  `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	RequestId     int64  `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ResponseIndex int64  `protobuf:"varint,3,opt,name=response_index,json=responseIndex,proto3" json:"response_index,omitempty"`
	Data          []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Set when this is the last message responsive to the given request.
	Fin                  bool     `protobuf:"varint,5,opt,name=fin,proto3" json:"fin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficResponseData) Reset()         { *m = TrafficResponseData{} }
func (m *TrafficResponseData) String() string { return proto.CompactTextString(m) }
func (*TrafficResponseData) ProtoMessage()    {}
func (*TrafficResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{2}
}

func (m *TrafficResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficResponseData.Unmarshal(m, b)
}
func (m *TrafficResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficResponseData.Marshal(b, m, deterministic)
}
func (m *TrafficResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficResponseData.Merge(m, src)
}
func (m *TrafficResponseData) XXX_Size() int {
	return xxx_messageInfo_TrafficResponseData.Size(m)
}
func (m *TrafficResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficResponseData proto.InternalMessageInfo

func (m *TrafficResponseData) GetMasterId() int64 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *TrafficResponseData) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *TrafficResponseData) GetResponseIndex() int64 {
	if m != nil {
		return m.ResponseIndex
	}
	return 0
}

func (m *TrafficResponseData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TrafficResponseData) GetFin() bool {
	if m != nil {
		return m.Fin
	}
	return false
}

// A FileRequest is sent from the server to the client and tells
// the client to attempt to download a file from the server.
type FileRequestData struct {
	// An identifier used to identify the frr master instance
	// responsible for this.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The name of the file to download.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequestData) Reset()         { *m = FileRequestData{} }
func (m *FileRequestData) String() string { return proto.CompactTextString(m) }
func (*FileRequestData) ProtoMessage()    {}
func (*FileRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{3}
}

func (m *FileRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequestData.Unmarshal(m, b)
}
func (m *FileRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequestData.Marshal(b, m, deterministic)
}
func (m *FileRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequestData.Merge(m, src)
}
func (m *FileRequestData) XXX_Size() int {
	return xxx_messageInfo_FileRequestData.Size(m)
}
func (m *FileRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequestData proto.InternalMessageInfo

func (m *FileRequestData) GetMasterId() int64 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *FileRequestData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A FileResponse is sent from the client to the server and
// reports that the client successfully downloaded a file from the
// server.
type FileResponseData struct {
	// An identifier used to identify the frr master instance
	// responsible for the underlying request.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The name of the file that was downloaded.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The size of the file that was downloaded.
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponseData) Reset()         { *m = FileResponseData{} }
func (m *FileResponseData) String() string { return proto.CompactTextString(m) }
func (*FileResponseData) ProtoMessage()    {}
func (*FileResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{4}
}

func (m *FileResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponseData.Unmarshal(m, b)
}
func (m *FileResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponseData.Marshal(b, m, deterministic)
}
func (m *FileResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponseData.Merge(m, src)
}
func (m *FileResponseData) XXX_Size() int {
	return xxx_messageInfo_FileResponseData.Size(m)
}
func (m *FileResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponseData proto.InternalMessageInfo

func (m *FileResponseData) GetMasterId() int64 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *FileResponseData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileResponseData) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MessageInfo struct {
	ClientId             []byte               `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Data                 *TrafficResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{5}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *MessageInfo) GetData() *TrafficResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileResponseInfo struct {
	ClientId             []byte            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Data                 *FileResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FileResponseInfo) Reset()         { *m = FileResponseInfo{} }
func (m *FileResponseInfo) String() string { return proto.CompactTextString(m) }
func (*FileResponseInfo) ProtoMessage()    {}
func (*FileResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb0d47701b8c91b, []int{6}
}

func (m *FileResponseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponseInfo.Unmarshal(m, b)
}
func (m *FileResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponseInfo.Marshal(b, m, deterministic)
}
func (m *FileResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponseInfo.Merge(m, src)
}
func (m *FileResponseInfo) XXX_Size() int {
	return xxx_messageInfo_FileResponseInfo.Size(m)
}
func (m *FileResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponseInfo proto.InternalMessageInfo

func (m *FileResponseInfo) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *FileResponseInfo) GetData() *FileResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "fleetspeak.frr.Config")
	proto.RegisterType((*TrafficRequestData)(nil), "fleetspeak.frr.TrafficRequestData")
	proto.RegisterType((*TrafficResponseData)(nil), "fleetspeak.frr.TrafficResponseData")
	proto.RegisterType((*FileRequestData)(nil), "fleetspeak.frr.FileRequestData")
	proto.RegisterType((*FileResponseData)(nil), "fleetspeak.frr.FileResponseData")
	proto.RegisterType((*MessageInfo)(nil), "fleetspeak.frr.MessageInfo")
	proto.RegisterType((*FileResponseInfo)(nil), "fleetspeak.frr.FileResponseInfo")
}

func init() {
	proto.RegisterFile("fleetspeak/src/inttesting/frr/proto/fleetspeak_frr/frr.proto", fileDescriptor_7fb0d47701b8c91b)
}

var fileDescriptor_7fb0d47701b8c91b = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xdd, 0x26, 0x86, 0xe4, 0x24, 0x8d, 0xe1, 0x88, 0xb2, 0xb4, 0x08, 0x71, 0x8b, 0x90,
	0x1b, 0xb3, 0xd0, 0x0a, 0xde, 0x78, 0xa5, 0x55, 0xd8, 0x8b, 0x80, 0x4c, 0x05, 0x2f, 0x97, 0x71,
	0xf7, 0x6c, 0x18, 0xcd, 0xce, 0xc6, 0x99, 0x89, 0xd8, 0x3e, 0x8c, 0x8f, 0xe0, 0xf3, 0xf8, 0x38,
	0x32, 0x7f, 0x62, 0x93, 0xb5, 0xd8, 0x82, 0x77, 0x93, 0xdf, 0xf9, 0xe6, 0x3b, 0xe7, 0x3b, 0x93,
	0x85, 0x57, 0xd5, 0x8a, 0xc8, 0xe8, 0x35, 0xf1, 0x2f, 0xa9, 0x56, 0x45, 0x2a, 0xa4, 0x31, 0xa4,
	0x8d, 0x90, 0xcb, 0xb4, 0x52, 0x2a, 0x5d, 0xab, 0xc6, 0x34, 0xe9, 0xb5, 0x26, 0xb7, 0xb0, 0x52,
	0x6a, 0xee, 0x0a, 0x38, 0xbe, 0xae, 0xcc, 0x2b, 0xa5, 0x8e, 0xce, 0x5a, 0x6e, 0x45, 0x53, 0xd7,
	0x8d, 0xfc, 0xcb, 0x25, 0x70, 0x6f, 0x92, 0x3c, 0x87, 0xde, 0x9b, 0x46, 0x56, 0x62, 0x89, 0x27,
	0x70, 0x58, 0x73, 0x6d, 0x48, 0xe5, 0x9a, 0xd4, 0x37, 0x52, 0x71, 0x34, 0x8d, 0x66, 0x03, 0x36,
	0xf2, 0xf0, 0xc2, 0xb1, 0xe4, 0x57, 0x04, 0xf8, 0x41, 0xf1, 0xaa, 0x12, 0x05, 0xa3, 0xaf, 0x1b,
	0xd2, 0xe6, 0x9c, 0x1b, 0x8e, 0xc7, 0x30, 0x08, 0x77, 0x45, 0xe9, 0xee, 0x75, 0x58, 0xdf, 0x83,
	0xac, 0xc4, 0x27, 0x00, 0xca, 0x6b, 0x6d, 0xf5, 0xc0, 0x55, 0x07, 0x81, 0x64, 0x25, 0x3e, 0x85,
	0x91, 0xdc, 0xd4, 0x79, 0x4d, 0x5a, 0xf3, 0x25, 0xe9, 0xb8, 0xe3, 0x04, 0x43, 0xb9, 0xa9, 0x17,
	0x01, 0x59, 0x49, 0x28, 0xe7, 0x5a, 0x5c, 0x51, 0xdc, 0xf5, 0x92, 0xc0, 0x2e, 0xc4, 0x15, 0xe1,
	0x0c, 0x26, 0x5b, 0x49, 0x49, 0x2b, 0x7e, 0x99, 0xd7, 0x3a, 0xbe, 0xef, 0x64, 0xe3, 0xc0, 0xcf,
	0x2d, 0x5e, 0x68, 0x7c, 0x0c, 0xbd, 0xcf, 0xc2, 0x18, 0x52, 0x71, 0x6f, 0x1a, 0xcd, 0x0e, 0x58,
	0xf8, 0x95, 0xfc, 0x88, 0xe0, 0xe1, 0x9f, 0x68, 0x7a, 0xdd, 0x48, 0x4d, 0xff, 0x9d, 0xed, 0x19,
	0x8c, 0x55, 0xf0, 0xca, 0x85, 0x2c, 0xe9, 0x7b, 0x48, 0x77, 0xb8, 0xa5, 0x99, 0x85, 0x88, 0xd0,
	0x2d, 0xb9, 0xe1, 0x2e, 0xd7, 0x88, 0xb9, 0x33, 0x4e, 0xa0, 0x53, 0x09, 0xe9, 0x32, 0xf4, 0x99,
	0x3d, 0x26, 0xaf, 0xe1, 0xc1, 0x3b, 0xb1, 0xa2, 0x3b, 0xef, 0x1d, 0xa1, 0x2b, 0x79, 0x4d, 0x6e,
	0xaa, 0x01, 0x73, 0xe7, 0xe4, 0x23, 0x4c, 0xbc, 0xc7, 0x5d, 0x03, 0xde, 0x60, 0x62, 0x99, 0x7b,
	0x06, 0x9b, 0xa5, 0xcb, 0xdc, 0x39, 0x29, 0x60, 0x18, 0x9e, 0x2b, 0x93, 0x55, 0x63, 0x3d, 0x8b,
	0x95, 0x20, 0x69, 0xb6, 0x9e, 0x23, 0xd6, 0xf7, 0x20, 0x2b, 0xf1, 0x65, 0x88, 0x6b, 0x3d, 0x87,
	0xa7, 0x27, 0xf3, 0xfd, 0xff, 0xf1, 0xfc, 0x86, 0x47, 0xf0, 0x3b, 0x49, 0x68, 0x7f, 0xfa, 0xdb,
	0x3b, 0xbd, 0xd8, 0xeb, 0x34, 0x6d, 0x77, 0x6a, 0xaf, 0xc2, 0xb7, 0x39, 0xfd, 0x19, 0x41, 0x6f,
	0xe1, 0x16, 0x80, 0xef, 0xe1, 0x11, 0xa3, 0xa2, 0x51, 0x65, 0x6b, 0x28, 0x3c, 0x6e, 0x7b, 0xed,
	0xa4, 0x3f, 0x8a, 0x77, 0x8b, 0x6f, 0xeb, 0xb5, 0xb9, 0x0c, 0xd5, 0xe4, 0x1e, 0x32, 0x40, 0xef,
	0xb8, 0xdb, 0x1c, 0xff, 0x39, 0xda, 0x6d, 0x9e, 0x9f, 0x7a, 0xee, 0x5b, 0x3e, 0xfb, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x80, 0xac, 0x45, 0x50, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	// RecordMessage records that a TrafficResponse message was received by
	// the FS server.
	RecordTrafficResponse(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// RecordFileResponse records that a FileResponse message was received
	// by the FS server.
	RecordFileResponse(ctx context.Context, in *FileResponseInfo, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) RecordTrafficResponse(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.frr.Master/RecordTrafficResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) RecordFileResponse(ctx context.Context, in *FileResponseInfo, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.frr.Master/RecordFileResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	// RecordMessage records that a TrafficResponse message was received by
	// the FS server.
	RecordTrafficResponse(context.Context, *MessageInfo) (*fleetspeak.EmptyMessage, error)
	// RecordFileResponse records that a FileResponse message was received
	// by the FS server.
	RecordFileResponse(context.Context, *FileResponseInfo) (*fleetspeak.EmptyMessage, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_RecordTrafficResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RecordTrafficResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.frr.Master/RecordTrafficResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RecordTrafficResponse(ctx, req.(*MessageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_RecordFileResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileResponseInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RecordFileResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.frr.Master/RecordFileResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RecordFileResponse(ctx, req.(*FileResponseInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fleetspeak.frr.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordTrafficResponse",
			Handler:    _Master_RecordTrafficResponse_Handler,
		},
		{
			MethodName: "RecordFileResponse",
			Handler:    _Master_RecordFileResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fleetspeak/src/inttesting/frr/proto/fleetspeak_frr/frr.proto",
}
