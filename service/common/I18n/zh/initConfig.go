package zh

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
	"service/common/constant"
)

func InitZh() {

	logx.Info(constant.UseI18nZn)
}

type Tran struct {
	Validate *validator.Validate
}

func (tran Tran) GetTransData() (func() (trans ut.Translator, found bool), func(trans ut.Translator) error) {
	return func() (trans ut.Translator, found bool) {
			zhCh := zh.New()
			// 万能翻译器，保存所有的语言环境和翻译数据
			uni := ut.New(zhCh)
			// 翻译器
			return uni.GetTranslator(constant.UseTranslateZh)
		}, func(trans ut.Translator) error {
			return zhTranslations.RegisterDefaultTranslations(tran.Validate, trans)
		}
}
