package middleware

import (
	"github.com/kaantecik/key-value-store/internal/logging"
	"net/http"
	"time"
)

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

// Write function.
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

// WriteHeader function.
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

// HttpLogger function returns a http.Handler.
// When a user send a http request, HttpLogger function logs detail of request.
// Format: [statusCode] METHOD URL took "duration"
func HttpLogger(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lrw := loggingResponseWriter{
			ResponseWriter: rw,
			responseData:   responseData,
		}
		h.ServeHTTP(&lrw, req)

		duration := time.Since(start)

		logging.HttpLogger.Infof("[%d] %v %v took %v", responseData.status, req.Method, req.RequestURI, duration)
	}
	return http.HandlerFunc(loggingFn)
}
