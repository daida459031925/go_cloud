package validatorUtil

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// 饿汉式
var vali *validator.Validate

const ()

func init() {
	fmt.Println("验证是不是全局只走一次")
	vali = validator.New()
}

func getVali() *validator.Validate {
	return vali
}

// 校验完整结构体
func ValiStruct(data any) (error, any) {
	err := vali.Struct(data)
	errMap := make(map[string]string)

	if err == nil {
		return nil, errMap
	}

	//断言为：validator.ValidationErrors，类型为：[]FieldError
	for _, e := range err.(validator.ValidationErrors) {
		//写入
		errMap[e.Field()] = e.ActualTag()
		//fmt.Println("Namespace:", e.Namespace())
		//fmt.Println("Field:", e.Field())
		//fmt.Println("StructNamespace:", e.StructNamespace())
		//fmt.Println("StructField:", e.StructField())
		//fmt.Println("Tag:", e.Tag())
		//fmt.Println("ActualTag:", e.ActualTag())
		//fmt.Println("Kind:", e.Kind())
		//fmt.Println("Type:", e.Type())
		//fmt.Println("Value:", e.Value())
		//fmt.Println("Param:", e.Param())
		//fmt.Println()
	}

	return err, errMap
}

// 校验单个字段
func ValiVariable(str string, validate string) bool {
	err := vali.Var(str, validate)
	if err != nil {
		return false
	}
	return true
}
