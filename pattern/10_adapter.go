package pattern

// Интерфейс, к которому мы адаптируем нашу структуру.
type Target interface {
	Process()
}

// Несовместимуая структура
type Adaptee struct {
}

// Метод, который необходимо адаптировать
func (a *Adaptee) AdapteeProcess() {
}

// Структура Адаптер
type Adapter struct {
	*Adaptee
}

// Конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// Вызывает адаптированный метод AdapteeProcess
func (a *Adapter) Process() {
	a.AdapteeProcess()
}

func main() {

	// Создаем адаптер
	adapter := NewAdapter(&Adaptee{})

	adapter.Process()

}
