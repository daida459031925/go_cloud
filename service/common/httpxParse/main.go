package httpxParse

import (
	"github.com/daida459031925/common/result"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"service/common/constant"
	"service/common/validatorUtil"
)

func Parse(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := httpx.Parse(r, v); err != nil {
		logx.Error(err)
		httpx.OkJson(w, result.Error(constant.ErrhttpxParse00_01))
		return true
	}

	//自定义模板？添加内容验证validate
	validate := validatorUtil.NewValidate()
	errList, e := validate.ValStruct(v)
	if e != nil {
		httpx.OkJson(w, result.ErrorData(e.Error(), errList))
		return true
	}

	return false
}
