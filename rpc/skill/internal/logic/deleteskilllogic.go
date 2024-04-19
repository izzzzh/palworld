package logic

import (
	"context"

	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkillLogic {
	return &DeleteSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSkillLogic) DeleteSkill(in *skill.DeleteSkillReq) (*skill.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &skill.CommonResp{}, nil
}
