package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.

	Шаблон "фабричный метод" полезен:
		1. Когда мы заранее не знаем тип объекта, который будет сгенерирован.
		2. Когда мы хотим сгенерировать объект на лету на основе некоторых критериев.
		3. Когда мы хотим дать возможность пользователям ресширять асти вашего фреймворка или библиотеки.

	+ Избавляет класс от привязки к конкретным классам продуктов.
	+ Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	+ Упрощает добавление новых продуктов в программу.
	+ Реализует принцип открытость/закрытость.

	- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
	- Фабрика рискует стать "супер-объектом", привязанным ко всем классам программы.
*/

// Order - интерфейс продукта
type Order interface {
	GetDetails() string
}

// BookOrder - заказ на книгу
type BookOrder struct {
	Title  string
	Author string
	Price  float64
}

// GetDetails - возвращает информацию о заказ на книгу
func (bo *BookOrder) GetDetails() string {
	return fmt.Sprintf("Book Order: Title - %s, Author - %s, Price - $%.2f", bo.Title, bo.Author, bo.Price)
}

// ElectronicOrder - заказ на электронное устройство
type ElectronicOrder struct {
	Product string
	Model   string
	Price   float64
}

// GetDetails - возвращает информацию о заказ на электронное устройство
func (eo *ElectronicOrder) GetDetails() string {
	return fmt.Sprintf("Electronic Order: Product - %s, Model - %s, Price - $%.2f", eo.Product, eo.Model, eo.Price)
}

// OrderFactory - интерфейс фабрики
type OrderFactory interface {
	CreateOrder() Order
}

// BookOrderFactory реализует фабрику для создания заказа на книгу
type BookOrderFactory struct{}

// CreateOrder - создает заказ на книгу
func (bof *BookOrderFactory) CreateOrder() Order {
	return &BookOrder{
		Title:  "Go идиомы и паттерны проектирования",
		Author: "Джон Боднер",
		Price:  10.99,
	}
}

// BookOrderFactory - конструктор фаблрики заказов на книги
func NewBookOrderFactory() *BookOrderFactory {
	return &BookOrderFactory{}
}

// ElectronicOrderFactory реализует фабрику для создания заказа на электронное устройство
type ElectronicOrderFactory struct{}

// CreateOrder - создает заказ на электронное устройство
func (eof *ElectronicOrderFactory) CreateOrder() Order {
	return &ElectronicOrder{
		Product: "Smartphone",
		Model:   "iPhone 15 ProMax",
		Price:   1299.99,
	}
}

// NewElectronicOrderFactory - конструктор фаблрики заказаов на электронные устройства
func NewElectronicOrderFactory() *ElectronicOrderFactory {
	return &ElectronicOrderFactory{}
}

// func main() {
// 	bookOrderFactory := NewBookOrderFactory()
// 	bookOrder := bookOrderFactory.CreateOrder()
// 	fmt.Println(bookOrder.GetDetails())

// 	electronicOrderFactory := NewElectronicOrderFactory()
// 	electronicOrder := electronicOrderFactory.CreateOrder()
// 	fmt.Println(electronicOrder.GetDetails())
// }
