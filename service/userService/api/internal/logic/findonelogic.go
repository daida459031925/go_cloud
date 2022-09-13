package logic

import (
	"context"

	"service/userService/api/internal/svc"
	"service/userService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Find_oneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFind_oneLogic(ctx context.Context, svcCtx *svc.ServiceContext) Find_oneLogic {
	return Find_oneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Find_oneLogic) Find_one(req types.UpdDBSysUser) (resp *types.RSysUser, err error) {
	// todo: add your logic here and delete this line

	return
}
