package goTest

import (
	"fmt"
	err "github.com/daida459031925/common/error"
	"github.com/daida459031925/common/translate"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"testing"
)

// 结构体验证
type User1 struct {
	Name    string `json:"name" validate:"min=2,max=3"`
	Age     int    `json:"Age" validate:"lt=-1|lt=-10"`
	Social1 Social `json:"social"`
}

type Social1 struct {
	Zhihu string `json:"zhihu"`
	Weibo string `json:"weibo"`
	Xss   string `json:"xss"`
}

func TestValidatorPast(t *testing.T) {
	validate = validator.New() //初始化（赋值）
	validateStruct()           //结构体校验
	validateVariable()         //变量校验

	user := &User1{
		Name:    "代打",
		Age:     0,
		Social1: Social{},
	}
	validate := validator.New() //初始化（赋值）
	err := validate.Struct(user)
	if err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Namespace:", e.Namespace())
			fmt.Println("Field:", e.Field())
			fmt.Println("StructNamespace:", e.StructNamespace())
			fmt.Println("StructField:", e.StructField())
			fmt.Println("Tag:", e.Tag())
			fmt.Println("ActualTag:", e.ActualTag())
			fmt.Println("Kind:", e.Kind())
			fmt.Println("Type:", e.Type())
			fmt.Println("Value:", e.Value())
			fmt.Println("Param:", e.Param())
			fmt.Println()
		}

		fmt.Println("*", err)
	}

	//zhCh := zh.New()
	//uni := ut.New(zhCh)                 // 万能翻译器，保存所有的语言环境和翻译数据
	//trans, _ := uni.GetTranslator("zh") // 翻译器
	//var str = "one"
	//err := trans.Add(str, "{0} 为必填字段!", true)
	//if err == nil {
	//	return
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
	e := trans.Add(str, "{0} 为必填字段!", true)
	if e == nil {
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
	zhCh := zh.New()
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zhTrans.RegisterDefaultTranslations(validate, trans)
	validateErr := validate.Struct(dataStruct)
	if validateErr != nil {
		for _, rangeErr := range validateErr.(validator.ValidationErrors) {
			return err.New(rangeErr.Translate(trans))
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
	_ = zhTrans.RegisterDefaultTranslations(Validate, Trans)
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
		validatorErr := errs.(validator.ValidationErrors)
		fmt.Println(validatorErr.Translate(nil))
	}
}

func TestValiTranslate(t *testing.T) {
	// 1.新建需要的语言
	en := en.New() //英文翻译器
	zh := zh.New() //中文翻译器

	// 2.将语言配置到翻译器中
	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)

	// 3.获取需要的语言
	trans, _ := uni.GetTranslator("zh")
	user := User1{
		Name: "tom",
		Age:  25,
	}
	validate := validator.New()
	// 4.把翻译器注册到validate中
	zhTrans.RegisterDefaultTranslations(validate, trans)
	StructErr := validate.Struct(user)
	if StructErr != nil {
		// fmt.Println(err)

		errs := StructErr.(validator.ValidationErrors)
		// 5.翻译错误信息
		fmt.Println()
		fmt.Println()
		s1 := fmt.Sprintf("%s", removeStructName(errs.Translate(trans)))
		s2 := fmt.Sprintf("%s", errs.Translate(trans))
		str, e := translate.TranslateConversion(s1)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(str)
		str, e = translate.TranslateConversion(s2)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(str)
	}

}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, e := range fields {
		result[field[strings.Index(field, ".")+1:]] = e
	}
	return result
}

func validateStruct() {
	var validate *validator.Validate //定义

	type Address struct {
		Province string `validate:"required"`       //非空
		City     string `validate:"required"`       //非空
		Phone    string `validate:"numeric,len=11"` //数字类型，长度为11
	}

	type User struct {
		Name  string `validate:"required"`       //非空
		Age   uint8  `validate:"gte=0,lte=130"`  //  0<=Age<=130
		Email string `validate:"required,email"` //非空，email格式
		//dive关键字代表 进入到嵌套结构体进行判断
		Address []*Address `validate:"dive"` //  可以拥有多个地址
	}

	address := Address{
		Province: "重庆",
		City:     "重庆",
		Phone:    "13366663333x",
	}
	user := User{
		Name:    "江洲",
		Age:     23,
		Email:   "jz@163.com",
		Address: []*Address{&address},
	}
	err := validate.Struct(user)
	if err != nil {
		//断言为：validator.ValidationErrors，类型为：[]FieldError
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Namespace:", e.Namespace())
			fmt.Println("Field:", e.Field())
			fmt.Println("StructNamespace:", e.StructNamespace())
			fmt.Println("StructField:", e.StructField())
			fmt.Println("Tag:", e.Tag())
			fmt.Println("ActualTag:", e.ActualTag())
			fmt.Println("Kind:", e.Kind())
			fmt.Println("Type:", e.Type())
			fmt.Println("Value:", e.Value())
			fmt.Println("Param:", e.Param())
			fmt.Println()
		}

		fmt.Println("结构体输入数据类型错误！")
		return
	} else {
		fmt.Println("结构体校验通过")
	}
}

// 变量校验
func validateVariable() {
	myEmail := "123@qq.com" //邮箱地址：xx@xx.com
	err := validate.Var(myEmail, "required,email")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("变量校验通过！")
	}
}
