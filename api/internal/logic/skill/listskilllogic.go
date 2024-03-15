package skill

import (
	"context"
	"net/http"
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

func (l *ListSkillLogic) ListSkill(req *types.ListSkillReq) (*types.ListSkillResp, error) {
	ret := &types.ListSkillResp{}
	in := &skill.ListSkillReq{
		AttributeId: req.AttributeID,
	}
	resp, err := l.svcCtx.SkillRpc.ListSkill(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return ret, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	skills := make([]types.Skill, 0)
	for _, val := range resp.Data {
		skills = append(skills, types.Skill{
			ID:          val.Id,
			Name:        val.Name,
			Desc:        val.Description,
			CT:          val.Ct,
			Power:       val.Power,
			AttributeID: val.AttributeId,
		})
	}
	ret.Data = skills
	return ret, nil
}
