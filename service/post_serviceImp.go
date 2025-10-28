package service

import (
	"context"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
)

type postServiceImp struct {
	PostRepository repository.PostRepository
}

func (service *postServiceImp) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	// tampung model domain dalam sebuah variabel

	post := domain.Post{
		User_Id:  request.User_Id,
		Title:    request.Title,
		Content:  request.Content,
		CreateAt: request.CreateAt,
	}

	//panggil service
	post = service.PostRepository.Save(ctx, post)

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return postResponse

}

func (service *postServiceImp) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	// tampung model domain dalam sebuah variabel
	post := domain.Post{
		Id:      request.Id,
		User_Id: request.User_Id,
		Title:   request.Title,
		Content: request.Content,
	}

}
