package logic

import (
	"context"
	"fmt"
	"github.com/daida459031925/common/result"
	"github.com/zeromicro/go-zero/core/logx"
	"gocv.io/x/gocv"
	"golang.org/x/net/html/atom"
	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"
	"service/userService/rpc/user"
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
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
	// todo: add your logic here and delete this line
	//if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
	//	return result.Error("参数错误")
	//}
	var i uint64 = 1
	var u user.IdReq
	u.Id = i
	fmt.Println(u)
	u1, e := l.svcCtx.UserRpc.GetUser(l.ctx, &u)
	if e != nil {
		return nil, e
	}
	fmt.Println(u1)

	atom.Time.String()
	a, b := 1, 0
	fmt.Println(a / b)

	userInfo, err := l.svcCtx.SysUserModel.FindOne(1)
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
	ru := &types.RUserToken{
		//Id:           userInfo.Id,
		//Name:         userInfo.Name,
		//Gender:       userInfo.Gender,
		//AccessToken:  jwtToken,
		//AccessExpire: now + accessExpire,
		//RefreshAfter: now + accessExpire/2,
	}
	return result.Ok(ru)
}
