package container_services

import (
	"context"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContainerLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerLogLogic {
	return &ContainerLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainerLogLogic) ContainerLog(req *types.ContainerLogReq) error {
	// todo: add your logic here and delete this line

	return nil
}
