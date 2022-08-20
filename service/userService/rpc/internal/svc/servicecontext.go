package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	sysImpl "service/common/generalSql/impl/sys"
	"service/userService/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	SysUserModel sysImpl.SysUserImpl
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化数据库
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//初始化缓存
		SysUserModel: sysImpl.GetSysUserModel(conn, c.CacheRedis),
	}
}
