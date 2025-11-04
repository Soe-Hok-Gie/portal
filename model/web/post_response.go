package web

type PostResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	User_Id  int    `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	// User    domain.User `json:"user"` //belongs to user
}

// type PostResponseTanpaUserId struct {
// 	Id      int    `json:"id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }
