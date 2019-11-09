// Code generated by protoc-gen-go. DO NOT EDIT.
// source: package.proto

package pack

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

type CallType int32

const (
	CallType_CallTypeUnknown CallType = 0
	CallType_Sync            CallType = 1
	CallType_Async           CallType = 2
	CallType_Push            CallType = 3
)

var CallType_name = map[int32]string{
	0: "CallTypeUnknown",
	1: "Sync",
	2: "Async",
	3: "Push",
}

var CallType_value = map[string]int32{
	"CallTypeUnknown": 0,
	"Sync":            1,
	"Async":           2,
	"Push":            3,
}

func (x CallType) String() string {
	return proto.EnumName(CallType_name, int32(x))
}

func (CallType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{0}
}

type RequestMeta struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	MethodName           string   `protobuf:"bytes,2,opt,name=methodName,proto3" json:"methodName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMeta) Reset()         { *m = RequestMeta{} }
func (m *RequestMeta) String() string { return proto.CompactTextString(m) }
func (*RequestMeta) ProtoMessage()    {}
func (*RequestMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{0}
}

func (m *RequestMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMeta.Unmarshal(m, b)
}
func (m *RequestMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMeta.Marshal(b, m, deterministic)
}
func (m *RequestMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMeta.Merge(m, src)
}
func (m *RequestMeta) XXX_Size() int {
	return xxx_messageInfo_RequestMeta.Size(m)
}
func (m *RequestMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMeta proto.InternalMessageInfo

func (m *RequestMeta) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RequestMeta) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

type ResponseMeta struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrText              string   `protobuf:"bytes,2,opt,name=errText,proto3" json:"errText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMeta) Reset()         { *m = ResponseMeta{} }
func (m *ResponseMeta) String() string { return proto.CompactTextString(m) }
func (*ResponseMeta) ProtoMessage()    {}
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{1}
}

func (m *ResponseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMeta.Unmarshal(m, b)
}
func (m *ResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMeta.Marshal(b, m, deterministic)
}
func (m *ResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMeta.Merge(m, src)
}
func (m *ResponseMeta) XXX_Size() int {
	return xxx_messageInfo_ResponseMeta.Size(m)
}
func (m *ResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMeta proto.InternalMessageInfo

func (m *ResponseMeta) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResponseMeta) GetErrText() string {
	if m != nil {
		return m.ErrText
	}
	return ""
}

//todo:加入压缩方式，校验和
type Header struct {
	Request              *RequestMeta  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *ResponseMeta `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Seq                  int64         `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{2}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetRequest() *RequestMeta {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Header) GetResponse() *ResponseMeta {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Header) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func init() {
	proto.RegisterEnum("intrepid.package.CallType", CallType_name, CallType_value)
	proto.RegisterType((*RequestMeta)(nil), "intrepid.package.RequestMeta")
	proto.RegisterType((*ResponseMeta)(nil), "intrepid.package.ResponseMeta")
	proto.RegisterType((*Header)(nil), "intrepid.package.Header")
}

func init() { proto.RegisterFile("package.proto", fileDescriptor_ae8103ff0e06fb71) }

var fileDescriptor_ae8103ff0e06fb71 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x0d, 0x6d, 0xd3, 0x17, 0x10, 0x96, 0x59, 0xba, 0x50, 0x45, 0x99, 0x2a, 0x86, 0x0c,
	0x65, 0x40, 0x62, 0x40, 0xa2, 0x5d, 0x58, 0xf8, 0x90, 0x29, 0x0b, 0x9b, 0x49, 0x9e, 0x68, 0xd4,
	0xd6, 0x76, 0x6d, 0x17, 0xe8, 0xaf, 0xe0, 0x2f, 0xa3, 0xb8, 0x31, 0x44, 0x88, 0xed, 0xde, 0xf9,
	0xde, 0xf9, 0x4e, 0x0f, 0x8e, 0xb5, 0x28, 0x96, 0xe2, 0x0d, 0x73, 0x6d, 0x94, 0x53, 0x8c, 0x56,
	0xd2, 0x19, 0xd4, 0x55, 0x99, 0x37, 0x7c, 0xf6, 0x00, 0x09, 0xc7, 0xcd, 0x16, 0xad, 0xbb, 0x43,
	0x27, 0x58, 0x0a, 0x89, 0x45, 0xf3, 0x5e, 0x15, 0x78, 0x2f, 0xd6, 0x38, 0x24, 0x29, 0x19, 0x0f,
	0x78, 0x9b, 0x62, 0x23, 0x80, 0x35, 0xba, 0x85, 0x2a, 0xbd, 0xa0, 0xe3, 0x05, 0x2d, 0x26, 0x9b,
	0xc2, 0x11, 0x47, 0xab, 0x95, 0xb4, 0xe8, 0x1d, 0x87, 0xd0, 0x47, 0x63, 0x66, 0xaa, 0xdc, 0xbb,
	0x75, 0x79, 0x18, 0x9b, 0x97, 0x39, 0x7e, 0xba, 0xc6, 0x26, 0x8c, 0xd9, 0x17, 0x81, 0xde, 0x2d,
	0x8a, 0x12, 0x0d, 0xbb, 0x84, 0xbe, 0xd9, 0xe7, 0xf3, 0xeb, 0xc9, 0xe4, 0x2c, 0xff, 0xdb, 0x21,
	0x6f, 0x15, 0xe0, 0x41, 0xcd, 0xae, 0x20, 0x36, 0x4d, 0x0e, 0x6f, 0x9f, 0x4c, 0x46, 0xff, 0x6d,
	0xfe, 0x26, 0xe5, 0x3f, 0x7a, 0x46, 0x21, 0xb2, 0xb8, 0x19, 0x46, 0x29, 0x19, 0x47, 0xbc, 0x86,
	0xe7, 0xd7, 0x10, 0xcf, 0xc4, 0x6a, 0x35, 0xdf, 0x69, 0x64, 0xa7, 0x70, 0x12, 0xf0, 0xb3, 0x5c,
	0x4a, 0xf5, 0x21, 0xe9, 0x01, 0x8b, 0xe1, 0xf0, 0x69, 0x27, 0x0b, 0x4a, 0xd8, 0x00, 0xba, 0x37,
	0xb6, 0x86, 0x9d, 0x9a, 0x7c, 0xdc, 0xda, 0x05, 0x8d, 0xa6, 0xf0, 0x12, 0x87, 0xcf, 0x5f, 0x7b,
	0xfe, 0x16, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0x7d, 0x76, 0xc7, 0x9c, 0x01, 0x00,
	0x00,
}