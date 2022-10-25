// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: services/cache/cache.proto

package cache

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *Listing `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
	Id      uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddListingRequest) Reset() {
	*x = AddListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListingRequest) ProtoMessage() {}

func (x *AddListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListingRequest.ProtoReflect.Descriptor instead.
func (*AddListingRequest) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{0}
}

func (x *AddListingRequest) GetListing() *Listing {
	if x != nil {
		return x.Listing
	}
	return nil
}

func (x *AddListingRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetListingByTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetListingByTitleRequest) Reset() {
	*x = GetListingByTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByTitleRequest) ProtoMessage() {}

func (x *GetListingByTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingByTitleRequest.ProtoReflect.Descriptor instead.
func (*GetListingByTitleRequest) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{1}
}

func (x *GetListingByTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetListingByTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *Listing `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
}

func (x *GetListingByTitleResponse) Reset() {
	*x = GetListingByTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByTitleResponse) ProtoMessage() {}

func (x *GetListingByTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingByTitleResponse.ProtoReflect.Descriptor instead.
func (*GetListingByTitleResponse) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{2}
}

func (x *GetListingByTitleResponse) GetListing() *Listing {
	if x != nil {
		return x.Listing
	}
	return nil
}

type GetListingByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetListingByIDRequest) Reset() {
	*x = GetListingByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByIDRequest) ProtoMessage() {}

func (x *GetListingByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingByIDRequest.ProtoReflect.Descriptor instead.
func (*GetListingByIDRequest) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{3}
}

func (x *GetListingByIDRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetListingByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *Listing `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
}

func (x *GetListingByIDResponse) Reset() {
	*x = GetListingByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByIDResponse) ProtoMessage() {}

func (x *GetListingByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingByIDResponse.ProtoReflect.Descriptor instead.
func (*GetListingByIDResponse) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{4}
}

func (x *GetListingByIDResponse) GetListing() *Listing {
	if x != nil {
		return x.Listing
	}
	return nil
}

type Listing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,3,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
}

func (x *Listing) Reset() {
	*x = Listing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cache_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing) ProtoMessage() {}

func (x *Listing) ProtoReflect() protoreflect.Message {
	mi := &file_services_cache_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listing.ProtoReflect.Descriptor instead.
func (*Listing) Descriptor() ([]byte, []int) {
	return file_services_cache_cache_proto_rawDescGZIP(), []int{5}
}

func (x *Listing) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Listing) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Listing) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

var File_services_cache_cache_proto protoreflect.FileDescriptor

var file_services_cache_cache_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x70, 0x61,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x66, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0xc1, 0x03,
	0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x5c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x2e, 0x70, 0x61,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x12, 0x38, 0x2e, 0x70, 0x61,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x23, 0x5a, 0x21, 0x70, 0x61, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_cache_cache_proto_rawDescOnce sync.Once
	file_services_cache_cache_proto_rawDescData = file_services_cache_cache_proto_rawDesc
)

func file_services_cache_cache_proto_rawDescGZIP() []byte {
	file_services_cache_cache_proto_rawDescOnce.Do(func() {
		file_services_cache_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_cache_cache_proto_rawDescData)
	})
	return file_services_cache_cache_proto_rawDescData
}

var file_services_cache_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_cache_cache_proto_goTypes = []interface{}{
	(*AddListingRequest)(nil),         // 0: pad.services.cache.services.cache.AddListingRequest
	(*GetListingByTitleRequest)(nil),  // 1: pad.services.cache.services.cache.GetListingByTitleRequest
	(*GetListingByTitleResponse)(nil), // 2: pad.services.cache.services.cache.GetListingByTitleResponse
	(*GetListingByIDRequest)(nil),     // 3: pad.services.cache.services.cache.GetListingByIDRequest
	(*GetListingByIDResponse)(nil),    // 4: pad.services.cache.services.cache.GetListingByIDResponse
	(*Listing)(nil),                   // 5: pad.services.cache.services.cache.Listing
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_services_cache_cache_proto_depIdxs = []int32{
	5, // 0: pad.services.cache.services.cache.AddListingRequest.listing:type_name -> pad.services.cache.services.cache.Listing
	5, // 1: pad.services.cache.services.cache.GetListingByTitleResponse.listing:type_name -> pad.services.cache.services.cache.Listing
	5, // 2: pad.services.cache.services.cache.GetListingByIDResponse.listing:type_name -> pad.services.cache.services.cache.Listing
	0, // 3: pad.services.cache.services.cache.Cache.AddListing:input_type -> pad.services.cache.services.cache.AddListingRequest
	1, // 4: pad.services.cache.services.cache.Cache.GetListingByTitle:input_type -> pad.services.cache.services.cache.GetListingByTitleRequest
	3, // 5: pad.services.cache.services.cache.Cache.GetListingByID:input_type -> pad.services.cache.services.cache.GetListingByIDRequest
	6, // 6: pad.services.cache.services.cache.Cache.Heartbeat:input_type -> google.protobuf.Empty
	6, // 7: pad.services.cache.services.cache.Cache.AddListing:output_type -> google.protobuf.Empty
	2, // 8: pad.services.cache.services.cache.Cache.GetListingByTitle:output_type -> pad.services.cache.services.cache.GetListingByTitleResponse
	4, // 9: pad.services.cache.services.cache.Cache.GetListingByID:output_type -> pad.services.cache.services.cache.GetListingByIDResponse
	6, // 10: pad.services.cache.services.cache.Cache.Heartbeat:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_cache_cache_proto_init() }
func file_services_cache_cache_proto_init() {
	if File_services_cache_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_cache_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListingRequest); i {
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
		file_services_cache_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingByTitleRequest); i {
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
		file_services_cache_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingByTitleResponse); i {
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
		file_services_cache_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingByIDRequest); i {
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
		file_services_cache_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingByIDResponse); i {
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
		file_services_cache_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listing); i {
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
			RawDescriptor: file_services_cache_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_cache_cache_proto_goTypes,
		DependencyIndexes: file_services_cache_cache_proto_depIdxs,
		MessageInfos:      file_services_cache_cache_proto_msgTypes,
	}.Build()
	File_services_cache_cache_proto = out.File
	file_services_cache_cache_proto_rawDesc = nil
	file_services_cache_cache_proto_goTypes = nil
	file_services_cache_cache_proto_depIdxs = nil
}
