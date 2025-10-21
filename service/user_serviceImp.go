package service

import (
	"context"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
)

//parameternya userrepo

type userServiceImp struct {
	UserRepository repository.UserRepository
}

// polimerisme
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImp{UserRepository: userRepository}
}

// implementasi bisnis logic
func (service *userServiceImp) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
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
