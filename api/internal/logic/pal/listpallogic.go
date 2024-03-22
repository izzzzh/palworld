package pal

import (
	"context"
	"net/http"
	"palworld/rpc/pal/pb/pal"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPalLogic {
	return &ListPalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPalLogic) ListPal(req *types.ListPalReq) (*types.ListPalResp, error) {
	ret := &types.ListPalResp{}
	in := &pal.ListPalReq{
		Name:        req.Name,
		AttributeId: req.AttributeId,
	}
	resp, err := l.svcCtx.PalRpc.ListPal(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return ret, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	pals := make([]types.ListPal, 0)
	for _, val := range resp.Data {
		pals = append(pals, types.ListPal{
			ID:           val.Id,
			Number:       val.Number,
			Name:         val.Name,
			Icon:         val.Icon,
			AttributeIds: val.AttributeIds,
		})
	}
	ret.Data = pals
	return ret, nil
}
