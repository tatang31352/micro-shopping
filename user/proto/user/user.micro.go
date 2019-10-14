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

func RegisterUserServiceHandler(s server.Server,hdlr UserServiceHandler,opts ...server.HandlerOption) error{

}
