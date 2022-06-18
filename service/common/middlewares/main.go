package middlewares

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

//消息中间件服务

type handlerMake struct {
	makes []rest.Middleware
}

func SetHandlers(hand ...func(next http.HandlerFunc) http.HandlerFunc) handlerMake {
	makes := make([]rest.Middleware, 0)
	if len(hand) > 0 {
		for i := range hand {
			makes = append(makes, hand[i])
		}
	}
	return handlerMake{
		makes: makes,
	}
}

// 全局异常初始化
func (h *handlerMake) InitAll(service *rest.Server) {
	for i := range h.makes {
		service.Use(h.makes[i])
	}
}
