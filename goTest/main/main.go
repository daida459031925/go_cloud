package main

/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h>

void pri(){
	printf("hey");
}

int add(int a,int b){
	return a+b;
}
*/
import "C" // 切勿换行再写这个

import (
	"fmt"
	c "github.com/daida459031925/go_cloud/goTest/captcha"
	"github.com/dchest/captcha"
	"log"
	"net/http"
)

func main() {
	fmt.Println(C.add(1, 2))

	// 简单设置log参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	http.HandleFunc("/", c.ShowCaptcha)
	http.HandleFunc("/processCapcha", c.ResultPage)

	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	log.Println("starting server : 8888")

	if err := http.ListenAndServe("localhost:8888", nil); err != nil {
		log.Fatal(err)
	}
}
