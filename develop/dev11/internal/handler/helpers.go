package handler

import (
	"develop/dev11/internal/errors"
	"develop/dev11/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// checkGetRequrst - функция для проверки GET-запроса на наличие в URL Query обязательных параметров
func checkGetRequrst(r *http.Request, keys ...string) error {
	for _, key := range keys {
		value := r.URL.Query().Get(key)
		if value == "" {
			return errors.NewBadRequestError("empty parameter: " + key)
		}
	}
	return nil
}

// checkPostRequrst - функция для проверки POST-запроса на наличие в body обязательных form параметров
func checkPostRequrst(r *http.Request, keys ...string) error {
	for _, key := range keys {
		value := r.PostFormValue(key)
		if value == "" {
			return errors.NewBadRequestError("empty parameter: " + key)
		}
	}
	return nil
}

// parseEventFromRequest - функция для парсинга данных о событии (Event) из HTTP-запроса
func parseEventFromRequest(r *http.Request, checkEventID bool) (*models.Event, error) {
	userID, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil {
		return nil, errors.NewBadRequestError("invalid user_id: use only numbers")
	}
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil && checkEventID {
		return nil, errors.NewBadRequestError("invalid id: use only numbers")
	}

	date, err := time.Parse(time.DateTime, r.PostFormValue("date"))
	if err != nil {
		return nil, errors.NewBadRequestError("invalid date format: correct format 2006-01-02 15:04:05")
	}

	title := strings.TrimSpace(r.PostFormValue("title"))
	description := strings.TrimSpace(r.PostFormValue("description"))

	return models.NewEvent(userID, id, date, title, description), nil
}
