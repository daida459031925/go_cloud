package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"service/common/httpxParse"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//首先模板自带检查是否包含对应的字段
		//Username string `json:"username" validate:"required,min=8,max=30"`
		//Password string `json:"password" validate:"required,min=8,max=30"`
		//只验证了json 和 form
		var req types.Login
		if v := httpxParse.Parse(w, r, &req); v {
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp := l.Login(req)
		httpx.OkJson(w, resp)
	}
}
