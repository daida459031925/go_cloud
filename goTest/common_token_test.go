package goTest

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	token2 "github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/**
结论：go-zero中使用jwt自动集成模式
1.服务器密钥采用完全随机模式即可，若需要使用每个客户单独appid的密钥管理需要自行实现
2.同一个用户获取jwt认证时候只要参数改变就会发放新的认证，服务器需要自行实现只能认证最新的，最简单认证方法数据库中加入发放jwt时间使用悲观锁来保证jwt唯一
3.jwt只是发放是否可以访问系统，系统权限需要使用casbin实现
*/
func TestTokenParser(t *testing.T) {
	const (
		key     = "14F17379-EB8F-411B-8F12-6929002DCA76"
		prevKey = "B63F477D-BBA3-4E52-96D3-C0034C27694A"
	)
	keys := []struct {
		key     string
		prevKey string
	}{
		{
			key,
			"",
		},
		{
			key,
			"",
		},
	}

	for _, pair := range keys {
		req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		token, err := buildToken("1", map[string]interface{}{}, 3600)
		assert.Nil(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		parser := token2.NewTokenParser(token2.WithResetDuration(time.Minute))
		tok, err := parser.ParseToken(req, pair.key, pair.prevKey)
		assert.Nil(t, err)
		assert.Equal(t, "1", tok.Claims.(jwt.MapClaims)["key"])
	}
}

func TestTokenParser_Expired(t *testing.T) {
	const (
		key     = "14F17379-EB8F-411B-8F12-6929002DCA76"
		prevKey = "B63F477D-BBA3-4E52-96D3-C0034C27694A"
	)
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	token, err := buildToken(key, map[string]interface{}{
		"key":  "value",
		"time": "2022-04-22 00:00:00",
	}, 50000)
	assert.Nil(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	req1 := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	token1, err1 := buildToken(key, map[string]interface{}{
		"key":  "value",
		"time": "2022-04-22 00:00:00",
	}, 50000)
	assert.Nil(t, err1)
	req1.Header.Set("Authorization", "Bearer "+token1)
	tt := time.Now().String()
	fmt.Println(tt)
	parser := token2.NewTokenParser(token2.WithResetDuration(time.Second))
	tok, err := parser.ParseToken(req, key, prevKey)
	tok1, err1 := parser.ParseToken(req1, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	tok, err = parser.ParseToken(req, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	//parser.resetTime = timex.Now() - time.Hour
	token, err = buildToken(key, map[string]interface{}{
		"key": "value",
	}, 5)
	parser = token2.NewTokenParser(token2.WithResetDuration(time.Second))
	req.Header.Set("Authorization", "Bearer "+token)
	tok, err = parser.ParseToken(req, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	assert.Equal(t, "value", tok1.Claims.(jwt.MapClaims)["key"])
}

func buildToken(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
