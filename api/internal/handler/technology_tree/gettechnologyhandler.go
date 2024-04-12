package technology_tree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/technology_tree"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	result "palworld/common/httpx"
)

func GetTechnologyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTechnologyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := technology_tree.NewGetTechnologyLogic(r.Context(), svcCtx)
		resp, err := l.GetTechnology(&req)
		result.HttpResult(r, w, resp, err)
	}
}
