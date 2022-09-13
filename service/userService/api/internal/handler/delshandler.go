package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
)

func delsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDelsLogic(r.Context(), svcCtx)
		err := l.Dels()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
