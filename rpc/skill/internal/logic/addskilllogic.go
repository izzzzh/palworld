package logic

import (
	"context"

	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSkillLogic {
	return &AddSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSkillLogic) AddSkill(in *skill.AddSkillReq) (*skill.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &skill.CommonResp{}, nil
}
