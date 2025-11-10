package service

import (
	"context"
	"medsos/helper"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
)

// parameternya userrepo
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

func (service *userServiceImp) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	// tampung model domain dalam sebuah variabel
	user := domain.User{
		Id:       request.Id,
		Username: request.Username,
	}

	//panggil service
	user = service.UserRepository.Update(ctx, user)

	// tampung model web response dalam sebuah variabel
	userResponse := web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
	return userResponse

}

func (service *userServiceImp) FindById(ctx context.Context, userId int) web.UserResponse {
	user, err := service.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)

	userResponse := web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
	return userResponse
}

func (service *userServiceImp) Delete(ctx context.Context, userId int) {

	service.UserRepository.Delete(ctx, userId)

}

func (service *userServiceImp) FindAll(ctx context.Context) []web.UserResponse {

	users := service.UserRepository.FindAll(ctx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse(user))

	}
	return userResponses
}

func (service *userServiceImp) FindUserPost(ctx context.Context, userId int) web.UserWithPostsResponse {
	user := service.UserRepository.FindUserPost(ctx, userId)

	var response web.UserWithPostsResponse
	response.Id = user.Id
	response.Username = user.Username
	response.Posts = user.Posts

	return response
}
