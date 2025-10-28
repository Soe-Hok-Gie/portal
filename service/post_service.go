package service

import (
	"context"
	"medsos/model/web"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
}
