package controller

import (
	"medsos/service"
	"net/http"
)

type postControllerImp struct {
	postService service.PostService
}

func (controller *postControllerImp) Save(writer http.ResponseWriter, request *http.Request) {

}
