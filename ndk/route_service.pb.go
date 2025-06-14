//*********************************************************************************************************************
//  Description: interface between router agents and SDK service manager
//
//  Copyright (c) 2018 Nokia
//********************************************************************************************************************

// NDK Version: v0.5.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: ndk/route_service.proto

package ndk

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Represents route key.
type RouteKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkInstanceName string           `protobuf:"bytes,1,opt,name=network_instance_name,json=networkInstanceName,proto3" json:"network_instance_name,omitempty"` // Network instance name
	IpPrefix            *IpAddrPrefLenPb `protobuf:"bytes,2,opt,name=ip_prefix,json=ipPrefix,proto3" json:"ip_prefix,omitempty"`                                    // IP prefix
}

func (x *RouteKey) Reset() {
	*x = RouteKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteKey) ProtoMessage() {}

func (x *RouteKey) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteKey.ProtoReflect.Descriptor instead.
func (*RouteKey) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{0}
}

func (x *RouteKey) GetNetworkInstanceName() string {
	if x != nil {
		return x.NetworkInstanceName
	}
	return ""
}

func (x *RouteKey) GetIpPrefix() *IpAddrPrefLenPb {
	if x != nil {
		return x.IpPrefix
	}
	return nil
}

// *
// Represents route data.
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NexthopGroupName string     `protobuf:"bytes,1,opt,name=nexthop_group_name,json=nexthopGroupName,proto3" json:"nexthop_group_name,omitempty"` // Next hop group name
	Preference       uint32     `protobuf:"varint,2,opt,name=preference,proto3" json:"preference,omitempty"`                                      // Preference
	Metric           uint32     `protobuf:"varint,3,opt,name=metric,proto3" json:"metric,omitempty"`                                              // Metric
	Nexthops         []*NextHop `protobuf:"bytes,4,rep,name=nexthops,proto3" json:"nexthops,omitempty"`                                           // List of next hops
	OwnerId          uint32     `protobuf:"varint,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`                             // Next hop owner identifier returned only on notification.
	NexthopGroupId   uint64     `protobuf:"varint,6,opt,name=nexthop_group_id,json=nexthopGroupId,proto3" json:"nexthop_group_id,omitempty"`      // Next-hop group identifier returned only on notification.
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetNexthopGroupName() string {
	if x != nil {
		return x.NexthopGroupName
	}
	return ""
}

func (x *Route) GetPreference() uint32 {
	if x != nil {
		return x.Preference
	}
	return 0
}

func (x *Route) GetMetric() uint32 {
	if x != nil {
		return x.Metric
	}
	return 0
}

func (x *Route) GetNexthops() []*NextHop {
	if x != nil {
		return x.Nexthops
	}
	return nil
}

func (x *Route) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Route) GetNexthopGroupId() uint64 {
	if x != nil {
		return x.NexthopGroupId
	}
	return 0
}

// *
// Represents route information.
type RouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  *RouteKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`   // Route key
	Data *Route    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // Route data
}

func (x *RouteInfo) Reset() {
	*x = RouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteInfo) ProtoMessage() {}

func (x *RouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteInfo.ProtoReflect.Descriptor instead.
func (*RouteInfo) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{2}
}

func (x *RouteInfo) GetKey() *RouteKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *RouteInfo) GetData() *Route {
	if x != nil {
		return x.Data
	}
	return nil
}

// *
// Represents route add request; can contain more than one route.
type RouteAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*RouteInfo `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"` // IP routes
}

func (x *RouteAddRequest) Reset() {
	*x = RouteAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAddRequest) ProtoMessage() {}

func (x *RouteAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAddRequest.ProtoReflect.Descriptor instead.
func (*RouteAddRequest) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{3}
}

func (x *RouteAddRequest) GetRoutes() []*RouteInfo {
	if x != nil {
		return x.Routes
	}
	return nil
}

// *
// Represents route add response.
type RouteAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of route add operation
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error string
}

func (x *RouteAddResponse) Reset() {
	*x = RouteAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAddResponse) ProtoMessage() {}

func (x *RouteAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAddResponse.ProtoReflect.Descriptor instead.
func (*RouteAddResponse) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{4}
}

func (x *RouteAddResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_SDK_MGR_STATUS_SUCCESS
}

func (x *RouteAddResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

// *
// Represents route delete request; can contain more than one route.
type RouteDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*RouteKey `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"` // IP routes
}

func (x *RouteDeleteRequest) Reset() {
	*x = RouteDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDeleteRequest) ProtoMessage() {}

func (x *RouteDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDeleteRequest.ProtoReflect.Descriptor instead.
func (*RouteDeleteRequest) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{5}
}

func (x *RouteDeleteRequest) GetRoutes() []*RouteKey {
	if x != nil {
		return x.Routes
	}
	return nil
}

// *
// Represents route delete response.
type RouteDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of route delete operation
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error string
}

func (x *RouteDeleteResponse) Reset() {
	*x = RouteDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDeleteResponse) ProtoMessage() {}

func (x *RouteDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDeleteResponse.ProtoReflect.Descriptor instead.
func (*RouteDeleteResponse) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{6}
}

func (x *RouteDeleteResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_SDK_MGR_STATUS_SUCCESS
}

func (x *RouteDeleteResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

// *
// Represents IP route subscription request.
type IpRouteSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *RouteKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Optional, to filter on name
}

func (x *IpRouteSubscriptionRequest) Reset() {
	*x = IpRouteSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpRouteSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpRouteSubscriptionRequest) ProtoMessage() {}

func (x *IpRouteSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpRouteSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*IpRouteSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{7}
}

func (x *IpRouteSubscriptionRequest) GetKey() *RouteKey {
	if x != nil {
		return x.Key
	}
	return nil
}

// *
// Represents IP route notification.
type IpRouteNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   SdkMgrOperation `protobuf:"varint,1,opt,name=op,proto3,enum=srlinux.sdk.SdkMgrOperation" json:"op,omitempty"` // Operation such as create, delete, or update
	Key  *RouteKey       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                 // IP route key
	Data *Route          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                               // IP route data
}

func (x *IpRouteNotification) Reset() {
	*x = IpRouteNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_route_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpRouteNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpRouteNotification) ProtoMessage() {}

func (x *IpRouteNotification) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_route_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpRouteNotification.ProtoReflect.Descriptor instead.
func (*IpRouteNotification) Descriptor() ([]byte, []int) {
	return file_ndk_route_service_proto_rawDescGZIP(), []int{8}
}

func (x *IpRouteNotification) GetOp() SdkMgrOperation {
	if x != nil {
		return x.Op
	}
	return SdkMgrOperation_SDK_MGR_OPERATION_CREATE
}

func (x *IpRouteNotification) GetKey() *RouteKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *IpRouteNotification) GetData() *Route {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ndk_route_service_proto protoreflect.FileDescriptor

var file_ndk_route_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x64, 0x6b, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x6e, 0x64, 0x6b, 0x2f, 0x6e, 0x65, 0x78, 0x74,
	0x68, 0x6f, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6e, 0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a,
	0x08, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x4c, 0x65, 0x6e, 0x50, 0x62, 0x52, 0x08,
	0x69, 0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xe4, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x30, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x68, 0x6f, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x72, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70,
	0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x5c, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x72, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a,
	0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x62, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x74, 0x72, 0x22, 0x43, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x72, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x13, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53,
	0x64, 0x6b, 0x4d, 0x67, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72,
	0x22, 0x45, 0x0a, 0x1a, 0x49, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x49, 0x70, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x27, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x72, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc1,
	0x02, 0x0a, 0x12, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x72, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75,
	0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75,
	0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x72, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2f, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2d, 0x6e,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x6e, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ndk_route_service_proto_rawDescOnce sync.Once
	file_ndk_route_service_proto_rawDescData = file_ndk_route_service_proto_rawDesc
)

func file_ndk_route_service_proto_rawDescGZIP() []byte {
	file_ndk_route_service_proto_rawDescOnce.Do(func() {
		file_ndk_route_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ndk_route_service_proto_rawDescData)
	})
	return file_ndk_route_service_proto_rawDescData
}

var file_ndk_route_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ndk_route_service_proto_goTypes = []interface{}{
	(*RouteKey)(nil),                   // 0: srlinux.sdk.RouteKey
	(*Route)(nil),                      // 1: srlinux.sdk.Route
	(*RouteInfo)(nil),                  // 2: srlinux.sdk.RouteInfo
	(*RouteAddRequest)(nil),            // 3: srlinux.sdk.RouteAddRequest
	(*RouteAddResponse)(nil),           // 4: srlinux.sdk.RouteAddResponse
	(*RouteDeleteRequest)(nil),         // 5: srlinux.sdk.RouteDeleteRequest
	(*RouteDeleteResponse)(nil),        // 6: srlinux.sdk.RouteDeleteResponse
	(*IpRouteSubscriptionRequest)(nil), // 7: srlinux.sdk.IpRouteSubscriptionRequest
	(*IpRouteNotification)(nil),        // 8: srlinux.sdk.IpRouteNotification
	(*IpAddrPrefLenPb)(nil),            // 9: srlinux.sdk.IpAddrPrefLenPb
	(*NextHop)(nil),                    // 10: srlinux.sdk.NextHop
	(SdkMgrStatus)(0),                  // 11: srlinux.sdk.SdkMgrStatus
	(SdkMgrOperation)(0),               // 12: srlinux.sdk.SdkMgrOperation
	(*SyncRequest)(nil),                // 13: srlinux.sdk.SyncRequest
	(*SyncResponse)(nil),               // 14: srlinux.sdk.SyncResponse
}
var file_ndk_route_service_proto_depIdxs = []int32{
	9,  // 0: srlinux.sdk.RouteKey.ip_prefix:type_name -> srlinux.sdk.IpAddrPrefLenPb
	10, // 1: srlinux.sdk.Route.nexthops:type_name -> srlinux.sdk.NextHop
	0,  // 2: srlinux.sdk.RouteInfo.key:type_name -> srlinux.sdk.RouteKey
	1,  // 3: srlinux.sdk.RouteInfo.data:type_name -> srlinux.sdk.Route
	2,  // 4: srlinux.sdk.RouteAddRequest.routes:type_name -> srlinux.sdk.RouteInfo
	11, // 5: srlinux.sdk.RouteAddResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	0,  // 6: srlinux.sdk.RouteDeleteRequest.routes:type_name -> srlinux.sdk.RouteKey
	11, // 7: srlinux.sdk.RouteDeleteResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	0,  // 8: srlinux.sdk.IpRouteSubscriptionRequest.key:type_name -> srlinux.sdk.RouteKey
	12, // 9: srlinux.sdk.IpRouteNotification.op:type_name -> srlinux.sdk.SdkMgrOperation
	0,  // 10: srlinux.sdk.IpRouteNotification.key:type_name -> srlinux.sdk.RouteKey
	1,  // 11: srlinux.sdk.IpRouteNotification.data:type_name -> srlinux.sdk.Route
	3,  // 12: srlinux.sdk.SdkMgrRouteService.RouteAddOrUpdate:input_type -> srlinux.sdk.RouteAddRequest
	5,  // 13: srlinux.sdk.SdkMgrRouteService.RouteDelete:input_type -> srlinux.sdk.RouteDeleteRequest
	13, // 14: srlinux.sdk.SdkMgrRouteService.SyncStart:input_type -> srlinux.sdk.SyncRequest
	13, // 15: srlinux.sdk.SdkMgrRouteService.SyncEnd:input_type -> srlinux.sdk.SyncRequest
	4,  // 16: srlinux.sdk.SdkMgrRouteService.RouteAddOrUpdate:output_type -> srlinux.sdk.RouteAddResponse
	6,  // 17: srlinux.sdk.SdkMgrRouteService.RouteDelete:output_type -> srlinux.sdk.RouteDeleteResponse
	14, // 18: srlinux.sdk.SdkMgrRouteService.SyncStart:output_type -> srlinux.sdk.SyncResponse
	14, // 19: srlinux.sdk.SdkMgrRouteService.SyncEnd:output_type -> srlinux.sdk.SyncResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_ndk_route_service_proto_init() }
func file_ndk_route_service_proto_init() {
	if File_ndk_route_service_proto != nil {
		return
	}
	file_ndk_nexthop_group_service_proto_init()
	file_ndk_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ndk_route_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAddRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAddResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteDeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpRouteSubscriptionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ndk_route_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpRouteNotification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ndk_route_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ndk_route_service_proto_goTypes,
		DependencyIndexes: file_ndk_route_service_proto_depIdxs,
		MessageInfos:      file_ndk_route_service_proto_msgTypes,
	}.Build()
	File_ndk_route_service_proto = out.File
	file_ndk_route_service_proto_rawDesc = nil
	file_ndk_route_service_proto_goTypes = nil
	file_ndk_route_service_proto_depIdxs = nil
}
