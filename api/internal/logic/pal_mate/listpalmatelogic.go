package pal_mate

import (
	"context"
	"fmt"
	"net/http"
	"palworld/rpc/pal_mate/pb/pal_mate"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPalMateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPalMateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPalMateLogic {
	return &ListPalMateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPalMateLogic) ListPalMate(req *types.ListPalMateReq) (*types.ListPalMateResp, error) {
	ret := &types.ListPalMateResp{}
	fmt.Println(req)
	in := &pal_mate.ListPalMateReq{
		ParentOne: req.ParentOne,
		ParentTwo: req.ParentTwo,
		Result:    req.Result,
		Page:      int32(req.Page),
	}
	resp, err := l.svcCtx.PalMateRpc.ListPalMate(l.ctx, in)
	if err != nil {
		ret.Message = err.Error()
		ret.Code = http.StatusInternalServerError
		return ret, nil
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	palMates := make([]*types.PalMate, 0)
	for _, val := range resp.Data {
		palMates = append(palMates, &types.PalMate{
			ParentOne: types.MateInfo{
				ID:   val.ParentOne.Id,
				Name: val.ParentOne.Name,
				Icon: val.ParentOne.Icon,
			},
			ParentTwo: types.MateInfo{
				ID:   val.ParentTwo.Id,
				Name: val.ParentTwo.Name,
				Icon: val.ParentTwo.Icon,
			},
			Result: types.MateInfo{
				ID:   val.Result.Id,
				Name: val.Result.Name,
				Icon: val.Result.Icon,
			},
		})
	}
	ret.Data = palMates
	return ret, nil
}
