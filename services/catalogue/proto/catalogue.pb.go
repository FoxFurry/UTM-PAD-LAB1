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

type GetListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing []*Listing `protobuf:"bytes,1,rep,name=listing,proto3" json:"listing,omitempty"`
}

func (x *GetListingsResponse) Reset() {
	*x = GetListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalogue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingsResponse) ProtoMessage() {}

func (x *GetListingsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetListingsResponse.ProtoReflect.Descriptor instead.
func (*GetListingsResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalogue_proto_rawDescGZIP(), []int{1}
}

func (x *GetListingsResponse) GetListing() []*Listing {
	if x != nil {
		return x.Listing
	}
	return nil
}

type Listing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceHolder int32 `protobuf:"varint,1,opt,name=place_holder,json=placeHolder,proto3" json:"place_holder,omitempty"`
}

func (x *Listing) Reset() {
	*x = Listing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalogue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing) ProtoMessage() {}

func (x *Listing) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Listing.ProtoReflect.Descriptor instead.
func (*Listing) Descriptor() ([]byte, []int) {
	return file_proto_catalogue_proto_rawDescGZIP(), []int{2}
}

func (x *Listing) GetPlaceHolder() int32 {
	if x != nil {
		return x.PlaceHolder
	}
	return 0
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
	0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x2c, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x32, 0xb9,
	0x01, 0x0a, 0x09, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x70, 0x61, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x70, 0x61,
	0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_catalogue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_catalogue_proto_goTypes = []interface{}{
	(*AddListingRequest)(nil),   // 0: pad.services.catalogue.proto.AddListingRequest
	(*GetListingsResponse)(nil), // 1: pad.services.catalogue.proto.GetListingsResponse
	(*Listing)(nil),             // 2: pad.services.catalogue.proto.Listing
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_proto_catalogue_proto_depIdxs = []int32{
	2, // 0: pad.services.catalogue.proto.AddListingRequest.listing:type_name -> pad.services.catalogue.proto.Listing
	2, // 1: pad.services.catalogue.proto.GetListingsResponse.listing:type_name -> pad.services.catalogue.proto.Listing
	2, // 2: pad.services.catalogue.proto.Catalogue.AddListing:input_type -> pad.services.catalogue.proto.Listing
	3, // 3: pad.services.catalogue.proto.Catalogue.GetAllListings:input_type -> google.protobuf.Empty
	3, // 4: pad.services.catalogue.proto.Catalogue.AddListing:output_type -> google.protobuf.Empty
	1, // 5: pad.services.catalogue.proto.Catalogue.GetAllListings:output_type -> pad.services.catalogue.proto.GetListingsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			switch v := v.(*GetListingsResponse); i {
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
			RawDescriptor: file_proto_catalogue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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