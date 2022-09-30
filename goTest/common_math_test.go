package goTest

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"testing"
)

func TestMath(t *testing.T) {
	a, e := strconv.ParseFloat("", 64)
	if e != nil {
		logx.Info(a)
	}

	fff := -8.0497183772403904e-17

	logx.Info(decimal.NewFromFloat(fff))

	logx.Info(fmt.Sprintf("%f", fff))

	d, _ := getDecimal("-2.22222222222222222222222222222222222222")
	b, _ := getDecimal("-0.00000000000000000000000000000000000001")
	b0, _ := getDecimal("2")
	d = d.Add(b)

	logx.Info(d.Add(d).DivRound(b0, 40))
}

func getDecimal(str string) (decimal.Decimal, error) {
	s := strings.TrimSpace(str)
	_, e := strconv.ParseFloat(str, 64)
	if e != nil {
		s = "0"
		e = errors.New(fmt.Sprintf("转换float64失败: %s", e))
	}

	return decimal.RequireFromString(s), e
}
