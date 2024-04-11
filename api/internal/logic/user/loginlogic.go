package user

import (
	"context"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	in := &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	resp, err := l.svcCtx.UserRpc.Login(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var ret = &types.LoginResp{}

	ret.Token = &types.Token{
		AccessToken:  resp.Data.Token.AccessToken,
		AccessExpire: resp.Data.Token.AccessExpire,
		RefreshAfter: resp.Data.Token.RefreshAfter,
	}

	ret.UserInfo = &types.UserInfo{
		Id:        resp.Data.User.Id,
		Username:  resp.Data.User.Username,
		Role:      resp.Data.User.Role,
		Avatar:    resp.Data.User.Avatar,
		CreatedAt: resp.Data.User.CreatedAt,
	}

	return ret, nil
}
