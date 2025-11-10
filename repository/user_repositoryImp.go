package repository

import (
	"context"
	"database/sql"
	"errors"
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

	script := "SELECT id, username FROM user WHERE id=?"
	rows, err := tx.QueryContext(ctx, script, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		//ada datanya
		rows.Scan(&user.Id, &user.Username)
		return user, nil
	} else {
		return user, errors.New("node found")
	}

}

// delete
func (repository *userRepositoryImp) Delete(ctx context.Context, id int) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	script := "DELETE FROM user  WHERE id=?"
	if _, err := tx.ExecContext(ctx, script, id); err != nil {
		panic(err)
	}
}

// findAll
func (repository *userRepositoryImp) FindAll(ctx context.Context) []domain.User {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "SELECT id, username FROM user"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)

	defer rows.Close()

	var users []domain.User

	//looping
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
		)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

func (repository *userRepositoryImp) FindUserPost(ctx context.Context, userId int) domain.UserPosts {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	script := "SELECT u.id AS user_id, u.username, p.id AS post_id, p.title AS post_title, p.content AS post_content FROM user u LEFT JOIN post p ON u.id = p.user_id WHERE u.id = ?;"
	rows, err := tx.QueryContext(ctx, script, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	//buat sebuah var user untuk menampung data dari domain.UserPosts
	var user domain.UserPosts
	//lalu panggil user.Posts untuk menampung slice PostWithoutUserId
	user.Posts = []domain.PostWithoutUserId{}

	for rows.Next() {
		var (
			postId      sql.NullInt64
			postTitle   sql.NullString
			postContent sql.NullString
		)

		if err := rows.Scan(&user.Id, &user.Username, &postId, &postTitle, &postContent); err != nil {
			panic(err)
		}
		if postId.Valid {
			user.Posts = append(user.Posts, domain.PostWithoutUserId{
				Id:      int(postId.Int64),
				Title:   postTitle.String,
				Content: postContent.String,
			})
		}
	}

	return user
}
