package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
	PalRpcConf            zrpc.RpcClientConf
	GoodsRpcConf          zrpc.RpcClientConf
	PalMateRpcConf        zrpc.RpcClientConf
	SkillRpcConf          zrpc.RpcClientConf
	TechnologyTreeRpcConf zrpc.RpcClientConf
	UserRpcConf           zrpc.RpcClientConf
}
