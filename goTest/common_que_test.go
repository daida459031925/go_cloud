package goTest

import (
	err "github.com/daida459031925/common/error"
	"github.com/daida459031925/common/fmt"
	result "github.com/daida459031925/common/result"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

// 编写工具链表执行器
type abc struct {
	a int
}

func (t abc) getData() {
	logx.Info(t.a)
}

// 目的为了测试步骤器的前置条件 不管指针还是非指针类型返还自己想要的
func TestParseUnPointer(t *testing.T) {
	var ad = abc{1}
	var asd any = ad
	a, _ := fmt.ParseUnPointer[abc](asd)
	//fmt.Println(e)
	//fmt.Println(a)
	a.a = 2
	logx.Info(a)
	logx.Info(asd)
	logx.Info("*************************************")

	ad = abc{1}
	asd = &ad
	a1, _ := fmt.ParseUnPointer[abc](asd)
	//fmt.Println(e1)
	//fmt.Println(a1)
	a1.a = 2
	logx.Info(a1)
	logx.Info(asd)
	logx.Info("*************************************")

	var aa = 11
	var bb = &aa

	cc, _ := fmt.ParseUnPointer[int](bb)
	cc = 2
	logx.Info(cc)
	logx.Info(*bb)
	logx.Info("*************************************")

	var bb1 = aa

	cc1, _ := fmt.ParseUnPointer[int](bb1)
	cc1 = 22
	logx.Info(cc1)
	logx.Info(bb1)
	logx.Info("*************************************")

}

// 测试工具步骤器的返还正常和非正常是否达到我要的效果
func TestResult(t *testing.T) {
	r := result.Ok().
		SetFuncErr(func(a any) any {
			d := 111
			//i := 1
			//b := 0
			//c := i / b
			//logx.Info(c)
			return &d
		}, err.New("打印")).
		SetFuncErr(func(a any) any {
			d := 222
			//i := 1
			//b := 0
			//c := i / b
			//logx.Info(c)
			return &d
		}, err.New("cursor打印1")).
		Exec()
	logx.Info(r.Date)
	logx.Info(r.Msg)
	logx.Info(r.Status)
	logx.Info(fmt.ParseUnPointer[int](r.Data))

}
