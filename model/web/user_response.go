package web

//// parameter kedua dari service (model dari response)

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
