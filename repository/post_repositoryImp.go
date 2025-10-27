package repository

import (
	"context"
	"database/sql"
	"medsos/helper"
	"medsos/model/domain"
)

type postRepositoryImp struct {
	DB *sql.DB
}

func (repository *postRepositoryImp) Save(ctx context.Context, post domain.Post) domain.Post {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

}
