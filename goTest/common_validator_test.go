package goTest

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"testing"
)

// 结构体验证
type User1 struct {
	Name    string `json:"name" validate:"min=2,max=3"`
	Age     int    `json:"Age" validate:"eqcsfield"`
	Social1 Social `json:"social"`
}

type Social1 struct {
	Zhihu string `json:"zhihu"`
	Weibo string `json:"weibo"`
	Xss   string `json:"xss"`
}

func TestValidatorPast(t *testing.T) {
	//user := &User1{
	//	Name:    "",
	//	Age:     0,
	//	Social1: Social{},
	//}
	//validate := validator.New() //初始化（赋值）
	//err := validate.Struct(user)
	//if err != nil {
	//	fmt.Println("*")
	//}
	zhCh := zh.New()
	uni := ut.New(zhCh)                 // 万能翻译器，保存所有的语言环境和翻译数据
	trans, _ := uni.GetTranslator("zh") // 翻译器
	var str = "one"
	err := trans.Add(str, "{0} 为必填字段!", true)
	if err == nil {
		return
	}
	fmt.Println(str)
	fmt.Println(trans.Locale())

}

const (
	BadRequest = "请求参数错误"
)

// ValidateData 全局model数据验证器
func Validate(dataStruct interface{}) error {

	//验证
	zh_ch := zh.New()
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}

type user struct {
	Id   int64  `validate:"required" label:"id" ignoreRequired:"true"`
	Name string `validate:"required" label:"姓名"`
}

func TestValidator(t *testing.T) {
	//var u user
	//u.Id = 0
	//u.Name = ""
	//if err := validator.Validate(user); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//return
}

// 初始化翻译器
func validateInit() {
	zhCh := zh.New()
	uni := ut.New(zhCh)                 // 万能翻译器，保存所有的语言环境和翻译数据
	Trans, _ := uni.GetTranslator("zh") // 翻译器
	Validate := validator.New()
	_ = zh_translations.RegisterDefaultTranslations(Validate, Trans)
	// 添加额外翻译
	_ = Validate.RegisterTranslation("required_without", Trans, func(ut ut.Translator) error {
		return ut.Add("required_without", "{0} 为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without", fe.Field())
		return t
	})
}

// 实例化验证对象
var validate = validator.New()

func TestValidator1(t *testing.T) {
	validateInit()
	// 结构体验证
	type Inner struct {
		String string `validate:"contains=111"`
	}
	inner := &Inner{String: "11@"}
	errs := validate.Struct(inner)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	// 变量验证
	m := map[string]string{"": "", "val3": "val3"}
	errs = validate.Var(m, "required,dive,keys,required,endkeys,required")
	if errs != nil {
		fmt.Println(errs.Error())
	}
}
