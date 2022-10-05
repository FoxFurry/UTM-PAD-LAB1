// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: proto/catalogue.proto

package catalogue

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
}

func (x *AddListingRequest) Reset() {
	*x = AddListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalogue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListingRequest) ProtoMessage() {}

func (x *AddListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[0]
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
	return file_proto_catalogue_proto_rawDescGZIP(), []int{0}
}

func (x *AddListingRequest) GetListing() *Listing {
	if x != nil {
		return x.Listing
	}
	return nil
}

type GetAllListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings []*Listing `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
}

func (x *GetAllListingsResponse) Reset() {
	*x = GetAllListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalogue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllListingsResponse) ProtoMessage() {}

func (x *GetAllListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllListingsResponse.ProtoReflect.Descriptor instead.
func (*GetAllListingsResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalogue_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllListingsResponse) GetListings() []*Listing {
	if x != nil {
		return x.Listings
	}
	return nil
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
		mi := &file_proto_catalogue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByTitleRequest) ProtoMessage() {}

func (x *GetListingByTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[2]
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
	return file_proto_catalogue_proto_rawDescGZIP(), []int{2}
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
		mi := &file_proto_catalogue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingByTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingByTitleResponse) ProtoMessage() {}

func (x *GetListingByTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[3]
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
	return file_proto_catalogue_proto_rawDescGZIP(), []int{3}
}

func (x *GetListingByTitleResponse) GetListing() *Listing {
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
		mi := &file_proto_catalogue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing) ProtoMessage() {}

func (x *Listing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[4]
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
	return file_proto_catalogue_proto_rawDescGZIP(), []int{4}
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

type ListingDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,3,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	AuthorName   string `protobuf:"bytes,4,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	AuthorPhone  string `protobuf:"bytes,5,opt,name=author_phone,json=authorPhone,proto3" json:"author_phone,omitempty"`
}

func (x *ListingDetails) Reset() {
	*x = ListingDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalogue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListingDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListingDetails) ProtoMessage() {}

func (x *ListingDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalogue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListingDetails.ProtoReflect.Descriptor instead.
func (*ListingDetails) Descriptor() ([]byte, []int) {
	return file_proto_catalogue_proto_rawDescGZIP(), []int{5}
}

func (x *ListingDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListingDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListingDetails) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *ListingDetails) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *ListingDetails) GetAuthorPhone() string {
	if x != nil {
		return x.AuthorPhone
	}
	return ""
}

var File_proto_catalogue_proto protoreflect.FileDescriptor

var file_proto_catalogue_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x66, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0xb1, 0x01,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x32, 0xcf, 0x02, 0x0a, 0x09, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x12,
	0x57, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x2e,
	0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x34, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x36, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x70, 0x61, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_catalogue_proto_rawDescOnce sync.Once
	file_proto_catalogue_proto_rawDescData = file_proto_catalogue_proto_rawDesc
)

func file_proto_catalogue_proto_rawDescGZIP() []byte {
	file_proto_catalogue_proto_rawDescOnce.Do(func() {
		file_proto_catalogue_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_catalogue_proto_rawDescData)
	})
	return file_proto_catalogue_proto_rawDescData
}

var file_proto_catalogue_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_catalogue_proto_goTypes = []interface{}{
	(*AddListingRequest)(nil),         // 0: pad.services.catalogue.proto.AddListingRequest
	(*GetAllListingsResponse)(nil),    // 1: pad.services.catalogue.proto.GetAllListingsResponse
	(*GetListingByTitleRequest)(nil),  // 2: pad.services.catalogue.proto.GetListingByTitleRequest
	(*GetListingByTitleResponse)(nil), // 3: pad.services.catalogue.proto.GetListingByTitleResponse
	(*Listing)(nil),                   // 4: pad.services.catalogue.proto.Listing
	(*ListingDetails)(nil),            // 5: pad.services.catalogue.proto.ListingDetails
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_proto_catalogue_proto_depIdxs = []int32{
	4, // 0: pad.services.catalogue.proto.AddListingRequest.listing:type_name -> pad.services.catalogue.proto.Listing
	4, // 1: pad.services.catalogue.proto.GetAllListingsResponse.listings:type_name -> pad.services.catalogue.proto.Listing
	4, // 2: pad.services.catalogue.proto.GetListingByTitleResponse.listing:type_name -> pad.services.catalogue.proto.Listing
	0, // 3: pad.services.catalogue.proto.Catalogue.AddListing:input_type -> pad.services.catalogue.proto.AddListingRequest
	6, // 4: pad.services.catalogue.proto.Catalogue.GetAllListings:input_type -> google.protobuf.Empty
	2, // 5: pad.services.catalogue.proto.Catalogue.GetListingByTitle:input_type -> pad.services.catalogue.proto.GetListingByTitleRequest
	6, // 6: pad.services.catalogue.proto.Catalogue.AddListing:output_type -> google.protobuf.Empty
	1, // 7: pad.services.catalogue.proto.Catalogue.GetAllListings:output_type -> pad.services.catalogue.proto.GetAllListingsResponse
	3, // 8: pad.services.catalogue.proto.Catalogue.GetListingByTitle:output_type -> pad.services.catalogue.proto.GetListingByTitleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_catalogue_proto_init() }
func file_proto_catalogue_proto_init() {
	if File_proto_catalogue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_catalogue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_catalogue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllListingsResponse); i {
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
		file_proto_catalogue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_catalogue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_catalogue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_catalogue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListingDetails); i {
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
			RawDescriptor: file_proto_catalogue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_catalogue_proto_goTypes,
		DependencyIndexes: file_proto_catalogue_proto_depIdxs,
		MessageInfos:      file_proto_catalogue_proto_msgTypes,
	}.Build()
	File_proto_catalogue_proto = out.File
	file_proto_catalogue_proto_rawDesc = nil
	file_proto_catalogue_proto_goTypes = nil
	file_proto_catalogue_proto_depIdxs = nil
}
