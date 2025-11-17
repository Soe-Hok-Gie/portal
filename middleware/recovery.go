package middleware

import "net/http"

// Ini adalah middleware.
// Middleware selalu menerima next, yaitu handler berikutnya (route yang sebenarnya).
// Dan harus mengembalikan http.Handler.
func Recovery(Next http.Handler) http.Handler {

}
