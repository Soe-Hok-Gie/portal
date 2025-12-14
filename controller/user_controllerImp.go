package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"medsos/helper"
	"medsos/model/web"
	"medsos/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type userControllerImp struct {
	userService service.UserService
}

// polimerisme
func NewUserController(userService service.UserService) UserController {
	return &userControllerImp{userService: userService}
}

// create
func (controller *userControllerImp) Create(writer http.ResponseWriter, request *http.Request) {
	// //membaca request body
	decoder := json.NewDecoder(request.Body)
	// membuat sebuah Var,  diambil dari model web
	var userReq web.UserCreateRequest
	err := decoder.Decode(&userReq)
	if err != nil {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.Response{
			//membuat standart respose
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   "invalid",
		})
		return
	}
	//memanggil service dan mengembalikan response
	response, err := controller.userService.Create(request.Context(), userReq)
	if err != nil {
		fmt.Println("err", err)
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.Response{
			//membuat standart respose
			Code:   http.StatusInternalServerError,
			Status: "internal server error",
			Data:   nil,
		})
		return
	}

	writer.Header().Set("content-type", "application/jso")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(web.Response{
		//membuat standart respose
		Code:   http.StatusCreated,
		Status: "created",
		Data:   response,
	})
}

// update
func (controller *userControllerImp) Update(writer http.ResponseWriter, request *http.Request) {

	//ambil Id
	vars := mux.Vars(request)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	helper.PanicIfError(err)

	//console.log
	fmt.Printf("user:  %v", id)

	// 2. Dekode data JSON dari body permintaan
	var user web.UserUpdateRequest
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", user)
	//simpan id sebelum passing
	user.Id = id
	fmt.Printf("%v", user)

	response := controller.userService.Update(request.Context(), user)
	fmt.Printf("%v", user)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

// findbyid
func (controller *userControllerImp) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["id"]
	id, err := strconv.Atoi(userId)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(writer)
		encoder.Encode("internal server error")
		return
	}
	response, err := controller.userService.FindById(request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			encoder := json.NewEncoder(writer)
			webResponse := web.Response{
				Code:   http.StatusNotFound,
				Status: "Not Found",
				Data:   nil,
			}
			encoder.Encode(webResponse)
		} else {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			encoder := json.NewEncoder(writer)
			webResponse := web.Response{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Data:   nil,
			}
			encoder.Encode(webResponse)
		}
		return
	}
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

// delete
func (controller *userControllerImp) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["id"]
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.userService.Delete(request.Context(), id)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

// findAll
func (controller *userControllerImp) FindAll(writer http.ResponseWriter, request *http.Request) {

	userResponses := controller.userService.FindAll(request.Context())

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}

// FindUserPost
func (controller *userControllerImp) FindUserPost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["id"]
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	response := controller.userService.FindUserPost(request.Context(), id)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)
}
