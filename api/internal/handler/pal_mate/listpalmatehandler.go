package pal_mate

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/pal_mate"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func ListPalMateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPalMateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pal_mate.NewListPalMateLogic(r.Context(), svcCtx)
		resp, err := l.ListPalMate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
