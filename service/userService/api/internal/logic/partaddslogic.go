package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type Part_addsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPart_addsLogic(ctx context.Context, svcCtx *svc.ServiceContext) Part_addsLogic {
	return Part_addsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Part_addsLogic) Part_adds() error {
	// todo: add your logic here and delete this line

	return nil
}
