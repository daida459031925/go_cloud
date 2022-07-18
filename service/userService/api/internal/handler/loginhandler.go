package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"service/userService/api/internal/logic"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//首先模板自带检查是否包含对应的字段
		//Username string `json:"username" form:"username" validate:"required,min=8,max=30"`
		//Password string `json:"password" form:"password" validate:"required,min=8,max=30"`
		//只验证了json 和 form
		var req types.Login
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//自定义模板？添加内容验证validate

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp := l.Login(req)
		httpx.OkJson(w, resp)
	}
}
