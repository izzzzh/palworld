package logic

import (
	"context"
	"net/http"
	"palworld/rpc/goods/dao"

	"palworld/rpc/goods/internal/svc"
	"palworld/rpc/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdsLogic {
	return &GetGoodsByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIdsLogic) GetGoodsByIds(in *goods.GetGoodsByIdsReq) (*goods.GetGoodsByIdsResp, error) {
	p := dao.Goods
	goodsRet, err := p.WithContext(l.ctx).Where(p.ID.In(in.Ids...)).Find()
	if err != nil {
		return nil, err
	}

	var ret = &goods.GetGoodsByIdsResp{}

	var data = make([]*goods.Goods, 0)
	for _, v := range goodsRet {
		data = append(data, &goods.Goods{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Image:       v.Image,
			Quality:     v.Quality,
		})
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = data

	return ret, nil
}
