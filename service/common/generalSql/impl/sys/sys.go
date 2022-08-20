package generalSql

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/common/generalSql"
	"service/common/generalSql/model/sys"
)

type (
	SysDictModel struct {
		*generalSql.DefaultModel[sys.SysDict]
	}

	SysResourcesModel struct {
		*generalSql.DefaultModel[sys.SysResources]
	}

	SysRoleModel struct {
		*generalSql.DefaultModel[sys.SysRole]
	}

	SysUserModel struct {
		*generalSql.DefaultModel[sys.SysUser]
	}

	SysDictImpl interface {
		generalSql.TkMybatisModel
	}

	SysResourcesImpl interface {
		generalSql.TkMybatisModel
	}

	SysRoleImpl interface {
		generalSql.TkMybatisModel
	}

	SysUserImpl interface {
		generalSql.TkMybatisModel
		FindOneAndDeleted(id uint64, deleted ...int) (any, error)
	}
)

/**
编写此类的目的是统一维护数据库连接防止我要改表或则表名到处修改
这样进行修改只有这一处变动其他位置不必变动
除非你换方法名
*/

const (
	sysUser      = "sys_user"
	sysDict      = "sys_dict"
	sysResources = "sys_resources"
	sysRole      = "sys_role"
)

func GetSysUserModel(conn sqlx.SqlConn, cache cache.CacheConf) SysUserImpl {
	return &SysUserModel{generalSql.CreateModel[sys.SysUser](conn, cache, sysUser)}
}

func GetSysDictModel(conn sqlx.SqlConn, cache cache.CacheConf) SysDictImpl {
	return &SysDictModel{generalSql.CreateModel[sys.SysDict](conn, cache, sysDict)}
}

func GetSysResourcesModel(conn sqlx.SqlConn, cache cache.CacheConf) SysResourcesImpl {
	return &SysResourcesModel{generalSql.CreateModel[sys.SysResources](conn, cache, sysResources)}
}

func GetSysRoleModel(conn sqlx.SqlConn, cache cache.CacheConf) SysRoleImpl {
	return &SysRoleModel{generalSql.CreateModel[sys.SysRole](conn, cache, sysRole)}
}
