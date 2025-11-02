package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
func (repository *postRepositoryImp) FindById(ctx context.Context, postId int) (domain.Post, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "SELECT id, title, content  FROM post WHERE id=?"
	rows, err := tx.QueryContext(ctx, script, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.Content)
		return post, nil
	} else {
		return post, errors.New("node found")
	}
}

func (repository *postRepositoryImp) FindAll(ctx context.Context) []domain.Post {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "SELECT id, user_id, title, content FROM post"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	fmt.Printf("\nrows : %v", rows)

	var posts []domain.Post

	//looping
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(
			&post.Id,
			&post.User_Id,
			&post.Title,
			&post.Content,
		)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}
	return posts

}

func (repository *postRepositoryImp) Delete(ctx context.Context, post domain.Post) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "DELETE FROM user WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, post.Id); err != nil {
		panic(err)
	}

}
