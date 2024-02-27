package data

import (
	"develop/dev11/internal/errors"
	"develop/dev11/internal/models"
	"fmt"
	"sync"
	"time"
)

// Eventer - интерфейс для работы с событиями
type Eventer interface {
	// Create создает новое событие
	Create(newEvent *models.Event) (int, error)
	// Update обновляет существующее событие
	Update(updataEvent *models.Event) error
	// Delete удаляет событие
	Delete(userID, id int) error
	// GetFor возвращает события для за заданный период
	GetFor(userID int, fromDate, toDate time.Time) ([]*models.Event, error)
}

// EventsData - структура для хранения событий
type EventsData struct {
	mu   sync.RWMutex                   // mu - мьютекс для безопасного доступа к данным
	data map[int]map[uint]*models.Event // data - хранит события для каждого пользователя
	id   uint                           // id - уникальный идентификатор события
}

// New - конструктор EventsData
func New() Eventer {
	return &EventsData{data: make(map[int]map[uint]*models.Event)}
}

// Create - создает новое событие
func (eventsData *EventsData) Create(newEvent *models.Event) (int, error) {
	eventsData.mu.Lock()
	defer eventsData.mu.Unlock()
	// Задаем id для события
	newEvent.ID = int(eventsData.id)
	// Создаем map для user если у него ещё не событий
	if _, ok := eventsData.data[newEvent.UserID]; !ok {
		eventsData.data[newEvent.UserID] = make(map[uint]*models.Event)
	}

	// Добавляем событие
	eventsData.data[newEvent.UserID][eventsData.id] = newEvent
	// Инкрементим id
	eventsData.id++
	return newEvent.ID, nil
}

// Update - обновляет существующее событие
func (eventsData *EventsData) Update(updataEvent *models.Event) error {
	eventsData.mu.Lock()
	defer eventsData.mu.Unlock()

	// Если пользователся или события нет, то возвращаем ошибку
	if err := eventsData.checkEvent(updataEvent.UserID, updataEvent.ID); err != nil {
		return err
	}

	// Обновляем событие
	eventsData.data[updataEvent.UserID][uint(updataEvent.ID)] = updataEvent
	return nil
}

// Delete - удаляет событие
func (eventsData *EventsData) Delete(userID, id int) error {
	eventsData.mu.Lock()
	defer eventsData.mu.Unlock()

	// Если пользователся или события нет, то возвращаем ошибку
	if err := eventsData.checkEvent(userID, id); err != nil {
		return err
	}

	// Удаляем событие
	delete(eventsData.data[userID], uint(id))
	return nil
}

// GetFor - возвращает события для указанного пользователя за заданный период
func (eventsData *EventsData) GetFor(userID int, fromDate, toDate time.Time) ([]*models.Event, error) {
	eventsData.mu.RLock()
	defer eventsData.mu.RUnlock()

	// Если пользователся нет, то возвращаем ошибку
	if _, ok := eventsData.data[userID]; !ok {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user_id %d", userID))
	}
	events := make([]*models.Event, 0)
	// Прохоимся по всем событиям пользователя
	for _, event := range eventsData.data[userID] {
		// Добавляем событие в слайс если date не привышает макс date и не меньше min date
		if event.Date.Before(toDate) && !event.Date.Before(fromDate) {
			events = append(events, event)
		}
	}
	return events, nil
}

// checkEvent - проверяет существование события
func (eventsData *EventsData) checkEvent(userID, id int) error {
	// Возвращаем ошибку если пользователя нет
	if _, ok := eventsData.data[userID]; !ok {
		return errors.NewNotFoundError(fmt.Sprintf("user_id %d", userID))
	}
	// Возвращаем ошибку если события нет
	if _, ok := eventsData.data[userID][uint(id)]; !ok {
		return errors.NewNotFoundError(fmt.Sprintf("event id %d", id))
	}
	return nil
}
