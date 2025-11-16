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
func (repository *postRepositoryImp) FindById(ctx context.Context, postId int) (domain.UserPost, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	//P alias post,U alias User,
	//AS p memberikan alias "p" untuk tabel post agar penulisan kode selanjutnya lebih singkat.
	//INNER JOIN  hanya baris yang cocok dikedua tabel yang akan dikembalikan.
	// ON p.user_id = u.id adalah kondisi untuk menggabungkan tabel.
	//LIMIT 1: membatasi jumlah baris hasil yang dikembalikan menjadi maksimal 1 baris.

	script := "SELECT p.id, p.user_id, p.title, p.content, u.username FROM post AS p INNER JOIN `user` AS u ON p.user_id = u.id WHERE p.id =? LIMIT 1;"
	rows, err := tx.QueryContext(ctx, script, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.UserPost{}
	if rows.Next() {
		rows.Scan(&post.Id, &post.User_Id, &post.Title, &post.Content, &post.Username)
		return post, nil
	} else {
		return post, errors.New("node found")
	}
}

func (repository *postRepositoryImp) FindAll(ctx context.Context, filter domain.PostFilter) []domain.Post {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	//langsung panggil filter sort
	script := fmt.Sprintf("SELECT id, user_id, title, content, created_at FROM post ORDER BY post.created_at %s", filter.Sort)
	fmt.Printf("Query: %q\n", script)

	rows, err := tx.QueryContext(ctx, script)
	fmt.Println("Query error:", err)
	defer rows.Close()

	var posts []domain.Post

	//looping
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(
			&post.Id,
			&post.User_Id,
			&post.Title,
			&post.Content,
			&post.Created_At,
		)
		fmt.Printf("\nrows : %v", rows)

		helper.PanicIfError(err)
		posts = append(posts, post)
		fmt.Printf("Total data dari DB: %d\n", len(posts))

	}

	return posts

}

func (repository *postRepositoryImp) Delete(ctx context.Context, id int) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "DELETE FROM user WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, id); err != nil {
		panic(err)
	}

}
