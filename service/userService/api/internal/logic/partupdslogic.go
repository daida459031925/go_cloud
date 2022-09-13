package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type Part_updsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPart_updsLogic(ctx context.Context, svcCtx *svc.ServiceContext) Part_updsLogic {
	return Part_updsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Part_updsLogic) Part_upds() error {
	// todo: add your logic here and delete this line

	return nil
}
