package goTest

import (
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestId(t *testing.T) {
	var i uint16 = 1
	var options *idgen.IdGeneratorOptions
	if i > 0 && i < 64 {
		options = idgen.NewIdGeneratorOptions(i)
	}
	idgen.SetIdGenerator(options)

	logx.Info(idgen.NextId())
}
