package bqueue

import (
	"github.com/daida459031925/common/error/try"
	"github.com/daida459031925/common/reflex"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
	"service/common/constant"
)

type (
	BQueue struct {
		// GetFunc 需要关闭程序的集合对象，目标是关闭操作对象集合，由最后一个添加进入的方法进行关闭
		funcList *glist.List
	}

	//为了防止编写两个很重复的方法使用组合对象实现防止重复代码过多
	pop interface {
		getFunc(l *glist.List) any
	}

	popFront struct {
	}

	popBack struct {
	}
)

func NewFuncList() *BQueue {
	return &BQueue{glist.New()}
}

// SetPopBackFunc 从屁股开始加
func (bq BQueue) SetPopBackFunc(f func()) {
	bq.funcList.PushBack(f)
}

// SetPopFrontFunc 从头开始插入
func (bq BQueue) SetPopFrontFunc(f func()) {
	bq.funcList.PushFront(f)
}

// ExecutePopBack 以栈的模式先进后出
func (bq BQueue) ExecutePopBack() {
	execute(bq.funcList, new(popBack))
}

// ExecutePopFront 以队列的模式先进先出
func (bq BQueue) ExecutePopFront() {
	execute(bq.funcList, new(popFront))
}

func (popFront) getFunc(l *glist.List) any {
	return l.PopFront()
}

func (popBack) getFunc(l *glist.List) any {
	return l.PopBack()
}

// 为了减少重复代码编写以及维护难度提高方法只写一处，执行方法使用pop组合对象方式减少代码重复度
func execute(g *glist.List, p pop) {
	try.Try(func() {
		v := reflex.GetRef(p).GetPointerData(false).Type()
		fx.From(func(source chan<- any) {
			if g != nil && g.Len() > 0 {
				for i := 0; i < g.Len(); i++ {
					source <- p.getFunc(g)
				}
			}
		}).ForEach(func(item interface{}) {
			r, e := item.(func())
			if e {
				try.Try(r).CatchAll(func(err error) {
					logx.Errorf(constant.ErrFuncBQueue02vs_01, v, err.Error())
				})
			} else {
				logx.Errorf(constant.ErrFuncBQueue01v_01, v)
			}
		})
	}).CatchAll(func(err error) {
		logx.Error(err)
	})
}
