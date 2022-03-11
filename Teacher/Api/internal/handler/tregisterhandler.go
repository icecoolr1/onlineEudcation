package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Teacher/Api/internal/logic"
	"onlineEudcation/Teacher/Api/internal/svc"
	"onlineEudcation/Teacher/Api/internal/types"
)

func tRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeacherInfoForRegister
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTRegisterLogic(r.Context(), svcCtx)
		resp, err := l.TRegister(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
