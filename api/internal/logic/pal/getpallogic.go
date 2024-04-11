package pal

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetPalLogic) GetPal(req *types.GetPalReq) (*types.Pal, error) {
	in := &pal.GetPalReq{
		Id: req.ID,
	}
	resp, err := l.svcCtx.PalRpc.GetPal(l.ctx, in)
	if err != nil {
		return nil, err
	}
	abilities := make([]types.Ability, 0)
	for _, val := range resp.Data.Abilities {
		abilities = append(abilities, types.Ability{
			Icon:  val.Icon,
			Name:  val.Name,
			Level: val.Level,
		})
	}
	skills := make([]types.Skill, 0)
	for _, val := range resp.Data.Skills {
		skills = append(skills, types.Skill{
			ID:          val.Id,
			Name:        val.Name,
			Desc:        val.Description,
			CT:          val.Ct,
			Power:       val.Power,
			AttributeID: val.AttributeId,
		})
	}
	ret := &types.Pal{
		ID:               resp.Data.Id,
		Name:             resp.Data.Name,
		Number:           resp.Data.Number,
		Icon:             resp.Data.Icon,
		Description:      resp.Data.Description,
		AttributeIds:     resp.Data.AttributeIds,
		HP:               resp.Data.Hp,
		Energy:           resp.Data.Energy,
		Defensively:      resp.Data.Defensively,
		Abilities:        abilities,
		Eat:              resp.Data.Eat,
		PassiveSkill:     resp.Data.Passive,
		PassiveSkillDesc: resp.Data.PassiveDesc,
		ActiveSkills:     skills,
	}
	return ret, nil
}
