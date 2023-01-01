package goTest

import (
	"database/sql"
	rsql "github.com/daida459031925/common/reflex/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"
)

// • EXPLAIN不会告诉你关于触发器、存储过程的信息或用户自定义函数对查询的影响情况
// • EXPLAIN不考虑各种Cache
// • EXPLAIN不能显示MySQL在执行查询时所作的优化工作
// • 部分统计信息是估算的，并非精确值
// • EXPLAIN只能解释SELECT操作，其他操作要重写为SELECT后查看执行计划。
type explain struct {
	Id           sql.NullInt64  `db:"id"`            //查询的序列号；查询中执行 select 子句或操作表的顺序,id值越大优先级越高,越先被执行。id相同,执行顺序由上至下。
	SelectType   string         `db:"select_type"`   //查询的类型，主要包括普通查询、联合查询、子查询。
	Table        string         `db:"table"`         //所访问数据库的表的名称
	Partitions   sql.NullString `db:"partitions"`    //如果查询是基于分区表的话, 会显示查询访问的分区
	Types        string         `db:"type"`          //联合查询使用的类型显示的访问类型是重要指标，结果从好到坏
	PossibleKeys sql.NullString `db:"possible_keys"` //指出mysql能使用哪个索引有助于在该表中找到该行，如果这个值为空，则表示没有相关的索引。
	Key          sql.NullString `db:"key"`           //显示mysql实际决定使用的建，如果没有索引被选择，键是null
	KeyLen       sql.NullInt64  `db:"key_len"`       //显示mysql决定使用索引的长度，在不损失精确性的情况下,长度越短越好。如果key是空，则长度就是null。
	Ref          sql.NullString `db:"ref"`           //显示哪个字段或常数与key一起使用。如果使用了索引则显示const，否则，显示null。
	Rows         sql.NullInt64  `db:"rows"`          //MYSQL认为必须检查的用来返回请求数据的行数。
	Filtered     sql.NullString `db:"filtered"`      //满足查询的记录数量的比例，注意是百分比，不是具体记录数 . 值越大越好
	Extra        sql.NullString `db:"extra"`         //出现以下2项意味着MYSQL根本不能使用索引,效率会受到重大影响。应尽可能对此进行优化。
	//TypeLv       int            //用来记录type等级分数，从十分到零分
}

func getSelectType(selectType string) string {
	s := ""
	switch selectType {
	case "SIMPLE":
		s = "简单的select查询,不使用union及子查询"
	case "PRIMARY":
		s = "最外层的select查询"
	case "UNION":
		s = "UNION中的第二个或随后的select查询,不依赖于外部查询的结果集"
	case "DEPENDENT UNION":
		s = "UNION中的第二个或随后的select查询,依赖于外部查询的结果集"
	case "SUBQUERY":
		s = "子查询中的第一个select查询,不依赖于外部查询的结果集"
	case "DEPENDENT SUBQUERY":
		s = "子查询中的第一个select查询,依赖于外部查询的结果集"
	case "DERIVED":
		s = "用于from子句里有子查询的情况。MySQL会递归执行这些子查询,把结果放在临时表里"
	case "UNCACHEABLE SUBQUERY":
		s = "结果集不能被缓存的子查询,必须重新为外 层查询的每一行进行评估"
	case "UNCACHEABLE UNION":
		s = "UNION中的第二个或随后的select 查询,属于不可缓存的子查询"
	default:
		s = selectType
	}
	return s
}

func getTypes(types string) (string, int) {
	s := ""
	i := 0
	switch types {
	case "system":
		s = "系统表；表仅有一行(=系统表)。这是const连接类型的一个特例"
		i = 12
	case "const":
		s = "读常量；const用于用常数值比较PRIMARY_KEY时。当查询的表仅有一行时,使用System"
		i = 11
	case "eq_ref":
		s = "最多一条匹配结果，通常是通过主键访问"
		i = 10
	case "ref":
		s = "连接不能基于关键字选择单个行,可能查找到多个符合条件的行。叫做ref是因为索引要跟某个参考值相比较。这个参考值或者是一个常数,或者是来自一个表里的多表查询的结果值"
		i = 9
	case "fulltext":
		s = "全文索引检索"
		i = 8
	case "ref_or_null":
		s = "如同ref,但是MySQL必须在初次查找的结果里找出null条目,然后进行二次查找"
		i = 7
	case "index_merge":
		s = "合并索引结果集；说明索引合并优化被使用了"
		i = 6
	case "unique_subquery":
		s = "在某些IN查询中使用此种类型,而不是常规的 ref:value IN (SELECT primary_key FROM single_table WHERE some_expr)"
		i = 5
	case "index_subquery":
		s = "子查询返回的是索引，但非主键；在某些IN查询中使用此种类型,与unique_subquery类似,但是查询的是非唯一性索引:value IN (SELECT key_column FROM single_table WHERE some_expr)"
		i = 4
	case "range":
		s = "索引范围扫描；只检索给定范围的行,使用一个索引来选择行。key列显示使用了哪个索引。当使用=、 <>、>、>=、<、<=、IS NULL、<=>、BETWEEN或者IN操作符,用常量比较关键字列时,可以使用range。"
		i = 3
	case "index":
		s = "全索引扫描；全表扫描,只是扫描表的时候按照索引次序进行,而不是行。主要优点就是避免了排序, 但是开销仍然非常大。"
		i = 2
	case "all":
		s = "全表扫描；最坏的情况,从头到尾全表扫描"
		i = 1
	default:
		s = types
	}
	return s, i
}

func getExtra(extra string) string {
	s := ""
	switch extra {
	case "Using filesort":
		s = "表示MySQL会对结果使用一个外部索引排序,而不是从表里按索引次序读到相关内容。可能在内存或者磁盘上进行排序。MySQL中无法利用索引完成的排序操作称为“文件排序”。"
	case "Using temporary":
		s = "表示MySQL在对查询结果排序时使用临时表。常见于排序order by和分组查询group by。"
	case "impossible where":
		s = "条件永远不成立"
	default:
		s = extra
	}
	return s
}

// 已经可以返还sql执行计划的内容，现在需要设计如何提示用户sql的合格度
func TestSql(t *testing.T) {
	conn := sqlx.NewMysql("root:MYSQL@tcp(192.168.0.100:3306)/iotcloud_xj?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	var exp []explain
	e := conn.QueryRows(&exp, "EXPLAIN SELECT * from user")
	if e == nil {
		logx.Info(rsql.RawField(exp, 4))
	}

}
