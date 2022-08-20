package I18n

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"service/common/constant"
)

var i18nMap = make(map[string]*message.Printer)

func initNewPrinter() {
	i18nMap[constant.I18nZh] = message.NewPrinter(language.Chinese)
	i18nMap[constant.I18nEn] = message.NewPrinter(language.English)
	logx.Info("国际化翻译器初始化成功")
}

func GetI18nMap() map[string]*message.Printer {
	return i18nMap
}
