package controller

import (
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
	//membaca request body
	decoder := json.NewDecoder(request.Body)
	//mengembalikan result, result diambil dari model web
	result := web.UserCreateRequest{}
	err := decoder.Decode(&result)
	helper.PanicIfError(err)

	//memanggil service dan mengembalikan response
	response := controller.userService.Create(request.Context(), result)

	//membuat standart respose
	webResponse := web.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	//proses mengubah data menjadi format lain
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

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
	helper.PanicIfError(err)
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
