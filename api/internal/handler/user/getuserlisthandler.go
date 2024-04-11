package user

import (
	"net/http"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/user"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
