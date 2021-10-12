// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ndk

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

// SdkMgrServiceClient is the client API for SdkMgrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdkMgrServiceClient interface {
	/// Register agent
	AgentRegister(ctx context.Context, in *AgentRegistrationRequest, opts ...grpc.CallOption) (*AgentRegistrationResponse, error)
	/// Unregister agent
	AgentUnRegister(ctx context.Context, in *AgentRegistrationRequest, opts ...grpc.CallOption) (*AgentRegistrationResponse, error)
	/// Register for event notifications
	NotificationRegister(ctx context.Context, in *NotificationRegisterRequest, opts ...grpc.CallOption) (*NotificationRegisterResponse, error)
	/// Returns current or specific notification subscription details
	NotificationQuery(ctx context.Context, in *NotificationQueryRequest, opts ...grpc.CallOption) (*NotificationQueryResponse, error)
	/// Send periodic keepalive message
	KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error)
	/// Get application name from application identifier
	GetAppId(ctx context.Context, in *AppIdRequest, opts ...grpc.CallOption) (*AppIdResponse, error)
}

type sdkMgrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkMgrServiceClient(cc grpc.ClientConnInterface) SdkMgrServiceClient {
	return &sdkMgrServiceClient{cc}
}

func (c *sdkMgrServiceClient) AgentRegister(ctx context.Context, in *AgentRegistrationRequest, opts ...grpc.CallOption) (*AgentRegistrationResponse, error) {
	out := new(AgentRegistrationResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/AgentRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrServiceClient) AgentUnRegister(ctx context.Context, in *AgentRegistrationRequest, opts ...grpc.CallOption) (*AgentRegistrationResponse, error) {
	out := new(AgentRegistrationResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/AgentUnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrServiceClient) NotificationRegister(ctx context.Context, in *NotificationRegisterRequest, opts ...grpc.CallOption) (*NotificationRegisterResponse, error) {
	out := new(NotificationRegisterResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/NotificationRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrServiceClient) NotificationQuery(ctx context.Context, in *NotificationQueryRequest, opts ...grpc.CallOption) (*NotificationQueryResponse, error) {
	out := new(NotificationQueryResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/NotificationQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrServiceClient) KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error) {
	out := new(KeepAliveResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrServiceClient) GetAppId(ctx context.Context, in *AppIdRequest, opts ...grpc.CallOption) (*AppIdResponse, error) {
	out := new(AppIdResponse)
	err := c.cc.Invoke(ctx, "/ndk.SdkMgrService/GetAppId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdkMgrServiceServer is the server API for SdkMgrService service.
// All implementations must embed UnimplementedSdkMgrServiceServer
// for forward compatibility
type SdkMgrServiceServer interface {
	/// Register agent
	AgentRegister(context.Context, *AgentRegistrationRequest) (*AgentRegistrationResponse, error)
	/// Unregister agent
	AgentUnRegister(context.Context, *AgentRegistrationRequest) (*AgentRegistrationResponse, error)
	/// Register for event notifications
	NotificationRegister(context.Context, *NotificationRegisterRequest) (*NotificationRegisterResponse, error)
	/// Returns current or specific notification subscription details
	NotificationQuery(context.Context, *NotificationQueryRequest) (*NotificationQueryResponse, error)
	/// Send periodic keepalive message
	KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error)
	/// Get application name from application identifier
	GetAppId(context.Context, *AppIdRequest) (*AppIdResponse, error)
	mustEmbedUnimplementedSdkMgrServiceServer()
}

// UnimplementedSdkMgrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSdkMgrServiceServer struct {
}

func (UnimplementedSdkMgrServiceServer) AgentRegister(context.Context, *AgentRegistrationRequest) (*AgentRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgentRegister not implemented")
}
func (UnimplementedSdkMgrServiceServer) AgentUnRegister(context.Context, *AgentRegistrationRequest) (*AgentRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgentUnRegister not implemented")
}
func (UnimplementedSdkMgrServiceServer) NotificationRegister(context.Context, *NotificationRegisterRequest) (*NotificationRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationRegister not implemented")
}
func (UnimplementedSdkMgrServiceServer) NotificationQuery(context.Context, *NotificationQueryRequest) (*NotificationQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationQuery not implemented")
}
func (UnimplementedSdkMgrServiceServer) KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedSdkMgrServiceServer) GetAppId(context.Context, *AppIdRequest) (*AppIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppId not implemented")
}
func (UnimplementedSdkMgrServiceServer) mustEmbedUnimplementedSdkMgrServiceServer() {}

// UnsafeSdkMgrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdkMgrServiceServer will
// result in compilation errors.
type UnsafeSdkMgrServiceServer interface {
	mustEmbedUnimplementedSdkMgrServiceServer()
}

func RegisterSdkMgrServiceServer(s grpc.ServiceRegistrar, srv SdkMgrServiceServer) {
	s.RegisterService(&SdkMgrService_ServiceDesc, srv)
}

func _SdkMgrService_AgentRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).AgentRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/AgentRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).AgentRegister(ctx, req.(*AgentRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrService_AgentUnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).AgentUnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/AgentUnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).AgentUnRegister(ctx, req.(*AgentRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrService_NotificationRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).NotificationRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/NotificationRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).NotificationRegister(ctx, req.(*NotificationRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrService_NotificationQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).NotificationQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/NotificationQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).NotificationQuery(ctx, req.(*NotificationQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).KeepAlive(ctx, req.(*KeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrService_GetAppId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrServiceServer).GetAppId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ndk.SdkMgrService/GetAppId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrServiceServer).GetAppId(ctx, req.(*AppIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdkMgrService_ServiceDesc is the grpc.ServiceDesc for SdkMgrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdkMgrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ndk.SdkMgrService",
	HandlerType: (*SdkMgrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AgentRegister",
			Handler:    _SdkMgrService_AgentRegister_Handler,
		},
		{
			MethodName: "AgentUnRegister",
			Handler:    _SdkMgrService_AgentUnRegister_Handler,
		},
		{
			MethodName: "NotificationRegister",
			Handler:    _SdkMgrService_NotificationRegister_Handler,
		},
		{
			MethodName: "NotificationQuery",
			Handler:    _SdkMgrService_NotificationQuery_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _SdkMgrService_KeepAlive_Handler,
		},
		{
			MethodName: "GetAppId",
			Handler:    _SdkMgrService_GetAppId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ndk/sdk_service.proto",
}

// SdkNotificationServiceClient is the client API for SdkNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdkNotificationServiceClient interface {
	/// Send stream of event notifications based on the agent subscriptions
	NotificationStream(ctx context.Context, in *NotificationStreamRequest, opts ...grpc.CallOption) (SdkNotificationService_NotificationStreamClient, error)
}

type sdkNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkNotificationServiceClient(cc grpc.ClientConnInterface) SdkNotificationServiceClient {
	return &sdkNotificationServiceClient{cc}
}

func (c *sdkNotificationServiceClient) NotificationStream(ctx context.Context, in *NotificationStreamRequest, opts ...grpc.CallOption) (SdkNotificationService_NotificationStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SdkNotificationService_ServiceDesc.Streams[0], "/ndk.SdkNotificationService/NotificationStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sdkNotificationServiceNotificationStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SdkNotificationService_NotificationStreamClient interface {
	Recv() (*NotificationStreamResponse, error)
	grpc.ClientStream
}

type sdkNotificationServiceNotificationStreamClient struct {
	grpc.ClientStream
}

func (x *sdkNotificationServiceNotificationStreamClient) Recv() (*NotificationStreamResponse, error) {
	m := new(NotificationStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SdkNotificationServiceServer is the server API for SdkNotificationService service.
// All implementations must embed UnimplementedSdkNotificationServiceServer
// for forward compatibility
type SdkNotificationServiceServer interface {
	/// Send stream of event notifications based on the agent subscriptions
	NotificationStream(*NotificationStreamRequest, SdkNotificationService_NotificationStreamServer) error
	mustEmbedUnimplementedSdkNotificationServiceServer()
}

// UnimplementedSdkNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSdkNotificationServiceServer struct {
}

func (UnimplementedSdkNotificationServiceServer) NotificationStream(*NotificationStreamRequest, SdkNotificationService_NotificationStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NotificationStream not implemented")
}
func (UnimplementedSdkNotificationServiceServer) mustEmbedUnimplementedSdkNotificationServiceServer() {
}

// UnsafeSdkNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdkNotificationServiceServer will
// result in compilation errors.
type UnsafeSdkNotificationServiceServer interface {
	mustEmbedUnimplementedSdkNotificationServiceServer()
}

func RegisterSdkNotificationServiceServer(s grpc.ServiceRegistrar, srv SdkNotificationServiceServer) {
	s.RegisterService(&SdkNotificationService_ServiceDesc, srv)
}

func _SdkNotificationService_NotificationStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SdkNotificationServiceServer).NotificationStream(m, &sdkNotificationServiceNotificationStreamServer{stream})
}

type SdkNotificationService_NotificationStreamServer interface {
	Send(*NotificationStreamResponse) error
	grpc.ServerStream
}

type sdkNotificationServiceNotificationStreamServer struct {
	grpc.ServerStream
}

func (x *sdkNotificationServiceNotificationStreamServer) Send(m *NotificationStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SdkNotificationService_ServiceDesc is the grpc.ServiceDesc for SdkNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdkNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ndk.SdkNotificationService",
	HandlerType: (*SdkNotificationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotificationStream",
			Handler:       _SdkNotificationService_NotificationStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ndk/sdk_service.proto",
}