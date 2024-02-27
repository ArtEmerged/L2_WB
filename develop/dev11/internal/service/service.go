package service

import (
	"develop/dev11/internal/data"
	"develop/dev11/internal/models"
	"time"
)

// Константы для определения режимов просмотра событий
const (
	ModeForDay   = "day"
	ModeForWeek  = "week"
	ModeForMonth = "month"
)

// Eventer - интерфейс для работы с событиями
type Eventer interface {
	// Create создает новое событие
	Create(createEvent *models.Event) (int, error)
	// Update обновляет существующее событие
	Update(updateEvent *models.Event) error
	// Delete удаляет событие
	Delete(userID, id int) error
	// GetFor возвращает события для указанного пользователя
	GetFor(userID int, date time.Time, mode string) ([]*models.Event, error)
}

// Service - структура сервиса
type Service struct {
	Eventer
}

// New - конструктор для Service
func New(data data.Eventer) *Service {
	return &Service{NewEventService(data)}
}
