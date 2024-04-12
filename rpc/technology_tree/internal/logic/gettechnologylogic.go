package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/goods/pb/goods"
	"palworld/rpc/technology_tree/dao"

	"palworld/rpc/technology_tree/internal/svc"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyLogic {
	return &GetTechnologyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTechnologyLogic) GetTechnology(in *technology_tree.GetTechnologyReq) (*technology_tree.GetTechnologyResp, error) {
	if in.Id <= 0 {
		return nil, errors.New("invalid id")
	}
	p := dao.TechnologyTree
	pm := dao.TechnologyMaterial
	tech, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	tms, err := pm.WithContext(l.ctx).Where(pm.TechnologyID.Eq(tech.ID)).Select(pm.MaterialID).Find()
	if err != nil {
		return nil, err
	}
	var ids []int64
	mapMaterial := make(map[int64]int32)
	for _, val := range tms {
		ids = append(ids, val.MaterialID)
		mapMaterial[val.MaterialID] = val.Cnt
	}
	materials, err := l.GetTechMaterial(ids)
	if err != nil {
		return nil, err
	}
	for i := range materials {
		materials[i].Count = mapMaterial[materials[i].Id]
	}
	var ret = &technology_tree.GetTechnologyResp{}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = &technology_tree.Technology{
		Id:          tech.ID,
		Name:        tech.Name,
		Icon:        tech.Icon,
		Description: tech.Description,
		Cost:        tech.Cost,
		Ancient:     tech.Ancient,
		Level:       tech.Level,
		Materials:   materials,
	}

	return ret, nil
}

func (l *GetTechnologyLogic) GetTechMaterial(ids []int64) ([]*technology_tree.Material, error) {

	var ret []*technology_tree.Material

	goodsResp, err := l.svcCtx.GoodsRpc.GetGoodsByIds(l.ctx, &goods.GetGoodsByIdsReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, val := range goodsResp.Data {
		ret = append(ret, &technology_tree.Material{
			Id:    val.Id,
			Name:  val.Name,
			Image: val.Image,
		})
	}

	return ret, nil
}
