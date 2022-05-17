package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"onlineEudcation/Admin/api/internal/logic"
	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"
)

func getCensorListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetCensorListLogic(r.Context(), svcCtx)
		resp, err := l.GetCensorList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
