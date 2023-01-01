package goTest

import (
	"github.com/yangtizi/cz88"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

// 都是本地调用可能失效
func TestIp(t *testing.T) {
	cip := cz88.GetAddress("14.215.177.39")
	logx.Info(cip)
}
