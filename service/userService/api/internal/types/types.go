// Code generated by goctl. DO NOT EDIT.
package types

type Login struct {
	Username string `json:"username" validate:"required,min=8,max=30"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type LogoutUserToken struct {
	Id uint64 `json:"id"`
}

type RUserToken struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire uint64 `json:"accessExpire"`
	RefreshAfter uint64 `json:"refreshAfter"`
}

type AddDBSysUser struct {
	Deleted         int    `json:"deleted,default=1"`                                                 // 标识是否删除1存在 2删除
	Account         string `json:"account" validate:"required,min=8,max=30"`                          // 账号
	Password        string `json:"password" validate:"required,min=8,max=30,eqfield=comparePassword"` // 密码
	ComparePassword string `json:"comparePassword" validate:"required,min=8,max=30,eqfield=password"` // 密码
	Gender          int    `json:"gender,default=0,options=0|1|2"`                                    // 性别（未知|男|女）
	TokenExpire     uint64 `json:"tokenExpire"`                                                       // token时间 单位秒
}

type UpdDBSysUser struct {
	Id              uint64 `json:"id"`                                                                // 唯一主键id
	Deleted         int    `json:"deleted"`                                                           // 标识是否删除1存在 2删除
	Password        string `json:"password" validate:"required,min=8,max=30,eqfield=comparePassword"` // 密码
	ComparePassword string `json:"comparePassword" validate:"required,min=8,max=30,eqfield=password"` // 密码
	Gender          int    `json:"gender,default=0,options=0|1|2"`                                    // 性别（未知|男|女）
	TokenExpire     uint64 `json:"tokenExpire"`                                                       // token时间 单位秒
}

type RSysUser struct {
	Id             uint64 `json:"id"`             // 唯一主键id
	Type           int    `json:"type"`           // 类型，用于区分用户加入来源：1.手动添加，2批量添加，3微信注册，4pc注册 5，同步,
	Account        string `json:"account"`        // 账号
	Gender         int    `json:"gender"`         // 性别（未知|男|女）
	CreateTime     string `json:"createTime"`     // 创建时间
	UpdateTime     string `json:"updateTime"`     // 最后更新时间
	CreateBy       uint64 `json:"createBy"`       // 创建人
	UpdateBy       uint64 `json:"updateBy"`       // 最后更新人
	CreateByUserId uint64 `json:"createByUserId"` // 为了查询用户，保存id，空间换时间的备用,
	UpdateByUserId uint64 `json:"updateByUserId"` // 为了查询用户，保存id，空间换时间的备用,
}
