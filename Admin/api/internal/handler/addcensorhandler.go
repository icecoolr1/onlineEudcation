package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Admin/api/internal/logic"
	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"
)

func addCensorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCensorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddCensorLogic(r.Context(), svcCtx)
		resp, err := l.AddCensor(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
