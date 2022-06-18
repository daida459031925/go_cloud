package main

import (
	"fmt"
	"github.com/gogf/gf/i18n/gi18n"
	"github.com/gogf/gf/os/gctx"
)

func main() {
	var (
		ctx  = gctx.New()
		i18n = gi18n.New()
	)

	i18n.SetLanguage("zh-CN")
	s := i18n.Translate(ctx, "December")
	fmt.Println(s)
}
