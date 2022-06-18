package middlewares

import (
	"fmt"
	"github.com/daida459031925/common/result"
	"github.com/daida459031925/common/runtimeStatus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateError(i int) func(next http.HandlerFunc) http.HandlerFunc {
	//全局异常返回 数字填写3  针对接口错误
	logx.Info("全局err打印位置为：", i)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				//内置函数，可以捕捉到函数异常
				err := recover()
				if err != nil {
					logx.Info("异常过滤器打印错误异常：")
					//这里是打印错误，还可以进行报警处理，例如微信，邮箱通知
					//panic(err)//退出程序
					//这个地方包装了一个方法基本值需要加1
					logx.Info(runtimeStatus.GetErrorStatus(i, err))
					httpx.OkJson(w, result.Error(fmt.Sprintf("%s", err)))
					return
				}
			}()

			logx.Info("异常过滤器进行中")
			next(w, r)
		}
	}
}
