package service

import (
	"context"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
)

//parameternya userrepo

type UserServiceImp struct {
	UserRepository repository.UserRepository
}

// implementasi bisnis logic
func (service *UserServiceImp) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	// tampung model domain dalam sebuah variabel
	user := domain.User{
		Username: request.Username,
	}

	//panggil service
	user = service.UserRepository.Save(ctx, user)

	// tampung model web response dalam sebuah variabel
	userResponse := web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
	return userResponse
}
