package skill

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/skill"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func ListSkillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListSkillReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := skill.NewListSkillLogic(r.Context(), svcCtx)
		resp, err := l.ListSkill(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
