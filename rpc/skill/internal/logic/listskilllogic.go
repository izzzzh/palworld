package logic

import (
	"context"
	"net/http"
	"palworld/rpc/skill/dao"
	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSkillLogic {
	return &ListSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSkillLogic) ListSkill(in *skill.ListSkillReq) (*skill.ListSkillResp, error) {
	ret := &skill.ListSkillResp{}
	p := dao.Skill
	pd := p.WithContext(l.ctx)
	if in.AttributeId > 0 {
		pd.Where(p.AttributeID.Eq(in.AttributeId))
	}
	skills, err := pd.Find()
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	listSkill := make([]*skill.Skill, 0)
	for _, val := range skills {
		listSkill = append(listSkill, &skill.Skill{
			Id:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Ct:          val.Ct,
			Power:       val.Power,
			AttributeId: val.AttributeID,
		})
	}
	ret.Data = listSkill
	return ret, nil
}
