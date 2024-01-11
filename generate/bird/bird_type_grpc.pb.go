// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bird/bird_type.proto

package bird

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "protoRepo/generate/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BirdType_CreateBirdType_FullMethodName  = "/api.bird.BirdType/CreateBirdType"
	BirdType_UpdateBirdType_FullMethodName  = "/api.bird.BirdType/UpdateBirdType"
	BirdType_DeleteBirdType_FullMethodName  = "/api.bird.BirdType/DeleteBirdType"
	BirdType_GetBirdType_FullMethodName     = "/api.bird.BirdType/GetBirdType"
	BirdType_ListBirdType_FullMethodName    = "/api.bird.BirdType/ListBirdType"
	BirdType_GetBirdTypeTree_FullMethodName = "/api.bird.BirdType/GetBirdTypeTree"
)

// BirdTypeClient is the client API for BirdType service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BirdTypeClient interface {
	CreateBirdType(ctx context.Context, in *CreateBirdTypeRequest, opts ...grpc.CallOption) (*common.Response, error)
	UpdateBirdType(ctx context.Context, in *UpdateBirdTypeRequest, opts ...grpc.CallOption) (*UpdateBirdTypeReply, error)
	DeleteBirdType(ctx context.Context, in *DeleteBirdTypeRequest, opts ...grpc.CallOption) (*DeleteBirdTypeReply, error)
	GetBirdType(ctx context.Context, in *GetBirdTypeRequest, opts ...grpc.CallOption) (*GetBirdTypeReply, error)
	ListBirdType(ctx context.Context, in *ListBirdTypeRequest, opts ...grpc.CallOption) (*ListBirdTypeReply, error)
	GetBirdTypeTree(ctx context.Context, in *GetBirdTypeTreeRequest, opts ...grpc.CallOption) (*GetBirdTypeTreeReply, error)
}

type birdTypeClient struct {
	cc grpc.ClientConnInterface
}

func NewBirdTypeClient(cc grpc.ClientConnInterface) BirdTypeClient {
	return &birdTypeClient{cc}
}

func (c *birdTypeClient) CreateBirdType(ctx context.Context, in *CreateBirdTypeRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, BirdType_CreateBirdType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdTypeClient) UpdateBirdType(ctx context.Context, in *UpdateBirdTypeRequest, opts ...grpc.CallOption) (*UpdateBirdTypeReply, error) {
	out := new(UpdateBirdTypeReply)
	err := c.cc.Invoke(ctx, BirdType_UpdateBirdType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdTypeClient) DeleteBirdType(ctx context.Context, in *DeleteBirdTypeRequest, opts ...grpc.CallOption) (*DeleteBirdTypeReply, error) {
	out := new(DeleteBirdTypeReply)
	err := c.cc.Invoke(ctx, BirdType_DeleteBirdType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdTypeClient) GetBirdType(ctx context.Context, in *GetBirdTypeRequest, opts ...grpc.CallOption) (*GetBirdTypeReply, error) {
	out := new(GetBirdTypeReply)
	err := c.cc.Invoke(ctx, BirdType_GetBirdType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdTypeClient) ListBirdType(ctx context.Context, in *ListBirdTypeRequest, opts ...grpc.CallOption) (*ListBirdTypeReply, error) {
	out := new(ListBirdTypeReply)
	err := c.cc.Invoke(ctx, BirdType_ListBirdType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdTypeClient) GetBirdTypeTree(ctx context.Context, in *GetBirdTypeTreeRequest, opts ...grpc.CallOption) (*GetBirdTypeTreeReply, error) {
	out := new(GetBirdTypeTreeReply)
	err := c.cc.Invoke(ctx, BirdType_GetBirdTypeTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirdTypeServer is the server API for BirdType service.
// All implementations must embed UnimplementedBirdTypeServer
// for forward compatibility
type BirdTypeServer interface {
	CreateBirdType(context.Context, *CreateBirdTypeRequest) (*common.Response, error)
	UpdateBirdType(context.Context, *UpdateBirdTypeRequest) (*UpdateBirdTypeReply, error)
	DeleteBirdType(context.Context, *DeleteBirdTypeRequest) (*DeleteBirdTypeReply, error)
	GetBirdType(context.Context, *GetBirdTypeRequest) (*GetBirdTypeReply, error)
	ListBirdType(context.Context, *ListBirdTypeRequest) (*ListBirdTypeReply, error)
	GetBirdTypeTree(context.Context, *GetBirdTypeTreeRequest) (*GetBirdTypeTreeReply, error)
	mustEmbedUnimplementedBirdTypeServer()
}

// UnimplementedBirdTypeServer must be embedded to have forward compatible implementations.
type UnimplementedBirdTypeServer struct {
}

func (UnimplementedBirdTypeServer) CreateBirdType(context.Context, *CreateBirdTypeRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBirdType not implemented")
}
func (UnimplementedBirdTypeServer) UpdateBirdType(context.Context, *UpdateBirdTypeRequest) (*UpdateBirdTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBirdType not implemented")
}
func (UnimplementedBirdTypeServer) DeleteBirdType(context.Context, *DeleteBirdTypeRequest) (*DeleteBirdTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBirdType not implemented")
}
func (UnimplementedBirdTypeServer) GetBirdType(context.Context, *GetBirdTypeRequest) (*GetBirdTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBirdType not implemented")
}
func (UnimplementedBirdTypeServer) ListBirdType(context.Context, *ListBirdTypeRequest) (*ListBirdTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBirdType not implemented")
}
func (UnimplementedBirdTypeServer) GetBirdTypeTree(context.Context, *GetBirdTypeTreeRequest) (*GetBirdTypeTreeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBirdTypeTree not implemented")
}
func (UnimplementedBirdTypeServer) mustEmbedUnimplementedBirdTypeServer() {}

// UnsafeBirdTypeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BirdTypeServer will
// result in compilation errors.
type UnsafeBirdTypeServer interface {
	mustEmbedUnimplementedBirdTypeServer()
}

func RegisterBirdTypeServer(s grpc.ServiceRegistrar, srv BirdTypeServer) {
	s.RegisterService(&BirdType_ServiceDesc, srv)
}

func _BirdType_CreateBirdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBirdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).CreateBirdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_CreateBirdType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).CreateBirdType(ctx, req.(*CreateBirdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdType_UpdateBirdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBirdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).UpdateBirdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_UpdateBirdType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).UpdateBirdType(ctx, req.(*UpdateBirdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdType_DeleteBirdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBirdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).DeleteBirdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_DeleteBirdType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).DeleteBirdType(ctx, req.(*DeleteBirdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdType_GetBirdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).GetBirdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_GetBirdType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).GetBirdType(ctx, req.(*GetBirdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdType_ListBirdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBirdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).ListBirdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_ListBirdType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).ListBirdType(ctx, req.(*ListBirdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdType_GetBirdTypeTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirdTypeTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdTypeServer).GetBirdTypeTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdType_GetBirdTypeTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdTypeServer).GetBirdTypeTree(ctx, req.(*GetBirdTypeTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BirdType_ServiceDesc is the grpc.ServiceDesc for BirdType service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BirdType_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.bird.BirdType",
	HandlerType: (*BirdTypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBirdType",
			Handler:    _BirdType_CreateBirdType_Handler,
		},
		{
			MethodName: "UpdateBirdType",
			Handler:    _BirdType_UpdateBirdType_Handler,
		},
		{
			MethodName: "DeleteBirdType",
			Handler:    _BirdType_DeleteBirdType_Handler,
		},
		{
			MethodName: "GetBirdType",
			Handler:    _BirdType_GetBirdType_Handler,
		},
		{
			MethodName: "ListBirdType",
			Handler:    _BirdType_ListBirdType_Handler,
		},
		{
			MethodName: "GetBirdTypeTree",
			Handler:    _BirdType_GetBirdTypeTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bird/bird_type.proto",
}
