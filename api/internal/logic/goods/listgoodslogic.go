package goods

import (
	"context"
	"net/http"
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

func (l *ListGoodsLogic) ListGoods(req *types.ListGoodsreq) (*types.ListGoodsResp, error) {
	ret := &types.ListGoodsResp{}
	in := &goods.ListGoodsReq{
		Name:     req.Name,
		Types:    req.Types,
		Page:     req.Page,
		PageSize: req.PageSize,
		Quality:  req.Quality,
	}
	resp, err := l.svcCtx.GoodsRpc.ListGoods(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return nil, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	goodsResult := make([]types.ListGoods, 0)
	for _, val := range resp.Data {
		goodsResult = append(goodsResult, types.ListGoods{
			ID:          val.Id,
			Name:        val.Name,
			Image:       val.Image,
			Description: val.Description,
			Workload:    val.Workload,
			Quality:     val.Quality,
			Types:       val.Types,
		})
	}
	ret.Data = goodsResult
	return ret, nil
}
