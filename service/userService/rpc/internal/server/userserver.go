// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"service/userService/rpc/internal/logic"
	"service/userService/rpc/internal/svc"
	"service/userService/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}