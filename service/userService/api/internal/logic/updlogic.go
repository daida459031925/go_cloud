package logic

import (
	"context"

	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdLogic {
	return UpdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdLogic) Upd(req types.UpdDBSysUser) (resp *types.RSysUser, err error) {
	// todo: add your logic here and delete this line

	return
}
