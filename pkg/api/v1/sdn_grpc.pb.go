// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SdnServiceClient is the client API for SdnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdnServiceClient interface {
	// StartLocalNetwork starts a Network on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the sdn/config.yaml
	//   3. all bytes constituting the Network YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalNetwork(ctx context.Context, opts ...grpc.CallOption) (SdnService_StartLocalNetworkClient, error)
	// StartFromPreviousNetwork starts a new Network based on a previous one.
	// If the previous Network does not have the can-replay condition set this call will result in an error.
	StartFromPreviousNetwork(ctx context.Context, in *StartFromPreviousNetworkRequest, opts ...grpc.CallOption) (*StartNetworkResponse, error)
	// StartNetworkRequest starts a new Network based on its specification.
	StartNetwork(ctx context.Context, in *StartNetworkRequest, opts ...grpc.CallOption) (*StartNetworkResponse, error)
	// Searches for Network(s) known to this instance
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	// Subscribe listens to new Network(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SdnService_SubscribeClient, error)
	// GetNetwork retrieves details of a single Network
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error)
	// Listen listens to Network updates and log output of a running Network
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (SdnService_ListenClient, error)
	// StopNetwork stops a currently running Network
	StopNetwork(ctx context.Context, in *StopNetworkRequest, opts ...grpc.CallOption) (*StopNetworkResponse, error)
}

type sdnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdnServiceClient(cc grpc.ClientConnInterface) SdnServiceClient {
	return &sdnServiceClient{cc}
}

func (c *sdnServiceClient) StartLocalNetwork(ctx context.Context, opts ...grpc.CallOption) (SdnService_StartLocalNetworkClient, error) {
	stream, err := c.cc.NewStream(ctx, &SdnService_ServiceDesc.Streams[0], "/v1.SdnService/StartLocalNetwork", opts...)
	if err != nil {
		return nil, err
	}
	x := &sdnServiceStartLocalNetworkClient{stream}
	return x, nil
}

type SdnService_StartLocalNetworkClient interface {
	Send(*StartLocalNetworkRequest) error
	CloseAndRecv() (*StartNetworkResponse, error)
	grpc.ClientStream
}

type sdnServiceStartLocalNetworkClient struct {
	grpc.ClientStream
}

func (x *sdnServiceStartLocalNetworkClient) Send(m *StartLocalNetworkRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sdnServiceStartLocalNetworkClient) CloseAndRecv() (*StartNetworkResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartNetworkResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sdnServiceClient) StartFromPreviousNetwork(ctx context.Context, in *StartFromPreviousNetworkRequest, opts ...grpc.CallOption) (*StartNetworkResponse, error) {
	out := new(StartNetworkResponse)
	err := c.cc.Invoke(ctx, "/v1.SdnService/StartFromPreviousNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdnServiceClient) StartNetwork(ctx context.Context, in *StartNetworkRequest, opts ...grpc.CallOption) (*StartNetworkResponse, error) {
	out := new(StartNetworkResponse)
	err := c.cc.Invoke(ctx, "/v1.SdnService/StartNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdnServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/v1.SdnService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdnServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SdnService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SdnService_ServiceDesc.Streams[1], "/v1.SdnService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &sdnServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SdnService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type sdnServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *sdnServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sdnServiceClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error) {
	out := new(GetNetworkResponse)
	err := c.cc.Invoke(ctx, "/v1.SdnService/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdnServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (SdnService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &SdnService_ServiceDesc.Streams[2], "/v1.SdnService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &sdnServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SdnService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type sdnServiceListenClient struct {
	grpc.ClientStream
}

func (x *sdnServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sdnServiceClient) StopNetwork(ctx context.Context, in *StopNetworkRequest, opts ...grpc.CallOption) (*StopNetworkResponse, error) {
	out := new(StopNetworkResponse)
	err := c.cc.Invoke(ctx, "/v1.SdnService/StopNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdnServiceServer is the server API for SdnService service.
// All implementations must embed UnimplementedSdnServiceServer
// for forward compatibility
type SdnServiceServer interface {
	// StartLocalNetwork starts a Network on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the sdn/config.yaml
	//   3. all bytes constituting the Network YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalNetwork(SdnService_StartLocalNetworkServer) error
	// StartFromPreviousNetwork starts a new Network based on a previous one.
	// If the previous Network does not have the can-replay condition set this call will result in an error.
	StartFromPreviousNetwork(context.Context, *StartFromPreviousNetworkRequest) (*StartNetworkResponse, error)
	// StartNetworkRequest starts a new Network based on its specification.
	StartNetwork(context.Context, *StartNetworkRequest) (*StartNetworkResponse, error)
	// Searches for Network(s) known to this instance
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	// Subscribe listens to new Network(s) updates
	Subscribe(*SubscribeRequest, SdnService_SubscribeServer) error
	// GetNetwork retrieves details of a single Network
	GetNetwork(context.Context, *GetNetworkRequest) (*GetNetworkResponse, error)
	// Listen listens to Network updates and log output of a running Network
	Listen(*ListenRequest, SdnService_ListenServer) error
	// StopNetwork stops a currently running Network
	StopNetwork(context.Context, *StopNetworkRequest) (*StopNetworkResponse, error)
	mustEmbedUnimplementedSdnServiceServer()
}

// UnimplementedSdnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSdnServiceServer struct {
}

func (UnimplementedSdnServiceServer) StartLocalNetwork(SdnService_StartLocalNetworkServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalNetwork not implemented")
}
func (UnimplementedSdnServiceServer) StartFromPreviousNetwork(context.Context, *StartFromPreviousNetworkRequest) (*StartNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousNetwork not implemented")
}
func (UnimplementedSdnServiceServer) StartNetwork(context.Context, *StartNetworkRequest) (*StartNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNetwork not implemented")
}
func (UnimplementedSdnServiceServer) ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedSdnServiceServer) Subscribe(*SubscribeRequest, SdnService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSdnServiceServer) GetNetwork(context.Context, *GetNetworkRequest) (*GetNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetwork not implemented")
}
func (UnimplementedSdnServiceServer) Listen(*ListenRequest, SdnService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedSdnServiceServer) StopNetwork(context.Context, *StopNetworkRequest) (*StopNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopNetwork not implemented")
}
func (UnimplementedSdnServiceServer) mustEmbedUnimplementedSdnServiceServer() {}

// UnsafeSdnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdnServiceServer will
// result in compilation errors.
type UnsafeSdnServiceServer interface {
	mustEmbedUnimplementedSdnServiceServer()
}

func RegisterSdnServiceServer(s grpc.ServiceRegistrar, srv SdnServiceServer) {
	s.RegisterService(&SdnService_ServiceDesc, srv)
}

func _SdnService_StartLocalNetwork_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SdnServiceServer).StartLocalNetwork(&sdnServiceStartLocalNetworkServer{stream})
}

type SdnService_StartLocalNetworkServer interface {
	SendAndClose(*StartNetworkResponse) error
	Recv() (*StartLocalNetworkRequest, error)
	grpc.ServerStream
}

type sdnServiceStartLocalNetworkServer struct {
	grpc.ServerStream
}

func (x *sdnServiceStartLocalNetworkServer) SendAndClose(m *StartNetworkResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sdnServiceStartLocalNetworkServer) Recv() (*StartLocalNetworkRequest, error) {
	m := new(StartLocalNetworkRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SdnService_StartFromPreviousNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdnServiceServer).StartFromPreviousNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SdnService/StartFromPreviousNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdnServiceServer).StartFromPreviousNetwork(ctx, req.(*StartFromPreviousNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdnService_StartNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdnServiceServer).StartNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SdnService/StartNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdnServiceServer).StartNetwork(ctx, req.(*StartNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdnService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdnServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SdnService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdnServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdnService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SdnServiceServer).Subscribe(m, &sdnServiceSubscribeServer{stream})
}

type SdnService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type sdnServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *sdnServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SdnService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdnServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SdnService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdnServiceServer).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdnService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SdnServiceServer).Listen(m, &sdnServiceListenServer{stream})
}

type SdnService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type sdnServiceListenServer struct {
	grpc.ServerStream
}

func (x *sdnServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SdnService_StopNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdnServiceServer).StopNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SdnService/StopNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdnServiceServer).StopNetwork(ctx, req.(*StopNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdnService_ServiceDesc is the grpc.ServiceDesc for SdnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SdnService",
	HandlerType: (*SdnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousNetwork",
			Handler:    _SdnService_StartFromPreviousNetwork_Handler,
		},
		{
			MethodName: "StartNetwork",
			Handler:    _SdnService_StartNetwork_Handler,
		},
		{
			MethodName: "ListNetworks",
			Handler:    _SdnService_ListNetworks_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _SdnService_GetNetwork_Handler,
		},
		{
			MethodName: "StopNetwork",
			Handler:    _SdnService_StopNetwork_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalNetwork",
			Handler:       _SdnService_StartLocalNetwork_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _SdnService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _SdnService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdn.proto",
}
