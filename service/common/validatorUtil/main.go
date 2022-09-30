package validatorUtil

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"service/common/I18n"
	"service/common/constant"
	"sync"
)

// go 中的互斥锁 还有读写锁RWMutex
var mutex sync.Mutex
var myVal *val = nil

type val struct {
	validate *validator.Validate //参数验证
}

func NewValidate() *val {
	if myVal != nil {
		return myVal
	}
	mutex.Lock()
	defer mutex.Unlock()
	if myVal != nil {
		return myVal
	}
	myVal = &val{validator.New()}
	return myVal
}

// ValStruct 校验完整结构体
func (v val) ValStruct(data any) (any, error) {
	err := v.validate.Struct(data)
	errMap := make(map[string]string)

	if err == nil {
		return errMap, nil
	}

	transStruct, i18ne := I18n.NewZh(v.validate)
	//断言为：validator.ValidationErrors，类型为：[]FieldError
	for _, e := range err.(validator.ValidationErrors) {
		//写入
		errString := e.Error()
		if i18ne == nil {
			errString = e.Translate(transStruct.Trans)
		}
		errMap[e.Field()] = errString
		//logx.Info("Namespace:", e.Namespace())
		//logx.Info("Field:", e.Field())
		//logx.Info("StructNamespace:", e.StructNamespace())
		//logx.Info("StructField:", e.StructField())
		//logx.Info("Tag:", e.Tag())
		//logx.Info("ActualTag:", e.ActualTag())
		//logx.Info("Kind:", e.Kind())
		//logx.Info("Type:", e.Type())
		//logx.Info("Value:", e.Value())
		//logx.Info("Param:", e.Param())
		//logx.Info("\\n")
	}

	return errMap, errors.New(constant.ErrValidator00_01)
}

// ValVariable 校验单个字段
func (v val) ValVariable(str string, validateString string) bool {
	err := v.validate.Var(str, validateString)
	if err != nil {
		return false
	}
	return true
}
