package model

import "database/sql"

type SysUser struct {
	Id             sql.NullInt64  `db:"id"`             // 唯一主键id
	Deleted        int            `db:"deleted"`        // 标识是否删除1存在 2删除
	Type           int            `db:"type"`           // 类型，用于区分用户加入来源：1.手动添加，2批量添加，3微信注册，4pc注册 5，同步,
	Account        sql.NullString `db:"account"`        // 账号
	Password       sql.NullString `db:"password"`       // 密码
	Salt           sql.NullString `db:"salt"`           // 盐值，用作生成密码加密
	Gender         sql.NullInt64  `db:"gender"`         // 性别（男|女|未知）
	CreateTime     sql.NullTime   `db:"createTime"`     // 创建时间
	UpdateTime     sql.NullTime   `db:"updateTime"`     // 最后更新时间
	CreateBy       sql.NullString `db:"createBy"`       // 创建人
	UpdateBy       sql.NullString `db:"updateBy"`       // 最后更新人
	CreateByUserId uint64         `db:"createByUserId"` // 为了查询用户，保存id，空间换时间的备用,
	UpdateByUserId uint64         `db:"updateByUserId"` // 为了查询用户，保存id，空间换时间的备用,
	Secret         sql.NullString `db:"secret"`         // 用户随机生成secret，用作token生成
	PrevSecret     sql.NullString `db:"prevSecret"`     // 用户私有secret
	TokenExpire    sql.NullInt64  `db:"tokenExpire"`    // token时间 单位秒
}
