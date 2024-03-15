// Code generated by goctl. DO NOT EDIT.
// Source: skill.proto

package server

import (
	"context"

	"palworld/rpc/skill/internal/logic"
	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"
)

type SkillServerServer struct {
	svcCtx *svc.ServiceContext
	skill.UnimplementedSkillServerServer
}

func NewSkillServerServer(svcCtx *svc.ServiceContext) *SkillServerServer {
	return &SkillServerServer{
		svcCtx: svcCtx,
	}
}

func (s *SkillServerServer) GetSkill(ctx context.Context, in *skill.GetSkillReq) (*skill.GetSkillResp, error) {
	l := logic.NewGetSkillLogic(ctx, s.svcCtx)
	return l.GetSkill(in)
}

func (s *SkillServerServer) ListSkill(ctx context.Context, in *skill.ListSkillReq) (*skill.ListSkillResp, error) {
	l := logic.NewListSkillLogic(ctx, s.svcCtx)
	return l.ListSkill(in)
}