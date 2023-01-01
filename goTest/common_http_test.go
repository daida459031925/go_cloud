package goTest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

const (
	HTTP = "http://"
	HTTPS = "https://"
)

// 工具类统一采用结构体（对象）形式，简单算法内容采用直接调用模型
type (
	//http.Client内部都存在协程
	httpx struct {
		client *http.Client
		url string
		param map[string]any
	}

	Resp struct {
		header http.Header
		result Result
		code int
	}
	// Result 其他工具类中返还的数据内容
	Result struct {
		Status    int16        `json:"status"` //状态类型
		Msg       string       `json:"msg"`    //错误时候的返回信息
		Data      any          `json:"data"`   //返还的数据
		Date      string       `json:"date"`   //记录数据返回时间
		funcSlice []func() any //记录当前需要执行的所有任务
	}
)

// Url 使用自定义的http进行初始化请求，添加数据模式采用链式编程
func Url(url string) *httpx {
	if len(url)<=0 {
		return nil
	}

	return &httpx{
		url: url,
	}
}

// Http 设置请求内容为http请求
func (h *httpx) Http() *httpx {
	h.url = join(HTTP,h.url)
	return h
}

// Https 设置请求为https请求
func (h *httpx) Https() *httpx {
	h.url = join(HTTPS,h.url)
	return h
}

// 判断传入的url中是否添加了http://或https://
func join(types,url string) string {
	b := bytes.Buffer{}
	if !(strings.Contains(url,HTTP)||strings.Contains(url,HTTPS)) {
		b.WriteString(types)
	}
	b.WriteString(url)
	return b.String()
}

// GetHeader 返回resp的header
func (r *Resp) GetHeader() http.Header {
	if r != nil {
		return r.header
	}
	return nil
}

// GetContent 返回response的[]byte格式内容
func (r *Resp) GetContent() Result {
	return r.result
}

// GetStatusCode 返回response的http code
func (r *Resp) GetStatusCode() int {
	if r != nil {
		return r.code
	}
	return 500
}

//获取返还值
func getRespBody(resp *http.Response) (response *Resp, e error) {
	if resp == nil {
		return nil, errors.New("response is nil")
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r Result
	err = json.Unmarshal(content, &r)
	if e!=nil {
		return nil, err
	}

	return &Resp{
		header:  resp.Header,
		result: r,
		code:    resp.StatusCode,
	}, nil
}

func main() {
	resp,e := http.Get(Url("www.baidu.com").Http().url)
	if e !=nil{
		return
	}
	defer resp.Body.Close()

	_,e = io.ReadAll(resp.Body)
	if e!=nil{
		return
	}

	//fmt.Println(string(body))

}


// go test -bench=方法名 -benchmem
// ns/op平均每次多少时间 1s=1000ms 1ms=1000us 1us=1000ns
// allocs/op进行多少次内存分配
// B/op标识每次操作分配多少字节
func BenchmarkName(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		a()
	}
}

func a()  {
	fmt.Print("1")
}



"https://blog.csdn.net/ydl1128/article/details/126259943"
"https://www.cnblogs.com/superhin/p/16332720.html"
"https://gitee.com/dianjiu/gokit"
"https://blog.csdn.net/asd1126163471/article/details/127020095"
"https://www.php.cn/be/go/475737.html"
"https://www.cnblogs.com/hsyw/p/16104591.html"
