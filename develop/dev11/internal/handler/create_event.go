package handler

import (
	"develop/dev11/internal/errors"
	"net/http"
)

// createEvent обрабатывает запрос на создание события
func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса на соответствие POST
	if r.Method != http.MethodPost {
		responsErrorJSON(w, errors.NewBadMethodError(r.Method, http.MethodPost), http.StatusMethodNotAllowed)
		return
	}

	// Проверяем наличие необходимых параметров в теле POST запроса
	if err := checkPostRequrst(r, "user_id", "date", "title"); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Парсим событие из тела запроса
	createEvent, err := parseEventFromRequest(r, false)
	if err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Проверяем валидность события
	if err := createEvent.Validate(); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Создаем событие через service
	eventID, err := h.service.Create(createEvent)
	if err != nil {
		// Если произошла ошибка, проверяем, соответствует ли она ошибке нашего интерфейса (HTTPError)
		if httpError, ok := err.(errors.HTTPError); ok {
			responsErrorJSON(w, httpError, httpError.StatusCode())
			return
		}
		// Если это другая ошибка, возвращаем внутреннюю серверную ошибку
		responsErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	// Возвращаем успешный ответ с ID созданного события
	responsJSON(w, map[string]any{"eventID": eventID}, http.StatusCreated)
}
