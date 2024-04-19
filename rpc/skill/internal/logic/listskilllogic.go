package logic

import (
	"context"
	"net/http"
	"palworld/rpc/pal/pb/pal"
	"palworld/rpc/skill/dao"
	"palworld/rpc/skill/internal/svc"
	"palworld/rpc/skill/pb/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSkillLogic {
	return &ListSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSkillLogic) ListSkill(in *skill.ListSkillReq) (*skill.ListSkillResp, error) {
	ret := &skill.ListSkillResp{}
	p := dao.Skill
	pd := p.WithContext(l.ctx)
	if in.AttributeId > 0 {
		pd.Where(p.AttributeID.Eq(in.AttributeId))
	}
	skills, err := pd.Find()
	if err != nil {
		return nil, err
	}
	psm := dao.PalSkillMap
	palSkill, err := psm.Find()
	if err != nil {
		return nil, err
	}

	palIds := make([]int64, 0)
	mapIds := make(map[int64]struct{})
	mapSkillPal := make(map[int64][]int64)

	for _, val := range palSkill {
		if _, ok := mapIds[val.PalID]; !ok {
			palIds = append(palIds, val.PalID)
		}
		mapSkillPal[val.SkillID] = append(mapSkillPal[val.SkillID], val.PalID)
	}

	palsResp, err := l.svcCtx.PalRpc.GetPalByIds(l.ctx, &pal.GetPalByIdsReq{Ids: palIds})
	if err != nil {
		return nil, err
	}
	mapPal := make(map[int64]*pal.PalIdsInfo)
	for _, val := range palsResp.Data {
		mapPal[val.Id] = val
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	listSkill := make([]*skill.Skill, 0)

	for _, val := range skills {
		skillPals := make([]*skill.SkillPals, 0)
		for _, palId := range mapSkillPal[val.ID] {
			if palInfo, ok := mapPal[palId]; ok {
				skillPals = append(skillPals, &skill.SkillPals{
					Id:    palInfo.Id,
					Name:  palInfo.Name,
					Image: palInfo.Icon,
				})
			}
		}

		listSkill = append(listSkill, &skill.Skill{
			Id:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Ct:          val.Ct,
			Power:       val.Power,
			AttributeId: val.AttributeID,
			Pals:        skillPals,
		})
	}
	ret.Data = listSkill
	return ret, nil
}
