package service

import (
	"context"
	"medsos/model/domain"
	"medsos/model/web"
)

type postServiceImp struct {
}

func (service postServiceImp) Create(ctx context.Context, request web.PostCreateRequest) web.PostCreateRequest {

	post := domain.Post{
		User_Id:  request.User_Id,
		Title:    request.Title,
		Content:  request.Content,
		CreateAt: request.CreateAt,
	}
	post := service.PostRepository.Create(post)
}
