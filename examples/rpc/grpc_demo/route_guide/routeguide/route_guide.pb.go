// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route_guide.proto

package routeguide

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

type Point struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_751d5f2517576e3b, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Rectangle struct {
	Low                  *Point   `protobuf:"bytes,1,opt,name=low,proto3" json:"low,omitempty"`
	High                 *Point   `protobuf:"bytes,2,opt,name=high,proto3" json:"high,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_751d5f2517576e3b, []int{1}
}
func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (dst *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(dst, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLow() *Point {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *Rectangle) GetHigh() *Point {
	if m != nil {
		return m.High
	}
	return nil
}

type Feature struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             *Point   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_751d5f2517576e3b, []int{2}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type RouteNote struct {
	Location             *Point   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_751d5f2517576e3b, []int{3}
}
func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (dst *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(dst, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RouteSummary struct {
	PointCount           int32    `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	FeatureCount         int32    `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	Distance             int32    `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	ElapsedTime          int32    `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSummary) Reset()         { *m = RouteSummary{} }
func (m *RouteSummary) String() string { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()    {}
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_751d5f2517576e3b, []int{4}
}
func (m *RouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSummary.Unmarshal(m, b)
}
func (m *RouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSummary.Marshal(b, m, deterministic)
}
func (dst *RouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSummary.Merge(dst, src)
}
func (m *RouteSummary) XXX_Size() int {
	return xxx_messageInfo_RouteSummary.Size(m)
}
func (m *RouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSummary proto.InternalMessageInfo

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "routeguide.Feature")
	proto.RegisterType((*RouteNote)(nil), "routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routeguide.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	ListFeature(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeatureClient, error)
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeature(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[0], "/routeguide.RouteGuide/ListFeature", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeatureClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeatureClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeatureClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[1], "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[2], "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	GetFeature(context.Context, *Point) (*Feature, error)
	ListFeature(*Rectangle, RouteGuide_ListFeatureServer) error
	RecordRoute(RouteGuide_RecordRouteServer) error
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeature(m, &routeGuideListFeatureServer{stream})
}

type RouteGuide_ListFeatureServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeatureServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeatureServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeature",
			Handler:       _RouteGuide_ListFeature_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route_guide.proto",
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor_route_guide_751d5f2517576e3b) }

var fileDescriptor_route_guide_751d5f2517576e3b = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x8b, 0xd3, 0x40,
	0x14, 0xdf, 0xd9, 0xed, 0xba, 0xcd, 0x4b, 0x45, 0xf6, 0x89, 0x10, 0x82, 0xa0, 0x66, 0x11, 0x7a,
	0x31, 0x2c, 0x2b, 0x78, 0x5b, 0xc5, 0x16, 0xec, 0xa5, 0x48, 0x89, 0x05, 0x8f, 0x61, 0x4c, 0x9e,
	0xe9, 0x40, 0x66, 0x26, 0x24, 0x13, 0xd4, 0xbb, 0xff, 0x81, 0xff, 0xb0, 0x64, 0xf2, 0xd1, 0x54,
	0x5b, 0xf6, 0x36, 0xef, 0xf7, 0xf1, 0x3e, 0x19, 0xb8, 0x2e, 0x75, 0x6d, 0x28, 0xce, 0x6a, 0x91,
	0x52, 0x58, 0x94, 0xda, 0x68, 0x04, 0x0b, 0x59, 0x24, 0xf8, 0x08, 0x97, 0x1b, 0x2d, 0x94, 0x41,
	0x1f, 0xa6, 0x39, 0x37, 0xc2, 0xd4, 0x29, 0x79, 0xec, 0x25, 0x9b, 0x5f, 0x46, 0x43, 0x8c, 0xcf,
	0xc1, 0xc9, 0xb5, 0xca, 0x5a, 0xf2, 0xdc, 0x92, 0x7b, 0x20, 0xf8, 0x0a, 0x4e, 0x44, 0x89, 0xe1,
	0x2a, 0xcb, 0x09, 0x6f, 0xe0, 0x22, 0xd7, 0x3f, 0x6c, 0x06, 0xf7, 0xee, 0x3a, 0xdc, 0x57, 0x0a,
	0x6d, 0x99, 0xa8, 0x61, 0xf1, 0x35, 0x4c, 0x76, 0x22, 0xdb, 0xd9, 0x54, 0x47, 0x55, 0x96, 0x0e,
	0xd6, 0x70, 0xf5, 0x89, 0xb8, 0xa9, 0x4b, 0x42, 0x84, 0x89, 0xe2, 0xb2, 0xed, 0xcc, 0x89, 0xec,
	0x1b, 0xdf, 0xc0, 0x34, 0xd7, 0x09, 0x37, 0x42, 0xab, 0xd3, 0x99, 0x06, 0x49, 0xb0, 0x05, 0x27,
	0x6a, 0xd8, 0xcf, 0xda, 0x1c, 0x7a, 0xd9, 0x83, 0x5e, 0xf4, 0xe0, 0x4a, 0x52, 0x55, 0xf1, 0xac,
	0x1d, 0xdf, 0x89, 0xfa, 0x30, 0xf8, 0xc3, 0x60, 0x66, 0xd3, 0x7e, 0xa9, 0xa5, 0xe4, 0xe5, 0x2f,
	0x7c, 0x01, 0x6e, 0xd1, 0xb8, 0xe3, 0x44, 0xd7, 0xca, 0x74, 0xab, 0x04, 0x0b, 0x2d, 0x1b, 0x04,
	0x6f, 0xe0, 0xf1, 0xf7, 0x76, 0xaa, 0x4e, 0xd2, 0x2e, 0x74, 0xd6, 0x81, 0xad, 0xc8, 0x87, 0x69,
	0x2a, 0x2a, 0xc3, 0x55, 0x42, 0xde, 0x45, 0x7b, 0x8d, 0x3e, 0xc6, 0x57, 0x30, 0xa3, 0x9c, 0x17,
	0x15, 0xa5, 0xb1, 0x11, 0x92, 0xbc, 0x89, 0xe5, 0xdd, 0x0e, 0xdb, 0x0a, 0x49, 0x77, 0xbf, 0xcf,
	0x01, 0x6c, 0x57, 0xab, 0x66, 0x1c, 0x7c, 0x07, 0xb0, 0x22, 0xd3, 0xef, 0xf2, 0xff, 0x49, 0xfd,
	0xa7, 0x63, 0xa8, 0xd3, 0x05, 0x67, 0x78, 0x0f, 0xee, 0x5a, 0x54, 0x83, 0xf1, 0xd9, 0x58, 0x35,
	0x9c, 0xfc, 0x84, 0xf9, 0x96, 0xe1, 0x7b, 0x70, 0x23, 0x4a, 0x74, 0x99, 0xda, 0x56, 0x8e, 0xd5,
	0xf5, 0x0e, 0x32, 0x8e, 0xd6, 0x18, 0x9c, 0xcd, 0x19, 0x7e, 0xe8, 0x2e, 0xb6, 0xdc, 0x71, 0xf3,
	0x4f, 0xf1, 0xfe, 0x90, 0xfe, 0x71, 0xb8, 0xb1, 0xdf, 0xb2, 0xc5, 0x3d, 0xcc, 0xb3, 0x92, 0xab,
	0x54, 0x2b, 0xa1, 0x55, 0x48, 0x3f, 0xb9, 0x2c, 0x72, 0xaa, 0xc2, 0xac, 0x2c, 0x92, 0x38, 0x25,
	0xa9, 0x47, 0xde, 0xc5, 0x93, 0xfd, 0xbe, 0x36, 0xcd, 0x2f, 0xd9, 0xb0, 0x6f, 0x8f, 0xec, 0x77,
	0x79, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x37, 0x66, 0xc9, 0x43, 0x03, 0x00, 0x00,
}