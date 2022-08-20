package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//数据库
	Mysql struct {
		DataSource string
	}
	//权限控制
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	//缓存
	CacheRedis cache.CacheConf
	//内部服务调用
	UserRpc zrpc.RpcClientConf
}
