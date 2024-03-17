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
	abilities := make([]types.Ability, 0)
	for _, val := range resp.Data.Abilities {
		abilities = append(abilities, types.Ability{
			Icon:  val.Icon,
			Name:  val.Name,
			Level: val.Level,
		})
	}
	ret.Data = types.Pal{
		ID:           resp.Data.Id,
		Name:         resp.Data.Name,
		Number:       resp.Data.Number,
		Icon:         resp.Data.Icon,
		AttributeIds: resp.Data.AttributeIds,
		HP:           resp.Data.Hp,
		Energy:       resp.Data.Energy,
		Defensively:  resp.Data.Defensively,
		Abilities:    abilities,
		Eat:          resp.Data.Eat,
	}
	return ret, nil
}
