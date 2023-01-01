package impl

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/common/generalSql"
	"service/common/generalSql/model/sys"
)

type (
	SysDictModel struct {
		*generalSql.DefaultModel[sys.Dict]
	}

	SysResourcesModel struct {
		*generalSql.DefaultModel[sys.Resources]
	}

	SysRoleModel struct {
		*generalSql.DefaultModel[sys.Role]
	}

	SysUserModel struct {
		*generalSql.DefaultModel[sys.User]
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
		FindOneLoginUser(account string) (*sys.User, error)
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
	return &SysUserModel{generalSql.NewModel[sys.User](conn, cache, sysUser)}
}

func GetSysDictModel(conn sqlx.SqlConn, cache cache.CacheConf) SysDictImpl {
	return &SysDictModel{generalSql.NewModel[sys.Dict](conn, cache, sysDict)}
}

func GetSysResourcesModel(conn sqlx.SqlConn, cache cache.CacheConf) SysResourcesImpl {
	return &SysResourcesModel{generalSql.NewModel[sys.Resources](conn, cache, sysResources)}
}

func GetSysRoleModel(conn sqlx.SqlConn, cache cache.CacheConf) SysRoleImpl {
	return &SysRoleModel{generalSql.NewModel[sys.Role](conn, cache, sysRole)}
}
