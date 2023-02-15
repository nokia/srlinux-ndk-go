//*********************************************************************************************************************
//  Description: interface between router agents and SDK service manager
//
//  Copyright (c) 2018 Nokia
//*********************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ndk/interface_service.proto

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

//*
// Represents interface subscription request.
type InterfaceSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *InterfaceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Optional, to filter on name
}

func (x *InterfaceSubscriptionRequest) Reset() {
	*x = InterfaceSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_interface_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceSubscriptionRequest) ProtoMessage() {}

func (x *InterfaceSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_interface_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*InterfaceSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_ndk_interface_service_proto_rawDescGZIP(), []int{0}
}

func (x *InterfaceSubscriptionRequest) GetKey() *InterfaceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

//*
//  Represents interface key.
type InterfaceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IfName string `protobuf:"bytes,1,opt,name=if_name,json=ifName,proto3" json:"if_name,omitempty"` // Interface name; for example, ethernet 1/1
}

func (x *InterfaceKey) Reset() {
	*x = InterfaceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_interface_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceKey) ProtoMessage() {}

func (x *InterfaceKey) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_interface_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceKey.ProtoReflect.Descriptor instead.
func (*InterfaceKey) Descriptor() ([]byte, []int) {
	return file_ndk_interface_service_proto_rawDescGZIP(), []int{1}
}

func (x *InterfaceKey) GetIfName() string {
	if x != nil {
		return x.IfName
	}
	return ""
}

//*
// Represents interface data.
type InterfaceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminIsUp   uint32        `protobuf:"varint,1,opt,name=admin_is_up,json=adminIsUp,proto3" json:"admin_is_up,omitempty"`                   // Admin state
	Mtu         uint32        `protobuf:"varint,2,opt,name=mtu,proto3" json:"mtu,omitempty"`                                                  // Maximum transmission unit
	IfType      IfMgrIfType   `protobuf:"varint,3,opt,name=if_type,json=ifType,proto3,enum=srlinux.sdk.IfMgrIfType" json:"if_type,omitempty"` // Interface type; for example, loopback, physical, or LAG
	PortId      *PortIdPb     `protobuf:"bytes,4,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`                               // Port identifier
	Description string        `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`                                   // Interface description
	MacAddr     *MacAddressPb `protobuf:"bytes,6,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`                            // MAC address
	AggregateId string        `protobuf:"bytes,7,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`                // associated aggregate id
	OperIsUp    uint32        `protobuf:"varint,8,opt,name=oper_is_up,json=operIsUp,proto3" json:"oper_is_up,omitempty"`                      // Operational state
}

func (x *InterfaceData) Reset() {
	*x = InterfaceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_interface_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceData) ProtoMessage() {}

func (x *InterfaceData) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_interface_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceData.ProtoReflect.Descriptor instead.
func (*InterfaceData) Descriptor() ([]byte, []int) {
	return file_ndk_interface_service_proto_rawDescGZIP(), []int{2}
}

func (x *InterfaceData) GetAdminIsUp() uint32 {
	if x != nil {
		return x.AdminIsUp
	}
	return 0
}

func (x *InterfaceData) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *InterfaceData) GetIfType() IfMgrIfType {
	if x != nil {
		return x.IfType
	}
	return IfMgrIfType_ETHERNET
}

func (x *InterfaceData) GetPortId() *PortIdPb {
	if x != nil {
		return x.PortId
	}
	return nil
}

func (x *InterfaceData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InterfaceData) GetMacAddr() *MacAddressPb {
	if x != nil {
		return x.MacAddr
	}
	return nil
}

func (x *InterfaceData) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *InterfaceData) GetOperIsUp() uint32 {
	if x != nil {
		return x.OperIsUp
	}
	return 0
}

//*
// Represents interface notification.
type InterfaceNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   SdkMgrOperation `protobuf:"varint,1,opt,name=op,proto3,enum=srlinux.sdk.SdkMgrOperation" json:"op,omitempty"` // Operation such as create, delete, or update
	Key  *InterfaceKey   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                 // Interface key
	Data *InterfaceData  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                               // Interface data
}

func (x *InterfaceNotification) Reset() {
	*x = InterfaceNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_interface_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceNotification) ProtoMessage() {}

func (x *InterfaceNotification) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_interface_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceNotification.ProtoReflect.Descriptor instead.
func (*InterfaceNotification) Descriptor() ([]byte, []int) {
	return file_ndk_interface_service_proto_rawDescGZIP(), []int{3}
}

func (x *InterfaceNotification) GetOp() SdkMgrOperation {
	if x != nil {
		return x.Op
	}
	return SdkMgrOperation_Create
}

func (x *InterfaceNotification) GetKey() *InterfaceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *InterfaceNotification) GetData() *InterfaceData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ndk_interface_service_proto protoreflect.FileDescriptor

var file_ndk_interface_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x64, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73,
	0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x14, 0x6e, 0x64, 0x6b, 0x2f,
	0x73, 0x64, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x1c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x27, 0x0a,
	0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x73, 0x55, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x31, 0x0a, 0x07, 0x69, 0x66,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x66, 0x4d, 0x67, 0x72, 0x49,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x69, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x50, 0x62, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x62, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x49, 0x73, 0x55, 0x70, 0x22, 0xa2, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73,
	0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67,
	0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x2b,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x72, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2f,
	0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2d, 0x6e, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x6e,
	0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ndk_interface_service_proto_rawDescOnce sync.Once
	file_ndk_interface_service_proto_rawDescData = file_ndk_interface_service_proto_rawDesc
)

func file_ndk_interface_service_proto_rawDescGZIP() []byte {
	file_ndk_interface_service_proto_rawDescOnce.Do(func() {
		file_ndk_interface_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ndk_interface_service_proto_rawDescData)
	})
	return file_ndk_interface_service_proto_rawDescData
}

var file_ndk_interface_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ndk_interface_service_proto_goTypes = []interface{}{
	(*InterfaceSubscriptionRequest)(nil), // 0: srlinux.sdk.InterfaceSubscriptionRequest
	(*InterfaceKey)(nil),                 // 1: srlinux.sdk.InterfaceKey
	(*InterfaceData)(nil),                // 2: srlinux.sdk.InterfaceData
	(*InterfaceNotification)(nil),        // 3: srlinux.sdk.InterfaceNotification
	(IfMgrIfType)(0),                     // 4: srlinux.sdk.IfMgrIfType
	(*PortIdPb)(nil),                     // 5: srlinux.sdk.PortIdPb
	(*MacAddressPb)(nil),                 // 6: srlinux.sdk.MacAddressPb
	(SdkMgrOperation)(0),                 // 7: srlinux.sdk.SdkMgrOperation
}
var file_ndk_interface_service_proto_depIdxs = []int32{
	1, // 0: srlinux.sdk.InterfaceSubscriptionRequest.key:type_name -> srlinux.sdk.InterfaceKey
	4, // 1: srlinux.sdk.InterfaceData.if_type:type_name -> srlinux.sdk.IfMgrIfType
	5, // 2: srlinux.sdk.InterfaceData.port_id:type_name -> srlinux.sdk.PortIdPb
	6, // 3: srlinux.sdk.InterfaceData.mac_addr:type_name -> srlinux.sdk.MacAddressPb
	7, // 4: srlinux.sdk.InterfaceNotification.op:type_name -> srlinux.sdk.SdkMgrOperation
	1, // 5: srlinux.sdk.InterfaceNotification.key:type_name -> srlinux.sdk.InterfaceKey
	2, // 6: srlinux.sdk.InterfaceNotification.data:type_name -> srlinux.sdk.InterfaceData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_ndk_interface_service_proto_init() }
func file_ndk_interface_service_proto_init() {
	if File_ndk_interface_service_proto != nil {
		return
	}
	file_ndk_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ndk_interface_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceSubscriptionRequest); i {
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
		file_ndk_interface_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceKey); i {
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
		file_ndk_interface_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceData); i {
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
		file_ndk_interface_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceNotification); i {
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
			RawDescriptor: file_ndk_interface_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ndk_interface_service_proto_goTypes,
		DependencyIndexes: file_ndk_interface_service_proto_depIdxs,
		MessageInfos:      file_ndk_interface_service_proto_msgTypes,
	}.Build()
	File_ndk_interface_service_proto = out.File
	file_ndk_interface_service_proto_rawDesc = nil
	file_ndk_interface_service_proto_goTypes = nil
	file_ndk_interface_service_proto_depIdxs = nil
}
