package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/user/dao"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	if in.Id <= 0 {
		return nil, errors.New("用户信息不存在")
	}
	p := dao.User
	userRet, err := p.WithContext(l.ctx).Where(p.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}
	var ret = &user.GetUserInfoResp{}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = &user.User{
		Id:        userRet.ID,
		Username:  userRet.Name,
		Role:      mapRole[int(userRet.Role)],
		Avatar:    userRet.Avatar,
		CreatedAt: userRet.CreatedAt.Format("2006-02-01 15:04:05"),
	}

	return ret, nil
}
