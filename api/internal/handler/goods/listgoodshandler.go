package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/goods"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func ListGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListGoodsreq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewListGoodsLogic(r.Context(), svcCtx)
		resp, err := l.ListGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
