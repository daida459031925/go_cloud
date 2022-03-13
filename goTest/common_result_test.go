package goTest

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	build := result.Build(200, "", nil)
	fmt.Println(build)
}
