package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Vedio/api/internal/logic"
	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"
)

func deleteVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteVideoLogic(r.Context(), svcCtx)
		resp, err := l.DeleteVideo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
