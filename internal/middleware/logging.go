package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrapperWritter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapperWritter) WriterHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapper := &wrapperWritter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		log.Println(wrapper.statusCode, r.URL.Path, time.Since(start))
	})
}
