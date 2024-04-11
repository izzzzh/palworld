package user

import (
	"context"
	"errors"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("两次密码不匹配")
	}

	in := &user.RegisterUserReq{
		Username:        req.Username,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		CaptchaId:       req.CaptchaId,
		Captcha:         req.CaptchaCode,
	}
	resp, err := l.svcCtx.UserRpc.RegisterUser(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var ret = &types.RegisterResp{}
	ret.Token = &types.Token{
		AccessToken:  resp.Data.Token.AccessToken,
		AccessExpire: resp.Data.Token.AccessExpire,
		RefreshAfter: resp.Data.Token.RefreshAfter,
	}
	ret.UserInfo = &types.UserInfo{
		Id:        resp.Data.User.Id,
		Avatar:    resp.Data.User.Avatar,
		Role:      resp.Data.User.Role,
		CreatedAt: resp.Data.User.CreatedAt,
		Username:  resp.Data.User.Username,
	}

	return ret, nil
}
