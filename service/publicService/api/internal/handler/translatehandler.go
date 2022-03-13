package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/publicService/api/internal/logic"
	"service/publicService/api/internal/svc"
)

func translateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTranslateLogic(r.Context(), svcCtx)
		resp, err := l.Translate()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
