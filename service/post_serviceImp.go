package service

import (
	"context"
	"fmt"
	"medsos/helper"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
	"strings"
)

type postServiceImp struct {
	PostRepository repository.PostRepository
}

// polimerisme
func NewPostService(postRepository repository.PostRepository) PostService {
	return &postServiceImp{PostRepository: postRepository}
}
func (service *postServiceImp) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	// tampung model domain dalam sebuah variabel
	post := domain.Post{
		User_Id:    request.User_Id,
		Title:      request.Title,
		Content:    request.Content,
		Created_At: request.CreateAt,
	}

	//panggil service
	post = service.PostRepository.Save(ctx, post)

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return postResponse

}

func (service *postServiceImp) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	// tampung model domain dalam sebuah variabel
	post := domain.Post{
		Id:      request.Id,
		User_Id: request.User_Id,
		Title:   request.Title,
		Content: request.Content,
	}

	//panggil service
	post = service.PostRepository.Update(ctx, post)

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return postResponse

}

func (service *postServiceImp) FindById(ctx context.Context, postId int) web.PostResponse {
	//panggil service
	post, err := service.PostRepository.FindById(ctx, postId)
	helper.PanicIfError(err)
	fmt.Println("post", post)

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:       post.Id,
		Username: post.Username,
		User_Id:  post.User_Id,
		Title:    post.Title,
		Content:  post.Content,
	}
	return postResponse

}

func (service *postServiceImp) FindAll(ctx context.Context, filter domain.PostFilter) []web.PostResponse {
	// validasi sort params
	filter.Sort = strings.ToLower(filter.Sort)
	if filter.Sort != "asc" && filter.Sort != "desc" {
		filter.Sort = "desc" //default
	}

	posts := service.PostRepository.FindAll(ctx, filter)
	var postResponses []web.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, helper.ToPostResponse(post))
	}
	return postResponses
}

func (service *postServiceImp) Delete(ctx context.Context, postId int) {

	service.PostRepository.Delete(ctx, postId)
}
