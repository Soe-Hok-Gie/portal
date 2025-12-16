package repository

import (
	"context"
	"medsos/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, post domain.Post) (domain.Post, error)
	Update(ctx context.Context, post domain.Post) (domain.Post, error)
	FindById(ctx context.Context, postId int) (domain.UserPost, error)
	FindAll(ctx context.Context, filter domain.PostFilter) []domain.Post //tambahkan filter di parameter 2
	Delete(ctx context.Context, id int)
}
