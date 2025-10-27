package web

import "time"

type PostCreateRequest struct {
	User_Id  int
	Title    string
	Content  string
	CreateAt time.Time
}
