package technology_tree

import (
	"context"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTechnologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTechnologyLogic {
	return &AddTechnologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTechnologyLogic) AddTechnology(req *types.AddTechnologyReq) error {

	material := make([]*technology_tree.TechnologyMaterial, 0)
	for _, v := range req.Material {
		material = append(material, &technology_tree.TechnologyMaterial{
			Id:    v.ID,
			Count: v.Count,
		})
	}
	in := &technology_tree.AddTechnologyTreeReq{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
		Ancient:     req.Ancient,
		Cost:        req.Cost,
		Level:       req.Level,
		Materials:   material,
	}
	_, err := l.svcCtx.TechnologyTreeRpc.AddTechnologyTree(l.ctx, in)

	if err != nil {
		return err
	}

	return nil
}
