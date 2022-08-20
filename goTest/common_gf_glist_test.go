package goTest

import (
	"fmt"
	"github.com/daida459031925/common/error/try"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestGlist(t *testing.T) {
	defer d()
	glist := glist.New()
	var index = 100
	for i := 0; i < index; i++ {
		var s = fmt.Sprintf("%d", i)
		glist.PushBack(func() { logx.Info(s) })
	}

	for i := 0; i < index; i++ {
		v := glist.PopBack()
		a, e := v.(func())
		if e {
			try.Try(a).CatchAll(func(err error) {
				logx.Errorf("server func error: %s", err.Error())
			})
		}
	}

}

func d() {
	logx.Info("****")
}
