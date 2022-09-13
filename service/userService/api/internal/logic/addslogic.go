package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type AddsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddsLogic {
	return AddsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddsLogic) Adds() error {
	// todo: add your logic here and delete this line

	return nil
}
