package service

import (
	"develop/dev11/internal/data"
	"develop/dev11/internal/models"
	"time"
)

// EventService - структура сервиса событий
type EventService struct {
	data data.Eventer
}

// NewEventService - конструктор для EventService
func NewEventService(data data.Eventer) *EventService {
	return &EventService{data: data}
}

// Create - метод для создания нового события
func (eventService *EventService) Create(createEvent *models.Event) (int, error) {
	return eventService.data.Create(createEvent)
}

// Update - метод для обновления существующего события
func (eventService *EventService) Update(updateEvent *models.Event) error {
	return eventService.data.Update(updateEvent)
}

// Delete - метод для удаления события
func (eventService *EventService) Delete(userID, id int) error {
	return eventService.data.Delete(userID, id)
}

// GetFor - метод для получения событий для указанного пользователя в определенном периоде времени
func (eventService *EventService) GetFor(userID int, fromDate time.Time, mode string) ([]*models.Event, error) {
	// toDate максимальная дата
	toDate := fromDate
	switch mode {
	case ModeForDay:
		// добавляем к максимальной дате 1 день от стартовой даты
		toDate = toDate.AddDate(0, 0, 1)
	case ModeForWeek:
		// добавляем к максимальной дате 7 дней от стартовой даты
		toDate = toDate.AddDate(0, 0, 7)
	case ModeForMonth:
		// добавляем к максимальной дате 1 месяц от стартовой даты
		toDate = toDate.AddDate(0, 1, 0)
	}

	return eventService.data.GetFor(userID, fromDate, toDate)
}
