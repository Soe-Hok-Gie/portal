package service

import "context"

type UserService interface {
	//parameter keduanya representasi dari request dan response
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
}
