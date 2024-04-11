package user

import (
	"net/http"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/user"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		result.HttpResult(r, w, struct{}{}, err)
	}
}
