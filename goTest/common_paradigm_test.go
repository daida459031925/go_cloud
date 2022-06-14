package goTest

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"testing"
)

//泛型测试
func TestId1(t *testing.T) {

	fmt.Printf("%v", Age[Name]())
}

func Age[T any]() []string {
	var t T
	return builder.RawFieldNames(t)
}
