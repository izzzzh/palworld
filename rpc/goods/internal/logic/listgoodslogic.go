package logic

import (
	"context"
	"net/http"
	"palworld/rpc/goods/dao"
	"palworld/rpc/goods/internal/svc"
	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGoodsLogic {
	return &ListGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGoodsLogic) ListGoods(in *goods.ListGoodsReq) (*goods.ListGoodsResp, error) {
	ret := &goods.ListGoodsResp{}
	g := dao.Goods
	pd := g.WithContext(l.ctx)
	if in.Name != "" {
		pd = pd.Where(g.Name.Eq(in.Name))
	}
	if in.Types != "" {
		pd = pd.Where(g.Types.Eq(in.Types))
	}
	if in.Quality != 0 {
		pd = pd.Where(g.Quality.Eq(in.Quality))
	}
	page, pageSize := in.Page, in.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	goodsResp, err := pd.Offset(int(offset)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	listGoods := make([]*goods.Goods, 0)
	for _, val := range goodsResp {
		listGoods = append(listGoods, &goods.Goods{
			Id:          val.ID,
			Name:        val.Name,
			Image:       val.Image,
			Description: val.Description,
			Workload:    val.Workload,
			Quality:     val.Quality,
			Types:       val.Types,
		})
	}
	ret.Data = listGoods
	return ret, nil
}
