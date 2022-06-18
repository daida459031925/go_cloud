package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	//Rpc
	zrpc.RpcServerConf

	//数据库
	Mysql struct {
		DataSource string
	}

	//redis缓存
	CacheRedis cache.CacheConf
}
