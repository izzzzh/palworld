package logic

import (
	"context"

	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkillLogic {
	return &GetSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkillLogic) GetSkill(in *skill.GetSkillReq) (*skill.GetSkillResp, error) {
	// todo: add your logic here and delete this line

	return &skill.GetSkillResp{}, nil
}
