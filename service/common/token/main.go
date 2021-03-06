package token

import (
	"errors"
	"github.com/daida459031925/common/result"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"service/common/constant"
	"strconv"
	"time"
)

type Auth struct {
	AccessToken string `json:"accessToken" validate:"required"`
	UserId     	string `json:"userId" validate:"required"`
}

/**
自定义实现jwt权限验证
*/
func HandlerAuthorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//放开登录请求拦截
		uri := r.RequestURI
		if("登录url" == uri){
			next(w,r)
			return
		}
		//解析请求中是否带有token和账号
		var auth Auth
		if err := httpx.Parse(r, &auth); err != nil {
			httpx.OkJson(w, result.Error("未获取到token和账号"))
			return
		}
		//拿到账号和token则需要在缓存中获取对应账户信息
		svc.ServiceContext
		//直接从数据库中获取对应的userId信息
		r.Context()
		//获取解析器
		tp := GetTokenParser(time.Second)
		jwtToken, err := ParseToken(tp, r, "secret", "prevSecret")

		if err != nil {
			httpx.Error(w, err)
		} else {

		}

		//如果权限未通过则直接返还
		if(){
			httpx.OkJson(w, result.Error("无效权限"))
			return
		}

		next(w,r)
	}
}

/**
根据传入内容生成jwt
secretKey：应该存放客户id，每次客户把id带入后台先根据id查询，查询完毕后验证私有和共有密钥，防止客户变更
payloads： 其他额外参数
	其他参数包括：userId、secret、
seconds：token有效时间
*/
func GetToken(userId int64, payloads map[string]interface{}, seconds int64) (string, error) {
	//后期改为从配置中获取
	mapkey :=[] string {constant.SYS_USER_ID,constant.TOKEN_SECRET}

	errString := constant.SYS_SPACE

	for key := range mapkey{
		if _,ok := payloads[mapkey[key]];!ok{
			errString += mapkey[key]
		}
	}

	if errString != constant.SYS_SPACE {
		return constant.SYS_SPACE,errors.New(constant.ERR_GET_TOKEN)
	}

	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	//失效时间
	claims[constant.TOKEN_EXPIRE] = now + seconds
	//创建时间
	claims[constant.TOKEN_IAT] = now

	for k, v := range payloads {
		claims[k] = v
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = claims

	//使用用户id作为key来
	return jwtToken.SignedString([]byte(strconv.FormatInt(userId,10)))
}

/**
生成解析器，用来解析token是否失效
*/
func GetTokenParser(duration time.Duration) *token.TokenParser {
	return token.NewTokenParser(token.WithResetDuration(duration))
}

func ParseToken(tp *token.TokenParser, r *http.Request, secret, prevSecret string) (*jwt.Token, error) {
	return tp.ParseToken(r, secret, prevSecret)
}
