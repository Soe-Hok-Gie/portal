package domain

import "time"

type Post struct {
	Id       int
	User_Id  int
	Title    string
	Content  string
	CreateAt time.Time
}

type PostWithoutUserId struct {
	Id       int
	Title    string
	Content  string
	CreateAt time.Time
}

type UserPost struct {
	Id       int
	User_Id  int
	Username string
	Title    string
	Content  string
	CreateAt time.Time
}
