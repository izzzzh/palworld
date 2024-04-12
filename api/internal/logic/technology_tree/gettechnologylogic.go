package technology_tree

import (
	"context"
	"errors"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyLogic {
	return &GetTechnologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTechnologyLogic) GetTechnology(req *types.GetTechnologyReq) (*types.Technology, error) {
	if req.ID <= 0 {
		return nil, errors.New("参数id无效")
	}
	in := &technology_tree.GetTechnologyReq{
		Id: req.ID,
	}

	resp, err := l.svcCtx.TechnologyTreeRpc.GetTechnology(l.ctx, in)
	if err != nil {
		return nil, err
	}
	materials := make([]types.TechnologyMaterial, 0)

	for _, v := range resp.Data.Materials {
		materials = append(materials, types.TechnologyMaterial{
			ID:    v.Id,
			Name:  v.Name,
			Image: v.Image,
			Count: v.Count,
		})
	}

	var ret = &types.Technology{
		ID:          resp.Data.Id,
		Name:        resp.Data.Name,
		Icon:        resp.Data.Icon,
		Level:       resp.Data.Level,
		Description: resp.Data.Description,
		Cost:        resp.Data.Cost,
		Ancient:     resp.Data.Ancient,
		Material:    materials,
	}

	return ret, nil
}
