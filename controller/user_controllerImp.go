package controller

import (
	"encoding/json"
	"medsos/helper"
	"medsos/model/web"
	"medsos/service"
	"net/http"

	"github.com/gorilla/mux"
)

type userControllerImp struct {
	userService service.UserService
}

// polimerisme
func NewUserController(userService service.UserService) UserController {
	return &userControllerImp{userService: userService}
}

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

func (controller *userControllerImp) Update(writer http.ResponseWriter, request *http.Request) {
	// 	decoder := json.NewDecoder(request.Body)

	// ProductUpdateRequest := web.ProductpdateRequest{}

	// productId := params.ByName("productId")
	// id, err := strconv.Atoi(productId)
	// ProductUpdateRequest.Id = id

	// decoder.Decode(&ProductUpdateRequest)
	// if err := controller.validate.Struct(ProductUpdateRequest); err != nil {
	// 	panic(err)

	// }

	// if err != nil {
	// 	panic(err)
	// }

	//ambil Id
	vars := mux.Vars(request)
	Id := vars["id"]
	// 	vars := mux.Vars(r)
	// 	idStr := vars["id"] // idStr is a string
	// 	id, err := strconv.Atoi(idStr)
	// 	if err != nil {
	// 		http.Error(writer, "Invalid ID format", http.StatusBadRequest)
	// 		return
	// 	}

	// 	fmt.Fprintf(writer, "Received item ID: %d", id)
	// }

	// 2. Dekode data JSON dari body permintaan
	var user web.UserUpdateRequest
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := controller.userService.Update(request.Context(), user)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   response,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)

}
