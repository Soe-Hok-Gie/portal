package repository

import (
	"context"
	"medsos/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, post domain.Post) domain.Post
}
