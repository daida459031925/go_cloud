package generalSql

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"service/common/constant"
	"service/common/generalSql/model/sys"
	"strings"
)

func (d *SysUserModel) FindOneAndDeleted(id uint64, deleted ...int) (any, error) {
	idKey := fmt.Sprintf("%s%v", d.CacheIdPrefix, id)
	var resp sys.User

	del := 1

	if deleted != nil && len(deleted) > 0 {
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

func (d *SysUserModel) FindOneLoginUser(account string) (*sys.User, error) {
	var resp sys.User
	var e error = nil
	//获取可以传入的key
	a := strings.TrimSpace(account)
	if len(a) <= 0 {
		return &resp, errors.New(constant.ErrSqlFind00_01)
	}

	args := []interface{}{a}
	query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", d.FieldNameRows, d.Table)

	err := d.QueryRowNoCache(&resp, query, args...)

	if err == sqlc.ErrNotFound {
		e = d.ErrNotFound
	}

	return &resp, e
}
