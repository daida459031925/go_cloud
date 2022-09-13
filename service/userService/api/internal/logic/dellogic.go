package logic

import (
	"context"

	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelLogic {
	return DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req types.UpdDBSysUser) (resp *types.RSysUser, err error) {
	// todo: add your logic here and delete this line

	return
}
