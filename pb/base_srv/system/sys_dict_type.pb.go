// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: base_srv/system/sys_dict_type.proto

package system

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	page "protoRepo/pb/common/page"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSysDictTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictName string `protobuf:"bytes,2,opt,name=dictName,proto3" json:"dictName,omitempty"`
	DictType string `protobuf:"bytes,3,opt,name=dictType,proto3" json:"dictType,omitempty"`
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateSysDictTypeRequest) Reset() {
	*x = CreateSysDictTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSysDictTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSysDictTypeRequest) ProtoMessage() {}

func (x *CreateSysDictTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSysDictTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateSysDictTypeRequest) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSysDictTypeRequest) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *CreateSysDictTypeRequest) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *CreateSysDictTypeRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateSysDictTypeRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type CreateSysDictTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSysDictTypeReply) Reset() {
	*x = CreateSysDictTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSysDictTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSysDictTypeReply) ProtoMessage() {}

func (x *CreateSysDictTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSysDictTypeReply.ProtoReflect.Descriptor instead.
func (*CreateSysDictTypeReply) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{1}
}

type UpdateSysDictTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId   int32  `protobuf:"varint,1,opt,name=dictId,proto3" json:"dictId,omitempty"`
	DictName string `protobuf:"bytes,2,opt,name=dictName,proto3" json:"dictName,omitempty"`
	DictType string `protobuf:"bytes,3,opt,name=dictType,proto3" json:"dictType,omitempty"`
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UpdateSysDictTypeRequest) Reset() {
	*x = UpdateSysDictTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSysDictTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSysDictTypeRequest) ProtoMessage() {}

func (x *UpdateSysDictTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSysDictTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateSysDictTypeRequest) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSysDictTypeRequest) GetDictId() int32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *UpdateSysDictTypeRequest) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *UpdateSysDictTypeRequest) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *UpdateSysDictTypeRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateSysDictTypeRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UpdateSysDictTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSysDictTypeReply) Reset() {
	*x = UpdateSysDictTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSysDictTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSysDictTypeReply) ProtoMessage() {}

func (x *UpdateSysDictTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSysDictTypeReply.ProtoReflect.Descriptor instead.
func (*UpdateSysDictTypeReply) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{3}
}

type DeleteSysDictTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId uint32 `protobuf:"varint,1,opt,name=dictId,proto3" json:"dictId,omitempty"`
}

func (x *DeleteSysDictTypeRequest) Reset() {
	*x = DeleteSysDictTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSysDictTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSysDictTypeRequest) ProtoMessage() {}

func (x *DeleteSysDictTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSysDictTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteSysDictTypeRequest) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSysDictTypeRequest) GetDictId() uint32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

type DeleteSysDictTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSysDictTypeReply) Reset() {
	*x = DeleteSysDictTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSysDictTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSysDictTypeReply) ProtoMessage() {}

func (x *DeleteSysDictTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSysDictTypeReply.ProtoReflect.Descriptor instead.
func (*DeleteSysDictTypeReply) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{5}
}

type GetSysDictTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId int32 `protobuf:"varint,1,opt,name=dictId,proto3" json:"dictId,omitempty"`
}

func (x *GetSysDictTypeRequest) Reset() {
	*x = GetSysDictTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysDictTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysDictTypeRequest) ProtoMessage() {}

func (x *GetSysDictTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysDictTypeRequest.ProtoReflect.Descriptor instead.
func (*GetSysDictTypeRequest) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{6}
}

func (x *GetSysDictTypeRequest) GetDictId() int32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

type GetSysDictTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId   int32  `protobuf:"varint,1,opt,name=dictId,proto3" json:"dictId,omitempty"`
	DictName string `protobuf:"bytes,2,opt,name=dictName,proto3" json:"dictName,omitempty"`
	DictType string `protobuf:"bytes,3,opt,name=dictType,proto3" json:"dictType,omitempty"`
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *GetSysDictTypeReply) Reset() {
	*x = GetSysDictTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysDictTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysDictTypeReply) ProtoMessage() {}

func (x *GetSysDictTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysDictTypeReply.ProtoReflect.Descriptor instead.
func (*GetSysDictTypeReply) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{7}
}

func (x *GetSysDictTypeReply) GetDictId() int32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *GetSysDictTypeReply) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *GetSysDictTypeReply) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *GetSysDictTypeReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetSysDictTypeReply) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type ListSysDictTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageQuery *page.PageQuery `protobuf:"bytes,1,opt,name=pageQuery,proto3" json:"pageQuery,omitempty"`
	DictId    int32           `protobuf:"varint,2,opt,name=DictId,proto3" json:"DictId,omitempty"`
	DictName  string          `protobuf:"bytes,3,opt,name=DictName,proto3" json:"DictName,omitempty"`
	DictType  string          `protobuf:"bytes,4,opt,name=DictType,proto3" json:"DictType,omitempty"`
	Status    int32           `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ListSysDictTypeRequest) Reset() {
	*x = ListSysDictTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysDictTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysDictTypeRequest) ProtoMessage() {}

func (x *ListSysDictTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysDictTypeRequest.ProtoReflect.Descriptor instead.
func (*ListSysDictTypeRequest) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{8}
}

func (x *ListSysDictTypeRequest) GetPageQuery() *page.PageQuery {
	if x != nil {
		return x.PageQuery
	}
	return nil
}

func (x *ListSysDictTypeRequest) GetDictId() int32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *ListSysDictTypeRequest) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *ListSysDictTypeRequest) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *ListSysDictTypeRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListSysDictTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo *page.PageInfo         `protobuf:"bytes,1,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	List     []*GetSysDictTypeReply `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListSysDictTypeReply) Reset() {
	*x = ListSysDictTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysDictTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysDictTypeReply) ProtoMessage() {}

func (x *ListSysDictTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_srv_system_sys_dict_type_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysDictTypeReply.ProtoReflect.Descriptor instead.
func (*ListSysDictTypeReply) Descriptor() ([]byte, []int) {
	return file_base_srv_system_sys_dict_type_proto_rawDescGZIP(), []int{9}
}

func (x *ListSysDictTypeReply) GetPageInfo() *page.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *ListSysDictTypeReply) GetList() []*GetSysDictTypeReply {
	if x != nil {
		return x.List
	}
	return nil
}

var File_base_srv_system_sys_dict_type_proto protoreflect.FileDescriptor

var file_base_srv_system_sys_dict_type_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x9a, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x18, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64,
	0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64,
	0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xaf, 0x01,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x80, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x32, 0xf5, 0x05, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22,
	0x17, 0x79, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64,
	0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x79, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x99,
	0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x79, 0x77, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x7b, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x79, 0x77, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x7b, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x7d, 0x12, 0x8a, 0x01,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x79, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x73, 0x72, 0x76, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_srv_system_sys_dict_type_proto_rawDescOnce sync.Once
	file_base_srv_system_sys_dict_type_proto_rawDescData = file_base_srv_system_sys_dict_type_proto_rawDesc
)

func file_base_srv_system_sys_dict_type_proto_rawDescGZIP() []byte {
	file_base_srv_system_sys_dict_type_proto_rawDescOnce.Do(func() {
		file_base_srv_system_sys_dict_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_srv_system_sys_dict_type_proto_rawDescData)
	})
	return file_base_srv_system_sys_dict_type_proto_rawDescData
}

var file_base_srv_system_sys_dict_type_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_base_srv_system_sys_dict_type_proto_goTypes = []interface{}{
	(*CreateSysDictTypeRequest)(nil), // 0: biz.base_srv.system.CreateSysDictTypeRequest
	(*CreateSysDictTypeReply)(nil),   // 1: biz.base_srv.system.CreateSysDictTypeReply
	(*UpdateSysDictTypeRequest)(nil), // 2: biz.base_srv.system.UpdateSysDictTypeRequest
	(*UpdateSysDictTypeReply)(nil),   // 3: biz.base_srv.system.UpdateSysDictTypeReply
	(*DeleteSysDictTypeRequest)(nil), // 4: biz.base_srv.system.DeleteSysDictTypeRequest
	(*DeleteSysDictTypeReply)(nil),   // 5: biz.base_srv.system.DeleteSysDictTypeReply
	(*GetSysDictTypeRequest)(nil),    // 6: biz.base_srv.system.GetSysDictTypeRequest
	(*GetSysDictTypeReply)(nil),      // 7: biz.base_srv.system.GetSysDictTypeReply
	(*ListSysDictTypeRequest)(nil),   // 8: biz.base_srv.system.ListSysDictTypeRequest
	(*ListSysDictTypeReply)(nil),     // 9: biz.base_srv.system.ListSysDictTypeReply
	(*page.PageQuery)(nil),           // 10: page.PageQuery
	(*page.PageInfo)(nil),            // 11: page.PageInfo
}
var file_base_srv_system_sys_dict_type_proto_depIdxs = []int32{
	10, // 0: biz.base_srv.system.ListSysDictTypeRequest.pageQuery:type_name -> page.PageQuery
	11, // 1: biz.base_srv.system.ListSysDictTypeReply.pageInfo:type_name -> page.PageInfo
	7,  // 2: biz.base_srv.system.ListSysDictTypeReply.list:type_name -> biz.base_srv.system.GetSysDictTypeReply
	0,  // 3: biz.base_srv.system.SysDictType.CreateSysDictType:input_type -> biz.base_srv.system.CreateSysDictTypeRequest
	2,  // 4: biz.base_srv.system.SysDictType.UpdateSysDictType:input_type -> biz.base_srv.system.UpdateSysDictTypeRequest
	4,  // 5: biz.base_srv.system.SysDictType.DeleteSysDictType:input_type -> biz.base_srv.system.DeleteSysDictTypeRequest
	6,  // 6: biz.base_srv.system.SysDictType.GetSysDictType:input_type -> biz.base_srv.system.GetSysDictTypeRequest
	8,  // 7: biz.base_srv.system.SysDictType.ListSysDictType:input_type -> biz.base_srv.system.ListSysDictTypeRequest
	1,  // 8: biz.base_srv.system.SysDictType.CreateSysDictType:output_type -> biz.base_srv.system.CreateSysDictTypeReply
	3,  // 9: biz.base_srv.system.SysDictType.UpdateSysDictType:output_type -> biz.base_srv.system.UpdateSysDictTypeReply
	5,  // 10: biz.base_srv.system.SysDictType.DeleteSysDictType:output_type -> biz.base_srv.system.DeleteSysDictTypeReply
	7,  // 11: biz.base_srv.system.SysDictType.GetSysDictType:output_type -> biz.base_srv.system.GetSysDictTypeReply
	9,  // 12: biz.base_srv.system.SysDictType.ListSysDictType:output_type -> biz.base_srv.system.ListSysDictTypeReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_base_srv_system_sys_dict_type_proto_init() }
func file_base_srv_system_sys_dict_type_proto_init() {
	if File_base_srv_system_sys_dict_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_srv_system_sys_dict_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSysDictTypeRequest); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSysDictTypeReply); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSysDictTypeRequest); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSysDictTypeReply); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSysDictTypeRequest); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSysDictTypeReply); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysDictTypeRequest); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysDictTypeReply); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysDictTypeRequest); i {
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
		file_base_srv_system_sys_dict_type_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysDictTypeReply); i {
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
			RawDescriptor: file_base_srv_system_sys_dict_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_base_srv_system_sys_dict_type_proto_goTypes,
		DependencyIndexes: file_base_srv_system_sys_dict_type_proto_depIdxs,
		MessageInfos:      file_base_srv_system_sys_dict_type_proto_msgTypes,
	}.Build()
	File_base_srv_system_sys_dict_type_proto = out.File
	file_base_srv_system_sys_dict_type_proto_rawDesc = nil
	file_base_srv_system_sys_dict_type_proto_goTypes = nil
	file_base_srv_system_sys_dict_type_proto_depIdxs = nil
}
