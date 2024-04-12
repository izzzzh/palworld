package technology_tree

import (
	"net/http"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/technology_tree"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func AddTechnologyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTechnologyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := technology_tree.NewAddTechnologyLogic(r.Context(), svcCtx)
		err := l.AddTechnology(&req)
		result.HttpResult(r, w, struct{}{}, err)
	}
}
