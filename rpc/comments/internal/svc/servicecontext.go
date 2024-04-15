package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"palworld/rpc/comments/internal/config"
	"palworld/rpc/user/userserver"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userserver.UserServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userserver.NewUserServer(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
