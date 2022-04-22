package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Courses/Api/internal/logic"
	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"
)

func updateCourseInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCourseInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateCourseInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCourseInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
