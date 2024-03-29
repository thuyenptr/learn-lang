// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

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

type Ping struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Pong struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MaxRequest struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxRequest) Reset()         { *m = MaxRequest{} }
func (m *MaxRequest) String() string { return proto.CompactTextString(m) }
func (*MaxRequest) ProtoMessage()    {}
func (*MaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{2}
}

func (m *MaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxRequest.Unmarshal(m, b)
}
func (m *MaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxRequest.Marshal(b, m, deterministic)
}
func (m *MaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxRequest.Merge(m, src)
}
func (m *MaxRequest) XXX_Size() int {
	return xxx_messageInfo_MaxRequest.Size(m)
}
func (m *MaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MaxRequest proto.InternalMessageInfo

func (m *MaxRequest) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type MaxResponse struct {
	Max                  int64    `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxResponse) Reset()         { *m = MaxResponse{} }
func (m *MaxResponse) String() string { return proto.CompactTextString(m) }
func (*MaxResponse) ProtoMessage()    {}
func (*MaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{3}
}

func (m *MaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxResponse.Unmarshal(m, b)
}
func (m *MaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxResponse.Marshal(b, m, deterministic)
}
func (m *MaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxResponse.Merge(m, src)
}
func (m *MaxResponse) XXX_Size() int {
	return xxx_messageInfo_MaxResponse.Size(m)
}
func (m *MaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaxResponse proto.InternalMessageInfo

func (m *MaxResponse) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func init() {
	proto.RegisterType((*Ping)(nil), "api.Ping")
	proto.RegisterType((*Pong)(nil), "api.Pong")
	proto.RegisterType((*MaxRequest)(nil), "api.MaxRequest")
	proto.RegisterType((*MaxResponse)(nil), "api.MaxResponse")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_6d51d96c3ad891f5) }

var fileDescriptor_6d51d96c3ad891f5 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xcd, 0x8a, 0xc2, 0x30,
	0x14, 0x85, 0x1b, 0x52, 0x86, 0xe9, 0x9d, 0xc5, 0x94, 0xac, 0x86, 0x32, 0x54, 0x09, 0x2e, 0x8a,
	0x8b, 0x20, 0x15, 0x7c, 0x8b, 0x82, 0x14, 0x5f, 0xe0, 0x0a, 0xa1, 0x66, 0x91, 0x1f, 0x4d, 0x0a,
	0x79, 0x7c, 0x49, 0xac, 0x74, 0xe9, 0xee, 0x72, 0xce, 0x07, 0xdf, 0xb9, 0x00, 0x4e, 0x99, 0x49,
	0xb8, 0x87, 0x0d, 0x96, 0x51, 0x74, 0x8a, 0xef, 0xa0, 0x3c, 0x2b, 0x33, 0xb1, 0x7f, 0xa8, 0x2e,
	0x4a, 0x4b, 0x1f, 0x50, 0xbb, 0x3f, 0xb2, 0x25, 0x1d, 0x1d, 0xd7, 0x20, 0x53, 0xf6, 0x23, 0xd5,
	0x02, 0x0c, 0x18, 0x47, 0x79, 0x9f, 0xa5, 0x0f, 0xac, 0x06, 0x6a, 0x66, 0xbd, 0x50, 0xe9, 0xe4,
	0x1b, 0xf8, 0xc9, 0xbd, 0x77, 0xd6, 0x78, 0x99, 0x00, 0x8d, 0xf1, 0x0d, 0x68, 0x8c, 0xfd, 0x1e,
	0xbe, 0xd3, 0x98, 0xac, 0x6a, 0xa1, 0x4c, 0x5b, 0x59, 0x25, 0xd0, 0x29, 0x91, 0xe2, 0x66, 0x39,
	0xad, 0x99, 0x78, 0xd1, 0x9f, 0xa0, 0x1c, 0x30, 0xdc, 0x98, 0x00, 0x3a, 0x60, 0x64, 0xbf, 0xb9,
	0x5b, 0xf5, 0x4d, 0xbd, 0x06, 0x2f, 0x1f, 0x2f, 0x3a, 0x72, 0x20, 0xd7, 0xaf, 0xfc, 0xfc, 0xf1,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xd0, 0x6c, 0xaf, 0x0a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/api.PingPong/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	Ping(context.Context, *Ping) (*Pong, error)
}

// UnimplementedPingPongServer can be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (*UnimplementedPingPongServer) Ping(ctx context.Context, req *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PingPong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Ping(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _PingPong_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}

// MathClient is the client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathClient interface {
	Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error)
}

type mathClient struct {
	cc *grpc.ClientConn
}

func NewMathClient(cc *grpc.ClientConn) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Math_serviceDesc.Streams[0], "/api.Math/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathMaxClient{stream}
	return x, nil
}

type Math_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type mathMaxClient struct {
	grpc.ClientStream
}

func (x *mathMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mathMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MathServer is the server API for Math service.
type MathServer interface {
	Max(Math_MaxServer) error
}

// UnimplementedMathServer can be embedded to have forward compatible implementations.
type UnimplementedMathServer struct {
}

func (*UnimplementedMathServer) Max(srv Math_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MathServer).Max(&mathMaxServer{stream})
}

type Math_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type mathMaxServer struct {
	grpc.ServerStream
}

func (x *mathMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mathMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Math",
	HandlerType: (*MathServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Max",
			Handler:       _Math_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ping.proto",
}
