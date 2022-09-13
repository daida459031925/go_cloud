package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
)

func part_updsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPart_updsLogic(r.Context(), svcCtx)
		err := l.Part_upds()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
