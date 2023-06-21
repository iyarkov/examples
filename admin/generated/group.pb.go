// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: group.proto

package generated

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

type GroupModificationResponse_ResponseCode int32

const (
	GroupModificationResponse_Ok            GroupModificationResponse_ResponseCode = 0
	GroupModificationResponse_NameNonUnique GroupModificationResponse_ResponseCode = 1
)

// Enum value maps for GroupModificationResponse_ResponseCode.
var (
	GroupModificationResponse_ResponseCode_name = map[int32]string{
		0: "Ok",
		1: "NameNonUnique",
	}
	GroupModificationResponse_ResponseCode_value = map[string]int32{
		"Ok":            0,
		"NameNonUnique": 1,
	}
)

func (x GroupModificationResponse_ResponseCode) Enum() *GroupModificationResponse_ResponseCode {
	p := new(GroupModificationResponse_ResponseCode)
	*p = x
	return p
}

func (x GroupModificationResponse_ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupModificationResponse_ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_group_proto_enumTypes[0].Descriptor()
}

func (GroupModificationResponse_ResponseCode) Type() protoreflect.EnumType {
	return &file_group_proto_enumTypes[0]
}

func (x GroupModificationResponse_ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupModificationResponse_ResponseCode.Descriptor instead.
func (GroupModificationResponse_ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5, 0}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt uint64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Group) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GroupListRequest) Reset() {
	*x = GroupListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupListRequest) ProtoMessage() {}

func (x *GroupListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupListRequest.ProtoReflect.Descriptor instead.
func (*GroupListRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupListRequest) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *GroupListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Group `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GroupListResponse) Reset() {
	*x = GroupListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupListResponse) ProtoMessage() {}

func (x *GroupListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupListResponse.ProtoReflect.Descriptor instead.
func (*GroupListResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupListResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GroupListResponse) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

type GroupCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GroupCreateRequest) Reset() {
	*x = GroupCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateRequest) ProtoMessage() {}

func (x *GroupCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateRequest.ProtoReflect.Descriptor instead.
func (*GroupCreateRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *GroupCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GroupUpdateRequest) Reset() {
	*x = GroupUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUpdateRequest) ProtoMessage() {}

func (x *GroupUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUpdateRequest.ProtoReflect.Descriptor instead.
func (*GroupUpdateRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *GroupUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupModificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   GroupModificationResponse_ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=GroupModificationResponse_ResponseCode" json:"code,omitempty"`
	Result *Group                                 `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GroupModificationResponse) Reset() {
	*x = GroupModificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupModificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupModificationResponse) ProtoMessage() {}

func (x *GroupModificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupModificationResponse.ProtoReflect.Descriptor instead.
func (*GroupModificationResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *GroupModificationResponse) GetCode() GroupModificationResponse_ResponseCode {
	if x != nil {
		return x.Code
	}
	return GroupModificationResponse_Ok
}

func (x *GroupModificationResponse) GetResult() *Group {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xa3, 0x01, 0x0a, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x6e, 0x55, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x10, 0x01, 0x32, 0xd1, 0x01, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_group_proto_goTypes = []interface{}{
	(GroupModificationResponse_ResponseCode)(0), // 0: GroupModificationResponse.ResponseCode
	(*Group)(nil),                     // 1: Group
	(*GroupListRequest)(nil),          // 2: GroupListRequest
	(*GroupListResponse)(nil),         // 3: GroupListResponse
	(*GroupCreateRequest)(nil),        // 4: GroupCreateRequest
	(*GroupUpdateRequest)(nil),        // 5: GroupUpdateRequest
	(*GroupModificationResponse)(nil), // 6: GroupModificationResponse
	(*PageRequest)(nil),               // 7: PageRequest
	(*GetRequest)(nil),                // 8: GetRequest
}
var file_group_proto_depIdxs = []int32{
	7, // 0: GroupListRequest.page:type_name -> PageRequest
	1, // 1: GroupListResponse.items:type_name -> Group
	0, // 2: GroupModificationResponse.code:type_name -> GroupModificationResponse.ResponseCode
	1, // 3: GroupModificationResponse.result:type_name -> Group
	8, // 4: Groups.Get:input_type -> GetRequest
	2, // 5: Groups.List:input_type -> GroupListRequest
	4, // 6: Groups.Create:input_type -> GroupCreateRequest
	5, // 7: Groups.Update:input_type -> GroupUpdateRequest
	1, // 8: Groups.Get:output_type -> Group
	3, // 9: Groups.List:output_type -> GroupListResponse
	6, // 10: Groups.Create:output_type -> GroupModificationResponse
	6, // 11: Groups.Update:output_type -> GroupModificationResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	file_standard_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupListRequest); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupListResponse); i {
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
		file_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateRequest); i {
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
		file_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUpdateRequest); i {
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
		file_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupModificationResponse); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		EnumInfos:         file_group_proto_enumTypes,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
