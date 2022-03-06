package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserIdPrefix = "cache:user:id:"
)

type (
	UserModel interface {
		Insert(data *User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		Update(data *User) error
		Delete(id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64          `db:"id"`         // 唯一主键id
		Deleted    int64          `db:"deleted"`    // 标识是否删除1存在 2删除
		Account    sql.NullString `db:"account"`    // 账号
		Password   sql.NullString `db:"password"`   // 密码
		Gender     sql.NullInt64  `db:"gender"`     // 性别（男|女|未知）
		CreateTime sql.NullTime   `db:"createTime"` // 创建时间
		UpdateTime sql.NullTime   `db:"updateTime"` // 最后更新时间
		CreateBy   sql.NullString `db:"createBy"`   // 创建人
		UpdateBy   sql.NullString `db:"updateBy"`   // 最后更新人
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Deleted, data.Account, data.Password, data.Gender, data.CreateBy, data.UpdateBy)

	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(data *User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.Deleted, data.Account, data.Password, data.Gender, data.CreateBy, data.UpdateBy, data.Id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
