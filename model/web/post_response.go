package web

import "time"

type PostResponse struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	User_Id    int       `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_At time.Time `json:"created_at"`

	// User    domain.User `json:"user"` //belongs to user
}

type PostWithoutUserResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
