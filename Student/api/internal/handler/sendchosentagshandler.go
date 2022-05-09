package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Student/api/internal/logic"
	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"
)

func sendChosenTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendChosenTags
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendChosenTagsLogic(r.Context(), svcCtx)
		resp, err := l.SendChosenTags(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
