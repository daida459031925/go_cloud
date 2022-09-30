package logic

import (
	"context"
	"github.com/daida459031925/common/encryption/pbkdf2"
	"github.com/daida459031925/common/result"
	"github.com/zeromicro/go-zero/core/logx"
	"service/common/constant"
	"service/common/token"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
	"strings"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.Login) (resp result.Result) {
	//去除左右的空格然后需要验证字段
	u := strings.TrimSpace(req.Username)
	p := strings.TrimSpace(req.Password)
	loginErrString := constant.ResultLoginErr01

	if len(u) == 0 || len(p) == 0 {
		return result.Error(loginErrString)
	}

	//使用布隆过滤器进行身份验证

	//找到用户密码需要进行加密处理
	user, e := l.svcCtx.SysUserModel.FindOneLoginUser(u)
	if e != nil {
		return result.Error(loginErrString)
	}

	tf := pbkdf2.CheckEncryptPwdMatch(p, user.Password)

	if !tf {
		return result.Error(loginErrString)
	}

	var payloads map[string]interface{} = make(map[string]interface{})
	payloads["prevSecret"] = user.PrevSecret
	payloads["secret"] = user.Secret
	payloads["userId"] = user.Id

	s, accessExpire, e := token.GetToken(user.Secret, payloads, user.TokenExpire)
	if e != nil {
		return result.Error(loginErrString)
	}

	var gender string

	switch user.DictId {
	case 0:
		gender = "未知"
	case 1:
		gender = "男"
	case 2:
		gender = "女"
	default:
		gender = "未知"
	}

	// ---end---
	ru := &types.RUserToken{
		Id:           user.Id,
		Name:         user.Account,
		Gender:       gender,
		AccessToken:  s,
		AccessExpire: accessExpire,
		RefreshAfter: accessExpire - 5*60, //建议刷新时间提前5分钟
	}
	return result.Ok(ru)
}
