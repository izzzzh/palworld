package pal

import (
	"context"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePalLogic {
	return &UpdatePalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePalLogic) UpdatePal(req *types.UpdatePalReq) (resp *types.UpdatePalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
