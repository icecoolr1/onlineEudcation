package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Teacher/Api/internal/logic"
	"onlineEudcation/Teacher/Api/internal/svc"
	"onlineEudcation/Teacher/Api/internal/types"
)

func tLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeacherInfoForLogin
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTLoginLogic(r.Context(), svcCtx)
		resp, err := l.TLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
