package service

import (
	"context"
	"medsos/repository"
)

//parameternya userrepo

type UserServiceImp struct {
	UserRepository repository.UserRepository
}

func (service *UserServiceImp) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {

}
