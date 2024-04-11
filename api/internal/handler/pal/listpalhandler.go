package pal

import (
	"net/http"

	"palworld/api/internal/logic/pal"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pal.NewListPalLogic(r.Context(), svcCtx)
		resp, err := l.ListPal(&req)
		result.HttpResult(r, w, resp, err)
	}
}
