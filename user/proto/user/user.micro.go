package  go_micro_srv_user

import (
	//"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	 "context"
)


type UserServiceHandler interface {
	Register(context.Context, *RegisterRequest, *Response) error
	Login(context.Context, *LoginRequest, *Response) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *Response) error
}

type userServiceHandler struct {
	UserServiceHandler
}

func RegisterUserServiceHandler(s server.Server,hdlr UserServiceHandler,opts ...server.HandlerOption) error{
	type userService interface {
		Register(ctx context.Context,in *RegisterRequest,out *Response)error
		Login(ctx context.Context,in *LoginRequest,out *Response) error
		UpdatePassword(ctx context.Context,in *UpdatePasswordRequest,out *Response) error
	}

	type UserService struct {
		userService
	}

	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h},opts...))
}
