// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/cache/cache.proto

package cache

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheClient interface {
	GetListingByTitle(ctx context.Context, in *GetListingByTitleRequest, opts ...grpc.CallOption) (*GetListingByTitleResponse, error)
	GetListingByID(ctx context.Context, in *GetListingByIDRequest, opts ...grpc.CallOption) (*GetListingByIDResponse, error)
}

type cacheClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheClient(cc grpc.ClientConnInterface) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) GetListingByTitle(ctx context.Context, in *GetListingByTitleRequest, opts ...grpc.CallOption) (*GetListingByTitleResponse, error) {
	out := new(GetListingByTitleResponse)
	err := c.cc.Invoke(ctx, "/pad.proto.Cache/GetListingByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetListingByID(ctx context.Context, in *GetListingByIDRequest, opts ...grpc.CallOption) (*GetListingByIDResponse, error) {
	out := new(GetListingByIDResponse)
	err := c.cc.Invoke(ctx, "/pad.proto.Cache/GetListingByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
// All implementations must embed UnimplementedCacheServer
// for forward compatibility
type CacheServer interface {
	GetListingByTitle(context.Context, *GetListingByTitleRequest) (*GetListingByTitleResponse, error)
	GetListingByID(context.Context, *GetListingByIDRequest) (*GetListingByIDResponse, error)
	mustEmbedUnimplementedCacheServer()
}

// UnimplementedCacheServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (UnimplementedCacheServer) GetListingByTitle(context.Context, *GetListingByTitleRequest) (*GetListingByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListingByTitle not implemented")
}
func (UnimplementedCacheServer) GetListingByID(context.Context, *GetListingByIDRequest) (*GetListingByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListingByID not implemented")
}
func (UnimplementedCacheServer) mustEmbedUnimplementedCacheServer() {}

// UnsafeCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServer will
// result in compilation errors.
type UnsafeCacheServer interface {
	mustEmbedUnimplementedCacheServer()
}

func RegisterCacheServer(s grpc.ServiceRegistrar, srv CacheServer) {
	s.RegisterService(&Cache_ServiceDesc, srv)
}

func _Cache_GetListingByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetListingByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pad.proto.Cache/GetListingByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetListingByTitle(ctx, req.(*GetListingByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetListingByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetListingByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pad.proto.Cache/GetListingByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetListingByID(ctx, req.(*GetListingByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cache_ServiceDesc is the grpc.ServiceDesc for Cache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pad.proto.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListingByTitle",
			Handler:    _Cache_GetListingByTitle_Handler,
		},
		{
			MethodName: "GetListingByID",
			Handler:    _Cache_GetListingByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cache/cache.proto",
}
