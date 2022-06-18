package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/common/generalSql"
	"service/userService/api/internal/config"
	"service/userService/model"
)

// svc时服务起来相当于初始化什么东西
type ServiceContext struct {
	Config config.Config
	//添加数据库支持
	UserModel generalSql.TkMybatisModel
	//远程服务调用初始化
	//UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//添加缓存
		UserModel: generalSql.CreateModel[model.SysUser](conn, c.CacheRedis, "sys_user"),
		//添加rpc依赖
		//UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
