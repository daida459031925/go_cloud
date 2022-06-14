package goTest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"reflect"
	"testing"
)

func TestXss(t *testing.T) {
	str1 := "<script>alert(2)\"'</script>"

	str2 := html.EscapeString(str1)
	fmt.Println(str2)
	str3 := html.UnescapeString(str2)
	fmt.Println(str3)
}

var jsonString = "{\n  \"users\": [\n    {\n      \"name\": \"Aric<>'\",\n      \"age\": 27,\n      \"social\": {\n        \"zhihu\": \"https://zhihu.com\",\n        \"weibo\": \"https://weibo.com\",\n      \"xss\":\"<script>alert(2)\\\"'</script>\"}\n    },\n    {\n      \"name\": \"gloria\",\n      \"age\": 22,\n      \"social\": {\n        \"zhihu\": \"https://zhihu.com\",\n        \"weibo\": \"https://weibo.com\",\n      \"xss\":\"<script>alert(2)\\\"'</script>\"}\n    }\n  ]\n}"

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Zhihu string `json:"zhihu"`
	Weibo string `json:"weibo"`
	Xss   string `json:"xss"`
}

func TestJsonPast(t *testing.T) {
	var users Users
	b := bytes.NewBufferString(jsonString)
	json.Unmarshal(b.Bytes(), &users)
	var a = "<>"
	fmt.Println(reflect.ValueOf(&a).Kind())
	x := XssUtilInit(XssEscape, Json, &users)
	x.xssCode()
	fmt.Println(users)
	x1 := XssUtilInit(XssUnEscape, Json, &users)
	x1.xssCode()
	fmt.Println(users)
}

//1. Reflection goes from interface value to reflection Object.
//2. Reflection goes from refelction object to interface value.
//3. To modify a reflection object, the value must be settable.

const (
	XssEscape   XssType = 0
	XssUnEscape XssType = 1

	Json TagType = "json"
	Form TagType = "form"
	Db   TagType = "db"
)

type (
	xssUtil struct {
		Data    any
		XssType XssType
		TagType TagType
		Filter  []string
	}

	XssType int

	TagType string
)

func XssUtilInit(xssType XssType, tagType TagType, data any, filter ...string) xssUtil {
	return xssUtil{Data: data, XssType: xssType, TagType: tagType, Filter: filter}
}

func (x xssUtil) xssCode() {
	//获取对象所有内容
	//t := reflect.TypeOf(data)
	v := reflect.ValueOf(x.Data)
	//目的是将指针类型的转换成实在的数据
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fieldStruct(v, x.XssType, x.TagType, x.Filter)
	fieldSlice(v, x.XssType, x.TagType, x.Filter)
	switch x.XssType {
	case XssEscape:
		escapeString(v, "", nil)
	case XssUnEscape:
		unEscapeString(v, "", nil)
	}
}

func fieldStruct(value reflect.Value, xssType XssType, tagType TagType, filter []string) {
	if value.Kind() == reflect.Struct {
		for j := 0; j < value.NumField(); j++ {
			v := value.Field(j)
			//获取当前对象名称key
			tagName := fmt.Sprintf("%s", tagType)
			name := value.Type().Field(j).Tag.Get(tagName)
			switch xssType {
			case XssEscape:
				escapeString(v, name, filter)
			case XssUnEscape:
				unEscapeString(v, name, filter)
			}
			fieldSlice(v, xssType, tagType, filter)
			fieldStruct(v, xssType, tagType, filter)
		}
	}
}

func fieldSlice(value reflect.Value, xssType XssType, tagType TagType, filter []string) {
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i)
			fieldStruct(v, xssType, tagType, filter)
		}
	}
}

//编码 " => &#34;
func escapeString(value reflect.Value, name string, filter []string) {
	if value.Kind() == reflect.String {
		if filter != nil && len(filter) > 0 {
			for i := range filter {
				if filter[i] == name {
					return
				}
			}
		}
		value.SetString(html.EscapeString(value.String()))
	}
}

//解码 &#34; => "
func unEscapeString(value reflect.Value, name string, filter []string) {
	if value.Kind() == reflect.String {
		if filter != nil && len(filter) > 0 {
			for i := range filter {
				if filter[i] == name {
					return
				}
			}
		}
		value.SetString(html.UnescapeString(value.String()))
	}
}
