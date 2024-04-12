package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"palworld/rpc/goods/goodsserver"
	"palworld/rpc/technology_tree/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goodsserver.GoodsServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goodsserver.NewGoodsServer(zrpc.MustNewClient(c.GoodsRpcConf)),
	}
}
