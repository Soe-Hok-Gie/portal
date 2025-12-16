package service

import (
	"context"
	"fmt"
	"medsos/helper"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/repository"
)

type postServiceImp struct {
	PostRepository repository.PostRepository
}

// polimerisme
func NewPostService(postRepository repository.PostRepository) PostService {
	return &postServiceImp{PostRepository: postRepository}
}
func (service *postServiceImp) Create(ctx context.Context, request web.PostCreateRequest) (web.PostResponse, error) {
	// tampung model domain dalam sebuah variabel
	post := domain.Post{
		User_Id:    request.User_Id,
		Title:      request.Title,
		Content:    request.Content,
		Created_At: request.CreateAt,
	}

	//panggil repository untuk mapping ke DB.
	post, err := service.PostRepository.Save(ctx, post)
	if err != nil {
		return web.PostResponse{}, fmt.Errorf("service create error :%w", err)
	}

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return postResponse, nil

}

func (service *postServiceImp) Update(ctx context.Context, request web.PostUpdateRequest) (web.PostResponse, error) {
	// tampung model domain dalam sebuah variabel
	post := domain.Post{
		Id:      request.Id,
		User_Id: request.User_Id,
		Title:   request.Title,
		Content: request.Content,
	}

	//panggil service
	post, err := service.PostRepository.Update(ctx, post)
	if err != nil {
		return web.PostResponse{}, fmt.Errorf("failed to update post: %w", err)
	}

	// tampung model web response dalam sebuah variabel
	postResponse := web.PostResponse{
		Id:      post.Id,
		User_Id: post.User_Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return postResponse, nil

}

func (service *postServiceImp) FindById(ctx context.Context, postId int) (web.PostResponse, error) {
	//deklarasi
	var postResponse web.PostResponse
	//panggil service
	post, err := service.PostRepository.FindById(ctx, postId)

	if err != nil {
		return postResponse, err
	}

	// tampung model web response dalam sebuah variabel
	postResponse = web.PostResponse{
		Id:       post.Id,
		Username: post.Username,
		User_Id:  post.User_Id,
		Title:    post.Title,
		Content:  post.Content,
	}
	return postResponse, nil

}

func (service *postServiceImp) FindAll(ctx context.Context, filter domain.PostFilter) []web.PostResponse {

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
