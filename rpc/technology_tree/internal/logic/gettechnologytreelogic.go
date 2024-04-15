package logic

import (
	"context"
	"net/http"
	"sort"
	"sync"

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
	pm := dao.TechnologyMaterial

	technology, err := p.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}

	mapsTechnology := make(map[int32][]*model.TechnologyTree)
	techMaterials := make(map[int64][]*model.TechnologyMaterial)
	mapIds := make(map[int64]bool)

	errChan := make(chan error, len(technology))
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, val := range technology {
		wg.Add(1)
		tech := val
		go func(tech *model.TechnologyTree) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			tms, err := pm.WithContext(l.ctx).Where(pm.TechnologyID.Eq(tech.ID)).Find()
			if err != nil {
				errChan <- err
			}
			techMaterials[tech.ID] = tms
			mapsTechnology[tech.Level] = append(mapsTechnology[tech.Level], tech)
		}(tech)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	var materialIds []int64
	for _, tms := range techMaterials {
		for _, tm := range tms {
			if !mapIds[tm.MaterialID] {
				mapIds[tm.MaterialID] = true
				materialIds = append(materialIds, tm.MaterialID)
			}
		}
	}

	goodsRet, err := l.GetTechMaterial(materialIds)
	mapGoods := make(map[int64]*technology_tree.Material)
	for i := range goodsRet {
		g := goodsRet[i]
		mapGoods[g.Id] = g
	}

	resData := make([]*technology_tree.TechnologyTree, 0)

	for key, val := range mapsTechnology {
		technologyTree := &technology_tree.TechnologyTree{}
		listTechnology := make([]*technology_tree.Technology, 0)
		for _, v := range val {
			var materials []*technology_tree.Material
			for _, tm := range techMaterials[v.ID] {
				materials = append(materials, &technology_tree.Material{
					Id:    mapGoods[tm.MaterialID].Id,
					Name:  mapGoods[tm.MaterialID].Name,
					Image: mapGoods[tm.MaterialID].Image,
					Count: tm.Cnt,
				})
			}

			listTechnology = append(listTechnology, &technology_tree.Technology{
				Id:          v.ID,
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

func (l *GetTechnologyTreeLogic) GetTechMaterial(ids []int64) ([]*technology_tree.Material, error) {

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
