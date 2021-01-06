// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package message_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Subtext              string   `protobuf:"bytes,2,opt,name=subtext,proto3" json:"subtext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Request) GetSubtext() string {
	if m != nil {
		return m.Subtext
	}
	return ""
}

type Response struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Subtext              string   `protobuf:"bytes,2,opt,name=subtext,proto3" json:"subtext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Response) GetSubtext() string {
	if m != nil {
		return m.Subtext
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "message_service.Request")
	proto.RegisterType((*Response)(nil), "message_service.Response")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x8d, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x32, 0xe7, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0xc0, 0xc2, 0x4c, 0x60, 0x61, 0x18, 0x57,
	0xc9, 0x82, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x95, 0x34, 0x9d, 0x46, 0x7e,
	0x5c, 0x7c, 0xbe, 0x10, 0x57, 0x04, 0x43, 0x1c, 0x21, 0x64, 0xc3, 0xc5, 0x1a, 0x9c, 0x58, 0xe9,
	0x59, 0x22, 0x24, 0xa1, 0x87, 0xe6, 0x3e, 0x3d, 0xa8, 0xe3, 0xa4, 0x24, 0xb1, 0xc8, 0x40, 0x6c,
	0x4f, 0x62, 0x03, 0x7b, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x63, 0x98, 0xe2, 0x43, 0xe7,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	SayIt(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SayIt(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message_service.MessageService/SayIt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	SayIt(context.Context, *Request) (*Response, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) SayIt(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayIt not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_SayIt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SayIt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_service.MessageService/SayIt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SayIt(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message_service.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayIt",
			Handler:    _MessageService_SayIt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
