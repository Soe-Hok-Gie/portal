package controller

import (
	"encoding/json"
	"medsos/helper"
	"medsos/model/web"
	"medsos/service"
	"net/http"
)

type postControllerImp struct {
	postService service.PostService
}

func (controller *postControllerImp) Save(writer http.ResponseWriter, request *http.Request) {
	//membaca request body
	decoder := json.NewDecoder(request.Body)
	//mengembalikan result, result diambil dari model web
	result := web.PostCreateRequest{}
	err := decoder.Decode(&result)
	helper.PanicIfError(err)

	//memanggil service dan mengembalikan response
	response := controller.postService.Create(request.Context(), result)

	//membuat standart response
	webResponse := web.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   response,
	}

}
