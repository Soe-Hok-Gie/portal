package service

import (
	"context"
	"medsos/model/web"
)

type UserService interface {
	//parameter keduanya representasi dari request dan response
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	FindById(ctx context.Context, request userId int) web.Response
	Delete(ctx context.Context, request userId int)
}
