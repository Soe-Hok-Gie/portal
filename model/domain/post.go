package domain

import "time"

type Post struct {
	Id int
	// Username string
	User_Id    int
	Title      string
	Content    string
	Created_At time.Time
}

type PostWithoutUserId struct {
	Id         int
	Title      string
	Content    string
	Created_At time.Time
}

type UserPost struct {
	Id         int
	User_Id    int
	Username   string
	Title      string
	Content    string
	Created_At time.Time
}

type PostFilter struct {
	Sort string
}
