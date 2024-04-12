// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package goodsserver

import (
	"context"

	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetGoodsByIdsReq  = goods.GetGoodsByIdsReq
	GetGoodsByIdsResp = goods.GetGoodsByIdsResp
	GetGoodsReq       = goods.GetGoodsReq
	GetGoodsResp      = goods.GetGoodsResp
	Goods             = goods.Goods
	ListGoodsReq      = goods.ListGoodsReq
	ListGoodsResp     = goods.ListGoodsResp
	Material          = goods.Material

	GoodsServer interface {
		ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsResp, error)
		GetGoods(ctx context.Context, in *GetGoodsReq, opts ...grpc.CallOption) (*GetGoodsResp, error)
		GetGoodsByIds(ctx context.Context, in *GetGoodsByIdsReq, opts ...grpc.CallOption) (*GetGoodsByIdsResp, error)
	}

	defaultGoodsServer struct {
		cli zrpc.Client
	}
)

func NewGoodsServer(cli zrpc.Client) GoodsServer {
	return &defaultGoodsServer{
		cli: cli,
	}
}

func (m *defaultGoodsServer) ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsResp, error) {
	client := goods.NewGoodsServerClient(m.cli.Conn())
	return client.ListGoods(ctx, in, opts...)
}

func (m *defaultGoodsServer) GetGoods(ctx context.Context, in *GetGoodsReq, opts ...grpc.CallOption) (*GetGoodsResp, error) {
	client := goods.NewGoodsServerClient(m.cli.Conn())
	return client.GetGoods(ctx, in, opts...)
}

func (m *defaultGoodsServer) GetGoodsByIds(ctx context.Context, in *GetGoodsByIdsReq, opts ...grpc.CallOption) (*GetGoodsByIdsResp, error) {
	client := goods.NewGoodsServerClient(m.cli.Conn())
	return client.GetGoodsByIds(ctx, in, opts...)
}
