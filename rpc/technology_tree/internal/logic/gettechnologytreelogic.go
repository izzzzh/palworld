package logic

import (
	"context"
	"net/http"
	"palworld/rpc/technology_tree/dao"
	"palworld/rpc/technology_tree/model"
	"sort"

	"palworld/rpc/technology_tree/internal/svc"
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
	ret.Code = http.StatusOK
	ret.Message = "ok"
	resData := make([]*technology_tree.TechnologyTree, 0)
	for key, val := range mapsTechnology {
		technologyTree := &technology_tree.TechnologyTree{}
		listTechnology := make([]*technology_tree.Technology, 0)
		for _, v := range val {
			listTechnology = append(listTechnology, &technology_tree.Technology{
				Name:        v.Name,
				Description: v.Description,
				Icon:        v.Icon,
				Ancient:     v.Ancient,
				Cost:        v.Cost,
			})
		}
		technologyTree.Level = key
		technologyTree.Data = listTechnology
		resData = append(resData, technologyTree)
	}
	sort.Slice(resData, func(i, j int) bool {
		return resData[i].Level < resData[j].Level
	})
	ret.Data = resData
	return ret, nil
}
