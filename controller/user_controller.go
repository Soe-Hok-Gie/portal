package controller

import "net/http"

// parameternya mengikuti http handler
type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request)
}
