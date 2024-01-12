// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bird_ai_srv/bird/bird_type.proto

package bird

import (
	page "github.com/samsaralc/proto_repo/pb/common/page"
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

type GetBirdTypeTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBirdTypeTreeRequest) Reset() {
	*x = GetBirdTypeTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirdTypeTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirdTypeTreeRequest) ProtoMessage() {}

func (x *GetBirdTypeTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirdTypeTreeRequest.ProtoReflect.Descriptor instead.
func (*GetBirdTypeTreeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{0}
}

func (x *GetBirdTypeTreeRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBirdTypeTreeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tree *GetBirdTypeTreeReply_TreeNode `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
}

func (x *GetBirdTypeTreeReply) Reset() {
	*x = GetBirdTypeTreeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirdTypeTreeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirdTypeTreeReply) ProtoMessage() {}

func (x *GetBirdTypeTreeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirdTypeTreeReply.ProtoReflect.Descriptor instead.
func (*GetBirdTypeTreeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{1}
}

func (x *GetBirdTypeTreeReply) GetTree() *GetBirdTypeTreeReply_TreeNode {
	if x != nil {
		return x.Tree
	}
	return nil
}

type CreateBirdTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pid            uint32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Classification uint32 `protobuf:"varint,3,opt,name=classification,proto3" json:"classification,omitempty"`
}

func (x *CreateBirdTypeRequest) Reset() {
	*x = CreateBirdTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBirdTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBirdTypeRequest) ProtoMessage() {}

func (x *CreateBirdTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBirdTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateBirdTypeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBirdTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBirdTypeRequest) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *CreateBirdTypeRequest) GetClassification() uint32 {
	if x != nil {
		return x.Classification
	}
	return 0
}

type CreateBirdTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBirdTypeReply) Reset() {
	*x = CreateBirdTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBirdTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBirdTypeReply) ProtoMessage() {}

func (x *CreateBirdTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBirdTypeReply.ProtoReflect.Descriptor instead.
func (*CreateBirdTypeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{3}
}

type UpdateBirdTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBirdTypeRequest) Reset() {
	*x = UpdateBirdTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBirdTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBirdTypeRequest) ProtoMessage() {}

func (x *UpdateBirdTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBirdTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateBirdTypeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{4}
}

type UpdateBirdTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBirdTypeReply) Reset() {
	*x = UpdateBirdTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBirdTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBirdTypeReply) ProtoMessage() {}

func (x *UpdateBirdTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBirdTypeReply.ProtoReflect.Descriptor instead.
func (*UpdateBirdTypeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{5}
}

type DeleteBirdTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBirdTypeRequest) Reset() {
	*x = DeleteBirdTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBirdTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBirdTypeRequest) ProtoMessage() {}

func (x *DeleteBirdTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBirdTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteBirdTypeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{6}
}

type DeleteBirdTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBirdTypeReply) Reset() {
	*x = DeleteBirdTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBirdTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBirdTypeReply) ProtoMessage() {}

func (x *DeleteBirdTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBirdTypeReply.ProtoReflect.Descriptor instead.
func (*DeleteBirdTypeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{7}
}

type GetBirdTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBirdTypeRequest) Reset() {
	*x = GetBirdTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirdTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirdTypeRequest) ProtoMessage() {}

func (x *GetBirdTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirdTypeRequest.ProtoReflect.Descriptor instead.
func (*GetBirdTypeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{8}
}

type GetBirdTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBirdTypeReply) Reset() {
	*x = GetBirdTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirdTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirdTypeReply) ProtoMessage() {}

func (x *GetBirdTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirdTypeReply.ProtoReflect.Descriptor instead.
func (*GetBirdTypeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{9}
}

type ListBirdTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pid            uint32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Classification uint32 `protobuf:"varint,4,opt,name=classification,proto3" json:"classification,omitempty"`
}

func (x *ListBirdTypeRequest) Reset() {
	*x = ListBirdTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBirdTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBirdTypeRequest) ProtoMessage() {}

func (x *ListBirdTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBirdTypeRequest.ProtoReflect.Descriptor instead.
func (*ListBirdTypeRequest) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{10}
}

func (x *ListBirdTypeRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListBirdTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListBirdTypeRequest) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ListBirdTypeRequest) GetClassification() uint32 {
	if x != nil {
		return x.Classification
	}
	return 0
}

type ListBirdTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ListBirdTypeReply_BirdType `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListBirdTypeReply) Reset() {
	*x = ListBirdTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBirdTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBirdTypeReply) ProtoMessage() {}

func (x *ListBirdTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBirdTypeReply.ProtoReflect.Descriptor instead.
func (*ListBirdTypeReply) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{11}
}

func (x *ListBirdTypeReply) GetList() []*ListBirdTypeReply_BirdType {
	if x != nil {
		return x.List
	}
	return nil
}

type GetBirdTypeTreeReply_TreeNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pid            uint32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Classification uint32 `protobuf:"varint,4,opt,name=classification,proto3" json:"classification,omitempty"`
	// 子节点，表示树的分支
	Children []*GetBirdTypeTreeReply_TreeNode `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *GetBirdTypeTreeReply_TreeNode) Reset() {
	*x = GetBirdTypeTreeReply_TreeNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirdTypeTreeReply_TreeNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirdTypeTreeReply_TreeNode) ProtoMessage() {}

func (x *GetBirdTypeTreeReply_TreeNode) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirdTypeTreeReply_TreeNode.ProtoReflect.Descriptor instead.
func (*GetBirdTypeTreeReply_TreeNode) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetBirdTypeTreeReply_TreeNode) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBirdTypeTreeReply_TreeNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBirdTypeTreeReply_TreeNode) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *GetBirdTypeTreeReply_TreeNode) GetClassification() uint32 {
	if x != nil {
		return x.Classification
	}
	return 0
}

func (x *GetBirdTypeTreeReply_TreeNode) GetChildren() []*GetBirdTypeTreeReply_TreeNode {
	if x != nil {
		return x.Children
	}
	return nil
}

type ListBirdTypeReply_BirdType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pid            uint32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Classification uint32 `protobuf:"varint,4,opt,name=classification,proto3" json:"classification,omitempty"`
}

func (x *ListBirdTypeReply_BirdType) Reset() {
	*x = ListBirdTypeReply_BirdType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBirdTypeReply_BirdType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBirdTypeReply_BirdType) ProtoMessage() {}

func (x *ListBirdTypeReply_BirdType) ProtoReflect() protoreflect.Message {
	mi := &file_bird_ai_srv_bird_bird_type_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBirdTypeReply_BirdType.ProtoReflect.Descriptor instead.
func (*ListBirdTypeReply_BirdType) Descriptor() ([]byte, []int) {
	return file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP(), []int{11, 0}
}

func (x *ListBirdTypeReply_BirdType) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListBirdTypeReply_BirdType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListBirdTypeReply_BirdType) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ListBirdTypeReply_BirdType) GetClassification() uint32 {
	if x != nil {
		return x.Classification
	}
	return 0
}

var File_bird_ai_srv_bird_bird_type_proto protoreflect.FileDescriptor

var file_bird_ai_srv_bird_bird_type_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2f, 0x62, 0x69,
	0x72, 0x64, 0x2f, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9b, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x72, 0x65, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x74, 0x72, 0x65, 0x65, 0x1a, 0xb9, 0x01, 0x0a, 0x08, 0x54,
	0x72, 0x65, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69,
	0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x73,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc3, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69,
	0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a,
	0x68, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9c, 0x05, 0x0a, 0x08, 0x42, 0x69,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62,
	0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x79, 0x77, 0x2f, 0x62, 0x69, 0x72, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x68, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62,
	0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64,
	0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x68, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69,
	0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61,
	0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7e, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69,
	0x72, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72,
	0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x79, 0x77, 0x2f, 0x62, 0x69, 0x72,
	0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x6b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x2c,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x62,
	0x69, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x54,
	0x72, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x46, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x69, 0x72, 0x64, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x73, 0x61, 0x72, 0x61, 0x6c, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x69, 0x72, 0x64, 0x5f,
	0x61, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2f, 0x62, 0x69, 0x72, 0x64, 0x3b, 0x62, 0x69, 0x72, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bird_ai_srv_bird_bird_type_proto_rawDescOnce sync.Once
	file_bird_ai_srv_bird_bird_type_proto_rawDescData = file_bird_ai_srv_bird_bird_type_proto_rawDesc
)

func file_bird_ai_srv_bird_bird_type_proto_rawDescGZIP() []byte {
	file_bird_ai_srv_bird_bird_type_proto_rawDescOnce.Do(func() {
		file_bird_ai_srv_bird_bird_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_bird_ai_srv_bird_bird_type_proto_rawDescData)
	})
	return file_bird_ai_srv_bird_bird_type_proto_rawDescData
}

var file_bird_ai_srv_bird_bird_type_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_bird_ai_srv_bird_bird_type_proto_goTypes = []interface{}{
	(*GetBirdTypeTreeRequest)(nil),        // 0: biz.bird_ai_srv.bird.GetBirdTypeTreeRequest
	(*GetBirdTypeTreeReply)(nil),          // 1: biz.bird_ai_srv.bird.GetBirdTypeTreeReply
	(*CreateBirdTypeRequest)(nil),         // 2: biz.bird_ai_srv.bird.CreateBirdTypeRequest
	(*CreateBirdTypeReply)(nil),           // 3: biz.bird_ai_srv.bird.CreateBirdTypeReply
	(*UpdateBirdTypeRequest)(nil),         // 4: biz.bird_ai_srv.bird.UpdateBirdTypeRequest
	(*UpdateBirdTypeReply)(nil),           // 5: biz.bird_ai_srv.bird.UpdateBirdTypeReply
	(*DeleteBirdTypeRequest)(nil),         // 6: biz.bird_ai_srv.bird.DeleteBirdTypeRequest
	(*DeleteBirdTypeReply)(nil),           // 7: biz.bird_ai_srv.bird.DeleteBirdTypeReply
	(*GetBirdTypeRequest)(nil),            // 8: biz.bird_ai_srv.bird.GetBirdTypeRequest
	(*GetBirdTypeReply)(nil),              // 9: biz.bird_ai_srv.bird.GetBirdTypeReply
	(*ListBirdTypeRequest)(nil),           // 10: biz.bird_ai_srv.bird.ListBirdTypeRequest
	(*ListBirdTypeReply)(nil),             // 11: biz.bird_ai_srv.bird.ListBirdTypeReply
	(*GetBirdTypeTreeReply_TreeNode)(nil), // 12: biz.bird_ai_srv.bird.GetBirdTypeTreeReply.TreeNode
	(*ListBirdTypeReply_BirdType)(nil),    // 13: biz.bird_ai_srv.bird.ListBirdTypeReply.BirdType
	(*page.Response)(nil),                 // 14: page.Response
}
var file_bird_ai_srv_bird_bird_type_proto_depIdxs = []int32{
	12, // 0: biz.bird_ai_srv.bird.GetBirdTypeTreeReply.tree:type_name -> biz.bird_ai_srv.bird.GetBirdTypeTreeReply.TreeNode
	13, // 1: biz.bird_ai_srv.bird.ListBirdTypeReply.list:type_name -> biz.bird_ai_srv.bird.ListBirdTypeReply.BirdType
	12, // 2: biz.bird_ai_srv.bird.GetBirdTypeTreeReply.TreeNode.children:type_name -> biz.bird_ai_srv.bird.GetBirdTypeTreeReply.TreeNode
	2,  // 3: biz.bird_ai_srv.bird.BirdType.CreateBirdType:input_type -> biz.bird_ai_srv.bird.CreateBirdTypeRequest
	4,  // 4: biz.bird_ai_srv.bird.BirdType.UpdateBirdType:input_type -> biz.bird_ai_srv.bird.UpdateBirdTypeRequest
	6,  // 5: biz.bird_ai_srv.bird.BirdType.DeleteBirdType:input_type -> biz.bird_ai_srv.bird.DeleteBirdTypeRequest
	8,  // 6: biz.bird_ai_srv.bird.BirdType.GetBirdType:input_type -> biz.bird_ai_srv.bird.GetBirdTypeRequest
	10, // 7: biz.bird_ai_srv.bird.BirdType.ListBirdType:input_type -> biz.bird_ai_srv.bird.ListBirdTypeRequest
	0,  // 8: biz.bird_ai_srv.bird.BirdType.GetBirdTypeTree:input_type -> biz.bird_ai_srv.bird.GetBirdTypeTreeRequest
	14, // 9: biz.bird_ai_srv.bird.BirdType.CreateBirdType:output_type -> page.Response
	5,  // 10: biz.bird_ai_srv.bird.BirdType.UpdateBirdType:output_type -> biz.bird_ai_srv.bird.UpdateBirdTypeReply
	7,  // 11: biz.bird_ai_srv.bird.BirdType.DeleteBirdType:output_type -> biz.bird_ai_srv.bird.DeleteBirdTypeReply
	9,  // 12: biz.bird_ai_srv.bird.BirdType.GetBirdType:output_type -> biz.bird_ai_srv.bird.GetBirdTypeReply
	11, // 13: biz.bird_ai_srv.bird.BirdType.ListBirdType:output_type -> biz.bird_ai_srv.bird.ListBirdTypeReply
	1,  // 14: biz.bird_ai_srv.bird.BirdType.GetBirdTypeTree:output_type -> biz.bird_ai_srv.bird.GetBirdTypeTreeReply
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_bird_ai_srv_bird_bird_type_proto_init() }
func file_bird_ai_srv_bird_bird_type_proto_init() {
	if File_bird_ai_srv_bird_bird_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirdTypeTreeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirdTypeTreeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBirdTypeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBirdTypeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBirdTypeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBirdTypeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBirdTypeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBirdTypeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirdTypeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirdTypeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBirdTypeRequest); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBirdTypeReply); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirdTypeTreeReply_TreeNode); i {
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
		file_bird_ai_srv_bird_bird_type_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBirdTypeReply_BirdType); i {
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
			RawDescriptor: file_bird_ai_srv_bird_bird_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bird_ai_srv_bird_bird_type_proto_goTypes,
		DependencyIndexes: file_bird_ai_srv_bird_bird_type_proto_depIdxs,
		MessageInfos:      file_bird_ai_srv_bird_bird_type_proto_msgTypes,
	}.Build()
	File_bird_ai_srv_bird_bird_type_proto = out.File
	file_bird_ai_srv_bird_bird_type_proto_rawDesc = nil
	file_bird_ai_srv_bird_bird_type_proto_goTypes = nil
	file_bird_ai_srv_bird_bird_type_proto_depIdxs = nil
}
