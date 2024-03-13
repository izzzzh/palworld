package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"palworld/rpc/pal/palserver"
	"palworld/rpc/pal_mate/internal/config"
)

type ServiceContext struct {
	Config config.Config
	PalRpc palserver.PalServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PalRpc: palserver.NewPalServer(zrpc.MustNewClient(c.PalRpcConf)),
	}
}
