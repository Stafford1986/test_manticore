// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: search-service.proto

package pb

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

type UserType int32

const (
	UserType_guest    UserType = 0
	UserType_admin    UserType = 1
	UserType_merchant UserType = 2
	UserType_customer UserType = 3
	UserType_otp      UserType = 4
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "guest",
		1: "admin",
		2: "merchant",
		3: "customer",
		4: "otp",
	}
	UserType_value = map[string]int32{
		"guest":    0,
		"admin":    1,
		"merchant": 2,
		"customer": 3,
		"otp":      4,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_search_service_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_search_service_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{0}
}

type OperationCode int32

const (
	OperationCode_checkBalance OperationCode = 0
	OperationCode_payIn        OperationCode = 1
	OperationCode_payOut       OperationCode = 2
	OperationCode_holdAdd      OperationCode = 3
	OperationCode_holdRemove   OperationCode = 4
)

// Enum value maps for OperationCode.
var (
	OperationCode_name = map[int32]string{
		0: "checkBalance",
		1: "payIn",
		2: "payOut",
		3: "holdAdd",
		4: "holdRemove",
	}
	OperationCode_value = map[string]int32{
		"checkBalance": 0,
		"payIn":        1,
		"payOut":       2,
		"holdAdd":      3,
		"holdRemove":   4,
	}
)

func (x OperationCode) Enum() *OperationCode {
	p := new(OperationCode)
	*p = x
	return p
}

func (x OperationCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationCode) Descriptor() protoreflect.EnumDescriptor {
	return file_search_service_proto_enumTypes[1].Descriptor()
}

func (OperationCode) Type() protoreflect.EnumType {
	return &file_search_service_proto_enumTypes[1]
}

func (x OperationCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationCode.Descriptor instead.
func (OperationCode) EnumDescriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{1}
}

// swagger:model AvailableProvider
type AvailableProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required: true
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// required: true
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// required: true
	ProviderType string `protobuf:"bytes,3,opt,name=provider_type,json=providerType,proto3" json:"provider_type,omitempty"`
}

func (x *AvailableProvider) Reset() {
	*x = AvailableProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableProvider) ProtoMessage() {}

func (x *AvailableProvider) ProtoReflect() protoreflect.Message {
	mi := &file_search_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableProvider.ProtoReflect.Descriptor instead.
func (*AvailableProvider) Descriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{0}
}

func (x *AvailableProvider) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AvailableProvider) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AvailableProvider) GetProviderType() string {
	if x != nil {
		return x.ProviderType
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Skip       int64 `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	TotalItems int64 `protobuf:"varint,3,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_search_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{1}
}

func (x *Pagination) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *Pagination) GetTotalItems() int64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

type CommentedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentedResponse) Reset() {
	*x = CommentedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentedResponse) ProtoMessage() {}

func (x *CommentedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentedResponse.ProtoReflect.Descriptor instead.
func (*CommentedResponse) Descriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{2}
}

func (x *CommentedResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CommentedResponse) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_search_service_proto protoreflect.FileDescriptor

var file_search_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2d, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x11, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x45, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x6f,
	0x74, 0x70, 0x10, 0x04, 0x2a, 0x55, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x49, 0x6e,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x61, 0x79, 0x4f, 0x75, 0x74, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x64, 0x64, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x68,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x10, 0x04, 0x32, 0xc2, 0x03, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0d, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x56, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63,
	0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_service_proto_rawDescOnce sync.Once
	file_search_service_proto_rawDescData = file_search_service_proto_rawDesc
)

func file_search_service_proto_rawDescGZIP() []byte {
	file_search_service_proto_rawDescOnce.Do(func() {
		file_search_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_service_proto_rawDescData)
	})
	return file_search_service_proto_rawDescData
}

var file_search_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_search_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_search_service_proto_goTypes = []interface{}{
	(UserType)(0),                 // 0: search.UserType
	(OperationCode)(0),            // 1: search.OperationCode
	(*AvailableProvider)(nil),     // 2: search.AvailableProvider
	(*Pagination)(nil),            // 3: search.Pagination
	(*CommentedResponse)(nil),     // 4: search.CommentedResponse
	(*VacancySearchEntity)(nil),   // 5: search.VacancySearchEntity
	(*VacancyEntity)(nil),         // 6: search.VacancyEntity
	(*ResumeSearchEntity)(nil),    // 7: search.ResumeSearchEntity
	(*ResumeEntity)(nil),          // 8: search.ResumeEntity
	(*VacancySearchResponse)(nil), // 9: search.VacancySearchResponse
	(*ResumeSearchResponse)(nil),  // 10: search.ResumeSearchResponse
}
var file_search_service_proto_depIdxs = []int32{
	5,  // 0: search.SearchService.VacancySearch:input_type -> search.VacancySearchEntity
	6,  // 1: search.SearchService.VacancyIndexUpdate:input_type -> search.VacancyEntity
	6,  // 2: search.SearchService.VacancyIndexCreate:input_type -> search.VacancyEntity
	7,  // 3: search.SearchService.ResumeSearch:input_type -> search.ResumeSearchEntity
	8,  // 4: search.SearchService.ResumeIndexUpdate:input_type -> search.ResumeEntity
	8,  // 5: search.SearchService.ResumeIndexCreate:input_type -> search.ResumeEntity
	9,  // 6: search.SearchService.VacancySearch:output_type -> search.VacancySearchResponse
	4,  // 7: search.SearchService.VacancyIndexUpdate:output_type -> search.CommentedResponse
	4,  // 8: search.SearchService.VacancyIndexCreate:output_type -> search.CommentedResponse
	10, // 9: search.SearchService.ResumeSearch:output_type -> search.ResumeSearchResponse
	4,  // 10: search.SearchService.ResumeIndexUpdate:output_type -> search.CommentedResponse
	4,  // 11: search.SearchService.ResumeIndexCreate:output_type -> search.CommentedResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_search_service_proto_init() }
func file_search_service_proto_init() {
	if File_search_service_proto != nil {
		return
	}
	file_search_proto_init()
	file_search_vacancy_proto_init()
	file_search_resume_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_search_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableProvider); i {
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
		file_search_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_search_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentedResponse); i {
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
			RawDescriptor: file_search_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_service_proto_goTypes,
		DependencyIndexes: file_search_service_proto_depIdxs,
		EnumInfos:         file_search_service_proto_enumTypes,
		MessageInfos:      file_search_service_proto_msgTypes,
	}.Build()
	File_search_service_proto = out.File
	file_search_service_proto_rawDesc = nil
	file_search_service_proto_goTypes = nil
	file_search_service_proto_depIdxs = nil
}
