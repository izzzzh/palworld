package user

import (
	"context"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (*types.Captcha, error) {
	in := &user.GetCaptchaReq{}
	resp, err := l.svcCtx.UserRpc.GetCaptcha(l.ctx, in)
	if err != nil {
		return nil, nil
	}

	ret := &types.Captcha{
		CaptchaId:     resp.CaptchaId,
		CaptchaBase64: resp.CaptchaBase64,
	}

	return ret, nil
}
