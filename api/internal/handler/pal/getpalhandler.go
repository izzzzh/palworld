package pal

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/pal"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func GetPalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pal.NewGetPalLogic(r.Context(), svcCtx)
		resp, err := l.GetPal(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
