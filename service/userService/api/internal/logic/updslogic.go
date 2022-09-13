package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"service/userService/api/internal/svc"
)

type UpdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdsLogic {
	return UpdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdsLogic) Upds() error {
	// todo: add your logic here and delete this line

	return nil
}
