package goTest

import (
	"fmt"
	"github.com/daida459031925/common/translate"
	"testing"
)

//翻译测试
func TestTranslate(t *testing.T) {
	str, err := translate.TranslateConversion("www.topgoer.com是个不错的go语言中文文档")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
