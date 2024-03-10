package pattern

/*
Декоратор  - это структурный паттерн проектирования, который позволяет динамичеки добавлять объектам новый функционал, оборачивая их в полезные “обёртки”

Декоратор помещаете целевой объект в другой объект-обёртку(декоратор), который запускает базовое поведение объекта, 
а затем добавляет к результату что-то своё.

Оба объекта имеют общий интерфейс, поэтому для пользователя нет никакой разницы, с каким объектом работать — чистым или обёрнутым.
Вы можете использовать несколько разных обёрток одновременно — результат будет иметь объединённое поведение всех обёрток сразу.

Шаблон "декоратор" полезен:
	1. Когда вам нужно добавлять обязанности объектам на лету, незаметно для кода, который их использует.
	2. Когда нельзя расширить обязанности объекта с помощью наследования.

+ Больше гибкости чем у наследования
+ Позволяет добавлять обязанности на лету
+ Можно добавлять несколько новых обязанностей сразу
+ Позволяет иметь несколько мелких объектов вместо одного объекта на все случаи жизни

- Трудно конфигурировать многократно обёрнутые объекты
- Обилие крошечных классов
*/

// pizza.go: Интерфейс компонента
type IPizza interface {
	getPrice() int
}

// veggieMania.go: Конкретный компонент
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// tomatoTopping.go: Конкретный декоратор
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// cheeseTopping.go: Конкретный декоратор
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// main.go: Клиентский код
/*
func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
*/
