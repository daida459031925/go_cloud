package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type Part_delsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPart_delsLogic(ctx context.Context, svcCtx *svc.ServiceContext) Part_delsLogic {
	return Part_delsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Part_delsLogic) Part_dels() error {
	// todo: add your logic here and delete this line

	return nil
}
