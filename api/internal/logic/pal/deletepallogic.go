package pal

import (
	"context"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePalLogic {
	return &DeletePalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePalLogic) DeletePal(req *types.DeletePalReq) (resp *types.UpdatePalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
