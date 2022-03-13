package goTest

import (
	"fmt"
	"go_cloud/common/error/try"
	"testing"
)

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
