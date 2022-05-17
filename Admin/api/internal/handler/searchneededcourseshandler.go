package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Admin/api/internal/logic"
	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"
)

func searchNeededCoursesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchCoursesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchNeededCoursesLogic(r.Context(), svcCtx)
		resp, err := l.SearchNeededCourses(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
