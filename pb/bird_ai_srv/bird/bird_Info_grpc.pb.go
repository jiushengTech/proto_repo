// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bird_ai_srv/bird/bird_Info.proto

package bird

import (
	context "context"
	common "github.com/samsaralc/proto_repo/pb/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BirdInfo_CreateBird_FullMethodName          = "/biz.bird_ai_srv.bird.BirdInfo/CreateBird"
	BirdInfo_UpdateBird_FullMethodName          = "/biz.bird_ai_srv.bird.BirdInfo/UpdateBird"
	BirdInfo_DeleteBird_FullMethodName          = "/biz.bird_ai_srv.bird.BirdInfo/DeleteBird"
	BirdInfo_GetBird_FullMethodName             = "/biz.bird_ai_srv.bird.BirdInfo/GetBird"
	BirdInfo_GetBirdDetailById_FullMethodName   = "/biz.bird_ai_srv.bird.BirdInfo/GetBirdDetailById"
	BirdInfo_ListBirdOrder_FullMethodName       = "/biz.bird_ai_srv.bird.BirdInfo/ListBirdOrder"
	BirdInfo_ListBirdBaseByOrder_FullMethodName = "/biz.bird_ai_srv.bird.BirdInfo/ListBirdBaseByOrder"
)

// BirdInfoClient is the client API for BirdInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BirdInfoClient interface {
	CreateBird(ctx context.Context, in *CreateBirdRequest, opts ...grpc.CallOption) (*CreateBirdReply, error)
	UpdateBird(ctx context.Context, in *UpdateBirdRequest, opts ...grpc.CallOption) (*UpdateBirdReply, error)
	DeleteBird(ctx context.Context, in *DeleteBirdRequest, opts ...grpc.CallOption) (*DeleteBirdReply, error)
	GetBird(ctx context.Context, in *GetBirdRequest, opts ...grpc.CallOption) (*GetBirdReply, error)
	GetBirdDetailById(ctx context.Context, in *GetBirdDetailByIdRequest, opts ...grpc.CallOption) (*GetBirdDetailByIdReply, error)
	ListBirdOrder(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBirdOrderReply, error)
	ListBirdBaseByOrder(ctx context.Context, in *ListBirdBaseByOrderRequest, opts ...grpc.CallOption) (*common.PageResult, error)
}

type birdInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewBirdInfoClient(cc grpc.ClientConnInterface) BirdInfoClient {
	return &birdInfoClient{cc}
}

func (c *birdInfoClient) CreateBird(ctx context.Context, in *CreateBirdRequest, opts ...grpc.CallOption) (*CreateBirdReply, error) {
	out := new(CreateBirdReply)
	err := c.cc.Invoke(ctx, BirdInfo_CreateBird_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) UpdateBird(ctx context.Context, in *UpdateBirdRequest, opts ...grpc.CallOption) (*UpdateBirdReply, error) {
	out := new(UpdateBirdReply)
	err := c.cc.Invoke(ctx, BirdInfo_UpdateBird_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) DeleteBird(ctx context.Context, in *DeleteBirdRequest, opts ...grpc.CallOption) (*DeleteBirdReply, error) {
	out := new(DeleteBirdReply)
	err := c.cc.Invoke(ctx, BirdInfo_DeleteBird_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) GetBird(ctx context.Context, in *GetBirdRequest, opts ...grpc.CallOption) (*GetBirdReply, error) {
	out := new(GetBirdReply)
	err := c.cc.Invoke(ctx, BirdInfo_GetBird_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) GetBirdDetailById(ctx context.Context, in *GetBirdDetailByIdRequest, opts ...grpc.CallOption) (*GetBirdDetailByIdReply, error) {
	out := new(GetBirdDetailByIdReply)
	err := c.cc.Invoke(ctx, BirdInfo_GetBirdDetailById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) ListBirdOrder(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBirdOrderReply, error) {
	out := new(ListBirdOrderReply)
	err := c.cc.Invoke(ctx, BirdInfo_ListBirdOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdInfoClient) ListBirdBaseByOrder(ctx context.Context, in *ListBirdBaseByOrderRequest, opts ...grpc.CallOption) (*common.PageResult, error) {
	out := new(common.PageResult)
	err := c.cc.Invoke(ctx, BirdInfo_ListBirdBaseByOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirdInfoServer is the server API for BirdInfo service.
// All implementations must embed UnimplementedBirdInfoServer
// for forward compatibility
type BirdInfoServer interface {
	CreateBird(context.Context, *CreateBirdRequest) (*CreateBirdReply, error)
	UpdateBird(context.Context, *UpdateBirdRequest) (*UpdateBirdReply, error)
	DeleteBird(context.Context, *DeleteBirdRequest) (*DeleteBirdReply, error)
	GetBird(context.Context, *GetBirdRequest) (*GetBirdReply, error)
	GetBirdDetailById(context.Context, *GetBirdDetailByIdRequest) (*GetBirdDetailByIdReply, error)
	ListBirdOrder(context.Context, *emptypb.Empty) (*ListBirdOrderReply, error)
	ListBirdBaseByOrder(context.Context, *ListBirdBaseByOrderRequest) (*common.PageResult, error)
	mustEmbedUnimplementedBirdInfoServer()
}

// UnimplementedBirdInfoServer must be embedded to have forward compatible implementations.
type UnimplementedBirdInfoServer struct {
}

func (UnimplementedBirdInfoServer) CreateBird(context.Context, *CreateBirdRequest) (*CreateBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBird not implemented")
}
func (UnimplementedBirdInfoServer) UpdateBird(context.Context, *UpdateBirdRequest) (*UpdateBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBird not implemented")
}
func (UnimplementedBirdInfoServer) DeleteBird(context.Context, *DeleteBirdRequest) (*DeleteBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBird not implemented")
}
func (UnimplementedBirdInfoServer) GetBird(context.Context, *GetBirdRequest) (*GetBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBird not implemented")
}
func (UnimplementedBirdInfoServer) GetBirdDetailById(context.Context, *GetBirdDetailByIdRequest) (*GetBirdDetailByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBirdDetailById not implemented")
}
func (UnimplementedBirdInfoServer) ListBirdOrder(context.Context, *emptypb.Empty) (*ListBirdOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBirdOrder not implemented")
}
func (UnimplementedBirdInfoServer) ListBirdBaseByOrder(context.Context, *ListBirdBaseByOrderRequest) (*common.PageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBirdBaseByOrder not implemented")
}
func (UnimplementedBirdInfoServer) mustEmbedUnimplementedBirdInfoServer() {}

// UnsafeBirdInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BirdInfoServer will
// result in compilation errors.
type UnsafeBirdInfoServer interface {
	mustEmbedUnimplementedBirdInfoServer()
}

func RegisterBirdInfoServer(s grpc.ServiceRegistrar, srv BirdInfoServer) {
	s.RegisterService(&BirdInfo_ServiceDesc, srv)
}

func _BirdInfo_CreateBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).CreateBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_CreateBird_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).CreateBird(ctx, req.(*CreateBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_UpdateBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).UpdateBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_UpdateBird_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).UpdateBird(ctx, req.(*UpdateBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_DeleteBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).DeleteBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_DeleteBird_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).DeleteBird(ctx, req.(*DeleteBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_GetBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).GetBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_GetBird_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).GetBird(ctx, req.(*GetBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_GetBirdDetailById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirdDetailByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).GetBirdDetailById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_GetBirdDetailById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).GetBirdDetailById(ctx, req.(*GetBirdDetailByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_ListBirdOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).ListBirdOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_ListBirdOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).ListBirdOrder(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdInfo_ListBirdBaseByOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBirdBaseByOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdInfoServer).ListBirdBaseByOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BirdInfo_ListBirdBaseByOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdInfoServer).ListBirdBaseByOrder(ctx, req.(*ListBirdBaseByOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BirdInfo_ServiceDesc is the grpc.ServiceDesc for BirdInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BirdInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biz.bird_ai_srv.bird.BirdInfo",
	HandlerType: (*BirdInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBird",
			Handler:    _BirdInfo_CreateBird_Handler,
		},
		{
			MethodName: "UpdateBird",
			Handler:    _BirdInfo_UpdateBird_Handler,
		},
		{
			MethodName: "DeleteBird",
			Handler:    _BirdInfo_DeleteBird_Handler,
		},
		{
			MethodName: "GetBird",
			Handler:    _BirdInfo_GetBird_Handler,
		},
		{
			MethodName: "GetBirdDetailById",
			Handler:    _BirdInfo_GetBirdDetailById_Handler,
		},
		{
			MethodName: "ListBirdOrder",
			Handler:    _BirdInfo_ListBirdOrder_Handler,
		},
		{
			MethodName: "ListBirdBaseByOrder",
			Handler:    _BirdInfo_ListBirdBaseByOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bird_ai_srv/bird/bird_Info.proto",
}
