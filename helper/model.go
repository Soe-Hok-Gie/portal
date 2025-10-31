package helper

import (
	"medsos/model/domain"
	"medsos/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
}
