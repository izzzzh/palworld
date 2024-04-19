package skill

import (
	"context"
	"palworld/rpc/skill/pb/skill"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSkillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSkillLogic {
	return &ListSkillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSkillLogic) ListSkill() ([]*types.SkillInfo, error) {
	in := &skill.ListSkillReq{}
	resp, err := l.svcCtx.SkillRpc.ListSkill(l.ctx, in)
	if err != nil {
		return nil, err
	}
	ret := make([]*types.SkillInfo, 0)
	for _, val := range resp.Data {
		pals := make([]types.SkillPal, 0)
		for _, pal := range val.Pals {
			pals = append(pals, types.SkillPal{
				ID:    pal.Id,
				Name:  pal.Name,
				Image: pal.Image,
			})
		}

		ret = append(ret, &types.SkillInfo{
			ID:          val.Id,
			Name:        val.Name,
			Desc:        val.Description,
			CT:          val.Ct,
			Power:       val.Power,
			AttributeID: val.AttributeId,
			Pals:        pals,
		})
	}
	return ret, nil
}
