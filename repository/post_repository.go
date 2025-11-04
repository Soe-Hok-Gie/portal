package repository

import (
	"context"
	"medsos/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, post domain.Post) domain.Post
	Update(ctx context.Context, post domain.Post) domain.Post
	FindById(ctx context.Context, postId int) (domain.UserPost, error)
	FindAll(ctx context.Context) []domain.Post
	Delete(ctx context.Context, id int)
}
