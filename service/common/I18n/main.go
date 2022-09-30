package I18n

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"service/common/I18n/en"
	"service/common/I18n/zh"
	"service/common/constant"
)

var i18nMap = make(map[string]*message.Printer)

func initNewPrinter() {
	i18nMap[constant.UseTranslateZh] = message.NewPrinter(language.Chinese)
	i18nMap[constant.UseTranslateEn] = message.NewPrinter(language.English)
	logx.Info(constant.UseI18nAll)
}

// TransInterface 语言的公用统一实现方法
type TransInterface interface {
	GetTransData() (func() (trans ut.Translator, found bool), func(trans ut.Translator) error)
}

// GetI18nMap 特定单词翻译 默认还是使用自带翻译
func GetI18nMap() map[string]*message.Printer {
	return i18nMap
}

type transStruct struct {
	//为了统一一个入口返还数据 进行抽象
	transInterface TransInterface
	//返还新这个新对象，未来还要结合i18n一起返还
	Trans ut.Translator
}

// GetTrans 需要一个参数验证器 那么返还一个已经制定好是语言的翻译器
func (transStruct transStruct) getTrans() (ut.Translator, error) {
	a, b := transStruct.transInterface.GetTransData()

	//不知道什么原因 写了zh 返还的Boolean值还是false 也就是不管正确还是错误都是错误，所以先注释
	Trans, _ := a()
	//if !tf {
	//	//未找到指定内容
	//	return nil, errors.New("未找到指定类型翻译器")
	//}
	e := b(Trans)
	if e != nil {
		return nil, errors.New("无法将翻译器进行内容翻译")
	}

	return Trans, nil
}

/**************************创建各类翻译器***********************************/
func NewEn(validate *validator.Validate) (*transStruct, error) {
	et := &en.Tran{Validate: validate}
	t1 := transStruct{transInterface: et}
	a, b := t1.getTrans()
	t1.Trans = a
	return &t1, b
}

func NewZh(validate *validator.Validate) (*transStruct, error) {
	zt := &zh.Tran{Validate: validate}
	t1 := transStruct{transInterface: zt}
	a, b := t1.getTrans()
	t1.Trans = a
	return &t1, b
}

/**************************创建各类翻译器***********************************/
