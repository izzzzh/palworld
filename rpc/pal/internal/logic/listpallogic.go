package logic

import (
	"context"
	"net/http"
	"palworld/rpc/pal/dao"
	"palworld/rpc/pal/pb/pal"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"palworld/rpc/pal/internal/svc"
)

type ListPalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPalLogic {
	return &ListPalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPalLogic) ListPal(in *pal.ListPalReq) (*pal.ListPalResp, error) {
	ret := &pal.ListPalResp{}
	p := dao.Pal
	pd := p.WithContext(l.ctx)
	if in.Name != "" {
		pd.Where(p.Name.Like(in.Name))
	}
	if in.AttributeId > 0 {
		pd.Where(p.AttributeIds.Like(strconv.Itoa(int(in.AttributeId))))
	}
	pals, err := pd.Find()
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	listPals := make([]*pal.ListPal, 0)
	for _, val := range pals {
		var attIds []int64
		for _, attId := range strings.Split(val.AttributeIds, ",") {
			id, _ := strconv.ParseInt(attId, 10, 64)
			attIds = append(attIds, id)
		}
		listPals = append(listPals, &pal.ListPal{
			Id:           val.ID,
			Number:       val.Number,
			Name:         val.Name,
			Icon:         val.Icon,
			AttributeIds: attIds,
		})
	}
	ret.Data = listPals
	return ret, nil
}
