package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/zVSciy/EventManager/Payment/internal/util"
)

type wrappedWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := util.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
