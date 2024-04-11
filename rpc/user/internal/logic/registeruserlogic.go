package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/user/model"

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
	}
	if err := AddUser(l.ctx, u); err != nil {
		return nil, err
	}
	var ret = &user.RegisterUserResp{}
	ret.Code = http.StatusOK
	ret.Message = "注册成功"

	return ret, nil
}
