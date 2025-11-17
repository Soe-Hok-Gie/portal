package middleware

import "net/http"

// Ini adalah middleware.
// Middleware selalu menerima next, yaitu handler berikutnya (route yang sebenarnya).
// Dan harus mengembalikan http.Handler.
func Recovery(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// middleware ini membungkus handler berikutnya dengan fungsi tambahan.
		// Jadi sebelum masuk ke handler asli, kode di dalam sini dijalankan.

	})

}
