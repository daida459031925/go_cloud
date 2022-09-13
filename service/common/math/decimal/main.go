package decimal

import (
	"errors"
	"fmt"
	deci "github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type Decimal struct {
	deci.Decimal
}

func GetDecimal(str string) (Decimal, error) {
	s := strings.TrimSpace(str)
	_, e := strconv.ParseFloat(str, 64)
	if e != nil {
		s = ""
		e = errors.New(fmt.Sprintf("转换float64失败: %s", e))
	}

	return Decimal{deci.RequireFromString(s)}, e
}
