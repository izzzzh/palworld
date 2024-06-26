package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"palworld/api/internal/config"
	"palworld/rpc/comments/commentsserver"
	"palworld/rpc/goods/goodsserver"
	"palworld/rpc/pal/palserver"
	"palworld/rpc/pal_mate/palmateserver"
	"palworld/rpc/skill/skillserver"
	"palworld/rpc/technology_tree/technologytreeserver"
	"palworld/rpc/user/userserver"
)

type ServiceContext struct {
	Config            config.Config
	PalRpc            palserver.PalServer
	GoodsRpc          goodsserver.GoodsServer
	PalMateRpc        palmateserver.PalMateServer
	SkillRpc          skillserver.SkillServer
	TechnologyTreeRpc technologytreeserver.TechnologyTreeServer
	UserRpc           userserver.UserServer
	CommentsRpc       commentsserver.CommentsServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		PalRpc:            palserver.NewPalServer(zrpc.MustNewClient(c.PalRpcConf)),
		GoodsRpc:          goodsserver.NewGoodsServer(zrpc.MustNewClient(c.GoodsRpcConf)),
		PalMateRpc:        palmateserver.NewPalMateServer(zrpc.MustNewClient(c.PalMateRpcConf)),
		SkillRpc:          skillserver.NewSkillServer(zrpc.MustNewClient(c.SkillRpcConf)),
		TechnologyTreeRpc: technologytreeserver.NewTechnologyTreeServer(zrpc.MustNewClient(c.TechnologyTreeRpcConf)),
		UserRpc:           userserver.NewUserServer(zrpc.MustNewClient(c.UserRpcConf)),
		CommentsRpc:       commentsserver.NewCommentsServer(zrpc.MustNewClient(c.CommentsRpcConf)),
	}
}
