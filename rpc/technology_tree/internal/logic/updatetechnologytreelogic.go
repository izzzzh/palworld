package logic

import (
	"context"

	"palworld/rpc/technology_tree/internal/svc"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTechnologyTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTechnologyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTechnologyTreeLogic {
	return &UpdateTechnologyTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTechnologyTreeLogic) UpdateTechnologyTree(in *technology_tree.UpdateTechnologyTreeReq) (*technology_tree.UpdateTechnologyTreeResp, error) {
	// todo: add your logic here and delete this line

	return &technology_tree.UpdateTechnologyTreeResp{}, nil
}
