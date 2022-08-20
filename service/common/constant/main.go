package constant

const (
	//网络标识
	NetHttp  = "http"
	NetHttps = "https"

	//系统访问权限内容
	AuthContents = "capi" //目录：打开菜单内容
	AuthMenu     = "mapi" //菜单：菜单页面跳转
	AuthButton   = "bapi" //按钮：按钮页面跳转
	AuthApi      = "iapi" //接口：数据内容

	//token
	TokenExpire       = "exp"          //有效时间
	TokenIat          = "iat"          //当前时间
	TokenRefreshAfter = "refreshAfter" //刷新时间
	TokenSecret       = "secret"       //共有密钥
	TokenPrevSecret   = "prevSecret"   //私有密钥

	//系统使用
	SysSpace  = ""
	SysUserId = "userId"

	//error
	ErrGetToken = "创建token失败,请检查参数"

	//DB 自己写的方法中 0是获取db中key 1是或value
	DbKey   = 0
	DbValue = 1

	//业务数据内容
	LogicLog = 5

	//I18n语言内容
	I18nEn = "en"
	I18nZh = "zh"
)
