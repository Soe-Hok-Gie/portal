package repository

import (
	"context"
	"database/sql"
	"fmt"
	"medsos/model/domain"
)

type userRepositoryImp struct {
	DB *sql.DB
}

// semacam polimerisme (agar bisa diinjek pada main function)
func NewUserRepository(DB *sql.DB) UserRepository {
	return &userRepositoryImp{DB: DB}
}

// create user
func (repository *userRepositoryImp) save(ctx context.Context, user domain.User) domain.User {
	fmt.Println()
}
