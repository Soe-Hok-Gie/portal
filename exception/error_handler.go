package exception

import (
	"encoding/json"
	"medsos/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		encoder.Encode(webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
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
