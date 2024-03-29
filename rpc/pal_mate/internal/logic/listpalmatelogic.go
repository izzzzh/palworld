package logic

import (
	"context"
	"net/http"
	"palworld/rpc/pal/pb/pal"
	"palworld/rpc/pal_mate/dao"
	"palworld/rpc/pal_mate/internal/svc"
	"palworld/rpc/pal_mate/pb/pal_mate"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPalMateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPalMateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPalMateLogic {
	return &ListPalMateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var PageSize = 20

func (l *ListPalMateLogic) ListPalMate(in *pal_mate.ListPalMateReq) (*pal_mate.ListPalMateResp, error) {
	// todo: add your logic here and delete this line
	ret := &pal_mate.ListPalMateResp{}
	p := dao.PalMateMap
	var pd = p.WithContext(l.ctx)
	if in.ParentOne > 0 && in.ParentTwo > 0 {
		pd = pd.Where(pd.Where(p.ParentOne.Eq(int32(in.ParentOne))).Where(p.ParentTwo.Eq(int32(in.ParentTwo))).
			Or(p.ParentOne.Eq(int32(in.ParentTwo))).Where(p.ParentTwo.Eq(int32(in.ParentOne))))
	} else if in.ParentOne > 0 {
		pd = pd.Where(pd.Where(p.ParentOne.Eq(int32(in.ParentOne))).Or(p.ParentTwo.Eq(int32(in.ParentOne))))
	} else if in.ParentTwo > 0 {
		pd = pd.Where(pd.Where(p.ParentOne.Eq(int32(in.ParentTwo))).Or(p.ParentTwo.Eq(int32(in.ParentTwo))))
	}
	if in.Result > 0 {
		pd = pd.Where(p.Result.Eq(int32(in.Result)))
	}
	page := int(in.Page)
	if in.Page <= 1 {
		page = 1
	}
	offset := (page - 1) * PageSize
	palMates, err := pd.Limit(PageSize).Offset(offset).Find()
	if err != nil {
		return nil, err
	}

	palResp, err := l.svcCtx.PalRpc.ListPal(l.ctx, &pal.ListPalReq{})
	if err != nil {
		return nil, err
	}
	palMap := make(map[int64]*pal.ListPal)
	for i, val := range palResp.Data {
		palMap[val.Id] = palResp.Data[i]
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	listPalMates := make([]*pal_mate.PalMate, 0)
	for _, val := range palMates {
		palOne := palMap[int64(val.ParentOne)]
		palTwo := palMap[int64(val.ParentTwo)]
		palResult := palMap[int64(val.Result)]
		listPalMates = append(listPalMates, &pal_mate.PalMate{
			ParentOne: &pal_mate.Pal{
				Id:   palOne.Id,
				Name: palOne.Name,
				Icon: palOne.Icon,
			},
			ParentTwo: &pal_mate.Pal{
				Id:   palTwo.Id,
				Name: palTwo.Name,
				Icon: palTwo.Icon,
			},
			Result: &pal_mate.Pal{
				Id:   palResult.Id,
				Name: palResult.Name,
				Icon: palResult.Icon,
			},
		})
	}
	ret.Data = listPalMates
	return ret, nil
}
