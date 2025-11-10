package service

import (
	"context"
	"medsos/model/web"
)

type UserService interface {
	//parameter keduanya representasi dari request dan response
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	FindById(ctx context.Context, userId int) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindAll(ctx context.Context) []web.UserResponse
	FindUserPost(ctx context.Context, userId int) web.UserWithPostsResponse
}
