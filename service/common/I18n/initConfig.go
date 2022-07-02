package I18n

import (
	"service/common/I18n/en"
	"service/common/I18n/zh"
)

func init() {
	en.InitEn()
	zh.InitZh()
	initNewPrinter()
}
