package goTest

import (
	"fmt"
	"go_cloud/common/result"
	"testing"
)

func TestBuild(t *testing.T) {
	build := result.Build(200, "", nil)
	fmt.Println(build)
}
