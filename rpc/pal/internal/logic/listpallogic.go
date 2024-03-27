package logic

import (
	"context"
	"encoding/json"
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
	listPals := make([]*pal.ListPal, 0)
	for _, val := range pals {
		ParseAttributeIDs(val.AttributeIds)
		var abilities = make([]*pal.Ability, 0)
		if err := json.Unmarshal([]byte(val.Abilities), &abilities); err != nil {
			ret.Code = http.StatusInternalServerError
			ret.Message = err.Error()
			return ret, nil
		}
		listPals = append(listPals, &pal.ListPal{
			Id:           val.ID,
			Number:       val.Number,
			Name:         val.Name,
			Icon:         val.Icon,
			AttributeIds: ParseAttributeIDs(val.AttributeIds),
			Abilities:    abilities,
		})
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = listPals
	return ret, nil
}

func ParseAttributeIDs(ids string) []int32 {
	var attIds []int32
	for _, attId := range strings.Split(ids, ",") {
		id, _ := strconv.Atoi(attId)
		attIds = append(attIds, int32(id))
	}
	return attIds
}
