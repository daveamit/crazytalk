// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	P
	C
	SayHiRequest
	SayHiResponse
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

type P struct {
	CFromP *C `protobuf:"bytes,1,opt,name=cFromP" json:"cFromP,omitempty"`
}

func (m *P) Reset()                    { *m = P{} }
func (m *P) String() string            { return proto.CompactTextString(m) }
func (*P) ProtoMessage()               {}
func (*P) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *P) GetCFromP() *C {
	if m != nil {
		return m.CFromP
	}
	return nil
}

type C struct {
	PFromC *P `protobuf:"bytes,1,opt,name=pFromC" json:"pFromC,omitempty"`
}

func (m *C) Reset()                    { *m = C{} }
func (m *C) String() string            { return proto.CompactTextString(m) }
func (*C) ProtoMessage()               {}
func (*C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *C) GetPFromC() *P {
	if m != nil {
		return m.PFromC
	}
	return nil
}

// SayHiRequest takes in a name
type SayHiRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	P    *P     `protobuf:"bytes,2,opt,name=p" json:"p,omitempty"`
}

func (m *SayHiRequest) Reset()                    { *m = SayHiRequest{} }
func (m *SayHiRequest) String() string            { return proto.CompactTextString(m) }
func (*SayHiRequest) ProtoMessage()               {}
func (*SayHiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SayHiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SayHiRequest) GetP() *P {
	if m != nil {
		return m.P
	}
	return nil
}

// SayHiResponse takes in a name
type SayHiResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SayHiResponse) Reset()                    { *m = SayHiResponse{} }
func (m *SayHiResponse) String() string            { return proto.CompactTextString(m) }
func (*SayHiResponse) ProtoMessage()               {}
func (*SayHiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SayHiResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*P)(nil), "test.P")
	proto.RegisterType((*C)(nil), "test.C")
	proto.RegisterType((*SayHiRequest)(nil), "test.SayHiRequest")
	proto.RegisterType((*SayHiResponse)(nil), "test.SayHiResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	// SayHi takes in a name and returns hello message
	SayHi(ctx context.Context, in *SayHiRequest, opts ...grpc.CallOption) (*SayHiResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHi(ctx context.Context, in *SayHiRequest, opts ...grpc.CallOption) (*SayHiResponse, error) {
	out := new(SayHiResponse)
	err := grpc.Invoke(ctx, "/test.Hello/SayHi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	// SayHi takes in a name and returns hello message
	SayHi(context.Context, *SayHiRequest) (*SayHiResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Hello/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHi(ctx, req.(*SayHiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Hello_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x86, 0x8d, 0x7c, 0x6d, 0xf1, 0xd4, 0xe5, 0x44, 0x29, 0x2e, 0x6a, 0x70, 0xd0, 0x25, 0x43,
	0x9d, 0xc4, 0xad, 0x05, 0xe9, 0x18, 0xe2, 0x2f, 0x88, 0xf5, 0xa8, 0x85, 0xa6, 0x89, 0x4d, 0x14,
	0xfc, 0xf7, 0x92, 0xb4, 0x05, 0x3b, 0xe5, 0xee, 0xe5, 0x79, 0xde, 0xc0, 0x01, 0x04, 0xf2, 0x41,
	0xb8, 0xd9, 0x06, 0x8b, 0x87, 0x38, 0xf3, 0x7b, 0x60, 0x12, 0x6f, 0x20, 0xef, 0x5e, 0x67, 0x6b,
	0x64, 0xc9, 0x6e, 0xd9, 0xc3, 0x69, 0x55, 0x88, 0xc4, 0x35, 0x6a, 0x8d, 0x23, 0xd5, 0x44, 0xca,
	0xc5, 0xb5, 0xd9, 0x53, 0x52, 0xad, 0x31, 0x7f, 0x86, 0xb3, 0x37, 0xfd, 0xdb, 0x0e, 0x8a, 0xbe,
	0xbe, 0xc9, 0x07, 0x44, 0x38, 0x4c, 0xda, 0x50, 0xc2, 0x4f, 0x54, 0x9a, 0xf1, 0x12, 0x98, 0x2b,
	0x8f, 0xf7, 0x3e, 0x73, 0xfc, 0x11, 0xce, 0x57, 0xd5, 0x3b, 0x3b, 0x79, 0xc2, 0x12, 0x0a, 0x43,
	0xde, 0xeb, 0x7e, 0xd3, 0xb7, 0xb5, 0x7a, 0x81, 0xac, 0xa5, 0x71, 0xb4, 0x58, 0x41, 0x96, 0x1c,
	0xc4, 0xa5, 0xe8, 0xff, 0xdf, 0xd7, 0x17, 0xbb, 0x6c, 0x29, 0xe5, 0x47, 0xf5, 0x1d, 0x5c, 0x0d,
	0x56, 0xf4, 0xb3, 0xeb, 0xc4, 0x87, 0xfe, 0x21, 0x6d, 0x86, 0x20, 0x3e, 0x63, 0x5b, 0x9d, 0xa5,
	0x47, 0xb2, 0xf7, 0x3c, 0x9d, 0xe7, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x77, 0xcb, 0x1f, 0x83,
	0x2c, 0x01, 0x00, 0x00,
}
