package controller

import (
	"net/http"
)

type userControllerImp struct {
}

func (controller *userControllerImp) Create(writer http.ResponseWriter, request *http.Request) UserController {
	return &userControllerImp
}
