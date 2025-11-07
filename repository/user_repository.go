package repository

import (
	"context"
	"medsos/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) domain.User
	Update(ctx context.Context, user domain.User) domain.User
	FindById(ctx context.Context, userId int) (domain.User, error)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []domain.User
	FindUserPost(ctx context.Context, userId int) domain.UserPosts
}
