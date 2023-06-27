package service

import (
	"context"
	"github.com/zutim/ztm-demo/common/app"
	"github.com/zutim/ztm-demo/common/pb/user"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (u *UserService) Create(ctx context.Context, req *user.UserCreateReq) (*user.UserCreateRes, error) {
	app.UserLog().Info("接收到了", req.Name)
	return &user.UserCreateRes{Id: 123}, nil
}

func (u *UserService) Get(ctx context.Context, req *user.UserGetReq) (*user.UserGetRes, error) {
	return &user.UserGetRes{}, nil
}

func NewUserService() user.UserServer {
	return &UserService{}
}
