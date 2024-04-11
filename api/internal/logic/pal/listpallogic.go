package pal

import (
	"context"
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

func (l *ListPalLogic) ListPal(req *types.ListPalReq) ([]*types.ListPal, error) {
	in := &pal.ListPalReq{
		Name:        req.Name,
		AttributeId: req.AttributeId,
	}
	resp, err := l.svcCtx.PalRpc.ListPal(l.ctx, in)
	if err != nil {
		return nil, err
	}
	pals := make([]*types.ListPal, 0)
	for _, val := range resp.Data {
		abilities := make([]types.Ability, 0)
		for _, ability := range val.Abilities {
			abilities = append(abilities, types.Ability{
				Icon:  ability.Icon,
				Name:  ability.Name,
				Level: ability.Level,
			})
		}
		pals = append(pals, &types.ListPal{
			ID:           val.Id,
			Number:       val.Number,
			Name:         val.Name,
			Icon:         val.Icon,
			AttributeIds: val.AttributeIds,
			Abilities:    abilities,
		})
	}
	return pals, nil
}
