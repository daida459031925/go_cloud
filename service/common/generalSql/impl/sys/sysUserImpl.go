package generalSql

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/common/generalSql/model/sys"
)

func (d *SysUserModel) FindOneAndDeleted(id uint64, deleted ...int) (any, error) {
	idKey := fmt.Sprintf("%s%v", d.CacheIdPrefix, id)
	var resp sys.SysUser

	del := 1

	if len(deleted) > 0 && fx.Just(1, 2).AnyMach(func(item interface{}) bool {
		return deleted[0] == item.(int)
	}) {
		del = deleted[0]
	}

	err := d.QueryRow(&resp, idKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `deleted` = ? limit 1", d.FieldNameRows, d.Table)
		//要想动态传入参数，需要将 args := []interface{}{id} 这样的对象   然后 使用 args... 作为参数传入才能正常执行sql
		//sql的?的个数必须与args相当 这就说明现在要做的是数据 解析器和 sql语句 问号生成器
		args := []interface{}{id, del}
		return conn.QueryRow(v, query, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, d.ErrNotFound
	default:
		return nil, err
	}
}
