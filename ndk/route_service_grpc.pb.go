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

// SdkMgrRouteServiceClient is the client API for SdkMgrRouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdkMgrRouteServiceClient interface {
	/// Add or update IP routes.
	RouteAddOrUpdate(ctx context.Context, in *RouteAddRequest, opts ...grpc.CallOption) (*RouteAddResponse, error)
	/// Delete IP routes.
	RouteDelete(ctx context.Context, in *RouteDeleteRequest, opts ...grpc.CallOption) (*RouteDeleteResponse, error)
	/// Synchronization start for IP routes
	SyncStart(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	/// Synchronization end for IP routes
	SyncEnd(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type sdkMgrRouteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkMgrRouteServiceClient(cc grpc.ClientConnInterface) SdkMgrRouteServiceClient {
	return &sdkMgrRouteServiceClient{cc}
}

func (c *sdkMgrRouteServiceClient) RouteAddOrUpdate(ctx context.Context, in *RouteAddRequest, opts ...grpc.CallOption) (*RouteAddResponse, error) {
	out := new(RouteAddResponse)
	err := c.cc.Invoke(ctx, "/srlinux.sdk.SdkMgrRouteService/RouteAddOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrRouteServiceClient) RouteDelete(ctx context.Context, in *RouteDeleteRequest, opts ...grpc.CallOption) (*RouteDeleteResponse, error) {
	out := new(RouteDeleteResponse)
	err := c.cc.Invoke(ctx, "/srlinux.sdk.SdkMgrRouteService/RouteDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrRouteServiceClient) SyncStart(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/srlinux.sdk.SdkMgrRouteService/SyncStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrRouteServiceClient) SyncEnd(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/srlinux.sdk.SdkMgrRouteService/SyncEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdkMgrRouteServiceServer is the server API for SdkMgrRouteService service.
// All implementations must embed UnimplementedSdkMgrRouteServiceServer
// for forward compatibility
type SdkMgrRouteServiceServer interface {
	/// Add or update IP routes.
	RouteAddOrUpdate(context.Context, *RouteAddRequest) (*RouteAddResponse, error)
	/// Delete IP routes.
	RouteDelete(context.Context, *RouteDeleteRequest) (*RouteDeleteResponse, error)
	/// Synchronization start for IP routes
	SyncStart(context.Context, *SyncRequest) (*SyncResponse, error)
	/// Synchronization end for IP routes
	SyncEnd(context.Context, *SyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedSdkMgrRouteServiceServer()
}

// UnimplementedSdkMgrRouteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSdkMgrRouteServiceServer struct {
}

func (UnimplementedSdkMgrRouteServiceServer) RouteAddOrUpdate(context.Context, *RouteAddRequest) (*RouteAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteAddOrUpdate not implemented")
}
func (UnimplementedSdkMgrRouteServiceServer) RouteDelete(context.Context, *RouteDeleteRequest) (*RouteDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteDelete not implemented")
}
func (UnimplementedSdkMgrRouteServiceServer) SyncStart(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncStart not implemented")
}
func (UnimplementedSdkMgrRouteServiceServer) SyncEnd(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncEnd not implemented")
}
func (UnimplementedSdkMgrRouteServiceServer) mustEmbedUnimplementedSdkMgrRouteServiceServer() {}

// UnsafeSdkMgrRouteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdkMgrRouteServiceServer will
// result in compilation errors.
type UnsafeSdkMgrRouteServiceServer interface {
	mustEmbedUnimplementedSdkMgrRouteServiceServer()
}

func RegisterSdkMgrRouteServiceServer(s grpc.ServiceRegistrar, srv SdkMgrRouteServiceServer) {
	s.RegisterService(&SdkMgrRouteService_ServiceDesc, srv)
}

func _SdkMgrRouteService_RouteAddOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrRouteServiceServer).RouteAddOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srlinux.sdk.SdkMgrRouteService/RouteAddOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrRouteServiceServer).RouteAddOrUpdate(ctx, req.(*RouteAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrRouteService_RouteDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrRouteServiceServer).RouteDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srlinux.sdk.SdkMgrRouteService/RouteDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrRouteServiceServer).RouteDelete(ctx, req.(*RouteDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrRouteService_SyncStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrRouteServiceServer).SyncStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srlinux.sdk.SdkMgrRouteService/SyncStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrRouteServiceServer).SyncStart(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrRouteService_SyncEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrRouteServiceServer).SyncEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srlinux.sdk.SdkMgrRouteService/SyncEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrRouteServiceServer).SyncEnd(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdkMgrRouteService_ServiceDesc is the grpc.ServiceDesc for SdkMgrRouteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdkMgrRouteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "srlinux.sdk.SdkMgrRouteService",
	HandlerType: (*SdkMgrRouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RouteAddOrUpdate",
			Handler:    _SdkMgrRouteService_RouteAddOrUpdate_Handler,
		},
		{
			MethodName: "RouteDelete",
			Handler:    _SdkMgrRouteService_RouteDelete_Handler,
		},
		{
			MethodName: "SyncStart",
			Handler:    _SdkMgrRouteService_SyncStart_Handler,
		},
		{
			MethodName: "SyncEnd",
			Handler:    _SdkMgrRouteService_SyncEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ndk/route_service.proto",
}