package logic

import (
	"context"
	"net/http"
	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var CaptchaStore = base64Captcha.DefaultMemStore

func (l *GetCaptchaLogic) GetCaptcha(in *user.GetCaptchaReq) (*user.GetCaptchaResp, error) {
	// todo: add your logic here and delete this line
	var ret = &user.GetCaptchaResp{}

	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, CaptchaStore)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.CaptchaId = id
	ret.CaptchaBase64 = b64s
	return ret, nil
}
