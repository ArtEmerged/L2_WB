package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	Шаблон "стратегия" полезен:
		1. Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
		2. Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
		3. Когда вы не хотите обнажать детали реализации алгоритмов для других классов.

	+ Горячая замена алгоритмов на лету.
	+ Изолирует код и данные алгоритмов от остальных классов.
	+ Уход от наследования к делегированию.
	+ Реализует принцип открытость/закрытость.

	- Усложняет программу за счёт дополнительных классов.
	- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

// Payment определяет интерфейс для всех способов оплаты
type Payment interface {
	Pay()
}

// cardPayment представляет оплату кредитной картой
type cardPayment struct{}

// NewCardPayment - конструктор для cardPayment
func NewCardPayment() Payment {
	return &cardPayment{}
}

// Pay - алгоритмом оплаты через карту
func (p *cardPayment) Pay() {
	fmt.Println("Card Payment")
}

// payPalPayment представляет оплату через PayPal
type payPalPayment struct{}

// NewPayPalPayment - конструктор для payPalPayment
func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

// Pay алгоритмом оплаты через PayPal
func (p *payPalPayment) Pay() {
	fmt.Println("PayPal Payment")
}

// sbpPayment представляет оплату через СБП
type sbpPayment struct{}

// NewSBPPayment - конструктор для sbpPayment
func NewSBPPayment() Payment {
	return &sbpPayment{}
}

// Pay алгоритмом оплаты через СБП
func (p *sbpPayment) Pay() {
	fmt.Println("СБП Payment")
}

// OrderW - заказ, который будет обрабатываться с помощью определенного способа оплаты
type OrderW struct {
	payment Payment
}

// setPayment устанавливает способ оплаты для заказа
func (order *OrderW) setPayment(payment Payment) {
	order.payment = payment
}

// processOrder обрабатывает заказ с помощью установленного способа оплаты
func (order *OrderW) processOrder() {
	order.payment.Pay()
}

// func main() {
// 	order := Order{}

// 	// Устанавливаем способ оплаты через карту
// 	order.setPayment(NewCardPayment())
// 	order.processOrder()

// 	// Устанавливаем способ оплаты через СБП
// 	order.setPayment(NewSBPPayment())
// 	order.processOrder()

// 	// Устанавливаем способ оплаты через PayPal
// 	order.setPayment(NewPayPalPayment())
// 	order.processOrder()
// }
