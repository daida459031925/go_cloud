# go mod 
出现异常使用  go mod tidy

# 业务开发分类
基于go-zero开源框架开发，业务分类如下
1.基本业务逻辑，包含了sql操作业务只有当前服务使用时则写入api中
2.业务涉及到其他业务共同访问时，同时包含了sql操作那么写入当前的rpc中
3.不设计到任意业务操作，并且不涉及到sql操作那么写入common中共同调用

# model
1.编写数据库sql文件
2.生成model文件
goctl model mysql ddl -src 文件位置/user.sql -dir ./执行后安装文件在哪里 -c

数据库连接模式
goctl model mysql datasource -url="$datasource" -table="user" -c -dir .

# api
1.编写api文件
2.使用命令生成代码
    命令示例：goctl api go -api .\service\userService\api\user.api -dir .\service\userService\api\
3.service/userService/api/user.go当前路径下handler.RegisterHandlers(server, ctx)这个方法逻辑不满足时需要重写
    示例：jwt: Auth
    ```go
        engine.AddRoutes(
            rest.WithMiddlewares(
                []rest.Middleware{serverCtx.Example},
                    []rest.Route{
                    {
                        Method:  http.MethodGet,
                        Path:    "/search/do",
                        Handler: searchHandler(serverCtx),
                    },
                }...,
            ),
            rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
        )
    ```
不满足我想要的一个客户一个auth内容，这个地方时写死的所以需要修改

4.修改config文件内容
5.完善yaml内容
6.编写svc下文件添加数据库、缓存支持等需要添加的支持
7.编写logic业务逻辑,并修改业务逻辑的返回值，返还对象改为统一的Result


# rpc
1.编写proto文件
2.使用命令生成代码
    命令示例(目前没发现别的写法)：
    $ cd service/user/rpc
    $ goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
3.添加Config配置
4.完善yaml内容
5.编写svc下初始化信息
6.添加logic业务逻辑
7.其他需要调用服务 重复步骤 3，4，5，6