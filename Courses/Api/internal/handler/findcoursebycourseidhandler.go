package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Courses/Api/internal/logic"
	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"
)

func findCourseByCourseIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindCourseByCourseId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFindCourseByCourseIdLogic(r.Context(), svcCtx)
		resp, err := l.FindCourseByCourseId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
