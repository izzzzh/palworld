package logic

import (
	"context"
	"net/http"
	"sort"

	"palworld/rpc/goods/pb/goods"
	"palworld/rpc/technology_tree/dao"
	"palworld/rpc/technology_tree/internal/svc"
	"palworld/rpc/technology_tree/model"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTechnologyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyTreeLogic {
	return &GetTechnologyTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTechnologyTreeLogic) GetTechnologyTree(in *technology_tree.GetTechnologyTreeReq) (*technology_tree.GetTechnologyTreeResp, error) {
	ret := &technology_tree.GetTechnologyTreeResp{}
	p := dao.TechnologyTree
	technology, err := p.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}
	mapsTechnology := make(map[int32][]*model.TechnologyTree)
	for _, val := range technology {
		mapsTechnology[val.Level] = append(mapsTechnology[val.Level], val)
	}
	resData := make([]*technology_tree.TechnologyTree, 0)

	for key, val := range mapsTechnology {
		technologyTree := &technology_tree.TechnologyTree{}
		listTechnology := make([]*technology_tree.Technology, 0)
		for _, v := range val {
			materials, err := l.GetTechMaterial(v.ID)
			if err != nil {
				return nil, err
			}
			listTechnology = append(listTechnology, &technology_tree.Technology{
				Name:        v.Name,
				Description: v.Description,
				Icon:        v.Icon,
				Ancient:     v.Ancient,
				Cost:        v.Cost,
				Materials:   materials,
				Level:       v.Level,
			})
		}
		technologyTree.Level = key
		technologyTree.Data = listTechnology
		resData = append(resData, technologyTree)
	}
	sort.Slice(resData, func(i, j int) bool {
		return resData[i].Level < resData[j].Level
	})

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = resData

	return ret, nil
}

func (l *GetTechnologyTreeLogic) GetTechMaterial(id int64) ([]*technology_tree.Material, error) {
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
