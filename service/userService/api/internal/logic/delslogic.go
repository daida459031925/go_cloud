package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type DelsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelsLogic {
	return DelsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelsLogic) Dels() error {
	// todo: add your logic here and delete this line

	return nil
}
