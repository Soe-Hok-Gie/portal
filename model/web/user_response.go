package web

//// parameter kedua dari service (model dari response)

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Posts    []PostResponseTanpaUserId `json:"posts"` //one to many
}

// type UsersResponse struct {
// 	Id       int    `json:"id"`
// 	Username string `json:"username"`
// 	// PostsCount int    `json:"posts_count"` //one to many
// }
