package logic

import (
	"context"
	"net/http"
	"palworld/rpc/user/dao"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByIdsLogic {
	return &GetUsersByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersByIdsLogic) GetUsersByIds(in *user.GetUsersByIdsReq) (*user.GetUsersByIdsResp, error) {
	var ret = &user.GetUsersByIdsResp{}
	if in.Ids == nil {
		ret.Code = http.StatusOK
		ret.Message = "ok"
		return ret, nil
	}
	p := dao.User
	userRet, err := p.Where(p.ID.In(in.Ids...)).Find()
	if err != nil {
		return nil, err
	}
	data := make([]*user.User, 0)

	for _, v := range userRet {
		data = append(data, &user.User{
			Id:       v.ID,
			Username: v.Name,
			Avatar:   v.Avatar,
		})
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = data

	return ret, nil
}
