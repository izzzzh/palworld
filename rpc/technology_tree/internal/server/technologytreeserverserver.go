// Code generated by goctl. DO NOT EDIT.
// Source: technology_tree.proto

package server

import (
	"context"

	"palworld/rpc/technology_tree/internal/logic"
	"palworld/rpc/technology_tree/internal/svc"
	"palworld/rpc/technology_tree/pb/technology_tree"
)

type TechnologyTreeServerServer struct {
	svcCtx *svc.ServiceContext
	technology_tree.UnimplementedTechnologyTreeServerServer
}

func NewTechnologyTreeServerServer(svcCtx *svc.ServiceContext) *TechnologyTreeServerServer {
	return &TechnologyTreeServerServer{
		svcCtx: svcCtx,
	}
}

func (s *TechnologyTreeServerServer) GetTechnologyTree(ctx context.Context, in *technology_tree.GetTechnologyTreeReq) (*technology_tree.GetTechnologyTreeResp, error) {
	l := logic.NewGetTechnologyTreeLogic(ctx, s.svcCtx)
	return l.GetTechnologyTree(in)
}

func (s *TechnologyTreeServerServer) AddTechnologyTree(ctx context.Context, in *technology_tree.AddTechnologyTreeReq) (*technology_tree.AddTechnologyTreeResp, error) {
	l := logic.NewAddTechnologyTreeLogic(ctx, s.svcCtx)
	return l.AddTechnologyTree(in)
}

func (s *TechnologyTreeServerServer) UpdateTechnologyTree(ctx context.Context, in *technology_tree.UpdateTechnologyTreeReq) (*technology_tree.UpdateTechnologyTreeResp, error) {
	l := logic.NewUpdateTechnologyTreeLogic(ctx, s.svcCtx)
	return l.UpdateTechnologyTree(in)
}

func (s *TechnologyTreeServerServer) GetTechnology(ctx context.Context, in *technology_tree.GetTechnologyReq) (*technology_tree.GetTechnologyResp, error) {
	l := logic.NewGetTechnologyLogic(ctx, s.svcCtx)
	return l.GetTechnology(in)
}
