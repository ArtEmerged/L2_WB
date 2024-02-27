package handler

import (
	"develop/dev11/internal/errors"
	"net/http"
)

// updateEvent обрабатывает запрос на обновление события
func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса на соответствие POST
	if r.Method != http.MethodPost {
		responsErrorJSON(w, errors.NewBadMethodError(r.Method, http.MethodPost), http.StatusMethodNotAllowed)
		return
	}

	// Проверяем наличие необходимых параметров в теле POST запроса
	if err := checkPostRequrst(r, "user_id", "id", "date", "title"); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Парсим событие из тела запроса
	updateEvent, err := parseEventFromRequest(r, true)
	if err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Проверяем валидность события
	if err := updateEvent.Validate(); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Обновляем событие через service
	err = h.service.Update(updateEvent)
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

	// Отправляем подтверждение об успешном обновлении события
	responsJSON(w, "OK", http.StatusOK)
}
