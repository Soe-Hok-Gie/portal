package exception

import (
	"encoding/json"
	"medsos/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if NotFoundError(writer, request, err) {
		return
	}

	InternalServerError(writer, request, err)
}
func NotFoundError(writer http.ResponseWriter, request http.Request, err interface{}) bool {

}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponse)
	// helper.WriteToResponseBody(writer, webResponse)//buat sebuah func dihelper untuk refactor decode and decode
}
