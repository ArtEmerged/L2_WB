package errors

import "fmt"

// HTTPError - интерфейс для представления ошибок HTTP
type HTTPError interface {
	StatusCode() int // StatusCode возвращает код состояния HTTP
	Error() string   // Error возвращает описание ошибки
}

// BadRequestError - ошибка "Неверный запрос"
type BadRequestError struct {
	err        string // Описание ошибки
	statusCode int    // Код состояния HTTP
}

// NewBadRequestError - конструктор для создания BadRequestError
func NewBadRequestError(err string) *BadRequestError {
	return &BadRequestError{
		err:        err,
		statusCode: 400,
	}
}

// Error возвращает текст ошибки 
func (b BadRequestError) Error() string {
	return fmt.Sprintf("bad request: %s", b.err)
}

// StatusCode возвращает код ошибки 
func (b BadRequestError) StatusCode() int {
	return b.statusCode
}

// BadMethodError - ошибка "Неверный метод"
type BadMethodError struct {
	wrongMethod string // Неверный метод
	rightMethod string // Правильный метод
	statusCode  int    // Код состояния HTTP
}

// NewBadMethodError - конструктор для создания BadMethodError
func NewBadMethodError(wrongMethod, rightMethod string) *BadMethodError {
	return &BadMethodError{
		wrongMethod: wrongMethod,
		rightMethod: rightMethod,
		statusCode:  405,
	}
}

// Error возвращает текст ошибки 
func (b BadMethodError) Error() string {
	return fmt.Sprintf("method not allowed: bad method %s, method must be %s", b.wrongMethod, b.rightMethod)
}

// StatusCode возвращает код ошибки 
func (e BadMethodError) StatusCode() int {
	return e.statusCode
}

// NotFoundError - ошибка "Не найдено"
type NotFoundError struct {
	err        string // Описание ошибки
	statusCode int    // Код состояния HTTP
}

// NewNotFoundError - конструктор для создания NotFoundError
func NewNotFoundError(err string) *NotFoundError {
	return &NotFoundError{
		err:        err,
		statusCode: 404,
	}
}

// Error возвращает текст ошибки 
func (n NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", n.err)
}

// StatusCode возвращает код ошибки 
func (n NotFoundError) StatusCode() int {
	return n.statusCode
}

// ServiceUnavailableError - ошибка "Служба недоступна"
type ServiceUnavailableError struct {
	statusCode int // Код состояния HTTP
}

// NewServiceUnavailableError - конструктор для создания ServiceUnavailableError
func NewServiceUnavailableError() *ServiceUnavailableError {
	return &ServiceUnavailableError{
		statusCode: 503,
	}
}

// Error возвращает текст ошибки 
func (s ServiceUnavailableError) Error() string {
	return "service unavailable"
}

// StatusCode возвращает код ошибки 
func (s ServiceUnavailableError) StatusCode() int {
	return s.statusCode
}

// InternalServerError - ошибка "Внутренняя ошибка сервера"
type InternalServerError struct {
	statusCode int    // Код состояния HTTP
	err        string // Описание ошибки
}

// NewInternalServerError - конструктор для создания InternalServerError
func NewInternalServerError(err string) *InternalServerError {
	return &InternalServerError{
		statusCode: 500,
		err:        err,
	}
}

// Error возвращает текст ошибки 
func (i InternalServerError) Error() string {
	return fmt.Sprintf("internal server error: %s", i.err)
}

// StatusCode возвращает код ошибки 
func (i InternalServerError) StatusCode() int {
	return i.statusCode
}
