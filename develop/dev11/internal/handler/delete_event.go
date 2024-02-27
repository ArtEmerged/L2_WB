package handler

import (
	"develop/dev11/internal/errors"
	"net/http"
	"strconv"
)

// deleteEvent обрабатывает запрос на удаление события
func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса на соответствие POST
	if r.Method != http.MethodPost {
		responsErrorJSON(w, errors.NewBadMethodError(r.Method, http.MethodPost), http.StatusMethodNotAllowed)
		return
	}

	// Проверяем наличие необходимых параметров в теле POST запроса
	if err := checkPostRequrst(r, "user_id", "id"); err != nil {
		responsErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Извлекаем и проверяем целочисленные значения user_id и id из тела POST запроса
	userID, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil {
		responsErrorJSON(w, errors.NewBadRequestError("invalid user_id: use only numbers"), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		responsErrorJSON(w, errors.NewBadRequestError("invalid id: use only numbers"), http.StatusBadRequest)
		return
	}

	// Удаляем событие через service
	err = h.service.Delete(userID, id)
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
	// Возвращаем успешный ответ
	responsJSON(w, "OK", http.StatusOK)
}
