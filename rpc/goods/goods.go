package main

import (
	"flag"
	"fmt"
	"palworld/rpc/goods/dao"

	"palworld/rpc/goods/internal/config"
	"palworld/rpc/goods/internal/server"
	"palworld/rpc/goods/internal/svc"
	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	err := dao.InitDB(c.DataSource)
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServerServer(grpcServer, server.NewGoodsServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
