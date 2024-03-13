package pal

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	"palworld/rpc/pal/pb/pal"
)

type GetPalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPalLogic {
	return &GetPalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPalLogic) GetPal(req *types.GetPalReq) (*types.GetPalResp, error) {
	ret := &types.GetPalResp{}
	in := &pal.GetPalReq{
		Id: req.ID,
	}
	resp, err := l.svcCtx.PalRpc.GetPal(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return ret, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = types.Pal{
		ID:           resp.Data.Id,
		Name:         resp.Data.Name,
		Number:       resp.Data.Number,
		Icon:         resp.Data.Icon,
		HP:           resp.Data.Hp,
		AttributeIds: resp.Data.AttributeIds,
		Energy:       resp.Data.Energy,
		Defensively:  resp.Data.Defensively,
	}
	return ret, nil
}
