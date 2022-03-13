package goTest

import (
	"fmt"

	"testing"
)

func TestTranslate(t *testing.T) {
	str, err := translate.TranslateConversion("www.topgoer.com是个不错的go语言中文文档")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
