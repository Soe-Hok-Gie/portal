package helper

import (
	"medsos/model/domain"
	"medsos/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:         post.Id,
		Username:   post.Username,
		User_Id:    post.User_Id,
		Title:      post.Title,
		Content:    post.Content,
		Created_At: post.Created_At,
	}
}
