package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
/*
	Часто используют шаблон проектирования фасада, когда система очень сложна или трудна для понимания,
	поскольку в системе имеется множество взаимозависимых классов или ее исходный код недоступен.
	Этот шаблон скрывает сложности более крупной системы и обеспечивает более простой интерфейс для клиента.

	+ Изолирует клиентов от компонентов сложной подсистемы.
	- Фасад рискует стать "супер-объектом", привязанным ко всем классам программы.
*/

// Фасад для уведомлений
type NotificationFacade struct {
	email *Email
	app   *App
	sms   *Sms
	user  *User
}

/*
Отправка уведомлений через все каналы связи.
Клиенту не нужно знать всю бизнес-логику, куда и как будут отправляться уведомления.
Для него все просто и понятно: вызывая этот метод, он отправит уведомление покупателю о готовности заказа.
*/
func (n *NotificationFacade) Send(userId int, orderId string) {
	user := n.user.GetUserById(userId)
	n.email.SendsNotice(user, orderId)
	n.app.SendsNotice(user, orderId)
	n.sms.SendsNotice(user, orderId)
}

type User struct {
	Id   int
	Name string
}

func (u *User) GetUserById(id int) *User {
	return &User{5, "Петя"}
}

type Email struct{}

// Отправляет уведомление пользователю по email
func (e *Email) SendsNotice(user *User, orderId string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления по email,
	// у структуры Email могут присутствовать подобные методы (GetEmailByUserId, EmailServiceConnection)
	fmt.Printf("[email] %s, ваш заказ:%s готов к выдачи\n", user.Name, orderId)
}

type App struct{}

// Отправляет уведомление пользователю в приложение
func (a *App) SendsNotice(user *User, orderId string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления в приложение,
	// у структуры App могут присутствовать подобные методы (AppServiceConnection)
	fmt.Printf("[app] %s, ваш заказ:%s готов к выдачи\n", user.Name, orderId)
}

type Sms struct{}

// Отправляет уведомление пользователю по SMS
func (s *Sms) SendsNotice(user *User, orderId string) {
	// Здесь может содержаться вся бизнес-логика для отправки уведомления по SMS,
	// у структуры Sms могут присутствовать подобные методы (GetPhoneByUserId, SmsServiceConnection)
	fmt.Printf("[sms] %s, ваш заказ:%s готов к выдачи\n", user.Name, orderId)
}
