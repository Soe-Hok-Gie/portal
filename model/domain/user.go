package domain

// representasi tabel user
type User struct {
	Id       int
	Username string
}

type UserPosts struct {
	Id       int
	Username string
	Posts    []PostWithoutUserId
}
