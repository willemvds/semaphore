// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package proto

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/jexia/maestro/annotations"

	context "golang.org/x/net/context"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_61adbf83eb3926e4, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_61adbf83eb3926e4, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "maestro.greeter.Request")
	proto.RegisterType((*Response)(nil), "maestro.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SayClient is the client API for Say service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayClient struct {
	cc *grpc.ClientConn
}

func NewSayClient(cc *grpc.ClientConn) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/maestro.greeter.Say/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayServer is the server API for Say service.
type SayServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterSayServer(s *grpc.Server, srv SayServer) {
	s.RegisterService(&_Say_serviceDesc, srv)
}

func _Say_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maestro.greeter.Say/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Say_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maestro.greeter.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Say_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_61adbf83eb3926e4) }

var fileDescriptor_hello_61adbf83eb3926e4 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x4d, 0x4c, 0x2d, 0x2e, 0x29, 0xca,
	0xd7, 0x4b, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x92, 0x12, 0x4c, 0xcc, 0xcb, 0xcb, 0x2f, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x51, 0x92, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0xa1, 0xd2, 0x20, 0xa6, 0x51, 0x26, 0x17, 0x73, 0x70, 0x62,
	0xa5, 0x90, 0x1d, 0x17, 0xab, 0x07, 0xc8, 0x5a, 0x21, 0x09, 0x3d, 0x34, 0x1b, 0xf5, 0xa0, 0x66,
	0x4b, 0x49, 0x62, 0x91, 0x81, 0x18, 0xab, 0xc4, 0x20, 0xa5, 0xf4, 0x68, 0xab, 0x84, 0x9c, 0x14,
	0xbf, 0xa1, 0x91, 0xb9, 0x9e, 0x81, 0x9e, 0x81, 0x9e, 0xa1, 0x95, 0xa9, 0x81, 0x81, 0xa9, 0xa1,
	0x12, 0x4b, 0x7a, 0x51, 0x41, 0xb2, 0x16, 0x2b, 0xd8, 0x9d, 0x4e, 0xec, 0x51, 0x10, 0x46, 0x12,
	0x1b, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe7, 0xd2, 0x63, 0xea, 0x00, 0x00,
	0x00,
}
