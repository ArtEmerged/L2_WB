package handler

import (
	"develop/dev11/internal/errors"
	"develop/dev11/internal/service"
	"net/http"
)

// Handler - структура обработчика HTTP-запросов
type Handler struct {
	service *service.Service 
}

// New - конструктор для Handler
func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

// InitRouter - метод инициализации роутера HTTP-запросов
func (h *Handler) InitRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.getEventsForDay)
	mux.HandleFunc("/events_for_week", h.getEventsForWeek)
	mux.HandleFunc("/events_for_month", h.getEventsForMonth)

	// Обработка запросов, которые не соответствуют ни одному из обработчиков
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			responsErrorJSON(w, errors.NewNotFoundError("path "+r.URL.Path), http.StatusNotFound)
		}
	})

	return httpLogger(mux)
}
