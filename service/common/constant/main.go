package constant

const (
	//网络标识
	NET_HTTP  = "http"
	NET_HTTPS = "https"

	//系统访问权限内容
	AUTH_CONTENTS = "capi" //目录：打开菜单内容
	AUTH_MENU     = "mapi" //菜单：菜单页面跳转
	AUTH_BUTTON   = "bapi" //按钮：按钮页面跳转
	AUTH_API      = "iapi" //接口：数据内容

	//token
	TOKEN_EXPIRE        = "exp"          //有效时间
	TOKEN_IAT           = "iat"          //当前时间
	TOKEN_REFRESH_AFTER = "refreshAfter" //刷新时间
	TOKEN_SECRET        = "secret"       //共有密钥
	TOKEN_PREV_SECRET   = "prevSecret"   //私有密钥

	//系统使用
	SYS_SPACE   = ""
	SYS_USER_ID = "userId"

	//error
	ERR_GET_TOKEN = "创建token失败,请检查参数"

	//DB 自己写的方法中 0是获取db中key 1是或value
	DB_KEY   = 0
	DB_VALUE = 1

	//业务数据内容
	LogicLog = 5
)
