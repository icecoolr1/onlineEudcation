package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Classes/api/internal/logic"
	"onlineEudcation/Classes/api/internal/svc"
	"onlineEudcation/Classes/api/internal/types"
)

func DeleteClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteClassRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteClassLogic(r.Context(), svcCtx)
		resp, err := l.DeleteClass(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
