package logic

import (
	"context"
	"net/http"
	"palworld/rpc/pal/dao"

	"palworld/rpc/pal/internal/svc"
	"palworld/rpc/pal/pb/pal"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPalByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPalByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPalByIdsLogic {
	return &GetPalByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPalByIdsLogic) GetPalByIds(in *pal.GetPalByIdsReq) (*pal.GetPalByIdsResp, error) {
	var ret = &pal.GetPalByIdsResp{}

	p := dao.Pal
	data := make([]*pal.PalIdsInfo, 0)
	err := p.Where(p.ID.In(in.Ids...)).Select(p.ID, p.Name, p.Icon).Scan(&data)
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = data

	return ret, nil
}
