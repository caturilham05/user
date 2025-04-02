package service

import (
	"caturilham05/user/model/web"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
}
