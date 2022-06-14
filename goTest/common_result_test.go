package goTest

import (
	"fmt"
	"github.com/daida459031925/common/result"
	"testing"
)

//返还对象
func TestBuild(t *testing.T) {
	build := result.Build(200, "", nil)
	fmt.Println(build)
}
