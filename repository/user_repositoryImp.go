package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
func (repository *userRepositoryImp) Save(ctx context.Context, user domain.User) (domain.User, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	script := "INSERT INTO user (id,username) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, script, user.Id, user.Username)
	if err != nil {
		return user, fmt.Errorf("repo :%w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("get id: %w", err)
	}

	user.Id = int(id)
	return user, nil
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
	var user domain.User
	tx, err := repository.DB.Begin()
	if err != nil {
		return user, err
	}
	defer helper.CommitOrRollBack(tx)

	script := "SELECT id, username FROM user WHERE id=?"
	rows, err := tx.QueryContext(ctx, script, userId)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	// user := domain.User{}
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

	//Metode rows.Next() memajukan kursor sql.Rows ke baris berikutnya yang tersedia dalam kumpulan hasil.
	for rows.Next() {
		var (
			postId      sql.NullInt64
			postTitle   sql.NullString //sql.NullString adalah membedakan antara string kosong ("") dan nilai yang benar-benar tidak ada (NULL) dari database
			postContent sql.NullString
		)

		//Metode rows.Scan() membaca nilai kolom baris saat ini (yang ditunjuk oleh kursor) dan menetapkannya ke variabel tujuan yang disediakan.
		if err := rows.Scan(&user.Id, &user.Username, &postId, &postTitle, &postContent); err != nil {
			panic(err)
		}
		if postId.Valid {
			//append() digunakan untuk menambahkan elemen ke akhir sebuah slice dan mengembalikan slice baru yang berisi elemen2 tambahan tsb..
			user.Posts = append(user.Posts, domain.PostWithoutUserId{
				Id:      int(postId.Int64),
				Title:   postTitle.String,
				Content: postContent.String,
			})
		}
	}

	return user
}
