package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение
	в зависимости от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

	Шаблон "состояние" полезен:
		1. Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
		причём типов состояний много, и их код часто меняется.
		2. Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
		которые выбирают поведения в зависимости от текущих значений полей класса.

	+ Избавляет от множества больших условных операторов машины состояний.
	+ Концентрирует в одном месте код, связанный с определённым состоянием.
	+ Упрощает код контекста.

	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

// OrderState - состояние заказа
type OrderState interface {
	ProcessOrder(order *Order)
}

// Order - заказ в интернет-магазине
type Order struct {
	state OrderState
}

// NewOrder конструктор для заказа
func NewOrder() *Order {
	return &Order{state: &NewOrderState{}}
}

// Process заказа
func (o *Order) Process() {
	o.state.ProcessOrder(o)
}

// SetState устанавливает новое состояние заказа
func (o *Order) SetState(state OrderState) {
	o.state = state
}

// NewOrderState - состояние нового заказа
type NewOrderState struct{}

// ProcessOrder обрабатывает заказ в состоянии нового заказа
func (nos *NewOrderState) ProcessOrder(order *Order) {
	fmt.Println("Поступил новый заказ. Обработка заказа...")
	time.Sleep(2 * time.Second)
	fmt.Println("Заказ обработан.")
	order.SetState(&ProcessingOrderState{})
}

// ProcessingOrderState - состояние обработки заказа
type ProcessingOrderState struct{}

// ProcessOrder обрабатывает заказ в состоянии обработки заказа
func (pos *ProcessingOrderState) ProcessOrder(order *Order) {
	fmt.Println("Заказ находится в состоянии обработки. Готовится к отправке...")
	time.Sleep(3 * time.Second)
	fmt.Println("Заказ готов к отправке.")
	order.SetState(&ReadyForShipmentState{})
}

// ReadyForShipmentState - состояние готовности к отправке заказа
type ReadyForShipmentState struct{}

// ProcessOrder обрабатывает заказ в состоянии готовности к отправке заказа
func (rfss *ReadyForShipmentState) ProcessOrder(order *Order) {
	fmt.Println("Заказ находится в состоянии готовности к отправке. Отправляем на доставку...")
	time.Sleep(1 * time.Second)
	fmt.Println("Заказ отправлен.")
	order.SetState(&DeliveredState{})
}

// DeliveredState - состояние доставленного заказа
type DeliveredState struct{}

// ProcessOrder обрабатывает заказ в состоянии доставленного заказа
func (ds *DeliveredState) ProcessOrder(order *Order) {
	fmt.Println(`Заказ находится в состоянии "Доставлен". Заказ успешно доставлен.`)
}

// func main() {
// 	order := NewOrder()
// 	order.Process()
// 	order.Process()
// 	order.Process()
// 	order.Process()
// }
