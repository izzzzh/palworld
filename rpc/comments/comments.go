package main

import (
	"flag"
	"fmt"
	"palworld/rpc/comments/dao"

	"palworld/rpc/comments/internal/config"
	"palworld/rpc/comments/internal/server"
	"palworld/rpc/comments/internal/svc"
	"palworld/rpc/comments/pb/comments"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comments.yaml", "the config file")

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
		comments.RegisterCommentsServerServer(grpcServer, server.NewCommentsServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
