package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"palworld/rpc/goods/dao"
	"palworld/rpc/goods/internal/svc"
	"palworld/rpc/goods/pb/goods"
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
		pd = pd.Where(g.Name.Like("%" + in.Name + "%"))
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
	listGoods := make([]*goods.Goods, 0)
	for _, val := range goodsResp {
		materials, err := GetMaterials(l.ctx, val.ID)
		if err != nil {
			return nil, err
		}
		listGoods = append(listGoods, &goods.Goods{
			Id:          val.ID,
			Name:        val.Name,
			Image:       val.Image,
			Description: val.Description,
			Workload:    val.Workload,
			Quality:     val.Quality,
			Types:       val.Types,
			Materials:   materials,
		})
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = listGoods
	return ret, nil
}

func GetMaterials(ctx context.Context, goodsID int64) ([]*goods.Material, error) {
	g := dao.Goods
	gm := dao.GoodsMaterial
	var ret []*goods.Material
	err := gm.WithContext(ctx).Select(g.ID, g.Name, g.Image, gm.Cnt.As("count")).
		Where(gm.GoodsID.Eq(goodsID)).LeftJoin(g, gm.MaterialID.EqCol(g.ID)).Scan(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
