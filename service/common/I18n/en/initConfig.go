package en

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/zeromicro/go-zero/core/logx"
	"service/common/constant"
)

func InitEn() {

	logx.Info(constant.UseI18nEn)
}

type Tran struct {
	Validate *validator.Validate
}

func (tran Tran) GetTransData() (func() (trans ut.Translator, found bool), func(trans ut.Translator) error) {
	return func() (trans ut.Translator, found bool) {
			enCh := en.New()
			// 万能翻译器，保存所有的语言环境和翻译数据
			uni := ut.New(enCh)
			// 翻译器
			return uni.GetTranslator(constant.UseTranslateEn)
		}, func(trans ut.Translator) error {
			return enTranslations.RegisterDefaultTranslations(tran.Validate, trans)
		}
}
