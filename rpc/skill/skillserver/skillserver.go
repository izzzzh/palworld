// Code generated by goctl. DO NOT EDIT.
// Source: skill.proto

package skillserver

import (
	"context"

	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSkillReq    = skill.AddSkillReq
	CommonResp     = skill.CommonResp
	DeleteSkillReq = skill.DeleteSkillReq
	GetSkillReq    = skill.GetSkillReq
	GetSkillResp   = skill.GetSkillResp
	ListSkillReq   = skill.ListSkillReq
	ListSkillResp  = skill.ListSkillResp
	Skill          = skill.Skill
	SkillPals      = skill.SkillPals
	UpdateSkillReq = skill.UpdateSkillReq

	SkillServer interface {
		GetSkill(ctx context.Context, in *GetSkillReq, opts ...grpc.CallOption) (*GetSkillResp, error)
		ListSkill(ctx context.Context, in *ListSkillReq, opts ...grpc.CallOption) (*ListSkillResp, error)
		AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*CommonResp, error)
		UpdateSkill(ctx context.Context, in *UpdateSkillReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteSkill(ctx context.Context, in *DeleteSkillReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultSkillServer struct {
		cli zrpc.Client
	}
)

func NewSkillServer(cli zrpc.Client) SkillServer {
	return &defaultSkillServer{
		cli: cli,
	}
}

func (m *defaultSkillServer) GetSkill(ctx context.Context, in *GetSkillReq, opts ...grpc.CallOption) (*GetSkillResp, error) {
	client := skill.NewSkillServerClient(m.cli.Conn())
	return client.GetSkill(ctx, in, opts...)
}

func (m *defaultSkillServer) ListSkill(ctx context.Context, in *ListSkillReq, opts ...grpc.CallOption) (*ListSkillResp, error) {
	client := skill.NewSkillServerClient(m.cli.Conn())
	return client.ListSkill(ctx, in, opts...)
}

func (m *defaultSkillServer) AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := skill.NewSkillServerClient(m.cli.Conn())
	return client.AddSkill(ctx, in, opts...)
}

func (m *defaultSkillServer) UpdateSkill(ctx context.Context, in *UpdateSkillReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := skill.NewSkillServerClient(m.cli.Conn())
	return client.UpdateSkill(ctx, in, opts...)
}

func (m *defaultSkillServer) DeleteSkill(ctx context.Context, in *DeleteSkillReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := skill.NewSkillServerClient(m.cli.Conn())
	return client.DeleteSkill(ctx, in, opts...)
}
