package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Vedio/api/internal/logic"
	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"
)

func updateVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateVideoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateVideo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
