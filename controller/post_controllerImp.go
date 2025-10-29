package controller

import (
	"encoding/json"
	"medsos/helper"
	"medsos/model/web"
	"medsos/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	//mencetak header json dan melakukan proses encoding
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *postControllerImp) Update(writer http.ResponseWriter, request *http.Request) {
	//ambil dan convert id
	vars := mux.Vars(request)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	//buat var untuk menampung web.postupdaterequest
	// Dekode data JSON dari body permintaan
	var post web.PostUpdateRequest
	err = json.NewDecoder(request.Body).Decode(&post)
	helper.PanicIfError(err)

	//simpan id sebelum dipassing
	post.Id = id

}
