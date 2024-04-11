package goods

import (
	"context"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGoodsLogic {
	return &ListGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListGoodsLogic) ListGoods(req *types.ListGoodsreq) ([]*types.ListGoods, error) {
	ret := make([]*types.ListGoods, 0)
	in := &goods.ListGoodsReq{
		Name:     req.Name,
		Types:    req.Types,
		Page:     req.Page,
		PageSize: req.PageSize,
		Quality:  req.Quality,
	}
	resp, err := l.svcCtx.GoodsRpc.ListGoods(l.ctx, in)
	if err != nil {
		return nil, err
	}
	goodsResult := make([]types.ListGoods, 0)
	for _, val := range resp.Data {
		materials := make([]types.Material, 0)
		for _, material := range val.Materials {
			materials = append(materials, types.Material{
				ID:    material.Id,
				Name:  material.Name,
				Image: material.Image,
				Count: material.Count,
			})
		}
		goodsResult = append(goodsResult, types.ListGoods{
			ID:          val.Id,
			Name:        val.Name,
			Image:       val.Image,
			Description: val.Description,
			Workload:    val.Workload,
			Quality:     val.Quality,
			Types:       val.Types,
			Materials:   materials,
		})
	}
	return ret, nil
}
