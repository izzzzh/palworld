package pal

import (
	"context"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePalLogic {
	return &CreatePalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePalLogic) CreatePal(req *types.CreatePalReq) (resp *types.CreatePalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
