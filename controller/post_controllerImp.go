package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"medsos/helper"
	"medsos/model/domain"
	"medsos/model/web"
	"medsos/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type postControllerImp struct {
	postService service.PostService
}

// polimerisme
func NewPostController(postService service.PostService) PostController {
	return &postControllerImp{postService: postService}
}

func (controller *postControllerImp) Create(writer http.ResponseWriter, request *http.Request) {
	//membaca request body
	decoder := json.NewDecoder(request.Body)
	//  membuat sebuah var , diambil dari model web
	var postReq web.PostCreateRequest
	err := decoder.Decode(&postReq)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.Response{
			//membuat standart response
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "invalid Request",
		})
		return
	}
	// //memanggil service dan mengembalikan response
	response, err := controller.postService.Create(request.Context(), postReq)
	if err != nil {
		fmt.Println("err", err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.Response{
			//membuat standart response
			Code:   http.StatusInternalServerError,
			Status: "internal server error",
			Data:   nil,
		})
		return
	}
	// //mencetak header json dan melakukan proses encoding
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(web.Response{
		//membuat standart response
		Code:   http.StatusCreated,
		Status: "Create",
		Data:   response,
	})
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

	//panggil service
	response := controller.postService.Update(request.Context(), post)

	//membuat standar response
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   response,
	}
	//mencetak header dan melakukan proses encoding
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *postControllerImp) FindById(writer http.ResponseWriter, request *http.Request) {
	//ambil dan convert id
	vars := mux.Vars(request)
	postId := vars["id"]
	id, err := strconv.Atoi(postId)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(writer)
		encoder.Encode("internal servel error")
		return
	}

	//panggil service
	response, err := controller.postService.FindById(request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			//sql.ErrNoRows adalah error khusus dalam paket database/sql pada Go
			// yang muncul ketika query tidak mengembalikan baris sama sekali.
			// Error ini bukan berarti query gagal, tapi hanya menandakan bahwa hasilnya kosong.
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			encoder := json.NewEncoder(writer)
			webResponse := web.Response{
				Code:   http.StatusNotFound,
				Status: "Not Ok",
				Data:   nil,
			}
			encoder.Encode(webResponse)
		} else {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			encoder := json.NewEncoder(writer)
			webResponse := web.Response{
				Code:   http.StatusInternalServerError,
				Status: "Not Ok",
				Data:   nil,
			}
			encoder.Encode(webResponse)
		}

		return
	}
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

func (controller *postControllerImp) FindAll(writer http.ResponseWriter, request *http.Request) {
	//panggil struct filter
	filter := domain.PostFilter{}
	//ambil query params
	filter.Sort = strings.ToUpper(request.URL.Query().Get("sort"))
	//looping
	if filter.Sort == "" {
		filter.Sort = "DESC"
	}
	//tambahkan flter pada parameter kedua
	postResponses := controller.postService.FindAll(request.Context(), filter)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   postResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

func (controller *postControllerImp) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	postId := vars["id"]
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	controller.postService.Delete(request.Context(), id)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}
