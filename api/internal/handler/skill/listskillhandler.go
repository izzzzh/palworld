package skill

import (
	"net/http"
	result "palworld/common/httpx"

	"palworld/api/internal/logic/skill"
	"palworld/api/internal/svc"
)

func ListSkillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := skill.NewListSkillLogic(r.Context(), svcCtx)
		resp, err := l.ListSkill()
		result.HttpResult(r, w, resp, err)
	}
}
