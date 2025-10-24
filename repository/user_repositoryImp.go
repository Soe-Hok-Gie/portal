package repository

import (
	"context"
	"database/sql"
	"medsos/helper"
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
func (repository *userRepositoryImp) Save(ctx context.Context, user domain.User) domain.User {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	script := "INSERT INTO user (id,username) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, script, user.Id, user.Username)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()

	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

// update user
func (repository *userRepositoryImp) Update(ctx context.Context, user domain.User) domain.User {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "UPDATE user SET username=? WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, user.Username, user.Id); err != nil {
		panic(err)
	}
	return user

}

// find by id
func (repository *userRepositoryImp) FindById(ctx context.Context, userId int) (domain.User, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

}

// delete
func (repository *userRepositoryImp) Delete(ctx context.Context, user domain.User) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	script := "DELETE FROM user  WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, user.Id); err != nil {
		panic(err)
	}
}
