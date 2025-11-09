package web

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserWithPostsResponse struct {
	Id       int                       `json:"id"`
	Username string                    `json:"username"`
	Posts    []PostWithoutUserResponse `json:"posts"`
}
