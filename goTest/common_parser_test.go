package goTest

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/zeromicro/go-zero/core/logx"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

var comments = []struct {
	list []string
	text string
}{
	{[]string{"//1"}, "1"},
	{[]string{"//"}, ""},
	{[]string{"//   "}, ""},
	{[]string{"//", "//", "//   "}, ""},
	{[]string{"// foo   "}, "foo\n"},
	{[]string{"//", "//", "// foo"}, "foo\n"},
	{[]string{"// foo  bar  "}, "foo  bar\n"},
	{[]string{"// foo", "// bar"}, "foo\nbar\n"},
	{[]string{"// foo", "//", "//", "//", "// bar"}, "foo\n\nbar\n"},
	{[]string{"// foo", "/* bar */"}, "foo\n bar\n"},
	{[]string{"//", "//", "//", "// foo", "//", "//", "//"}, "foo\n"},

	{[]string{"/**/"}, ""},
	{[]string{"/*   */"}, ""},
	{[]string{"/**/", "/**/", "/*   */"}, ""},
	{[]string{"/* Foo   */"}, " Foo\n"},
	{[]string{"/* Foo  Bar  */"}, " Foo  Bar\n"},
	{[]string{"/* Foo*/", "/* Bar*/"}, " Foo\n Bar\n"},
	{[]string{"/* Foo*/", "/**/", "/**/", "/**/", "// Bar"}, " Foo\n\nBar\n"},
	{[]string{"/* Foo*/", "/*\n*/", "//", "/*\n*/", "// Bar"}, " Foo\n\nBar\n"},
	{[]string{"/* Foo*/", "// Bar"}, " Foo\nBar\n"},
	{[]string{"/* Foo\n Bar*/"}, " Foo\n Bar\n"},

	{[]string{"// foo", "//go:noinline", "// bar", "//:baz"}, "foo\nbar\n:baz\n"},
	{[]string{"// foo", "//lint123:ignore", "// bar"}, "foo\nbar\n"},
}

func TestCommentText(t *testing.T) {
	for i, c := range comments {
		list := make([]*ast.Comment, len(c.list))
		for i, s := range c.list {
			list[i] = &ast.Comment{Text: s}
		}

		text := (&ast.CommentGroup{list}).Text()
		if text != c.text {
			logx.Infof("case %d: got %q; expected %q", i, text, c.text)
		}
	}
}

var isDirectiveTests = []struct {
	in string
	ok bool
}{
	{"abc", false},
	{"go:inline", true},
	{"Go:inline", false},
	{"go:Inline", false},
	{":inline", false},
	{"lint:ignore", true},
	{"lint:1234", true},
	{"1234:lint", true},
	{"go: inline", false},
	{"go:", false},
	{"go:*", false},
	{"go:x*", true},
	{"export foo", true},
	{"extern foo", true},
	{"expert foo", false},
}

func TestIsDirective(t *testing.T) {
	//for _, tt := range isDirectiveTests {
	//	if ok := ast.isDirective(tt.in); ok != tt.ok {
	//		t.Errorf("isDirective(%q) = %v, want %v", tt.in, ok, tt.ok)
	//	}
	//}

	fset := token.NewFileSet() // 相对于fset的position
	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// 解析src但在处理导入后停止。
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从文件CAST打印导入。
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}

func abcdefg() {
	logx.Info("我执行了")
}

/*
规则解析器
*/
func TestGZJXQ(t *testing.T) {
	functions := map[string]govaluate.ExpressionFunction{}
	functions["stolen"] = func(args ...any) (any, error) {
		abcdefg()
		return true, nil
	}

	//expression, _ := govaluate.NewEvaluableExpression("!(foo in (0,1,'0','1'))")
	expression, _ := govaluate.NewEvaluableExpressionWithFunctions("(foo in (0,1,'0','1',0.0,'0.0'))", functions)

	parameters := make(map[string]any, 8)
	parameters["foo"] = 1

	result, _ := expression.Evaluate(parameters)
	// result is now set to "false", the bool value.
	logx.Info(result)

	//expr := `a > 1 && b > 2`
	//exprAst, err := parser.ParseExpr(expr)
	//if err != nil {
	//	logx.Info(err)
	//	return
	//}
	//fest := token.NewFileSet()
	//ast.Print(fest, exprAst)
	//
	//logx.Info(exprAst.Pos())
	//logx.Info(exprAst.End())
	//expr := `a > 1 && b > 2`
	//m := map[string]int64{"a": 2, "b": 3}
	//logx.Info(Eval(m, expr))
}

// Eval : 计算 expr 的值
func Eval(m map[string]int64, expr string) (bool, error) {
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, err
	}
	// 打印 ast
	fset := token.NewFileSet()
	ast.Print(fset, exprAst)
	return judge(exprAst, m), nil
}

// dfs
func judge(bop ast.Node, m map[string]int64) bool {
	// 叶子结点
	if isLeaf(bop) {
		// 断言成二元表达式
		expr := bop.(*ast.BinaryExpr)
		// 左边
		x := expr.X.(*ast.Ident)
		logx.Info(x)
		// 右边
		y := expr.Y.(*ast.BasicLit)
		logx.Info(y)
		//如果是 ">" 符号
		if expr.Op == token.GTR {
			left := m[x.Name]
			right, _ := strconv.ParseInt(y.Value, 10, 64)
			return left > right
		}
		return false
	}
	// 不是叶子节点那么一定是 binary expression（我们目前只处理二元表达式）
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		println("this cannot be true")
		return false
	}
	// 递归地计算左节点和右节点的值
	switch expr.Op {
	case token.LAND:
		return judge(expr.X, m) && judge(expr.Y, m)
	case token.LOR:
		return judge(expr.X, m) || judge(expr.Y, m)
	}
	println("unsupported operator")
	return false
}

// 判断是否是叶子节点
func isLeaf(bop ast.Node) bool {
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	// 二元表达式的最小单位，左节点是标识符，右节点是值
	_, okL := expr.X.(*ast.Ident)
	_, okR := expr.Y.(*ast.BasicLit)
	if okL && okR {
		return true
	}
	return false
}
