package middlewares

import (
	"errors"
	"github.com/daida459031925/common/result"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"service/common/constant"
	"time"
)

type Auth struct {
	Id           uint64 `json:"id"`
	AccessToken  string `json:"accessToken"`
	AccessExpire uint64 `json:"accessExpire"`
}

/*
*
自定义实现jwt权限验证
*/
func NewAuthorize() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//放开登录请求拦截
			uri := r.URL.Path
			if "/user/login" == uri {
				next(w, r)
				return
			}
			//解析请求中是否带有token和账号
			var auth Auth
			if err := httpx.Parse(r, &auth); err != nil {
				httpx.OkJson(w, result.Error("未获取到token和账号"))
				return
			}
			//拿到账号和token则需要在缓存中获取对应账户信息
			//var UserRpc zrpc.RpcClientConf
			//u := user.NewUser(zrpc.MustNewClient(UserRpc))
			//u.GetUser()
			//l.svcCtx.SysUserModel.FindOne(1)
			//直接从数据库中获取对应的userId信息
			r.Context()
			//获取解析器
			tp := GetTokenParser(time.Second)
			_, err := ParseToken(tp, r, "secret", "prevSecret")

			if err != nil {
				httpx.Error(w, err)
			} else {

			}

			//如果权限未通过则直接返还
			//if  {
			//	httpx.OkJson(w, result.Error("无效权限"))
			//	return
			//}

			next(w, r)
		}
	}
}

/*
根据传入内容生成jwt
secretKey：应该存放私有或者共有密钥，共有密钥防止客户变更
payloads： 其他额外参数

其他参数包括：userId、secret、

seconds：token有效时间
*/
func GetToken(secretKey string, payloads map[string]interface{}, seconds uint64) (string, uint64, error) {
	//后期改为从配置中获取
	mapKey := []string{constant.SysUserId, constant.TokenSecret}
	//token错误批量信息
	errString := constant.SysSpace
	//定义返还的失效时间
	var rTokenExpire uint64 = 0

	for key := range mapKey {
		if _, ok := payloads[mapKey[key]]; !ok {
			errString += mapKey[key]
		}
	}

	if errString != constant.SysSpace {
		return constant.SysSpace, rTokenExpire, errors.New(constant.ErrGetToken00_01)
	}

	now := time.Now().Unix()
	claims := make(jwt.MapClaims)

	rTokenExpire = uint64(now) + seconds

	//失效时间 小的整数类型往大的证书类型合并
	claims[constant.TokenExpire] = rTokenExpire
	//创建时间
	claims[constant.TokenIat] = now

	for k, v := range payloads {
		claims[k] = v
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = claims

	//使用用户id作为key来
	//return jwtToken.SignedString([]byte(strconv.FormatInt(userId,10)))
	r, e := jwtToken.SignedString([]byte(secretKey))
	return r, rTokenExpire, e
}

/*
生成解析器，用来解析token是否失效
*/
func GetTokenParser(duration time.Duration) *token.TokenParser {
	return token.NewTokenParser(token.WithResetDuration(duration))
}

func GetTokenParserSecond() *token.TokenParser {
	return token.NewTokenParser(token.WithResetDuration(time.Second))
}

func ParseToken(tp *token.TokenParser, r *http.Request, secret, prevSecret string) (*jwt.Token, error) {
	return tp.ParseToken(r, secret, prevSecret)
}
