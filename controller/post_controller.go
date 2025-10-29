package controller

import "net/http"

type PostController interface {
	Save(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
}
