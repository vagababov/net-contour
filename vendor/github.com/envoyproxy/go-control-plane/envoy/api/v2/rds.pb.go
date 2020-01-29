// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/rds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RdsDummy) Reset()         { *m = RdsDummy{} }
func (m *RdsDummy) String() string { return proto.CompactTextString(m) }
func (*RdsDummy) ProtoMessage()    {}
func (*RdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_78812f46dcff924a, []int{0}
}

func (m *RdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RdsDummy.Unmarshal(m, b)
}
func (m *RdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RdsDummy.Marshal(b, m, deterministic)
}
func (m *RdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RdsDummy.Merge(m, src)
}
func (m *RdsDummy) XXX_Size() int {
	return xxx_messageInfo_RdsDummy.Size(m)
}
func (m *RdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_RdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_RdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RdsDummy)(nil), "envoy.api.v2.RdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/rds.proto", fileDescriptor_78812f46dcff924a) }

var fileDescriptor_78812f46dcff924a = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4f, 0xdb, 0x40,
	0x18, 0xc6, 0xeb, 0x54, 0x4d, 0xab, 0x4b, 0x96, 0x5a, 0x6d, 0x5a, 0x39, 0x69, 0x92, 0xa6, 0x7f,
	0xd5, 0xc1, 0xae, 0x9c, 0x2d, 0x63, 0x1a, 0x55, 0x4c, 0x28, 0x4a, 0x04, 0x23, 0xd2, 0xc5, 0xbe,
	0x38, 0x27, 0xd9, 0xbe, 0xe3, 0xfe, 0x18, 0xbc, 0x32, 0x21, 0x16, 0x06, 0x24, 0xc4, 0x07, 0x60,
	0xe3, 0xe3, 0xf0, 0x15, 0x98, 0x19, 0x60, 0x47, 0xc8, 0x67, 0x1b, 0x72, 0x44, 0x30, 0x65, 0xb3,
	0xef, 0xf7, 0xbc, 0xcf, 0x3d, 0xef, 0x7b, 0x2f, 0x68, 0xa0, 0x38, 0x21, 0xa9, 0x03, 0x29, 0x76,
	0x12, 0xd7, 0x61, 0x3e, 0xb7, 0x29, 0x23, 0x82, 0x98, 0x75, 0x75, 0x6e, 0x43, 0x8a, 0xed, 0xc4,
	0xb5, 0x5a, 0x9a, 0xca, 0xc7, 0xdc, 0x23, 0x09, 0x62, 0x69, 0xae, 0xb5, 0x5a, 0x01, 0x21, 0x41,
	0x88, 0x14, 0x86, 0x71, 0x4c, 0x04, 0x14, 0x98, 0xc4, 0x85, 0x93, 0xd5, 0x2e, 0xa8, 0xfa, 0x9b,
	0xc9, 0xb9, 0xb3, 0xc7, 0x20, 0xa5, 0x88, 0x95, 0xbc, 0x5b, 0x78, 0x3f, 0x16, 0x3a, 0x0c, 0x71,
	0x22, 0x99, 0x87, 0x4a, 0x07, 0xe9, 0x53, 0xa8, 0x09, 0x22, 0x1c, 0x30, 0x28, 0x4a, 0xfe, 0x29,
	0x81, 0x21, 0xf6, 0xa1, 0x40, 0x4e, 0xf9, 0x51, 0x80, 0xcf, 0x7a, 0x73, 0x44, 0x96, 0xa4, 0x07,
	0xc0, 0xbb, 0x89, 0xcf, 0x47, 0x32, 0x8a, 0x52, 0xf7, 0xba, 0x02, 0x3e, 0x4e, 0x32, 0x36, 0x2a,
	0xfb, 0x9a, 0x22, 0x96, 0x60, 0x0f, 0x99, 0x5b, 0xa0, 0x3e, 0x15, 0x0c, 0xc1, 0x48, 0x61, 0x6e,
	0xb6, 0xed, 0xe5, 0xa9, 0xd8, 0x0f, 0xfa, 0x09, 0xda, 0x95, 0x88, 0x0b, 0xab, 0xf3, 0x2c, 0xe7,
	0x94, 0xc4, 0x1c, 0xf5, 0x5e, 0xfd, 0x36, 0xfe, 0x1a, 0xe6, 0x0e, 0xa8, 0x8d, 0x50, 0x28, 0x60,
	0xe1, 0xfa, 0xed, 0x49, 0x55, 0x86, 0x56, 0xac, 0xbf, 0xbf, 0x2c, 0xd2, 0xfc, 0x25, 0xa8, 0xfd,
	0x47, 0xc2, 0x5b, 0xac, 0x2b, 0xf5, 0x8f, 0x83, 0xcb, 0xab, 0x93, 0x4a, 0xa3, 0xf7, 0x41, 0x7b,
	0xfc, 0x81, 0x9a, 0x27, 0x57, 0xec, 0xf5, 0xc0, 0xf8, 0x63, 0xfd, 0x3a, 0x3a, 0x3f, 0xbd, 0x7d,
	0xfb, 0x15, 0x74, 0x34, 0x3b, 0x15, 0xe2, 0x1f, 0x89, 0xe7, 0x38, 0x90, 0x4c, 0x3d, 0x9f, 0x7b,
	0x61, 0x80, 0xe6, 0x36, 0x66, 0x42, 0xc2, 0x70, 0x83, 0x70, 0xb1, 0x32, 0x76, 0x1f, 0xbc, 0x57,
	0xfd, 0x2d, 0x69, 0xd6, 0x3f, 0x25, 0xeb, 0xa7, 0x8a, 0xdb, 0x05, 0x6d, 0xad, 0x24, 0x5f, 0x92,
	0xa5, 0x9b, 0x87, 0x9b, 0x37, 0x67, 0x77, 0xc7, 0x6f, 0xbe, 0x98, 0xcd, 0x5c, 0xc7, 0xf3, 0x90,
	0x85, 0x30, 0xe9, 0xc3, 0x90, 0x2e, 0x20, 0xb0, 0x30, 0xc9, 0xaf, 0xa6, 0x8c, 0xec, 0xa7, 0x5a,
	0x8a, 0x61, 0xb6, 0x69, 0xe3, 0x6c, 0xeb, 0xc6, 0xc6, 0xa1, 0x61, 0x8c, 0xab, 0xb3, 0xaa, 0xda,
	0xc1, 0xfe, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x4f, 0x54, 0x74, 0x7c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.RouteDiscoveryService/DeltaRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceDeltaRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_DeltaRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceDeltaRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	DeltaRoutes(RouteDiscoveryService_DeltaRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedRouteDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRouteDiscoveryServiceServer struct {
}

func (*UnimplementedRouteDiscoveryServiceServer) StreamRoutes(srv RouteDiscoveryService_StreamRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) DeltaRoutes(srv RouteDiscoveryService_DeltaRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) FetchRoutes(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRoutes not implemented")
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_DeltaRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).DeltaRoutes(&routeDiscoveryServiceDeltaRoutesServer{stream})
}

type RouteDiscoveryService_DeltaRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceDeltaRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRoutes",
			Handler:       _RouteDiscoveryService_DeltaRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}

// VirtualHostDiscoveryServiceClient is the client API for VirtualHostDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHostDiscoveryServiceClient interface {
	DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error)
}

type virtualHostDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHostDiscoveryServiceClient(cc *grpc.ClientConn) VirtualHostDiscoveryServiceClient {
	return &virtualHostDiscoveryServiceClient{cc}
}

func (c *virtualHostDiscoveryServiceClient) DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VirtualHostDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.VirtualHostDiscoveryService/DeltaVirtualHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualHostDiscoveryServiceDeltaVirtualHostsClient{stream}
	return x, nil
}

type VirtualHostDiscoveryService_DeltaVirtualHostsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsClient struct {
	grpc.ClientStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualHostDiscoveryServiceServer is the server API for VirtualHostDiscoveryService service.
type VirtualHostDiscoveryServiceServer interface {
	DeltaVirtualHosts(VirtualHostDiscoveryService_DeltaVirtualHostsServer) error
}

// UnimplementedVirtualHostDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHostDiscoveryServiceServer struct {
}

func (*UnimplementedVirtualHostDiscoveryServiceServer) DeltaVirtualHosts(srv VirtualHostDiscoveryService_DeltaVirtualHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaVirtualHosts not implemented")
}

func RegisterVirtualHostDiscoveryServiceServer(s *grpc.Server, srv VirtualHostDiscoveryServiceServer) {
	s.RegisterService(&_VirtualHostDiscoveryService_serviceDesc, srv)
}

func _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualHostDiscoveryServiceServer).DeltaVirtualHosts(&virtualHostDiscoveryServiceDeltaVirtualHostsServer{stream})
}

type VirtualHostDiscoveryService_DeltaVirtualHostsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsServer struct {
	grpc.ServerStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VirtualHostDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.VirtualHostDiscoveryService",
	HandlerType: (*VirtualHostDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaVirtualHosts",
			Handler:       _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}