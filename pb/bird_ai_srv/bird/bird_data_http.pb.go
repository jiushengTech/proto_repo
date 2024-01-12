// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.1
// source: bird_ai_srv/bird/bird_data.proto

package bird

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBirdDataGetBirdDistribution = "/biz.bird_ai_srv.bird.BirdData/GetBirdDistribution"
const OperationBirdDataGetBirdEvolution = "/biz.bird_ai_srv.bird.BirdData/GetBirdEvolution"
const OperationBirdDataListBirdDynamic = "/biz.bird_ai_srv.bird.BirdData/ListBirdDynamic"

type BirdDataHTTPServer interface {
	// GetBirdDistribution 获取鸟类分布
	GetBirdDistribution(context.Context, *GetBirdDistributionRequest) (*GetBirdDistributionReply, error)
	GetBirdEvolution(context.Context, *GetBirdEvolutionRequest) (*GetBirdEvolutionReply, error)
	// ListBirdDynamic 获取鸟类动态
	ListBirdDynamic(context.Context, *emptypb.Empty) (*ListBirdDynamicReply, error)
}

func RegisterBirdDataHTTPServer(s *http.Server, srv BirdDataHTTPServer) {
	r := s.Route("/")
	r.GET("yw/bird/listBirdDynamic", _BirdData_ListBirdDynamic0_HTTP_Handler(srv))
	r.GET("yw/bird/getBirdEvolution", _BirdData_GetBirdEvolution0_HTTP_Handler(srv))
	r.GET("yw/bird/getBirdDistribution", _BirdData_GetBirdDistribution0_HTTP_Handler(srv))
}

func _BirdData_ListBirdDynamic0_HTTP_Handler(srv BirdDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBirdDataListBirdDynamic)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBirdDynamic(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBirdDynamicReply)
		return ctx.Result(200, reply)
	}
}

func _BirdData_GetBirdEvolution0_HTTP_Handler(srv BirdDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBirdEvolutionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBirdDataGetBirdEvolution)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBirdEvolution(ctx, req.(*GetBirdEvolutionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBirdEvolutionReply)
		return ctx.Result(200, reply)
	}
}

func _BirdData_GetBirdDistribution0_HTTP_Handler(srv BirdDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBirdDistributionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBirdDataGetBirdDistribution)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBirdDistribution(ctx, req.(*GetBirdDistributionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBirdDistributionReply)
		return ctx.Result(200, reply)
	}
}

type BirdDataHTTPClient interface {
	GetBirdDistribution(ctx context.Context, req *GetBirdDistributionRequest, opts ...http.CallOption) (rsp *GetBirdDistributionReply, err error)
	GetBirdEvolution(ctx context.Context, req *GetBirdEvolutionRequest, opts ...http.CallOption) (rsp *GetBirdEvolutionReply, err error)
	ListBirdDynamic(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListBirdDynamicReply, err error)
}

type BirdDataHTTPClientImpl struct {
	cc *http.Client
}

func NewBirdDataHTTPClient(client *http.Client) BirdDataHTTPClient {
	return &BirdDataHTTPClientImpl{client}
}

func (c *BirdDataHTTPClientImpl) GetBirdDistribution(ctx context.Context, in *GetBirdDistributionRequest, opts ...http.CallOption) (*GetBirdDistributionReply, error) {
	var out GetBirdDistributionReply
	pattern := "yw/bird/getBirdDistribution"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBirdDataGetBirdDistribution))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BirdDataHTTPClientImpl) GetBirdEvolution(ctx context.Context, in *GetBirdEvolutionRequest, opts ...http.CallOption) (*GetBirdEvolutionReply, error) {
	var out GetBirdEvolutionReply
	pattern := "yw/bird/getBirdEvolution"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBirdDataGetBirdEvolution))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BirdDataHTTPClientImpl) ListBirdDynamic(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListBirdDynamicReply, error) {
	var out ListBirdDynamicReply
	pattern := "yw/bird/listBirdDynamic"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBirdDataListBirdDynamic))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
