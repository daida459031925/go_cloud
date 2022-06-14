package goTest

import (
	"fmt"
	"github.com/daida459031925/common/error/try"
	"testing"
)

//try测试
func TestTry(t *testing.T) {

	try.Try(func() {
		a, b := 1, 0
		var c = a / b
		fmt.Sprintf("Try1 start %d ", c)
	}).CatchAll(func(err error) {
		fmt.Println("Try1 Err1 Catch:", err.Error())
	}).Finally(func() {
		fmt.Println("Try1 done")
	})
}
