package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
)

func refreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewToken
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRefreshLogic(r.Context(), svcCtx)
		err := l.Refresh(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
