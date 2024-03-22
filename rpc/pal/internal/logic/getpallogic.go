package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"palworld/rpc/pal/dao"
	"palworld/rpc/pal/internal/svc"
	"palworld/rpc/pal/pb/pal"
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
	if in.Id <= 0 {
		return nil, errors.New("无效id")
	}
	p := dao.Pal
	resp, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}
	ret := &pal.GetPalResp{}
	var abilities = make([]*pal.Ability, 0)
	if err := json.Unmarshal([]byte(resp.Abilities), &abilities); err != nil {
		ret.Code = http.StatusInternalServerError
		ret.Message = err.Error()
		return ret, nil
	}

	skills, err := getPalSkillsByID(l.ctx, in.Id)
	if err != nil {
		ret.Code = http.StatusInternalServerError
		ret.Message = err.Error()
		return ret, nil
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = &pal.Pal{
		Id:           resp.ID,
		Number:       resp.Number,
		Name:         resp.Name,
		Icon:         resp.Icon,
		AttributeIds: ParseAttributeIDs(resp.AttributeIds),
		Hp:           resp.Hp,
		Energy:       resp.Energy,
		Defensively:  resp.Defensively,
		Eat:          resp.Eat,
		Abilities:    abilities,
		Description:  resp.Description,
		Passive:      resp.Passive,
		PassiveDesc:  resp.PassiveDesc,
		Skills:       skills,
	}
	return ret, nil
}

func getPalSkillsByID(ctx context.Context, id int64) ([]*pal.Skill, error) {
	psm := dao.PalSkillMap
	s := dao.Skill
	skills := make([]*pal.Skill, 0)
	err := psm.WithContext(ctx).Select(s.ID, s.Name, s.Ct, s.Description, s.Power, s.AttributeID).
		LeftJoin(s, psm.SkillID.
			EqCol(s.ID)).Where(psm.PalID.Eq(id)).Scan(skills)
	if err != nil {
		return nil, err
	}

	return skills, nil
}
