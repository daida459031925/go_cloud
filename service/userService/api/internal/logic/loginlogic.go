package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/html/atom"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
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

func (l *LoginLogic) Login(req types.Login) (*types.UserToken, error) /*(resp result.Result)*/ {
	// todo: add your logic here and delete this line
	//if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
	//	return result.Error("参数错误")
	//}

	atom.Time.String()
	a, b := 1, 0
	fmt.Println(a / b)

	userInfo, err := l.svcCtx.UserModel.FindOne(1)
	fmt.Println(userInfo)
	fmt.Println(err)
	//switch err {
	//case nil:
	//case model.ErrNotFound:
	//	return result.Error("用户名不存在")
	//default:
	//	return result.Error("未知错误")
	//}

	//if "userInfo.Password" != req.Password {
	//	return result.Error("用户密码不正确")
	//}

	//User,err := l.svcCtx.UserRpc.GetUser(l.ctx,nil)

	// ---start---
	//now := time.Now().Unix()
	//Secret := userInfo.Secret
	//PrevSecret := userInfo.PrevSecret
	//TokenExpire := userInfo.TokenExpire
	//jwtToken, err := (l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	//if err != nil {
	//	return nil, err
	//}
	// ---end---

	return &types.UserToken{
		//Id:           userInfo.Id,
		//Name:         userInfo.Name,
		//Gender:       userInfo.Gender,
		//AccessToken:  jwtToken,
		//AccessExpire: now + accessExpire,
		//RefreshAfter: now + accessExpire/2,
	}, nil
	//return result.Ok(nil)
}
