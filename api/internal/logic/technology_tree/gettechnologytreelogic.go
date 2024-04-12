package technology_tree

import (
	"context"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTechnologyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyTreeLogic {
	return &GetTechnologyTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTechnologyTreeLogic) GetTechnologyTree() ([]*types.TechnologyTree, error) {

	in := &technology_tree.GetTechnologyTreeReq{}
	resp, err := l.svcCtx.TechnologyTreeRpc.GetTechnologyTree(l.ctx, in)
	if err != nil {
		return nil, err
	}
	retData := make([]*types.TechnologyTree, 0)
	for _, val := range resp.Data {
		technologyList := make([]types.Technology, 0)
		for _, technology := range val.Data {
			materials := make([]types.TechnologyMaterial, 0)
			for _, material := range technology.Materials {
				materials = append(materials, types.TechnologyMaterial{
					ID:    material.Id,
					Name:  material.Name,
					Image: material.Image,
					Count: material.Count,
				})
			}

			technologyList = append(technologyList, types.Technology{
				ID:          technology.Id,
				Name:        technology.Name,
				Description: technology.Description,
				Cost:        technology.Cost,
				Icon:        technology.Icon,
				Ancient:     technology.Ancient,
				Material:    materials,
				Level:       technology.Level,
			})
		}
		retData = append(retData, &types.TechnologyTree{
			Level: val.Level,
			Data:  technologyList,
		})
	}

	return retData, nil
}
