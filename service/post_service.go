package service

import (
	"context"
	"medsos/model/domain"
	"medsos/model/web"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	FindById(ctx context.Context, postId int) (web.PostResponse, error)
	FindAll(ctx context.Context, filter domain.PostFilter) []web.PostResponse
	Delete(ctx context.Context, postId int)
}
