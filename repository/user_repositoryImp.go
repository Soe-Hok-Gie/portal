package repository

import (
	"context"
	"database/sql"
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
	tx, err := repository.DB.Begin()

	if err != nil {
		panic(err)
	}
	script := "INSERT INTO user (id,username) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, script, user.Id, user.Username)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
	user.Id = int(id)
	return user
}
