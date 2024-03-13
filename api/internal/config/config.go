package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	PalRpcConf     zrpc.RpcClientConf
	GoodsRpcConf   zrpc.RpcClientConf
	PalMateRpcConf zrpc.RpcClientConf
}
