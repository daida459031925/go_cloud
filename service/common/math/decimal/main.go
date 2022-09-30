package decimal

import (
	"errors"
	"fmt"
	deci "github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"service/common/algorithm/queue/bqueue"
	"service/common/constant"
	"strconv"
	"strings"
)

/*
小数计算方式
*/
type decimal struct {
	value string
	deci.Decimal
	*bqueue.BQueue
}

func NewDecimal(number string) (*decimal, error) {
	s := strings.TrimSpace(number)
	_, e := strconv.ParseFloat(s, 64)
	if e != nil {
		s = ""
		e = errors.New(fmt.Sprintf(constant.ErrDecimal01s_01, e))
		logx.Error(e)
	}
	return &decimal{
		s,
		deci.RequireFromString(s),
		bqueue.NewFuncList(),
	}, e
}
