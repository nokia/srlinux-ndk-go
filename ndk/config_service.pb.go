//*********************************************************************************************************************
//  Description: interface between router agents and SDK service manager
//
//  Copyright (c) 2018 Nokia
//*********************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: ndk/config_service.proto

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
// Represents configuration subscription request.
type ConfigSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *ConfigKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Optional, to filter on name
}

func (x *ConfigSubscriptionRequest) Reset() {
	*x = ConfigSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSubscriptionRequest) ProtoMessage() {}

func (x *ConfigSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ConfigSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_ndk_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSubscriptionRequest) GetKey() *ConfigKey {
	if x != nil {
		return x.Key
	}
	return nil
}

//*
// Represents configuration key.
type ConfigKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsPath string   `protobuf:"bytes,1,opt,name=js_path,json=jsPath,proto3" json:"js_path,omitempty"` // JSON path formatted string from YANG; for example, interface{.name==ethernet1/1}.my_field
	Keys   []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`                   // Value for keys
}

func (x *ConfigKey) Reset() {
	*x = ConfigKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigKey) ProtoMessage() {}

func (x *ConfigKey) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigKey.ProtoReflect.Descriptor instead.
func (*ConfigKey) Descriptor() ([]byte, []int) {
	return file_ndk_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigKey) GetJsPath() string {
	if x != nil {
		return x.JsPath
	}
	return ""
}

func (x *ConfigKey) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

//*
// Represents configuration data.
type ConfigData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"` // Entire configuration fragment as JSON string
}

func (x *ConfigData) Reset() {
	*x = ConfigData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigData) ProtoMessage() {}

func (x *ConfigData) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigData.ProtoReflect.Descriptor instead.
func (*ConfigData) Descriptor() ([]byte, []int) {
	return file_ndk_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigData) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

//*
// Represents configuration notification message to subscribe to configuration events
type ConfigNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   SdkMgrOperation `protobuf:"varint,1,opt,name=op,proto3,enum=ndk.SdkMgrOperation" json:"op,omitempty"` // Operation indicating create, delete, or update
	Key  *ConfigKey      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                         // Configuration key
	Data *ConfigData     `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                       // Configuration data
}

func (x *ConfigNotification) Reset() {
	*x = ConfigNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigNotification) ProtoMessage() {}

func (x *ConfigNotification) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigNotification.ProtoReflect.Descriptor instead.
func (*ConfigNotification) Descriptor() ([]byte, []int) {
	return file_ndk_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigNotification) GetOp() SdkMgrOperation {
	if x != nil {
		return x.Op
	}
	return SdkMgrOperation_Create
}

func (x *ConfigNotification) GetKey() *ConfigKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ConfigNotification) GetData() *ConfigData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ndk_config_service_proto protoreflect.FileDescriptor

var file_ndk_config_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x64, 0x6b, 0x1a,
	0x14, 0x6e, 0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6e, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x6a, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x20,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e,
	0x22, 0x81, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6e, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x64, 0x6b,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6e, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2f, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x2d, 0x6e, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x6e, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ndk_config_service_proto_rawDescOnce sync.Once
	file_ndk_config_service_proto_rawDescData = file_ndk_config_service_proto_rawDesc
)

func file_ndk_config_service_proto_rawDescGZIP() []byte {
	file_ndk_config_service_proto_rawDescOnce.Do(func() {
		file_ndk_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ndk_config_service_proto_rawDescData)
	})
	return file_ndk_config_service_proto_rawDescData
}

var file_ndk_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ndk_config_service_proto_goTypes = []interface{}{
	(*ConfigSubscriptionRequest)(nil), // 0: ndk.ConfigSubscriptionRequest
	(*ConfigKey)(nil),                 // 1: ndk.ConfigKey
	(*ConfigData)(nil),                // 2: ndk.ConfigData
	(*ConfigNotification)(nil),        // 3: ndk.ConfigNotification
	(SdkMgrOperation)(0),              // 4: ndk.SdkMgrOperation
}
var file_ndk_config_service_proto_depIdxs = []int32{
	1, // 0: ndk.ConfigSubscriptionRequest.key:type_name -> ndk.ConfigKey
	4, // 1: ndk.ConfigNotification.op:type_name -> ndk.SdkMgrOperation
	1, // 2: ndk.ConfigNotification.key:type_name -> ndk.ConfigKey
	2, // 3: ndk.ConfigNotification.data:type_name -> ndk.ConfigData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ndk_config_service_proto_init() }
func file_ndk_config_service_proto_init() {
	if File_ndk_config_service_proto != nil {
		return
	}
	file_ndk_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ndk_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSubscriptionRequest); i {
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
		file_ndk_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigKey); i {
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
		file_ndk_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigData); i {
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
		file_ndk_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigNotification); i {
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
			RawDescriptor: file_ndk_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ndk_config_service_proto_goTypes,
		DependencyIndexes: file_ndk_config_service_proto_depIdxs,
		MessageInfos:      file_ndk_config_service_proto_msgTypes,
	}.Build()
	File_ndk_config_service_proto = out.File
	file_ndk_config_service_proto_rawDesc = nil
	file_ndk_config_service_proto_goTypes = nil
	file_ndk_config_service_proto_depIdxs = nil
}
