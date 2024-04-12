package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/goods/dao"

	"palworld/rpc/goods/internal/svc"
	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsLogic {
	return &GetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsLogic) GetGoods(in *goods.GetGoodsReq) (*goods.GetGoodsResp, error) {
	if in.Id <= 0 {
		return nil, errors.New("invalid id")
	}
	p := dao.Goods
	goodsRet, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	materials, err := GetMaterials(l.ctx, goodsRet.ID)
	if err != nil {
		return nil, err
	}
	var ret = &goods.GetGoodsResp{}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = &goods.Goods{
		Id:          goodsRet.ID,
		Name:        goodsRet.Name,
		Description: goodsRet.Description,
		Image:       goodsRet.Image,
		Quality:     goodsRet.Quality,
		Types:       goodsRet.Types,
		Materials:   materials,
	}

	return ret, nil
}
