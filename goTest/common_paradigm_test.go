package goTest

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"testing"
)

// 泛型测试
func TestId1(t *testing.T) {

	fmt.Printf("%v", Age[Name]())
}

func Age[T any]() []string {
	var t T
	return builder.RawFieldNames(t)
}

type in interface {
	ins
}

type ins interface {
	instf()
}

type inst struct {
}

func (i inst) instf() {
}

type ins1 interface {
	instf1()
}

type inst1 struct {
}

func (i inst1) instf1() {
	fmt.Println("*****************")
}

func TestAny(t *testing.T) {
	Push()

}

func Push(i ...int) *inst {
	fmt.Println(i)
	return &inst{}
}

type a1 struct {
}

type a2 struct {
	a1
}
