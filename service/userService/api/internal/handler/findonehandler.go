package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
)

func find_oneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdDBSysUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFind_oneLogic(r.Context(), svcCtx)
		resp, err := l.Find_one(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}