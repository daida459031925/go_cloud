package middlewares

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"service/common/constant"
)

func NewXSS(i int) func(next http.HandlerFunc) http.HandlerFunc {
	//全局异常返回 数字填写3  针对接口错误
	logx.Info(constant.ErrAllXssInit00_01, i)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m := map[any]interface{}{}
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				json.Unmarshal(body, &m)
				logx.Info(m)
			} else {
				logx.Error("无法获取Body内容，或者出错")
			}
			next(w, r)
		}
	}
}
