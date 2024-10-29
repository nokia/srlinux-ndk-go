//*********************************************************************************************************************
//  Description: Interface between SDK service manager and Next Hop Group agent
//
//  Copyright (c) 2018 Nokia
//*********************************************************************************************************************

// NDK Version: v0.4.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: ndk/nexthop_group_service.proto

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

const (
	SdkMgrNextHopGroupService_NextHopGroupAddOrUpdate_FullMethodName = "/srlinux.sdk.SdkMgrNextHopGroupService/NextHopGroupAddOrUpdate"
	SdkMgrNextHopGroupService_NextHopGroupDelete_FullMethodName      = "/srlinux.sdk.SdkMgrNextHopGroupService/NextHopGroupDelete"
	SdkMgrNextHopGroupService_SyncStart_FullMethodName               = "/srlinux.sdk.SdkMgrNextHopGroupService/SyncStart"
	SdkMgrNextHopGroupService_SyncEnd_FullMethodName                 = "/srlinux.sdk.SdkMgrNextHopGroupService/SyncEnd"
)

// SdkMgrNextHopGroupServiceClient is the client API for SdkMgrNextHopGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdkMgrNextHopGroupServiceClient interface {
	// / Add or update one or more next-hop groups.
	NextHopGroupAddOrUpdate(ctx context.Context, in *NextHopGroupRequest, opts ...grpc.CallOption) (*NextHopGroupResponse, error)
	// / Delete next-hop group.
	NextHopGroupDelete(ctx context.Context, in *NextHopGroupDeleteRequest, opts ...grpc.CallOption) (*NextHopGroupDeleteResponse, error)
	// / Synchronization start to open synchronization operation.
	SyncStart(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	// / Synchronization end to close synchronization operation.
	SyncEnd(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type sdkMgrNextHopGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkMgrNextHopGroupServiceClient(cc grpc.ClientConnInterface) SdkMgrNextHopGroupServiceClient {
	return &sdkMgrNextHopGroupServiceClient{cc}
}

func (c *sdkMgrNextHopGroupServiceClient) NextHopGroupAddOrUpdate(ctx context.Context, in *NextHopGroupRequest, opts ...grpc.CallOption) (*NextHopGroupResponse, error) {
	out := new(NextHopGroupResponse)
	err := c.cc.Invoke(ctx, SdkMgrNextHopGroupService_NextHopGroupAddOrUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrNextHopGroupServiceClient) NextHopGroupDelete(ctx context.Context, in *NextHopGroupDeleteRequest, opts ...grpc.CallOption) (*NextHopGroupDeleteResponse, error) {
	out := new(NextHopGroupDeleteResponse)
	err := c.cc.Invoke(ctx, SdkMgrNextHopGroupService_NextHopGroupDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrNextHopGroupServiceClient) SyncStart(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, SdkMgrNextHopGroupService_SyncStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkMgrNextHopGroupServiceClient) SyncEnd(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, SdkMgrNextHopGroupService_SyncEnd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdkMgrNextHopGroupServiceServer is the server API for SdkMgrNextHopGroupService service.
// All implementations should embed UnimplementedSdkMgrNextHopGroupServiceServer
// for forward compatibility
type SdkMgrNextHopGroupServiceServer interface {
	// / Add or update one or more next-hop groups.
	NextHopGroupAddOrUpdate(context.Context, *NextHopGroupRequest) (*NextHopGroupResponse, error)
	// / Delete next-hop group.
	NextHopGroupDelete(context.Context, *NextHopGroupDeleteRequest) (*NextHopGroupDeleteResponse, error)
	// / Synchronization start to open synchronization operation.
	SyncStart(context.Context, *SyncRequest) (*SyncResponse, error)
	// / Synchronization end to close synchronization operation.
	SyncEnd(context.Context, *SyncRequest) (*SyncResponse, error)
}

// UnimplementedSdkMgrNextHopGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSdkMgrNextHopGroupServiceServer struct {
}

func (UnimplementedSdkMgrNextHopGroupServiceServer) NextHopGroupAddOrUpdate(context.Context, *NextHopGroupRequest) (*NextHopGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextHopGroupAddOrUpdate not implemented")
}
func (UnimplementedSdkMgrNextHopGroupServiceServer) NextHopGroupDelete(context.Context, *NextHopGroupDeleteRequest) (*NextHopGroupDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextHopGroupDelete not implemented")
}
func (UnimplementedSdkMgrNextHopGroupServiceServer) SyncStart(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncStart not implemented")
}
func (UnimplementedSdkMgrNextHopGroupServiceServer) SyncEnd(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncEnd not implemented")
}

// UnsafeSdkMgrNextHopGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdkMgrNextHopGroupServiceServer will
// result in compilation errors.
type UnsafeSdkMgrNextHopGroupServiceServer interface {
	mustEmbedUnimplementedSdkMgrNextHopGroupServiceServer()
}

func RegisterSdkMgrNextHopGroupServiceServer(s grpc.ServiceRegistrar, srv SdkMgrNextHopGroupServiceServer) {
	s.RegisterService(&SdkMgrNextHopGroupService_ServiceDesc, srv)
}

func _SdkMgrNextHopGroupService_NextHopGroupAddOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextHopGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrNextHopGroupServiceServer).NextHopGroupAddOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SdkMgrNextHopGroupService_NextHopGroupAddOrUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrNextHopGroupServiceServer).NextHopGroupAddOrUpdate(ctx, req.(*NextHopGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrNextHopGroupService_NextHopGroupDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextHopGroupDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrNextHopGroupServiceServer).NextHopGroupDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SdkMgrNextHopGroupService_NextHopGroupDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrNextHopGroupServiceServer).NextHopGroupDelete(ctx, req.(*NextHopGroupDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrNextHopGroupService_SyncStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrNextHopGroupServiceServer).SyncStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SdkMgrNextHopGroupService_SyncStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrNextHopGroupServiceServer).SyncStart(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkMgrNextHopGroupService_SyncEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkMgrNextHopGroupServiceServer).SyncEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SdkMgrNextHopGroupService_SyncEnd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkMgrNextHopGroupServiceServer).SyncEnd(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdkMgrNextHopGroupService_ServiceDesc is the grpc.ServiceDesc for SdkMgrNextHopGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdkMgrNextHopGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "srlinux.sdk.SdkMgrNextHopGroupService",
	HandlerType: (*SdkMgrNextHopGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextHopGroupAddOrUpdate",
			Handler:    _SdkMgrNextHopGroupService_NextHopGroupAddOrUpdate_Handler,
		},
		{
			MethodName: "NextHopGroupDelete",
			Handler:    _SdkMgrNextHopGroupService_NextHopGroupDelete_Handler,
		},
		{
			MethodName: "SyncStart",
			Handler:    _SdkMgrNextHopGroupService_SyncStart_Handler,
		},
		{
			MethodName: "SyncEnd",
			Handler:    _SdkMgrNextHopGroupService_SyncEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ndk/nexthop_group_service.proto",
}
