package logic

import (
	"context"

	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkillLogic {
	return &UpdateSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSkillLogic) UpdateSkill(in *skill.UpdateSkillReq) (*skill.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &skill.CommonResp{}, nil
}
