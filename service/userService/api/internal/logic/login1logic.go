package logic

import (
	"context"

	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Login1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogin1Logic(ctx context.Context, svcCtx *svc.ServiceContext) Login1Logic {
	return Login1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Login1Logic) Login1(req types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}
