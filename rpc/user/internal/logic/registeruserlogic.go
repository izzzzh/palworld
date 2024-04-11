package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/user/model"
	"time"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var DefaultUserRole = 1

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterUserReq) (*user.RegisterUserResp, error) {
	// todo: add your logic here and delete this line
	if in.Password != in.ConfirmPassword {
		return nil, errors.New("两次密码不一致")
	}

	if !CaptchaStore.Verify(in.CaptchaId, in.Captcha, true) {
		return nil, errors.New("验证码错误")
	}
	enPassword, err := EncryptPassword(in.Password)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		Name:     in.Username,
		Password: enPassword,
		Role:     int32(DefaultUserRole),
		Avatar:   DefaultAvatar,
	}
	if err = AddUser(l.ctx, u); err != nil {
		return nil, err
	}
	var ret = &user.RegisterUserResp{}
	ret.Code = http.StatusOK
	ret.Message = "注册成功"

	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	now := time.Now().Unix()

	token, err := GetJwtToken(secret, now, expire, u.ID)
	if err != nil {
		return nil, err
	}

	loginInfo := &user.LoginInfo{}
	loginInfo.Token = &user.Token{
		AccessToken:  token,
		AccessExpire: now + expire,
		RefreshAfter: now + expire/2,
	}
	loginInfo.User = &user.User{
		Id:        u.ID,
		Username:  u.Name,
		Role:      mapRole[int(u.Role)],
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt.Format("2006-02-01 15:04:05"),
	}

	ret.Data = loginInfo

	return ret, nil
}
