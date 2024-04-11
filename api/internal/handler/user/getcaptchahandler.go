package user

import (
	"net/http"

	"palworld/api/internal/logic/user"
	"palworld/api/internal/svc"
	result "palworld/common/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		result.HttpResult(r, w, resp, err)
	}
}
