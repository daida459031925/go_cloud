package goTest

import (
	"database/sql"
	"fmt"
	rsql "github.com/daida459031925/common/reflex/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"reflect"
	"strings"
	"testing"
)

//sql通用化测试

type Name struct {
	Id              int64          `db:"id"`                  // 唯一主键id
	Deleted         int64          `db:"deleted"`             // 标识是否删除1存在 2删除
	Account         sql.NullString `db:"account"`             // 账号
	Password        sql.NullString `db:"password"`            // 密码
	Salt            sql.NullString `db:"salt"`                // 盐值，用作生成密码加密
	Gender          sql.NullInt64  `db:"gender"`              // 性别（男|女|未知）
	CreateTime      sql.NullTime   `db:"createTime"`          // 创建时间
	UpdateTime      sql.NullTime   `db:"updateTime"`          // 最后更新时间
	CreateBy        sql.NullString `db:"createBy"`            // 创建人
	UpdateBy        sql.NullString `db:"updateBy"`            // 最后更新人
	Secret          sql.NullString `db:"secret"`              // 用户随机生成secret，用作token生成
	PrevSecret      sql.NullString `db:"prevSecret"`          // 用户私有secret
	TokenExpire     sql.NullInt64  `db:"tokenExpire"`         // token时间 单位秒
	TokenExpireName string         `json:"token_expire_name"` // token时间 单位秒
}

// 可以拿到结构体中对应的数据，不管是什么结构体都能拿到
func TestRawFieldValues(tes *testing.T) {
	var n = Name{}

	t := reflect.TypeOf(n)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Printf("Check type eror not Struct")
	}

	v := reflect.ValueOf(n)

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		tes.Logf("field %d, type %s, key %s ,value %v",
			i, t.Field(i).Type, t.Field(i).Name, v.FieldByName(t.Field(i).Name))
	}

}

// 可以拿到结构体中对应的数据，不管是什么结构体都能拿到
func TestRawField(tes *testing.T) {
	var n = Name{TokenExpire: sql.NullInt64{Int64: 123, Valid: true}}
	in := make([]Name, 0)
	in = append(in, n)
	if s, err := rsql.RawField(in, 0); err == nil {
		fmt.Println(fmt.Sprintf("key:%s", s))
	}
	if s, err := rsql.RawField(in, 1); err == nil {
		fmt.Println(fmt.Sprintf("value:%s", s))
	}
	if s, err := rsql.RawField(in, 3); err == nil {
		fmt.Println(fmt.Sprintf("key:%s", s))
	}
	if s, err := rsql.RawField(in, 4); err == nil {
		fmt.Println(fmt.Sprintf("value:%s", s))
	}
}

// 根据拿到的key 和 value 进行sql语句动态执行
type (
	defaultModel struct {
		sqlc.CachedConn
		table string
		//缓存key "cache:sysUser:id:"
		cacheIdPrefix string
		ErrNotFound   error
	}

	//需要对接口下基本内容进行实现
	tkMybatisModel interface {
		//Insert(data any) (sql.Result, error)
		FindOne(id int64) (any, error)
		//Update(data any) error
		//Delete(id int64) error
		//
		//InsertList(datas []any) (sql.Result, error)
		//FindList() (any, error)
		//UpdateList(datas []any) error
		//DeleteList(id int64) error
	}
)

func newModel(conn sqlx.SqlConn, c cache.CacheConf, table string) tkMybatisModel {
	return &defaultModel{
		CachedConn:    sqlc.NewConn(conn, c),
		table:         table,
		cacheIdPrefix: fmt.Sprintf("cache:%s:id:", table),
		ErrNotFound:   sqlx.ErrNotFound,
	}
}

func (d *defaultModel) Insert(data any) (sql.Result, error) {
	return nil, nil
}

func (d *defaultModel) FindOne(id int64) (any, error) {
	sysUserIdKey := fmt.Sprintf("%s%v", d.cacheIdPrefix, id)
	var resp any

	query := ""

	if s, err := rsql.RawField(resp, 0); err == nil {
		query = strings.Join(s, ",")
	}

	err := d.QueryRow(&resp, sysUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", query, d.table)
		//要想动态传入参数，需要将 args := []interface{}{id} 这样的对象   然后 使用 args... 作为参数传入才能正常执行sql
		//sql的?的个数必须与args相当 这就说明现在要做的是数据 解析器和 sql语句 问号生成器
		args := []interface{}{id}
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

func (d *defaultModel) Update(data any) error {
	return nil
}

func (d *defaultModel) Delete(id int64) error {
	return nil
}

func (d *defaultModel) InsertList(datas []any) (sql.Result, error) {
	return nil, nil
}

func (d *defaultModel) FindList() (any, error) {
	return nil, nil
}

func (d *defaultModel) UpdateList(datas []any) error {
	return nil
}

func (d *defaultModel) DeleteList(id int64) error {
	return nil
}
