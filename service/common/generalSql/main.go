package generalSql

import (
	"database/sql"
	"fmt"
	rsql "github.com/daida459031925/common/reflex/sql"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strconv"
	"strings"
)

// 根据拿到的key 和 value 进行sql语句动态执行
type (
	//通用泛型接口可以从数据库基础信息传入 相当于初始化时候设置类型
	defaultModel[T any] struct {
		/*-----------数据库相关---------------*/

		//数据库连接
		sqlc.CachedConn
		//目的是为了解决通用批量问题
		sqlcon sqlx.SqlConn
		//表名
		table string
		//当前表全部信息
		fieldNames []string
		//将数组的key 编程字符串的key
		fieldNameRows string
		/*-----------缓存相关---------------*/

		//缓存key "cache:sysUser:id:"
		cacheIdPrefix string

		/*-----------错误相关---------------*/
		//错误类
		errNotFound error

		//下面是需要删除的字段
		//sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")

	}

	//需要对接口下基本内容进行实现
	TkMybatisModel interface {
		Insert(data any) (sql.Result, error)
		FindOne(id uint64) (any, error)
		Update(data any, id uint64) error
		Delete(id uint64) error

		InsertList(datas []any) (uint64, error)
		FindList() (any, error)
		UpdateList(datas []any) error
		DeleteList(ids []uint64) error
	}
)

// 创建时候指定类型
func CreateModel[T any](conn sqlx.SqlConn, c cache.CacheConf, table string) TkMybatisModel {
	var t T
	query := builder.RawFieldNames(t)
	return &defaultModel[T]{
		CachedConn:    sqlc.NewConn(conn, c),
		sqlcon:        conn,
		table:         table,
		cacheIdPrefix: fmt.Sprintf("cache:%s:id:", table),
		errNotFound:   sqlx.ErrNotFound,
		fieldNames:    query,
		fieldNameRows: strings.Join(query, ","),
	}
}

//通用化crud完成

func (d *defaultModel[T]) Insert(data any) (sql.Result, error) {
	//获取可以传入的key
	var key string
	var value []interface{}
	keys, err := rsql.RawField(data, 0)
	if err != nil {
		return nil, errors.New("Insert 获取key失败 ")
	}
	key = strings.Join(keys, ",")

	var valueDataString []string

	//获取传入的value
	values, err := rsql.RawField(data, 1)
	if err != nil {
		return nil, errors.New("Insert 获取value失败 ")
	}

	for i := range values {
		value = append(value, values[i])
		valueDataString = append(valueDataString, "?")
	}

	v := strings.Join(valueDataString, ",")

	query := fmt.Sprintf("insert into %s (%s) values (%s)", d.table, key, v)
	ret, err := d.ExecNoCache(query, value...)

	return ret, err
}

// 需要添加查询什么内容，例如我要查询user
func (d *defaultModel[T]) FindOne(id uint64) (any, error) {
	idKey := fmt.Sprintf("%s%v", d.cacheIdPrefix, id)
	var resp T

	err := d.QueryRow(&resp, idKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", d.fieldNameRows, d.table)
		//要想动态传入参数，需要将 args := []interface{}{id} 这样的对象   然后 使用 args... 作为参数传入才能正常执行sql
		//sql的?的个数必须与args相当 这就说明现在要做的是数据 解析器和 sql语句 问号生成器
		args := []interface{}{id}
		return conn.QueryRow(v, query, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, d.errNotFound
	default:
		return nil, err
	}
}

// 通过仔细思考更新时候必须强行代入id，防止数据修改错误
func (d *defaultModel[T]) Update(data any, id uint64) error {
	//sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	//获取可以传入的key
	var key string
	var value []interface{}
	keys, err := rsql.RawField(data, 0)
	if err != nil {
		return errors.New("Insert 获取key失败 ")
	}
	key = strings.Join(stringx.Remove(keys, "id"), "=?,") + "=?"

	//获取传入的value
	values, err := rsql.RawField(data, 1)
	if err != nil {
		return errors.New("Insert 获取value失败 ")
	}

	for i := range values {
		value = append(value, values[i])
	}

	value = append(value, id)

	idKey := fmt.Sprintf("%s%v", d.cacheIdPrefix, id)
	_, err = d.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", d.table, key)
		return conn.Exec(query, value...)
	}, idKey)
	return err
}

func (d *defaultModel[T]) Delete(id uint64) error {
	idKey := fmt.Sprintf("%s%v", d.cacheIdPrefix, id)
	_, err := d.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", d.table)
		return conn.Exec(query, id)
	}, idKey)
	return err
}

//批量的crud

func (d *defaultModel[T]) InsertList(datas []any) (uint64, error) {
	//使用最简单方式实现
	var value = make([]string, 0)
	keys, err := rsql.RawField(datas[0], 3)
	if err != nil {
		return 0, errors.New("批量添加，获取key错误")
	}

	for range keys {
		value = append(value, "?")
	}

	var insertsql = fmt.Sprintf(`insert into %s(%s) values (%s)`,
		d.table, strings.Join(keys, ","), strings.Join(value, ","))

	bulkInserter, err := sqlx.NewBulkInserter(d.sqlcon, insertsql)
	if err != nil {
		return 0, errors.New("批量添加，连接错误错误")
	}

	var reurnErr error = nil
	var reurnList uint64 = 0
	for i := range datas {
		val, err := rsql.RawField(datas[i], 4)
		if err != nil {
			reurnErr = errors.New("批量添加，获取对象参数错误")
			break
		}

		var valData []interface{}
		for i2 := range val {
			valData = append(valData, val[i2])
		}

		if err := bulkInserter.Insert(valData...); err != nil {
			reurnErr = errors.New(fmt.Sprintf("批量添加，添加失败: %s", valData))
			break
		}
		reurnList++
	}

	if reurnErr != nil {
		return 0, reurnErr
	}

	bulkInserter.Flush()
	fmt.Println("批量添加，添加成功,完毕")

	return reurnList, nil
}

func (d *defaultModel[T]) FindList() (any, error) {
	//sqlx.NewBulkInserter()
	//for i := range datas {
	//	//开启事务 Transact() 都会自动回滚事务
	//	err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//
	//		_, err := d.Insert(datas[i])
	//		// 返回任何错误都会回滚事务
	//		if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//			logx.Errorf("insert userinfo stmt exec: %s", err)
	//			return err
	//		}
	//
	//		// 还可以继续执行 insert/update/delete 相关操作
	//		return nil
	//	})
	//
	//}
	//
	////开启事务 出现错误直接回滚
	//err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//	for i := range datas {
	//		data := datas[i]
	//		//获取可以传入的key
	//		var key string
	//		var value []interface{}
	//		if keys, err := rsql.RawField(data, 0); err == nil {
	//			key = strings.Join(keys, ",")
	//		} else {
	//			return errors.New("Insert 获取key失败 ")
	//		}
	//
	//		var valueDataString []string
	//
	//		//获取传入的value
	//		if values, err := rsql.RawField(data, 1); err == nil {
	//
	//			for i := range values {
	//				value = append(value, values[i])
	//				var item = ""
	//
	//				if i == 0 {
	//					item += "("
	//				}
	//
	//				item += "?"
	//
	//				if i == len(values) {
	//					item += ")"
	//				}
	//
	//				valueDataString = append(valueDataString, item)
	//			}
	//		} else {
	//			return errors.New("Insert 获取value失败 ")
	//		}
	//
	//		v := strings.Join(valueDataString, ",")
	//
	//		query := fmt.Sprintf("insert into %s (%s) values %s", d.table, key, v)
	//
	//	}
	//	stmt, err := session.Prepare(insertsql)
	//	if err != nil {
	//		return err
	//	}
	//	defer stmt.Close()
	//
	//	sqlx.NewBulkInserter()
	//
	//	// 返回任何错误都会回滚事务
	//	if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//		logx.Errorf("insert userinfo stmt exec: %s", err)
	//		return err
	//	}
	//
	//	// 还可以继续执行 insert/update/delete 相关操作
	//	return nil
	//})

	//err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//	//一段完整逻辑
	//	stmt, err := session.Prepare(insertsql)
	//	if err != nil {
	//		return err
	//	}
	//	defer stmt.Close()
	//
	//	// 返回任何错误都会回滚事务
	//	if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//		logx.Errorf("insert userinfo stmt exec: %s", err)
	//		return err
	//	}
	//	//一段完整逻辑结束  可以开始下一段
	//
	//	session.QueryRowsPartialCtx()
	//
	//	// 还可以继续执行 insert/update/delete 相关操作
	//	return nil
	//})

	return nil, nil
}

func (d *defaultModel[T]) UpdateList(datas []any) error {
	//sqlx.NewBulkInserter()
	//for i := range datas {
	//	//开启事务 Transact() 都会自动回滚事务
	//	err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//
	//		_, err := d.Insert(datas[i])
	//		// 返回任何错误都会回滚事务
	//		if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//			logx.Errorf("insert userinfo stmt exec: %s", err)
	//			return err
	//		}
	//
	//		// 还可以继续执行 insert/update/delete 相关操作
	//		return nil
	//	})
	//
	//}
	//
	////开启事务 出现错误直接回滚
	//err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//	for i := range datas {
	//		data := datas[i]
	//		//获取可以传入的key
	//		var key string
	//		var value []interface{}
	//		if keys, err := rsql.RawField(data, 0); err == nil {
	//			key = strings.Join(keys, ",")
	//		} else {
	//			return errors.New("Insert 获取key失败 ")
	//		}
	//
	//		var valueDataString []string
	//
	//		//获取传入的value
	//		if values, err := rsql.RawField(data, 1); err == nil {
	//
	//			for i := range values {
	//				value = append(value, values[i])
	//				var item = ""
	//
	//				if i == 0 {
	//					item += "("
	//				}
	//
	//				item += "?"
	//
	//				if i == len(values) {
	//					item += ")"
	//				}
	//
	//				valueDataString = append(valueDataString, item)
	//			}
	//		} else {
	//			return errors.New("Insert 获取value失败 ")
	//		}
	//
	//		v := strings.Join(valueDataString, ",")
	//
	//		query := fmt.Sprintf("insert into %s (%s) values %s", d.table, key, v)
	//
	//	}
	//	stmt, err := session.Prepare(insertsql)
	//	if err != nil {
	//		return err
	//	}
	//	defer stmt.Close()
	//
	//	sqlx.NewBulkInserter()
	//
	//	// 返回任何错误都会回滚事务
	//	if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//		logx.Errorf("insert userinfo stmt exec: %s", err)
	//		return err
	//	}
	//
	//	// 还可以继续执行 insert/update/delete 相关操作
	//	return nil
	//})

	//err := d.CachedConn.Transact(func(session sqlx.Session) error {
	//	//一段完整逻辑
	//	stmt, err := session.Prepare(insertsql)
	//	if err != nil {
	//		return err
	//	}
	//	defer stmt.Close()
	//
	//	// 返回任何错误都会回滚事务
	//	if _, err := stmt.Exec(uid, username, mobilephone); err != nil {
	//		logx.Errorf("insert userinfo stmt exec: %s", err)
	//		return err
	//	}
	//	//一段完整逻辑结束  可以开始下一段
	//
	//	session.QueryRowsPartialCtx()
	//
	//	// 还可以继续执行 insert/update/delete 相关操作
	//	return nil
	//})
	return nil
}

// 批量删除数据库数据，同时删除缓存
func (d *defaultModel[T]) DeleteList(ids []uint64) error {
	var idkeys []string
	var idvalues []string
	var args []interface{}
	for i := range ids {
		idkeys = append(idkeys, fmt.Sprintf("%s%v", d.cacheIdPrefix, ids[i]))
		args = append(args, ids[i])
		idvalues = append(idvalues, strconv.FormatUint(ids[i], 10))
	}

	_, err := d.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` in (%s)", d.table, strings.Join(idvalues, ","))
		return conn.Exec(query, args...)
	}, idkeys...)
	return err
}

// 获取缓存key字符串
func (d *defaultModel[T]) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", d.cacheIdPrefix, primary)
}

// 根据主键查询查看是否存在数据库中
func (d *defaultModel[T]) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", d.fieldNameRows, d.table)
	return conn.QueryRow(v, query, primary)
}
