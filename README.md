# go_cloud

bin 存放二进制文件
pkg 存放打包文件
log 存放日志
gotest 测试项目

需要生成的文件，不管什么参数变动了，一定要变动对应的proto、api、model然后重新生成文件


## 项目启动
1.打包的文件只有go语言的文件，静态文件是没有进行打包的（例如生成的exe时候需要将etc/user-api.yaml放在exe同目录下）
2.运行go打包好的项目时候需要使用./XXXX项目

项目执行方法

1.编辑配置
2.设置软件包路径：可执行方法.go的可执行方法路径
3.设置工作目录：可执行方法.go的可执行方法路径
4.go工具实参：-i
5.设置程序实参数：可执行方法.go -f 配置文件.yaml


var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

- 三个参数分别对应
- f ： -f 
- etc/user-api.yaml ： etc/user-api.yaml
- go run user.go -f etc/user-api.yaml 


## api代码自动生成
使用下载好的 goctl 进行代码生成 windows下只能使用命令

cd book/service/user/api

goctl api go -api user.api -dir .

goctl api go -api 文件位置/user.api -dir ./执行后安装文件在哪里

## model代码自动生成
cd service/user/model

goctl model mysql ddl -src user.sql -dir . -c

goctl model mysql ddl -src 文件位置/user.sql -dir ./执行后安装文件在哪里 -c

## 数据库连接模式
goctl model mysql datasource -url="$datasource" -table="user" -c -dir .

## idea插件模式
在Goland中，右键user.sql，依次进入并点击New->Go Zero->Model Code即可生成，或者打开user.sql文件，
进入编辑区，使用快捷键Command+N（for mac OS）或者 alt+insert（for windows），选择Mode Code即可

## rpc服务代码自动生成
编译proto文件
cd service/user/rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

## 项目打包
goreleaser init
goreleaser --snapshot --skip-publish --rm-dist

我们的demo项目没有使用go generate，需要把 - go generate ./... 注释掉

解释说明

id：打包后目录前缀
goos：目标系统
goarch：目标CPU架构
snapshot.name_template：生成压缩包名称前缀
详细配置请参考官网
build配置： 编译配置
archives配置：打包配置
goos和goarch是乘积关系，至于为什么没有生成windows_arm64这样的可执行程序 是因为windows不能再arm上跑。

## 本地调试
编辑配置：
 - 将工作目录和软件包目录替换到需要执行的go文件目录下
 - go工具实参添加-i
 - 程序实参添加 xxx.go -f ./etc/user.yaml

## 当在代码中使用了第三方库 ，但是go.mod中并没有跟着更新的时候
在指定的go.mod文件所在目录运行
go mod tidy

若文件已经下载完毕还有不对的地方，则需要在go.mod中删除对应依赖，然后同步

## 启动顺序
有rpc服务的优先启动rpc 在启动api

## go-zero连接redis etc会自动重新连接


# gf文档
https://goframe.org/pages/viewpage.action?pageId=1114119
# go-zero文档
https://go-zero.dev/cn/