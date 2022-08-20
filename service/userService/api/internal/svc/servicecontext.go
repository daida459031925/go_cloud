package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	sysImpl "service/common/generalSql/impl/sys"
	"service/userService/api/internal/config"
	"service/userService/rpc/user"
)

// svc时服务起来相当于初始化什么东西
type ServiceContext struct {
	Config config.Config
	//添加数据库支持
	SysUserModel sysImpl.SysUserImpl
	SysDictModel sysImpl.SysDictImpl
	SysResModel  sysImpl.SysResourcesImpl
	SysRoleModel sysImpl.SysRoleImpl
	//远程服务调用初始化
	UserRpc user.User
	//消息队列服务

}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//添加缓存
		SysUserModel: sysImpl.GetSysUserModel(conn, c.CacheRedis),
		SysDictModel: sysImpl.GetSysDictModel(conn, c.CacheRedis),
		SysResModel:  sysImpl.GetSysResourcesModel(conn, c.CacheRedis),
		SysRoleModel: sysImpl.GetSysRoleModel(conn, c.CacheRedis),
		//添加rpc依赖
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
