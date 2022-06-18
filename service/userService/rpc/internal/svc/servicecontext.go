package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/userService/model"
	"service/userService/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化数据库
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//初始化缓存
		UserModel: model.NewSysUserModel(conn, c.CacheRedis),
	}
}
