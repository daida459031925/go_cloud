package logic

import (
	"context"

	"service/publicService/api/internal/svc"
	"service/publicService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranslateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTranslateLogic(ctx context.Context, svcCtx *svc.ServiceContext) TranslateLogic {
	return TranslateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TranslateLogic) Translate() (resp *types.Res, err error) {
	// todo: add your logic here and delete this line

	return
}
