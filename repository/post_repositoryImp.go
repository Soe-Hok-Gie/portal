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

func NewPostRepository(DB *sql.DB) PostRepository {
	return &postRepositoryImp{DB: DB}
}

func (repository *postRepositoryImp) Save(ctx context.Context, post domain.Post) domain.Post {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "INSERT INTO post (id, user_id, title, content) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, post.Id, post.User_Id, post.Title, post.Content)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post

}

func (repository *postRepositoryImp) Update(ctx context.Context, post domain.Post) domain.Post {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "UPDATE post SET user_id=?, title=?,content=? WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, post.User_Id, post.Title, post.Content, post.Id); err != nil {
		panic(err)
	}
	return post
}
func (repository *postRepositoryImp) FindById(ctx context.Context, postId int) domain.Post {

	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "SELECT id, title FROM post WHERE id=?"
	rows, err := tx.QueryContext(ctx, script, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {

	}

}
