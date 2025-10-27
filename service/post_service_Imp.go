package service

import (
	"context"
	"medsos/model/web"
)

type postServiceImp struct {
}

func (service postServiceImp) Create(ctx context.Context, request web.PostCreateRequest) web.PostCreateRequest
