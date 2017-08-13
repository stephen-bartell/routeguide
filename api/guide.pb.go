// Code generated by protoc-gen-go. DO NOT EDIT.
// source: guide.proto

/*
Package guide is a generated protocol buffer package.

It is generated from these files:
	guide.proto

It has these top-level messages:
	Point
	Rectangle
	Feature
	RouteNote
	RouteSummary
*/
package guide

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

// Point are lat-long pairs in E7 representation.
type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=Latitude" json:"Latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// The lat-long of a rectangle. lo and hi are diagonally opposite points.
type Rectangle struct {
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a point.
type Feature struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Message  *Point `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RouteNote) Reset()                    { *m = RouteNote{} }
func (m *RouteNote) String() string            { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()               {}
func (*RouteNote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() *Point {
	if m != nil {
		return m.Message
	}
	return nil
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	PointCount   int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	Distance     int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	ElapsedTime  int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
}

func (m *RouteSummary) Reset()                    { *m = RouteSummary{} }
func (m *RouteSummary) String() string            { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()               {}
func (*RouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*Point)(nil), "guide.Point")
	proto.RegisterType((*Rectangle)(nil), "guide.Rectangle")
	proto.RegisterType((*Feature)(nil), "guide.Feature")
	proto.RegisterType((*RouteNote)(nil), "guide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "guide.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Guide service

type GuideClient interface {
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (Guide_ListFeaturesClient, error)
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (Guide_RecordRouteClient, error)
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (Guide_RouteChatClient, error)
}

type guideClient struct {
	cc *grpc.ClientConn
}

func NewGuideClient(cc *grpc.ClientConn) GuideClient {
	return &guideClient{cc}
}

func (c *guideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := grpc.Invoke(ctx, "/guide.Guide/GetFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (Guide_ListFeaturesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Guide_serviceDesc.Streams[0], c.cc, "/guide.Guide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Guide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type guideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *guideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (Guide_RecordRouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Guide_serviceDesc.Streams[1], c.cc, "/guide.Guide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideRecordRouteClient{stream}
	return x, nil
}

type Guide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type guideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *guideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (Guide_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Guide_serviceDesc.Streams[2], c.cc, "/guide.Guide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideRouteChatClient{stream}
	return x, nil
}

type Guide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type guideRouteChatClient struct {
	grpc.ClientStream
}

func (x *guideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Guide service

type GuideServer interface {
	GetFeature(context.Context, *Point) (*Feature, error)
	ListFeatures(*Rectangle, Guide_ListFeaturesServer) error
	RecordRoute(Guide_RecordRouteServer) error
	RouteChat(Guide_RouteChatServer) error
}

func RegisterGuideServer(s *grpc.Server, srv GuideServer) {
	s.RegisterService(&_Guide_serviceDesc, srv)
}

func _Guide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guide.Guide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GuideServer).ListFeatures(m, &guideListFeaturesServer{stream})
}

type Guide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type guideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *guideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _Guide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuideServer).RecordRoute(&guideRecordRouteServer{stream})
}

type Guide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type guideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *guideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Guide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuideServer).RouteChat(&guideRouteChatServer{stream})
}

type Guide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type guideRouteChatServer struct {
	grpc.ServerStream
}

func (x *guideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Guide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guide.Guide",
	HandlerType: (*GuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _Guide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _Guide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _Guide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _Guide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "guide.proto",
}

func init() { proto.RegisterFile("guide.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xed, 0xf6, 0x6b, 0xbe, 0x36, 0x93, 0x28, 0x32, 0x5e, 0x4a, 0x10, 0xd4, 0x08, 0xd2, 0x83,
	0x94, 0x52, 0xeb, 0x0f, 0x90, 0x82, 0xb9, 0x14, 0x91, 0xe8, 0x55, 0xca, 0x9a, 0x8c, 0xe9, 0x42,
	0x92, 0x2d, 0xc9, 0xe6, 0xe0, 0xef, 0xf0, 0xcf, 0xf9, 0x73, 0x24, 0x9b, 0x4d, 0x6b, 0x6d, 0xf1,
	0x96, 0x79, 0xef, 0xcd, 0x9b, 0x9d, 0x37, 0x01, 0x27, 0xa9, 0x44, 0x4c, 0xe3, 0x75, 0x21, 0x95,
	0x44, 0x4b, 0x17, 0xfe, 0x3d, 0x58, 0x4f, 0x52, 0xe4, 0x0a, 0x3d, 0x18, 0x2c, 0xb8, 0x12, 0xaa,
	0x8a, 0x69, 0xc8, 0x2e, 0xd8, 0xc8, 0x0a, 0x37, 0x35, 0x9e, 0x81, 0x9d, 0xca, 0x3c, 0x69, 0xc8,
	0xae, 0x26, 0xb7, 0x80, 0x1f, 0x80, 0x1d, 0x52, 0xa4, 0x78, 0x9e, 0xa4, 0xb5, 0xb4, 0x9b, 0x4a,
	0x6d, 0xe0, 0x4c, 0xdd, 0x71, 0x33, 0x50, 0x0f, 0x08, 0xbb, 0xa9, 0xac, 0xd9, 0x95, 0xd0, 0x0e,
	0x7b, 0xec, 0x4a, 0xf8, 0x01, 0xf4, 0x1f, 0x88, 0xab, 0xaa, 0x20, 0x44, 0xe8, 0xe5, 0x3c, 0x6b,
	0x5e, 0x62, 0x87, 0xfa, 0x1b, 0x47, 0x30, 0x48, 0x65, 0xc4, 0x95, 0x90, 0xf9, 0x41, 0x8b, 0x0d,
	0xeb, 0xbf, 0x82, 0x1d, 0xca, 0x4a, 0xd1, 0xa3, 0x54, 0xbb, 0x6d, 0xec, 0xaf, 0x36, 0xbc, 0x86,
	0x7e, 0x46, 0x65, 0xc9, 0x13, 0x3a, 0xe8, 0xdf, 0x92, 0xfe, 0x27, 0x03, 0x57, 0xfb, 0x3f, 0x57,
	0x59, 0xc6, 0x8b, 0x0f, 0x3c, 0x07, 0x67, 0x5d, 0x4b, 0x96, 0x91, 0xac, 0x72, 0x65, 0xe2, 0x03,
	0x0d, 0xcd, 0x6b, 0x04, 0xaf, 0xe0, 0xe8, 0xbd, 0xd9, 0xcc, 0x48, 0x9a, 0x10, 0x5d, 0x03, 0x36,
	0x22, 0x0f, 0x06, 0xb1, 0x28, 0x15, 0xcf, 0x23, 0x1a, 0xfe, 0x6b, 0x2e, 0xd0, 0xd6, 0x78, 0x09,
	0x2e, 0xa5, 0x7c, 0x5d, 0x52, 0xbc, 0x54, 0x22, 0xa3, 0x61, 0x4f, 0xf3, 0x8e, 0xc1, 0x5e, 0x44,
	0x46, 0xd3, 0x2f, 0x06, 0x56, 0x50, 0x3f, 0x17, 0x6f, 0x00, 0x02, 0x52, 0x6d, 0x94, 0x3b, 0x4b,
	0x78, 0xc7, 0xa6, 0x32, 0xac, 0xdf, 0xc1, 0x19, 0xb8, 0x0b, 0x51, 0xb6, 0xf2, 0x12, 0x4f, 0x8c,
	0x62, 0x73, 0xd3, 0xfd, 0x9e, 0x09, 0xc3, 0x19, 0x38, 0x21, 0x45, 0xb2, 0x88, 0x75, 0x10, 0xbf,
	0x86, 0x9c, 0xb6, 0x16, 0x3f, 0x42, 0xf2, 0x3b, 0x23, 0x86, 0x77, 0xe6, 0x30, 0xf3, 0x15, 0x57,
	0xdb, 0x41, 0xed, 0xa9, 0xbc, 0x3d, 0xa4, 0x6e, 0x9a, 0xb0, 0xb7, 0xff, 0xfa, 0x97, 0xbd, 0xfd,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x1b, 0x49, 0x7c, 0xc1, 0x02, 0x00, 0x00,
}