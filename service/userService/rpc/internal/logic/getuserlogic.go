package logic

import (
	"context"
	"service/common/generalSql/model/sys"
	"service/userService/rpc/types/user"

	"service/userService/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line
	User, e := l.svcCtx.SysUserModel.FindOneAndDeleted(in.Id)
	if e != nil {
		return nil, e
	}

	var u, e1 = User.(sys.SysUser)

	if e1 {
		return nil, e
	}

	return &user.UserInfoReply{
		Id: u.Id,
		//Name:   User.Name,
		//Number: User.Number,
		//Gender: User.Gender,
	}, nil
}
