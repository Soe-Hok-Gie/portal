package web

type PostResponse struct {
	Id      int    `json:"id"`
	User_Id int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
