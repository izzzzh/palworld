package technology_tree

import (
	"context"
	"net/http"
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

func (l *GetTechnologyTreeLogic) GetTechnologyTree() (*types.GetTechnologyTreeResp, error) {
	ret := &types.GetTechnologyTreeResp{}
	in := &technology_tree.GetTechnologyTreeReq{}
	resp, err := l.svcCtx.TechnologyTreeRpc.GetTechnologyTree(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return ret, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	retData := make([]types.TechnologyTree, 0)
	for _, val := range resp.Data {
		technologyList := make([]types.Technology, 0)
		for _, technology := range val.Data {
			technologyList = append(technologyList, types.Technology{
				Name:        technology.Name,
				Description: technology.Description,
				Cost:        int(technology.Cost),
				Icon:        technology.Icon,
				Ancient:     technology.Ancient,
			})
		}
		retData = append(retData, types.TechnologyTree{
			Level: val.Level,
			Data:  technologyList,
		})
	}

	ret.Data = retData
	return ret, nil
}
