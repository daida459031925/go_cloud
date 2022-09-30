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
	SysComma  = ","
	SysDoubt  = "?"
	SysUserId = "userId"
	SysId     = "id"

	//error
	ErrGetToken00_01 = "创建token失败,请检查参数"

	ErrFuncBQueue02vs_01 = "当前类型是：%v,服务器执行错误: %s"
	ErrFuncBQueue01v_01  = "当前类型是：%v,当前任务不是方法无法执行"

	ErrClientPulsar00_01     = "pulsar连接失败"
	ErrClientPulsar01s_01    = "pulsar连接失败: %s"
	ErrPulsarProducer01s_01  = "pulsar创建生产者失败: %s"
	ErrPulsarSubscribe01s_01 = "pulsar创建消费者失败: %s"

	ErrSqlInsert00_01      = "Insert 获取key失败 "
	ErrSqlInsert00_02      = "Insert 获取value失败 "
	ErrSqlUpdate00_01      = "Update 获取key失败 "
	ErrSqlUpdate00_02      = "Update 获取value失败 "
	ErrSqlInsertList00_01  = "InsertList，获取key错误"
	ErrSqlInsertList00_02  = "InsertList，连接错误"
	ErrSqlInsertList00_03  = "InsertList，获取对象参数错误"
	ErrSqlInsertList00_04  = "InsertList，完毕"
	ErrSqlInsertList01s_05 = "InsertList，添加失败: %s"
	ErrSqlFind00_01        = "Find 获取value失败 "

	ErrDecimal01s_01 = "转换float64失败: %s"

	ErrValidator00_01 = "结构体参数规范验证失败"

	ErrAllErrorInit00_01 = "全局err打印位置为："
	ErrAllXssInit00_01   = "全局XSS打印位置为："
	ErrhttpxParse00_01   = "参数解析异常"

	//基本使用信息
	UseClientPulsar01s_01           = "pulsar://%s"
	UsePulsarConsumerMessage02vs_01 = "已经收到消息，消息id: %v -- 内容为: '%s'"

	UseI18nEn  = "英文翻译初始化完毕"
	UseI18nZn  = "中文翻译初始化完毕"
	UseI18nAll = "国际化翻译器初始化成功"

	//I18n语言内容
	UseTranslateEn = "en"
	UseTranslateZh = "zh"

	UseSqlCache = "cache:%s:id:"

	//业务正常错误返还
	ResultLoginErr01 = "账号或密码错误"

	//DB 自己写的方法中 0是获取db中key 1是或value
	DbKey   = 0
	DbValue = 1

	//业务数据内容
	LogicLog = 5
)
