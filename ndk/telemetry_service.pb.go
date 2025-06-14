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
// source: ndk/telemetry_service.proto

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
// Represents telemetry key.
type TelemetryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsPath string `protobuf:"bytes,1,opt,name=js_path,json=jsPath,proto3" json:"js_path,omitempty"` // JSON path referencing the key for telemetry data
}

func (x *TelemetryKey) Reset() {
	*x = TelemetryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryKey) ProtoMessage() {}

func (x *TelemetryKey) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryKey.ProtoReflect.Descriptor instead.
func (*TelemetryKey) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryKey) GetJsPath() string {
	if x != nil {
		return x.JsPath
	}
	return ""
}

// *
// Represents telemetry data.
type TelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonContent string `protobuf:"bytes,1,opt,name=json_content,json=jsonContent,proto3" json:"json_content,omitempty"` // Structured JSON telemetry data
}

func (x *TelemetryData) Reset() {
	*x = TelemetryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryData) ProtoMessage() {}

func (x *TelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryData.ProtoReflect.Descriptor instead.
func (*TelemetryData) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryData) GetJsonContent() string {
	if x != nil {
		return x.JsonContent
	}
	return ""
}

// *
// Represents telemetry information.
type TelemetryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  *TelemetryKey  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`   // Telemetry key
	Data *TelemetryData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // Telemetry data
}

func (x *TelemetryInfo) Reset() {
	*x = TelemetryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryInfo) ProtoMessage() {}

func (x *TelemetryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryInfo.ProtoReflect.Descriptor instead.
func (*TelemetryInfo) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{2}
}

func (x *TelemetryInfo) GetKey() *TelemetryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TelemetryInfo) GetData() *TelemetryData {
	if x != nil {
		return x.Data
	}
	return nil
}

// *
// Represents telemetry update request.
type TelemetryUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*TelemetryInfo `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"` // State of application
}

func (x *TelemetryUpdateRequest) Reset() {
	*x = TelemetryUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryUpdateRequest) ProtoMessage() {}

func (x *TelemetryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryUpdateRequest.ProtoReflect.Descriptor instead.
func (*TelemetryUpdateRequest) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{3}
}

func (x *TelemetryUpdateRequest) GetStates() []*TelemetryInfo {
	if x != nil {
		return x.States
	}
	return nil
}

// *
// Represents telemetry update response.
type TelemetryUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of telemetry update request
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error message
}

func (x *TelemetryUpdateResponse) Reset() {
	*x = TelemetryUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryUpdateResponse) ProtoMessage() {}

func (x *TelemetryUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryUpdateResponse.ProtoReflect.Descriptor instead.
func (*TelemetryUpdateResponse) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{4}
}

func (x *TelemetryUpdateResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_SDK_MGR_STATUS_SUCCESS
}

func (x *TelemetryUpdateResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

// *
// Represents telemetry delete request.
type TelemetryDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*TelemetryKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"` // Telemetry key
}

func (x *TelemetryDeleteRequest) Reset() {
	*x = TelemetryDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryDeleteRequest) ProtoMessage() {}

func (x *TelemetryDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryDeleteRequest.ProtoReflect.Descriptor instead.
func (*TelemetryDeleteRequest) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{5}
}

func (x *TelemetryDeleteRequest) GetKeys() []*TelemetryKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

// *
// Represents telemetry delete response.
type TelemetryDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of delete request
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error message
}

func (x *TelemetryDeleteResponse) Reset() {
	*x = TelemetryDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ndk_telemetry_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryDeleteResponse) ProtoMessage() {}

func (x *TelemetryDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ndk_telemetry_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryDeleteResponse.ProtoReflect.Descriptor instead.
func (*TelemetryDeleteResponse) Descriptor() ([]byte, []int) {
	return file_ndk_telemetry_service_proto_rawDescGZIP(), []int{6}
}

func (x *TelemetryDeleteResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_SDK_MGR_STATUS_SUCCESS
}

func (x *TelemetryDeleteResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

var File_ndk_telemetry_service_proto protoreflect.FileDescriptor

var file_ndk_telemetry_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x64, 0x6b, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73,
	0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x14, 0x6e, 0x64, 0x6b, 0x2f,
	0x73, 0x64, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x0c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x6a, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6a, 0x73, 0x50, 0x61, 0x74, 0x68, 0x22, 0x32, 0x0a, 0x0d, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a,
	0x0d, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x72, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x16, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x17, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x74, 0x72, 0x22, 0x47, 0x0a, 0x16, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x69, 0x0a,
	0x17, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x32, 0xdd, 0x01, 0x0a, 0x16, 0x53, 0x64, 0x6b,
	0x4d, 0x67, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2f, 0x73, 0x72, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2d, 0x6e, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x6e, 0x64, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ndk_telemetry_service_proto_rawDescOnce sync.Once
	file_ndk_telemetry_service_proto_rawDescData = file_ndk_telemetry_service_proto_rawDesc
)

func file_ndk_telemetry_service_proto_rawDescGZIP() []byte {
	file_ndk_telemetry_service_proto_rawDescOnce.Do(func() {
		file_ndk_telemetry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ndk_telemetry_service_proto_rawDescData)
	})
	return file_ndk_telemetry_service_proto_rawDescData
}

var file_ndk_telemetry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ndk_telemetry_service_proto_goTypes = []interface{}{
	(*TelemetryKey)(nil),            // 0: srlinux.sdk.TelemetryKey
	(*TelemetryData)(nil),           // 1: srlinux.sdk.TelemetryData
	(*TelemetryInfo)(nil),           // 2: srlinux.sdk.TelemetryInfo
	(*TelemetryUpdateRequest)(nil),  // 3: srlinux.sdk.TelemetryUpdateRequest
	(*TelemetryUpdateResponse)(nil), // 4: srlinux.sdk.TelemetryUpdateResponse
	(*TelemetryDeleteRequest)(nil),  // 5: srlinux.sdk.TelemetryDeleteRequest
	(*TelemetryDeleteResponse)(nil), // 6: srlinux.sdk.TelemetryDeleteResponse
	(SdkMgrStatus)(0),               // 7: srlinux.sdk.SdkMgrStatus
}
var file_ndk_telemetry_service_proto_depIdxs = []int32{
	0, // 0: srlinux.sdk.TelemetryInfo.key:type_name -> srlinux.sdk.TelemetryKey
	1, // 1: srlinux.sdk.TelemetryInfo.data:type_name -> srlinux.sdk.TelemetryData
	2, // 2: srlinux.sdk.TelemetryUpdateRequest.states:type_name -> srlinux.sdk.TelemetryInfo
	7, // 3: srlinux.sdk.TelemetryUpdateResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	0, // 4: srlinux.sdk.TelemetryDeleteRequest.keys:type_name -> srlinux.sdk.TelemetryKey
	7, // 5: srlinux.sdk.TelemetryDeleteResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	3, // 6: srlinux.sdk.SdkMgrTelemetryService.TelemetryAddOrUpdate:input_type -> srlinux.sdk.TelemetryUpdateRequest
	5, // 7: srlinux.sdk.SdkMgrTelemetryService.TelemetryDelete:input_type -> srlinux.sdk.TelemetryDeleteRequest
	4, // 8: srlinux.sdk.SdkMgrTelemetryService.TelemetryAddOrUpdate:output_type -> srlinux.sdk.TelemetryUpdateResponse
	6, // 9: srlinux.sdk.SdkMgrTelemetryService.TelemetryDelete:output_type -> srlinux.sdk.TelemetryDeleteResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ndk_telemetry_service_proto_init() }
func file_ndk_telemetry_service_proto_init() {
	if File_ndk_telemetry_service_proto != nil {
		return
	}
	file_ndk_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ndk_telemetry_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryKey); i {
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
		file_ndk_telemetry_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryData); i {
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
		file_ndk_telemetry_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryInfo); i {
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
		file_ndk_telemetry_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryUpdateRequest); i {
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
		file_ndk_telemetry_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryUpdateResponse); i {
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
		file_ndk_telemetry_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryDeleteRequest); i {
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
		file_ndk_telemetry_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryDeleteResponse); i {
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
			RawDescriptor: file_ndk_telemetry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ndk_telemetry_service_proto_goTypes,
		DependencyIndexes: file_ndk_telemetry_service_proto_depIdxs,
		MessageInfos:      file_ndk_telemetry_service_proto_msgTypes,
	}.Build()
	File_ndk_telemetry_service_proto = out.File
	file_ndk_telemetry_service_proto_rawDesc = nil
	file_ndk_telemetry_service_proto_goTypes = nil
	file_ndk_telemetry_service_proto_depIdxs = nil
}
