package domain

import "time"

type Post struct {
	Id       int
	User_Id  int
	Title    string
	Content  string
	CreateAt time.Time
}
