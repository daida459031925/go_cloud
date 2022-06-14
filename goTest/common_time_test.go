package goTest

import (
	"fmt"
	"testing"
	"time"
)

//时间测试
func TestMonth(t *testing.T) {
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Month().String())
}
