package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"palworld/rpc/pal/dao"
	"palworld/rpc/pal/internal/svc"
	"palworld/rpc/pal/pb/pal"
	"strconv"
	"strings"
)

type GetPalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPalLogic {
	return &GetPalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPalLogic) GetPal(in *pal.GetPalReq) (*pal.GetPalResp, error) {
	if in.Id < 0 {
		return nil, errors.New("无效id")
	}
	p := dao.Pal
	resp, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}
	ret := &pal.GetPalResp{}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	var attIds []int64
	for _, val := range strings.Split(resp.AttributeIds, ",") {
		attId, _ := strconv.Atoi(val)
		attIds = append(attIds, int64(attId))
	}
	ret.Data = &pal.Pal{
		Id:           resp.ID,
		Number:       resp.Number,
		Name:         resp.Name,
		Icon:         resp.Icon,
		AttributeIds: attIds,
		Hp:           int64(resp.Hp),
		Energy:       int64(resp.Energy),
		Defensively:  int64(resp.Defensively),
	}
	return ret, nil
}
