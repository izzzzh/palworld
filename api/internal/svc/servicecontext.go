package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"palworld/api/internal/config"
	"palworld/rpc/goods/goodsserver"
	"palworld/rpc/pal/palserver"
	"palworld/rpc/pal_mate/palmateserver"
)

type ServiceContext struct {
	Config     config.Config
	PalRpc     palserver.PalServer
	GoodsRpc   goodsserver.GoodsServer
	PalMateRpc palmateserver.PalMateServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		PalRpc:     palserver.NewPalServer(zrpc.MustNewClient(c.PalRpcConf)),
		GoodsRpc:   goodsserver.NewGoodsServer(zrpc.MustNewClient(c.GoodsRpcConf)),
		PalMateRpc: palmateserver.NewPalMateServer(zrpc.MustNewClient(c.PalMateRpcConf)),
	}
}
