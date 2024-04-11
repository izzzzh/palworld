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

func (l *LoginLogic) Login(req *types.LoginReq) (*types.Token, error) {
	in := &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	resp, err := l.svcCtx.UserRpc.Login(l.ctx, in)
	if err != nil {
		return nil, err
	}
	ret := &types.Token{
		AccessToken:  resp.Data.AccessToken,
		AccessExpire: resp.Data.AccessExpire,
		RefreshAfter: resp.Data.RefreshAfter,
	}

	return ret, nil
}
