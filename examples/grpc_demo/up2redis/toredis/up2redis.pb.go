// Code generated by protoc-gen-go. DO NOT EDIT.
// source: up2redis.proto

package redis

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

// redis上报数据格式
type ReportRequest struct {
	Delay                float32  `protobuf:"fixed32,2,opt,name=delay,proto3" json:"delay,omitempty"`
	Timestamp            uint32   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Loss                 float32  `protobuf:"fixed32,4,opt,name=loss,proto3" json:"loss,omitempty"`
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ExpireTime           uint32   `protobuf:"varint,6,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_up2redis_45184d504139c9ab, []int{0}
}
func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (dst *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(dst, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetDelay() float32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *ReportRequest) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ReportRequest) GetLoss() float32 {
	if m != nil {
		return m.Loss
	}
	return 0
}

func (m *ReportRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReportRequest) GetExpireTime() uint32 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

// 返回数据格式
type ReportResponse struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_up2redis_45184d504139c9ab, []int{1}
}
func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (dst *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(dst, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func (m *ReportResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*ReportRequest)(nil), "redis.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "redis.ReportResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RedisClient is the client API for Redis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RedisClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type redisClient struct {
	cc *grpc.ClientConn
}

func NewRedisClient(cc *grpc.ClientConn) RedisClient {
	return &redisClient{cc}
}

func (c *redisClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/redis.Redis/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisServer is the server API for Redis service.
type RedisServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterRedisServer(s *grpc.Server, srv RedisServer) {
	s.RegisterService(&_Redis_serviceDesc, srv)
}

func _Redis_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/redis.Redis/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Redis_serviceDesc = grpc.ServiceDesc{
	ServiceName: "redis.Redis",
	HandlerType: (*RedisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _Redis_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "up2redis.proto",
}

func init() { proto.RegisterFile("up2redis.proto", fileDescriptor_up2redis_45184d504139c9ab) }

var fileDescriptor_up2redis_45184d504139c9ab = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x14, 0x84, 0x4d, 0xff, 0xa0, 0x47, 0x5a, 0xe4, 0x50, 0x21, 0x88, 0x60, 0x29, 0x2e, 0xba, 0xea,
	0xa2, 0x2e, 0xdc, 0xfa, 0x0c, 0xc1, 0xbd, 0x54, 0x7a, 0x84, 0x60, 0xdb, 0xc4, 0x24, 0x05, 0xfb,
	0x0a, 0x3e, 0xb5, 0x24, 0x51, 0x2e, 0xf7, 0xee, 0xe6, 0x0c, 0x99, 0x8f, 0x99, 0x40, 0xbd, 0xeb,
	0xd1, 0xd0, 0x2c, 0xed, 0xa0, 0x8d, 0x72, 0x0a, 0xf3, 0x70, 0x74, 0x3f, 0x0c, 0x2a, 0x41, 0x5a,
	0x19, 0x27, 0xe8, 0x6b, 0x27, 0xeb, 0xb0, 0x81, 0x7c, 0xa6, 0x65, 0x3a, 0x78, 0xd2, 0xb2, 0x3e,
	0x11, 0xf1, 0xc0, 0x7b, 0x28, 0x9d, 0x5c, 0xc9, 0xba, 0x69, 0xd5, 0x3c, 0x6d, 0x59, 0x5f, 0x89,
	0x93, 0x81, 0x08, 0xd9, 0xa2, 0xac, 0xe5, 0x59, 0x88, 0x04, 0x8d, 0x37, 0x90, 0x7e, 0xd2, 0xc1,
	0x59, 0xcb, 0xfa, 0x52, 0x78, 0x89, 0x0f, 0x70, 0x4d, 0xdf, 0x5a, 0x1a, 0x7a, 0xf3, 0x49, 0x5e,
	0x04, 0x0a, 0x44, 0xeb, 0x55, 0xae, 0xd4, 0x3d, 0x42, 0xfd, 0xdf, 0xc5, 0x6a, 0xb5, 0x59, 0xf2,
	0x60, 0xb9, 0x7d, 0xa8, 0x3f, 0x4a, 0xd0, 0xe3, 0x0b, 0xe4, 0xc2, 0x77, 0xc7, 0x67, 0x28, 0xe2,
	0x73, 0x6c, 0x86, 0x38, 0xed, 0x6c, 0xc9, 0xdd, 0xed, 0x85, 0x1b, 0x99, 0xdd, 0xd5, 0x7b, 0x11,
	0xbe, 0xe0, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x04, 0xb9, 0x28, 0x6f, 0x14, 0x01, 0x00, 0x00,
}
