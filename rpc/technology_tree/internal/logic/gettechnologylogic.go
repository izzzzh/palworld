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
	tech, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}
	materials, err := l.GetTechMaterial(tech.ID)
	if err != nil {
		return nil, err
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

func (l *GetTechnologyLogic) GetTechMaterial(id int64) ([]*technology_tree.Material, error) {
	p := dao.TechnologyMaterial
	techMap, err := p.WithContext(l.ctx).Where(p.TechnologyID.Eq(id)).Select(p.MaterialID).Find()
	if err != nil {
		return nil, err
	}

	var ret []*technology_tree.Material
	for _, val := range techMap {
		goodsResp, err := l.svcCtx.GoodsRpc.GetGoods(l.ctx, &goods.GetGoodsReq{Id: val.MaterialID})
		if err != nil {
			return nil, err
		}
		ret = append(ret, &technology_tree.Material{
			Id:    goodsResp.Data.Id,
			Name:  goodsResp.Data.Name,
			Count: val.Cnt,
			Image: goodsResp.Data.Image,
		})

	}

	return ret, nil
}
