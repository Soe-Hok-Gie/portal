package web

import "medsos/model/domain"

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserWithPostsResponse struct {
	Id       int                        `json:"id"`
	Username string                     `json:"username"`
	Posts    []domain.PostWithoutUserId `json:"posts"`
}
