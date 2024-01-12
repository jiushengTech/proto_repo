// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: device_srv/device/device.proto

package device

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDeviceRequest) Reset() {
	*x = CreateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceRequest) ProtoMessage() {}

func (x *CreateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{0}
}

type CreateDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDeviceReply) Reset() {
	*x = CreateDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceReply) ProtoMessage() {}

func (x *CreateDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceReply.ProtoReflect.Descriptor instead.
func (*CreateDeviceReply) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{1}
}

type UpdateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDeviceRequest) Reset() {
	*x = UpdateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceRequest) ProtoMessage() {}

func (x *UpdateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{2}
}

type UpdateDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDeviceReply) Reset() {
	*x = UpdateDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceReply) ProtoMessage() {}

func (x *UpdateDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceReply.ProtoReflect.Descriptor instead.
func (*UpdateDeviceReply) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{3}
}

type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{4}
}

type DeleteDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDeviceReply) Reset() {
	*x = DeleteDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceReply) ProtoMessage() {}

func (x *DeleteDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceReply.ProtoReflect.Descriptor instead.
func (*DeleteDeviceReply) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{5}
}

type GetDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeviceRequest) Reset() {
	*x = GetDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRequest) ProtoMessage() {}

func (x *GetDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId     string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	DeviceName   string `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceStatus string `protobuf:"bytes,3,opt,name=deviceStatus,proto3" json:"deviceStatus,omitempty"`
}

func (x *GetDeviceReply) Reset() {
	*x = GetDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceReply) ProtoMessage() {}

func (x *GetDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceReply.ProtoReflect.Descriptor instead.
func (*GetDeviceReply) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{7}
}

func (x *GetDeviceReply) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *GetDeviceReply) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *GetDeviceReply) GetDeviceStatus() string {
	if x != nil {
		return x.DeviceStatus
	}
	return ""
}

type ListDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDeviceRequest) Reset() {
	*x = ListDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceRequest) ProtoMessage() {}

func (x *ListDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceRequest.ProtoReflect.Descriptor instead.
func (*ListDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{8}
}

type ListDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDeviceReply) Reset() {
	*x = ListDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_srv_device_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceReply) ProtoMessage() {}

func (x *ListDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_srv_device_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceReply.ProtoReflect.Descriptor instead.
func (*ListDeviceReply) Descriptor() ([]byte, []int) {
	return file_device_srv_device_device_proto_rawDescGZIP(), []int{9}
}

var File_device_srv_device_device_proto protoreflect.FileDescriptor

var file_device_srv_device_device_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x8d, 0x04, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x64, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2a, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x71, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x35, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x65,
	0x70, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x72, 0x76,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_srv_device_device_proto_rawDescOnce sync.Once
	file_device_srv_device_device_proto_rawDescData = file_device_srv_device_device_proto_rawDesc
)

func file_device_srv_device_device_proto_rawDescGZIP() []byte {
	file_device_srv_device_device_proto_rawDescOnce.Do(func() {
		file_device_srv_device_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_srv_device_device_proto_rawDescData)
	})
	return file_device_srv_device_device_proto_rawDescData
}

var file_device_srv_device_device_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_device_srv_device_device_proto_goTypes = []interface{}{
	(*CreateDeviceRequest)(nil), // 0: biz.device_srv.device.CreateDeviceRequest
	(*CreateDeviceReply)(nil),   // 1: biz.device_srv.device.CreateDeviceReply
	(*UpdateDeviceRequest)(nil), // 2: biz.device_srv.device.UpdateDeviceRequest
	(*UpdateDeviceReply)(nil),   // 3: biz.device_srv.device.UpdateDeviceReply
	(*DeleteDeviceRequest)(nil), // 4: biz.device_srv.device.DeleteDeviceRequest
	(*DeleteDeviceReply)(nil),   // 5: biz.device_srv.device.DeleteDeviceReply
	(*GetDeviceRequest)(nil),    // 6: biz.device_srv.device.GetDeviceRequest
	(*GetDeviceReply)(nil),      // 7: biz.device_srv.device.GetDeviceReply
	(*ListDeviceRequest)(nil),   // 8: biz.device_srv.device.ListDeviceRequest
	(*ListDeviceReply)(nil),     // 9: biz.device_srv.device.ListDeviceReply
}
var file_device_srv_device_device_proto_depIdxs = []int32{
	0, // 0: biz.device_srv.device.Device.CreateDevice:input_type -> biz.device_srv.device.CreateDeviceRequest
	2, // 1: biz.device_srv.device.Device.UpdateDevice:input_type -> biz.device_srv.device.UpdateDeviceRequest
	4, // 2: biz.device_srv.device.Device.DeleteDevice:input_type -> biz.device_srv.device.DeleteDeviceRequest
	6, // 3: biz.device_srv.device.Device.GetDevice:input_type -> biz.device_srv.device.GetDeviceRequest
	8, // 4: biz.device_srv.device.Device.ListDevice:input_type -> biz.device_srv.device.ListDeviceRequest
	1, // 5: biz.device_srv.device.Device.CreateDevice:output_type -> biz.device_srv.device.CreateDeviceReply
	3, // 6: biz.device_srv.device.Device.UpdateDevice:output_type -> biz.device_srv.device.UpdateDeviceReply
	5, // 7: biz.device_srv.device.Device.DeleteDevice:output_type -> biz.device_srv.device.DeleteDeviceReply
	7, // 8: biz.device_srv.device.Device.GetDevice:output_type -> biz.device_srv.device.GetDeviceReply
	9, // 9: biz.device_srv.device.Device.ListDevice:output_type -> biz.device_srv.device.ListDeviceReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_device_srv_device_device_proto_init() }
func file_device_srv_device_device_proto_init() {
	if File_device_srv_device_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_srv_device_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceRequest); i {
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
		file_device_srv_device_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceReply); i {
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
		file_device_srv_device_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceRequest); i {
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
		file_device_srv_device_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceReply); i {
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
		file_device_srv_device_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_device_srv_device_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceReply); i {
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
		file_device_srv_device_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceRequest); i {
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
		file_device_srv_device_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceReply); i {
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
		file_device_srv_device_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceRequest); i {
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
		file_device_srv_device_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceReply); i {
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
			RawDescriptor: file_device_srv_device_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_srv_device_device_proto_goTypes,
		DependencyIndexes: file_device_srv_device_device_proto_depIdxs,
		MessageInfos:      file_device_srv_device_device_proto_msgTypes,
	}.Build()
	File_device_srv_device_device_proto = out.File
	file_device_srv_device_device_proto_rawDesc = nil
	file_device_srv_device_device_proto_goTypes = nil
	file_device_srv_device_device_proto_depIdxs = nil
}
