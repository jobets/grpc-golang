// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello-service.proto

package api

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

//message Hello
type Hello struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5bbcd7ae0220f22, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Hello) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

//Hello Request
type HelloRequest struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5bbcd7ae0220f22, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

//Hello Request
type HelloRequestMultipleTimes struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequestMultipleTimes) Reset()         { *m = HelloRequestMultipleTimes{} }
func (m *HelloRequestMultipleTimes) String() string { return proto.CompactTextString(m) }
func (*HelloRequestMultipleTimes) ProtoMessage()    {}
func (*HelloRequestMultipleTimes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5bbcd7ae0220f22, []int{2}
}

func (m *HelloRequestMultipleTimes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequestMultipleTimes.Unmarshal(m, b)
}
func (m *HelloRequestMultipleTimes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequestMultipleTimes.Marshal(b, m, deterministic)
}
func (m *HelloRequestMultipleTimes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequestMultipleTimes.Merge(m, src)
}
func (m *HelloRequestMultipleTimes) XXX_Size() int {
	return xxx_messageInfo_HelloRequestMultipleTimes.Size(m)
}
func (m *HelloRequestMultipleTimes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequestMultipleTimes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequestMultipleTimes proto.InternalMessageInfo

func (m *HelloRequestMultipleTimes) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

//Hello Response
type HelloResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5bbcd7ae0220f22, []int{3}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

//Hello Response
type HelloResponseMultipleTimes struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponseMultipleTimes) Reset()         { *m = HelloResponseMultipleTimes{} }
func (m *HelloResponseMultipleTimes) String() string { return proto.CompactTextString(m) }
func (*HelloResponseMultipleTimes) ProtoMessage()    {}
func (*HelloResponseMultipleTimes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5bbcd7ae0220f22, []int{4}
}

func (m *HelloResponseMultipleTimes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponseMultipleTimes.Unmarshal(m, b)
}
func (m *HelloResponseMultipleTimes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponseMultipleTimes.Marshal(b, m, deterministic)
}
func (m *HelloResponseMultipleTimes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponseMultipleTimes.Merge(m, src)
}
func (m *HelloResponseMultipleTimes) XXX_Size() int {
	return xxx_messageInfo_HelloResponseMultipleTimes.Size(m)
}
func (m *HelloResponseMultipleTimes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponseMultipleTimes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponseMultipleTimes proto.InternalMessageInfo

func (m *HelloResponseMultipleTimes) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "api.Hello")
	proto.RegisterType((*HelloRequest)(nil), "api.HelloRequest")
	proto.RegisterType((*HelloRequestMultipleTimes)(nil), "api.HelloRequestMultipleTimes")
	proto.RegisterType((*HelloResponse)(nil), "api.HelloResponse")
	proto.RegisterType((*HelloResponseMultipleTimes)(nil), "api.HelloResponseMultipleTimes")
}

func init() { proto.RegisterFile("hello-service.proto", fileDescriptor_b5bbcd7ae0220f22) }

var fileDescriptor_b5bbcd7ae0220f22 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0xeb, 0xa2, 0x54, 0xe4, 0x83, 0x0e, 0x18, 0x89, 0x3f, 0x45, 0x94, 0xca, 0x0b, 0x59,
	0x88, 0xa2, 0xc2, 0xc0, 0xc2, 0xd2, 0x2e, 0x2c, 0x20, 0xd1, 0xc0, 0xc2, 0x82, 0x4c, 0xf9, 0x28,
	0x96, 0x1c, 0x27, 0xd8, 0x0e, 0x52, 0x9f, 0x94, 0xd7, 0x41, 0x71, 0x12, 0x91, 0xa8, 0x45, 0x74,
	0xcc, 0x9d, 0xef, 0x77, 0xe7, 0xc8, 0xb0, 0xff, 0x81, 0x52, 0xa6, 0x17, 0x06, 0xf5, 0x97, 0x98,
	0x63, 0x98, 0xe9, 0xd4, 0xa6, 0x74, 0x8b, 0x67, 0x82, 0x4d, 0xc1, 0xbb, 0x2d, 0x3c, 0x7a, 0x0a,
	0xf0, 0x2e, 0xb4, 0xb1, 0x2f, 0x8a, 0x27, 0x78, 0x44, 0x46, 0x24, 0xf0, 0x67, 0xbe, 0x53, 0xee,
	0x79, 0x82, 0xf4, 0x04, 0x7c, 0xc9, 0x6b, 0xb7, 0xeb, 0xdc, 0xed, 0x42, 0x28, 0x4c, 0x16, 0xc1,
	0xae, 0x83, 0xcc, 0xf0, 0x33, 0x47, 0x63, 0xe9, 0x08, 0x3c, 0x57, 0xe8, 0x30, 0x3b, 0x63, 0x08,
	0x79, 0x26, 0xc2, 0xf2, 0x44, 0x69, 0xb0, 0x1b, 0x38, 0x6e, 0x26, 0xee, 0x72, 0x69, 0x45, 0x26,
	0xf1, 0x51, 0x24, 0x68, 0x36, 0x88, 0x9f, 0x43, 0xbf, 0x8a, 0x9b, 0x2c, 0x55, 0x06, 0xe9, 0x01,
	0xf4, 0x34, 0x9a, 0x5c, 0xda, 0x6a, 0x79, 0xf5, 0xc5, 0xae, 0x60, 0xd0, 0x3a, 0xd8, 0x2e, 0xfa,
	0x23, 0x35, 0xfe, 0xee, 0x56, 0x17, 0x8a, 0xcb, 0x1f, 0x46, 0xaf, 0xa1, 0x1f, 0xf3, 0xa5, 0x93,
	0x9e, 0x14, 0xd7, 0x4b, 0xba, 0xd7, 0xd8, 0x54, 0x5e, 0x61, 0x40, 0x9b, 0x52, 0xd9, 0xc6, 0x3a,
	0xf4, 0x01, 0x0e, 0xeb, 0xe4, 0x54, 0x0a, 0x54, 0x36, 0xb6, 0x1a, 0x79, 0x22, 0xd4, 0x82, 0x0e,
	0x57, 0x18, 0xad, 0x75, 0xeb, 0x81, 0x01, 0x69, 0x22, 0x8b, 0x7d, 0xa8, 0x7f, 0x91, 0x6b, 0x66,
	0x9d, 0xad, 0x52, 0x5a, 0x35, 0xac, 0x13, 0x11, 0xba, 0x80, 0x61, 0x8d, 0x9c, 0x88, 0x37, 0xa1,
	0x71, 0x6e, 0x45, 0xaa, 0xb8, 0xdc, 0x7c, 0xec, 0xff, 0x35, 0x01, 0x89, 0xc8, 0xc4, 0x7b, 0x2e,
	0x5e, 0xdd, 0x6b, 0xcf, 0xbd, 0xc0, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x17, 0xfa,
	0xfe, 0x98, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	SayHelloUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	//Client Streaming RPC
	SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloClientStreamingClient, error)
	//Server Streaming RPC
	SayHelloServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayHelloServerStreamingClient, error)
	//Bidirectional Streaming RPC
	SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloBidirectionalStreamingClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHelloUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/api.HelloService/SayHelloUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/api.HelloService/SayHelloClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloClientStreamingClient{stream}
	return x, nil
}

type HelloService_SayHelloClientStreamingClient interface {
	Send(*HelloRequestMultipleTimes) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloClientStreamingClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloClientStreamingClient) Send(m *HelloRequestMultipleTimes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloClientStreamingClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayHelloServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[1], "/api.HelloService/SayHelloServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_SayHelloServerStreamingClient interface {
	Recv() (*HelloResponseMultipleTimes, error)
	grpc.ClientStream
}

type helloServiceSayHelloServerStreamingClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloServerStreamingClient) Recv() (*HelloResponseMultipleTimes, error) {
	m := new(HelloResponseMultipleTimes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[2], "/api.HelloService/SayHelloBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloBidirectionalStreamingClient{stream}
	return x, nil
}

type HelloService_SayHelloBidirectionalStreamingClient interface {
	Send(*HelloRequestMultipleTimes) error
	Recv() (*HelloResponseMultipleTimes, error)
	grpc.ClientStream
}

type helloServiceSayHelloBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloBidirectionalStreamingClient) Send(m *HelloRequestMultipleTimes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloBidirectionalStreamingClient) Recv() (*HelloResponseMultipleTimes, error) {
	m := new(HelloResponseMultipleTimes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Unary
	SayHelloUnary(context.Context, *HelloRequest) (*HelloResponse, error)
	//Client Streaming RPC
	SayHelloClientStreaming(HelloService_SayHelloClientStreamingServer) error
	//Server Streaming RPC
	SayHelloServerStreaming(*HelloRequest, HelloService_SayHelloServerStreamingServer) error
	//Bidirectional Streaming RPC
	SayHelloBidirectionalStreaming(HelloService_SayHelloBidirectionalStreamingServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHelloUnary(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloUnary not implemented")
}
func (*UnimplementedHelloServiceServer) SayHelloClientStreaming(srv HelloService_SayHelloClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStreaming not implemented")
}
func (*UnimplementedHelloServiceServer) SayHelloServerStreaming(req *HelloRequest, srv HelloService_SayHelloServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStreaming not implemented")
}
func (*UnimplementedHelloServiceServer) SayHelloBidirectionalStreaming(srv HelloService_SayHelloBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStreaming not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHelloUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHelloUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HelloService/SayHelloUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHelloUnary(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayHelloClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloClientStreaming(&helloServiceSayHelloClientStreamingServer{stream})
}

type HelloService_SayHelloClientStreamingServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequestMultipleTimes, error)
	grpc.ServerStream
}

type helloServiceSayHelloClientStreamingServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloClientStreamingServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloClientStreamingServer) Recv() (*HelloRequestMultipleTimes, error) {
	m := new(HelloRequestMultipleTimes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_SayHelloServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).SayHelloServerStreaming(m, &helloServiceSayHelloServerStreamingServer{stream})
}

type HelloService_SayHelloServerStreamingServer interface {
	Send(*HelloResponseMultipleTimes) error
	grpc.ServerStream
}

type helloServiceSayHelloServerStreamingServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloServerStreamingServer) Send(m *HelloResponseMultipleTimes) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_SayHelloBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloBidirectionalStreaming(&helloServiceSayHelloBidirectionalStreamingServer{stream})
}

type HelloService_SayHelloBidirectionalStreamingServer interface {
	Send(*HelloResponseMultipleTimes) error
	Recv() (*HelloRequestMultipleTimes, error)
	grpc.ServerStream
}

type helloServiceSayHelloBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloBidirectionalStreamingServer) Send(m *HelloResponseMultipleTimes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloBidirectionalStreamingServer) Recv() (*HelloRequestMultipleTimes, error) {
	m := new(HelloRequestMultipleTimes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloUnary",
			Handler:    _HelloService_SayHelloUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloClientStreaming",
			Handler:       _HelloService_SayHelloClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServerStreaming",
			Handler:       _HelloService_SayHelloServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStreaming",
			Handler:       _HelloService_SayHelloBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello-service.proto",
}
