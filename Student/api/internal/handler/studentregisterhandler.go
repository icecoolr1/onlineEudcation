package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Student/api/internal/logic"
	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"
)

func studentRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStudentRegisterLogic(r.Context(), svcCtx)
		resp, err := l.StudentRegister(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
