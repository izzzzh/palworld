package skill

import (
	"context"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSkillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSkillLogic {
	return &AddSkillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSkillLogic) AddSkill(req *types.AddSkillReq) error {
	// todo: add your logic here and delete this line

	return nil
}
