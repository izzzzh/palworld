package main

import (
	"flag"
	"fmt"
	"palworld/rpc/skill/dao"

	"palworld/rpc/skill/internal/config"
	"palworld/rpc/skill/internal/server"
	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/skill.yaml", "the config file")

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
		skill.RegisterSkillServerServer(grpcServer, server.NewSkillServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
