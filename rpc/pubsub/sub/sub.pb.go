// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sub.proto

package sub

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

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_168c8107357a6d2a, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "sub.String")
}

func init() { proto.RegisterFile("sub.proto", fileDescriptor_168c8107357a6d2a) }

var fileDescriptor_168c8107357a6d2a = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0x4d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2e, 0x4d, 0x52, 0x92, 0xe3, 0x62, 0x0b, 0x2e,
	0x29, 0xca, 0xcc, 0x4b, 0x17, 0x12, 0xe1, 0x62, 0x0d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c, 0x62, 0xb8, 0x78, 0x03, 0x4a, 0x93, 0x8a, 0x4b, 0x93,
	0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x94, 0xb9, 0xd8, 0x03, 0x4a, 0x93, 0x72, 0x32,
	0x8b, 0x33, 0x84, 0xb8, 0xf5, 0x40, 0x86, 0x41, 0xb4, 0x4b, 0x21, 0x73, 0x84, 0xd4, 0xb8, 0x38,
	0x82, 0x4b, 0x93, 0x8a, 0x8b, 0x32, 0x93, 0x52, 0x71, 0xab, 0x32, 0x60, 0x4c, 0x62, 0x03, 0xbb,
	0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xb4, 0xe8, 0x3f, 0x96, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	Subsribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubsribeClient, error)
}

type pubsubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServiceClient(cc *grpc.ClientConn) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/sub.PubsubService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServiceClient) Subsribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubsribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubService_serviceDesc.Streams[0], "/sub.PubsubService/Subsribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubsribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubsribeClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubsubServiceSubsribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubsribeClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServiceServer is the server API for PubsubService service.
type PubsubServiceServer interface {
	Publish(context.Context, *String) (*String, error)
	Subsribe(*String, PubsubService_SubsribeServer) error
}

// UnimplementedPubsubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubsubServiceServer struct {
}

func (*UnimplementedPubsubServiceServer) Publish(ctx context.Context, req *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedPubsubServiceServer) Subsribe(req *String, srv PubsubService_SubsribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subsribe not implemented")
}

func RegisterPubsubServiceServer(s *grpc.Server, srv PubsubServiceServer) {
	s.RegisterService(&_PubsubService_serviceDesc, srv)
}

func _PubsubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sub.PubsubService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubService_Subsribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subsribe(m, &pubsubServiceSubsribeServer{stream})
}

type PubsubService_SubsribeServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubsubServiceSubsribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubsribeServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sub.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubsubService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subsribe",
			Handler:       _PubsubService_Subsribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sub.proto",
}