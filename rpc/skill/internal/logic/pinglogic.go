package logic

import (
	"context"

	"skill/internal/svc"
	"skill/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *skill.Request) (*skill.Response, error) {
	// todo: add your logic here and delete this line

	return &skill.Response{}, nil
}
