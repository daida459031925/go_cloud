package sys

import (
	"database/sql"
	"time"
)

// 如果struct名称首字母是小写的，这个struct不会被导出。连同它里面的字段也不会导出，即使有首字母大写的字段名。
// 如果struct名称首字母大写，则struct会被导出，但只会导出它内部首字母大写的字段，那些小写首字母的字段不会被导出。
type (
	// Dict 字典表
	Dict struct {
		Id           uint64         `db:"id"`             // 字典主键
		Deleted      uint64         `db:"deleted"`        // 标识是否删除1存在 2删除
		Status       uint64         `db:"status"`         // 状态（1正常 2停用）
		IsEdit       uint64         `db:"is_edit"`        // 表示当前内容是否可以供普通用户修改（1可以2不行）
		Tp           string         `db:"type"`           // 字典类型
		Name         string         `db:"name"`           // 字典名称
		SuperId      uint64         `db:"super_id"`       // 上级id可以构建树形
		Describe     sql.NullString `db:"describe"`       // 备注
		CssClass     sql.NullString `db:"css_class"`      // 样式属性（其他样式扩展）
		ListClass    sql.NullString `db:"list_class"`     // 表格回显样式
		CreateUserId uint64         `db:"create_user_id"` // 创建者
		CreateTime   time.Time      `db:"create_time"`    // 创建时间
		UpdateUserId sql.NullInt64  `db:"update_user_id"` // 更新者
		UpdateTime   sql.NullTime   `db:"update_time"`    // 更新时间
	}
	// Resources 资源表
	Resources struct {
		Id           uint64         `db:"id"`             // 唯一主键id
		Deleted      uint64         `db:"deleted"`        // 标识是否删除1存在 2删除
		Name         string         `db:"name"`           // 资源名称
		Tp           string         `db:"type"`           // 资源的类型
		Url          sql.NullString `db:"url"`            // 资源路径
		Model        string         `db:"model"`          // 资源请求类型
		SuperId      uint64         `db:"super_id"`       // 上级id可以构建树形
		Describe     sql.NullString `db:"describe"`       // 请求方法内容描述
		OrderBy      uint64         `db:"order_by"`       // 排序
		CreateUserId uint64         `db:"create_user_id"` // 创建者
		CreateTime   time.Time      `db:"create_time"`    // 创建时间
		UpdateUserId sql.NullInt64  `db:"update_user_id"` // 更新者
		UpdateTime   sql.NullTime   `db:"update_time"`    // 更新时间
	}
	// Role 角色表
	Role struct {
		Id           uint64        `db:"id"`             // 唯一主键id
		Deleted      uint64        `db:"deleted"`        // 标识是否删除1存在 2删除
		Name         string        `db:"name"`           // 角色名字用作标识客户看到的名字，可根据不同业务来实现名字是否相同
		SuperId      uint64        `db:"super_id"`       // 上级id可以构建树形
		CreateUserId uint64        `db:"create_user_id"` // 创建者
		CreateTime   time.Time     `db:"create_time"`    // 创建时间
		UpdateUserId sql.NullInt64 `db:"update_user_id"` // 更新者
		UpdateTime   sql.NullTime  `db:"update_time"`    // 更新时间
	}
	// User 用户表
	User struct {
		Id           uint64        `db:"id"`             // 唯一主键id
		Deleted      uint64        `db:"deleted"`        // 标识是否删除1存在 2删除
		Account      string        `db:"account"`        // 账号
		Password     string        `db:"password"`       // 密码
		Salt         string        `db:"salt"`           // 盐值，用作生成密码加密
		DictId       uint64        `db:"dict_id"`        // 性别（男|女|未知）
		CreateUserId uint64        `db:"create_user_id"` // 创建人
		CreateTime   time.Time     `db:"create_time"`    // 创建时间
		UpdateUserId sql.NullInt64 `db:"update_user_id"` // 最后更新人
		UpdateTime   sql.NullTime  `db:"update_time"`    // 最后更新时间
		Secret       string        `db:"secret"`         // 用户随机生成secret，用作token生成
		PrevSecret   string        `db:"prev_secret"`    // 用户私有secret
		TokenExpire  uint64        `db:"token_expire"`   // token时间 单位秒
	}
)
