package logic

import (
	"context"
	"net/http"
	"palworld/rpc/technology_tree/dao"
	"palworld/rpc/technology_tree/model"

	"palworld/rpc/technology_tree/internal/svc"
	"palworld/rpc/technology_tree/pb/technology_tree"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTechnologyTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTechnologyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTechnologyTreeLogic {
	return &AddTechnologyTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTechnologyTreeLogic) AddTechnologyTree(in *technology_tree.AddTechnologyTreeReq) (*technology_tree.AddTechnologyTreeResp, error) {
	var ret = &technology_tree.AddTechnologyTreeResp{}

	tech := model.TechnologyTree{
		Name:        in.Name,
		Icon:        in.Icon,
		Level:       in.Level,
		Cost:        in.Cost,
		Ancient:     in.Ancient,
		Description: in.Description,
	}
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		p := dao.TechnologyTree
		q := tx.TechnologyMaterial
		err := p.WithContext(l.ctx).Create(&tech)
		if err != nil {
			return err
		}
		for _, m := range in.Materials {
			material := model.TechnologyMaterial{
				TechnologyID: tech.ID,
				MaterialID:   m.Id,
				Cnt:          m.Count,
			}
			err = q.WithContext(l.ctx).Create(&material)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"

	return ret, nil
}
