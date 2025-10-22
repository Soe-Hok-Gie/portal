package repository

import (
	"context"
	"medsos/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) domain.User
	Update(ctx context.Context, user domain.User) domain.User
}
