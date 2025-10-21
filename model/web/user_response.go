package web

//// parameter kedua dari service (model dari response)

type UserResponse struct {
	id       int    `json:"id"`
	username string `json:"username"`
}
