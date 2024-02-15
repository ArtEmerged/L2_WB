package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
/*
	Часто используют шаблон проектирования фасада, когда система очень сложна или трудна для понимания,
	поскольку в системе имеется множество взаимозависимых классов или ее исходный код недоступен.
	Этот шаблон скрывает сложности более крупной системы и обеспечивает более простой интерфейс для клиента.

	+ Изолирует клиентов от компонентов сложной подсистемы.
	- Фасад рискует стать "супер-объектом", привязанным ко всем классам программы.
*/

// NotificationFacade фасад для уведомлений
type NotificationFacade struct {
	email *Email
	app   *App
	sms   *SMS
	user  *User
}

// Send отправяет уведомления через все каналы связи.
// Клиенту не нужно знать всю бизнес-логику, куда и как будут отправляться уведомления.
// Для него все просто и понятно: вызывая этот метод, он отправит уведомление покупателю о готовности заказа.
func (n *NotificationFacade) Send(userID int, orderID string) {
	user := n.user.GetUserByID(userID)
	n.email.SendsNotice(user, orderID)
	n.app.SendsNotice(user, orderID)
	n.sms.SendsNotice(user, orderID)
}

// User обеспечивает взамодействие c User
type User struct {
	ID   int
	Name string
}

// GetUserByID находит ures по userID
func (u *User) GetUserByID(ID int) *User {
	return &User{ID, "Петя"}
}

// Email обеспечивает взамодействие с email сервисом
type Email struct{}

// SendsNotice отправляет уведомление пользователю по email
func (e *Email) SendsNotice(user *User, orderID string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления по email,
	// у структуры Email могут присутствовать подобные методы (GetEmailByUserID, EmailServiceConnection)
	fmt.Printf("[email] %s, ваш заказ:%s готов к выдаче!\n", user.Name, orderID)
}

// App обеспечивает взамодействие c приложением
type App struct{}

// SendsNotice отправляет уведомление пользователю в приложение
func (a *App) SendsNotice(user *User, orderID string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления в приложение,
	// у структуры App могут присутствовать подобные методы (AppServiceConnection)
	fmt.Printf("[app] %s, ваш заказ:%s готов к выдаче!\n", user.Name, orderID)
}

// SMS обеспечивает взамодействие c sms сервисом
type SMS struct{}

// SendsNotice отправляет уведомление пользователю по SMS
func (s *SMS) SendsNotice(user *User, orderID string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления по SMS,
	// у структуры Sms могут присутствовать подобные методы (GetPhoneByUserID, SmsServiceConnection)
	fmt.Printf("[sms] %s, ваш заказ:%s готов к выдаче!\n", user.Name, orderID)
}

// func main() {
// 	notice := NotificationFacade{}
// 	notice.Send(444321, "DfE42vD")
// }
