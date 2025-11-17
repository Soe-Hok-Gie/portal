package exception

import "net/http"

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	InternalServerError(writer, request, err)
}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

}
