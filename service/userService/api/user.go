package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"service/common/constant"
	"service/common/middlewares"
	"service/userService/api/internal/config"
	"service/userService/api/internal/handler"
	"service/userService/api/internal/svc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件 自定义中间 可以拦截所有错误，那么可以在里面自定义错误集
	var mid = middlewares.SetHandlers(
		middlewares.CreateError(constant.LogicLog),
		//middlewares.CreateXSS(constant.LogicLog),
	)
	mid.InitAll(server)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
