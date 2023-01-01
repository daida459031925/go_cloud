package ruleParser

import (
	"github.com/Knetic/govaluate"
)

// 根据需求加入指定的方法
var functions map[string]govaluate.ExpressionFunction

/*
规则执行器
由于内置的需要自己实现很多东西，所以使用开源规则执行器
*/
type execRule struct {
	//解析字符串后得到的解析器
	eval *govaluate.EvaluableExpression
	//规则字符串
	exp string
}

func NewExecRule(exp string) *execRule {
	r, e := govaluate.NewEvaluableExpression(exp)
	exec := &execRule{}
	if e != nil {
		exec.eval = r
		exec.exp = exp
	}
	return exec
}

func (exec execRule) addFunc(key string, function func(args ...any) (any, error)) {
	functions[key] = function
}

func (exec execRule) delFunc(key string) {
	delete(functions, key)
}

func (exec execRule) loadFunc() {
	if len(functions) > 0 {
		r, e := govaluate.NewEvaluableExpressionWithFunctions(exec.exp, functions)
		if e != nil {
			exec.eval = r
		}
	}
}

func (exec execRule) exec(parameters map[string]any) (bool, error) {
	r, e := exec.eval.Evaluate(parameters)
	if e != nil {
		return false, e
	}

	return r.(bool), e
}
