package repository

import (
	"context"
	"medsos/model/domain"
)

type UserRepository interface {
	save(ctx context.Context, user domain.User) domain.User
}
