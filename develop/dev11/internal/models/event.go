package models

import (
	"develop/dev11/internal/errors"
	"fmt"
	"time"
	"unicode/utf8"
)

// Event представляет событие в календаре пользователя.
// Являемся объектом доменной области
type Event struct {
	UserID      int       `json:"user_id"`     // ID пользователя, владельца события
	ID          int       `json:"id"`          // Уникальный идентификатор события
	Date        time.Time `json:"date"`        // Дата и время проведения события
	Title       string    `json:"title"`       // Заголовок события
	Description string    `json:"description"` // Описание события
}

// NewEvent - конструктор для Event
func NewEvent(userID, id int, date time.Time, title, description string) *Event {
	return &Event{
		UserID:      userID,
		ID:          id,
		Date:        date,
		Title:       title,
		Description: description,
	}
}

// Validate выполняет валидацию полей события
func (event *Event) Validate() error {
	if event.UserID < 0 {
		return errors.NewBadRequestError("user_id must be positive")
	}

	if event.ID < 0 {
		return errors.NewBadRequestError("id must be positive")
	}

	if time.Now().After(event.Date) {
		return errors.NewBadRequestError("event date cannot be in the past")
	}

	if err := validateText("title", event.Title, 20, true); err != nil {
		return err
	}

	if err := validateText("description", event.Description, 50, false); err != nil {
		return err
	}
	return nil
}

// validateText выполняет валидацию текстовых полей 
func validateText(parameter, text string, textLength int, noEmpty bool) error {
	if noEmpty && text == "" {
		return errors.NewBadRequestError("empty parameter: " + parameter)
	}
	if utf8.RuneCountInString(text) > textLength {
		return errors.NewBadRequestError(fmt.Sprintf("%s parameter is too long, maximum length %d symbols", parameter, textLength))
	}
	return nil
}
