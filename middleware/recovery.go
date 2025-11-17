package middleware

import (
	"medsos/exception"
	"net/http"
)

// Ini adalah middleware.
// Middleware selalu menerima next, yaitu handler berikutnya (route yang sebenarnya).
// Dan harus mengembalikan http.Handler.
func Recovery(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// middleware ini membungkus handler berikutnya dengan fungsi tambahan.
		// Jadi sebelum masuk ke handler asli, kode di dalam sini dijalankan.

		// defer = jalankan kode ini setelah fungsi selesai.
		defer func() {
			//looping
			if rec := recover(); rec != nil {
				exception.ErrorHandler(w, r, rec)
			}

		}()
		Next.ServeHTTP(w, r) //Ini menjalankan handler berikutnya (route asli yang dibuat).
	})

}
