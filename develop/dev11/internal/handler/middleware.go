package handler

import (
	"log"
	"net/http"
	"time"
)

// httpLogger - функция для логирования HTTP-запросов
func httpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rec := NewResponseRecorder(w)
		handler.ServeHTTP(rec, r)
		duration := time.Since(startTime)
		
		level := "[INFO]"
		switch {
		case rec.statusCode >= 500:
			level = "[ERROR]"
		case rec.statusCode >= 400:
			level = "[WARN]"
		}

		log.Printf("%s received a HTTP request Method=%s Status=%d Path=%s Duration=%s\n", level, r.Method, rec.statusCode, r.RequestURI, duration)

	})
}

// ResponseRecorder - структура для записи статус-кода HTTP-ответа
type ResponseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// NewResponseRecorder - конструктор для ResponseRecorder
func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{ResponseWriter: w}
}

// WriteHeader - метод для установки статус-кода HTTP-ответа и сохранения его в ResponseRecorder
func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}
