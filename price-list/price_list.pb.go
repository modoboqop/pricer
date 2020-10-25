// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: price-list/price_list.proto

package price_list

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SortingType int32

const (
	SortingType_SortByProductName  SortingType = 0
	SortingType_SortByPrice        SortingType = 1
	SortingType_SortByUpdatesCount SortingType = 2
	SortingType_SortByUpdateTime   SortingType = 3
)

// Enum value maps for SortingType.
var (
	SortingType_name = map[int32]string{
		0: "SortByProductName",
		1: "SortByPrice",
		2: "SortByUpdatesCount",
		3: "SortByUpdateTime",
	}
	SortingType_value = map[string]int32{
		"SortByProductName":  0,
		"SortByPrice":        1,
		"SortByUpdatesCount": 2,
		"SortByUpdateTime":   3,
	}
)

func (x SortingType) Enum() *SortingType {
	p := new(SortingType)
	*p = x
	return p
}

func (x SortingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingType) Descriptor() protoreflect.EnumDescriptor {
	return file_price_list_price_list_proto_enumTypes[0].Descriptor()
}

func (SortingType) Type() protoreflect.EnumType {
	return &file_price_list_price_list_proto_enumTypes[0]
}

func (x SortingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingType.Descriptor instead.
func (SortingType) EnumDescriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{0}
}

type SortingDirection int32

const (
	SortingDirection_SortAsc  SortingDirection = 0
	SortingDirection_SortDesc SortingDirection = 1
)

// Enum value maps for SortingDirection.
var (
	SortingDirection_name = map[int32]string{
		0: "SortAsc",
		1: "SortDesc",
	}
	SortingDirection_value = map[string]int32{
		"SortAsc":  0,
		"SortDesc": 1,
	}
)

func (x SortingDirection) Enum() *SortingDirection {
	p := new(SortingDirection)
	*p = x
	return p
}

func (x SortingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_price_list_price_list_proto_enumTypes[1].Descriptor()
}

func (SortingDirection) Type() protoreflect.EnumType {
	return &file_price_list_price_list_proto_enumTypes[1]
}

func (x SortingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingDirection.Descriptor instead.
func (SortingDirection) EnumDescriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{1}
}

// The request message containing the url of external service
type UpdatePriceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UpdatePriceListRequest) Reset() {
	*x = UpdatePriceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceListRequest) ProtoMessage() {}

func (x *UpdatePriceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceListRequest.ProtoReflect.Descriptor instead.
func (*UpdatePriceListRequest) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePriceListRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UpdatePriceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductPrice `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *UpdatePriceListResponse) Reset() {
	*x = UpdatePriceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceListResponse) ProtoMessage() {}

func (x *UpdatePriceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceListResponse.ProtoReflect.Descriptor instead.
func (*UpdatePriceListResponse) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePriceListResponse) GetProducts() []*ProductPrice {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName       string `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	ProductPriceCents int32  `protobuf:"varint,2,opt,name=productPriceCents,proto3" json:"productPriceCents,omitempty"`
}

func (x *ProductPrice) Reset() {
	*x = ProductPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPrice) ProtoMessage() {}

func (x *ProductPrice) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPrice.ProtoReflect.Descriptor instead.
func (*ProductPrice) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{2}
}

func (x *ProductPrice) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductPrice) GetProductPriceCents() int32 {
	if x != nil {
		return x.ProductPriceCents
	}
	return 0
}

type GetProductPricesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit            int64            `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset           int64            `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortingType      SortingType      `protobuf:"varint,3,opt,name=sortingType,proto3,enum=SortingType" json:"sortingType,omitempty"`
	SortingDirection SortingDirection `protobuf:"varint,4,opt,name=sortingDirection,proto3,enum=SortingDirection" json:"sortingDirection,omitempty"`
}

func (x *GetProductPricesRequest) Reset() {
	*x = GetProductPricesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductPricesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductPricesRequest) ProtoMessage() {}

func (x *GetProductPricesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductPricesRequest.ProtoReflect.Descriptor instead.
func (*GetProductPricesRequest) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductPricesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetProductPricesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetProductPricesRequest) GetSortingType() SortingType {
	if x != nil {
		return x.SortingType
	}
	return SortingType_SortByProductName
}

func (x *GetProductPricesRequest) GetSortingDirection() SortingDirection {
	if x != nil {
		return x.SortingDirection
	}
	return SortingDirection_SortAsc
}

type GetProductPricesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductPrices `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetProductPricesResponse) Reset() {
	*x = GetProductPricesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductPricesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductPricesResponse) ProtoMessage() {}

func (x *GetProductPricesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductPricesResponse.ProtoReflect.Descriptor instead.
func (*GetProductPricesResponse) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductPricesResponse) GetProducts() []*ProductPrices {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName       string `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	ProductPriceCents int32  `protobuf:"varint,2,opt,name=productPriceCents,proto3" json:"productPriceCents,omitempty"`
	UpdateCount       int64  `protobuf:"varint,3,opt,name=updateCount,proto3" json:"updateCount,omitempty"`
	UpdateTime        string `protobuf:"bytes,4,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ProductPrices) Reset() {
	*x = ProductPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_list_price_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPrices) ProtoMessage() {}

func (x *ProductPrices) ProtoReflect() protoreflect.Message {
	mi := &file_price_list_price_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPrices.ProtoReflect.Descriptor instead.
func (*ProductPrices) Descriptor() ([]byte, []int) {
	return file_price_list_price_list_proto_rawDescGZIP(), []int{5}
}

func (x *ProductPrices) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductPrices) GetProductPriceCents() int32 {
	if x != nil {
		return x.ProductPriceCents
	}
	return 0
}

func (x *ProductPrices) GetUpdateCount() int64 {
	if x != nil {
		return x.UpdateCount
	}
	return 0
}

func (x *ProductPrices) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

var File_price_list_price_list_proto protoreflect.FileDescriptor

var file_price_list_price_list_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x5e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x43, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xb6, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x2a, 0x63, 0x0a, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x03, 0x2a, 0x2d, 0x0a, 0x10, 0x53, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x6f, 0x72, 0x74, 0x41, 0x73, 0x63, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x6f,
	0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x10, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x18, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x66, 0x2f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x3b,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_price_list_price_list_proto_rawDescOnce sync.Once
	file_price_list_price_list_proto_rawDescData = file_price_list_price_list_proto_rawDesc
)

func file_price_list_price_list_proto_rawDescGZIP() []byte {
	file_price_list_price_list_proto_rawDescOnce.Do(func() {
		file_price_list_price_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_list_price_list_proto_rawDescData)
	})
	return file_price_list_price_list_proto_rawDescData
}

var file_price_list_price_list_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_price_list_price_list_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_price_list_price_list_proto_goTypes = []interface{}{
	(SortingType)(0),                 // 0: SortingType
	(SortingDirection)(0),            // 1: SortingDirection
	(*UpdatePriceListRequest)(nil),   // 2: UpdatePriceListRequest
	(*UpdatePriceListResponse)(nil),  // 3: UpdatePriceListResponse
	(*ProductPrice)(nil),             // 4: ProductPrice
	(*GetProductPricesRequest)(nil),  // 5: GetProductPricesRequest
	(*GetProductPricesResponse)(nil), // 6: GetProductPricesResponse
	(*ProductPrices)(nil),            // 7: ProductPrices
}
var file_price_list_price_list_proto_depIdxs = []int32{
	4, // 0: UpdatePriceListResponse.products:type_name -> ProductPrice
	0, // 1: GetProductPricesRequest.sortingType:type_name -> SortingType
	1, // 2: GetProductPricesRequest.sortingDirection:type_name -> SortingDirection
	7, // 3: GetProductPricesResponse.products:type_name -> ProductPrices
	2, // 4: PriceList.UpdatePrices:input_type -> UpdatePriceListRequest
	5, // 5: PriceList.GetProductPrices:input_type -> GetProductPricesRequest
	3, // 6: PriceList.UpdatePrices:output_type -> UpdatePriceListResponse
	6, // 7: PriceList.GetProductPrices:output_type -> GetProductPricesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_price_list_price_list_proto_init() }
func file_price_list_price_list_proto_init() {
	if File_price_list_price_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_list_price_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceListRequest); i {
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
		file_price_list_price_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceListResponse); i {
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
		file_price_list_price_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPrice); i {
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
		file_price_list_price_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductPricesRequest); i {
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
		file_price_list_price_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductPricesResponse); i {
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
		file_price_list_price_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPrices); i {
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
			RawDescriptor: file_price_list_price_list_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_price_list_price_list_proto_goTypes,
		DependencyIndexes: file_price_list_price_list_proto_depIdxs,
		EnumInfos:         file_price_list_price_list_proto_enumTypes,
		MessageInfos:      file_price_list_price_list_proto_msgTypes,
	}.Build()
	File_price_list_price_list_proto = out.File
	file_price_list_price_list_proto_rawDesc = nil
	file_price_list_price_list_proto_goTypes = nil
	file_price_list_price_list_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PriceListClient is the client API for PriceList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PriceListClient interface {
	// Request product prices from external URL
	UpdatePrices(ctx context.Context, in *UpdatePriceListRequest, opts ...grpc.CallOption) (*UpdatePriceListResponse, error)
	// Get product prices from db
	GetProductPrices(ctx context.Context, in *GetProductPricesRequest, opts ...grpc.CallOption) (*GetProductPricesResponse, error)
}

type priceListClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceListClient(cc grpc.ClientConnInterface) PriceListClient {
	return &priceListClient{cc}
}

func (c *priceListClient) UpdatePrices(ctx context.Context, in *UpdatePriceListRequest, opts ...grpc.CallOption) (*UpdatePriceListResponse, error) {
	out := new(UpdatePriceListResponse)
	err := c.cc.Invoke(ctx, "/PriceList/UpdatePrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceListClient) GetProductPrices(ctx context.Context, in *GetProductPricesRequest, opts ...grpc.CallOption) (*GetProductPricesResponse, error) {
	out := new(GetProductPricesResponse)
	err := c.cc.Invoke(ctx, "/PriceList/GetProductPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceListServer is the server API for PriceList service.
type PriceListServer interface {
	// Request product prices from external URL
	UpdatePrices(context.Context, *UpdatePriceListRequest) (*UpdatePriceListResponse, error)
	// Get product prices from db
	GetProductPrices(context.Context, *GetProductPricesRequest) (*GetProductPricesResponse, error)
}

// UnimplementedPriceListServer can be embedded to have forward compatible implementations.
type UnimplementedPriceListServer struct {
}

func (*UnimplementedPriceListServer) UpdatePrices(context.Context, *UpdatePriceListRequest) (*UpdatePriceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrices not implemented")
}
func (*UnimplementedPriceListServer) GetProductPrices(context.Context, *GetProductPricesRequest) (*GetProductPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductPrices not implemented")
}

func RegisterPriceListServer(s *grpc.Server, srv PriceListServer) {
	s.RegisterService(&_PriceList_serviceDesc, srv)
}

func _PriceList_UpdatePrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePriceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceListServer).UpdatePrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PriceList/UpdatePrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceListServer).UpdatePrices(ctx, req.(*UpdatePriceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceList_GetProductPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceListServer).GetProductPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PriceList/GetProductPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceListServer).GetProductPrices(ctx, req.(*GetProductPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PriceList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PriceList",
	HandlerType: (*PriceListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePrices",
			Handler:    _PriceList_UpdatePrices_Handler,
		},
		{
			MethodName: "GetProductPrices",
			Handler:    _PriceList_GetProductPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price-list/price_list.proto",
}