package handler

import (
	"develop/dev11/internal/errors"
	"develop/dev11/internal/service"
	"net/http"
	"strconv"
	"time"
)

// getEventsForDay обрабатывает запрос на получение событий на конкретный день
func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса на соответствие GET
	if r.Method != http.MethodGet {
		responsErrorJSON(w, errors.NewBadMethodError(r.Method, http.MethodGet), http.StatusMethodNotAllowed)
		return
	}
	
	// Проверяем наличие необходимых параметров в запросе
	if err := checkGetRequrst(r, "user_id", "date"); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Извлекаем и проверяем целочисленное значение user_id из строки запроса
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		responsErrorJSON(w, errors.NewBadRequestError("invalid user_id: use only numbers"), http.StatusBadRequest)
		return
	}

	// Извлекаем и парсим дату из строки запроса
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		responsErrorJSON(w, errors.NewBadRequestError("invalid date format: correct format 2006-01-02"), http.StatusBadRequest)
		return
	}

	// Получаем события для указанного пользователя и даты через service
	events, err := h.service.GetFor(userID, date, service.ModeForDay)
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

	// Возвращаем полученные события в формате JSON
	responsJSON(w, events, http.StatusOK)
}
