package pattern

import "fmt"

/*
	Абстрактная фабрика - это порождающий паттерн проектирования, который позволяет создавать семейства связанных объектов,
	не привязываясь к конкретным классам создаваемых объектов.

	Клиентский код должен работать как с фабриками, так и с продуктами только через их общие интерфейсы.
	Это позволит подавать в ваши классы любой тип фабрики и производить любые продукты, ничего не ломая.

	В реализации есть 5 элементов:
		1. Интерфейс абстрактной фабрики : интерфейс, реализуемый конкретными абстрактными классами фабрики.
		2. Конкретные абстрактные фабричные классы : классы, которые возвращают экземпляры конкретного класса.
		3. Общий интерфейс : интерфейс, реализованный конкретными классами.
		4. Конкретные классы : классы, реализующие фабричный интерфейс и содержащие фактические операции.
		Получение этих классов является конечной целью этой реализации.
		5. Factory Producer : возвращает экземпляры конкретной фабрики на основе некоторого заранее определенного условия.
		Factory Producer напрямую использует только абстрактный интерфейс фабрики и конкретные абстрактные классы фабрики.

	Шаблон "Абстрактная фабрика" полезен:
		1. Когда бизнес-логика программы должна работать с разными видами связанных друг с другом продуктов, не завися от конкретных классов продуктов.
		2. Когда в программе уже используется Фабричный метод, но очередные изменения предполагают введение новых типов продуктов.

	+ Гарантирует сочетаемость создаваемых продуктов.
	+ Избавляет клиентский код от привязки к конкретным классам продуктов.
	+ Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	+ Упрощает добавление новых продуктов в программу.
	+ Реализует принцип открытости/закрытости.

	- Усложняет код программы из-за введения множества дополнительных классов.
	- Требует наличия всех типов продуктов в каждой вариации.
*/

type Brand string

const (
	AdidasB Brand = "adidas"
	NikeB   Brand = "nike"
)

// iSportsFactory.go: Интерфейс абстрактной фабрики
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand Brand) (ISportsFactory, error) {
	switch brand {
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	case AdidasB:
		return &Adidas{}, nil
	case NikeB:
		return &Nike{}, nil

	}
}

// adidas.go: Конкретная фабрика
type Adidas struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "Adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "Adidas",
			size: 14,
		},
	}
}

// nike.go: Конкретная фабрика
type Nike struct{}

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 14,
		},
	}
}

func (a *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 14,
		},
	}
}

// iShoe.go: Абстрактный продукт
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) { s.logo = logo }

func (s *Shoe) getLogo() string { return s.logo }

func (s *Shoe) setSize(size int) { s.size = size }

func (s *Shoe) getSize() int { return s.size }

// adidasShoe.go: Конкретный продукт
type AdidasShoe struct {
	Shoe
}

// nikeShoe.go: Конкретный продукт
type NikeShoe struct {
	Shoe
}

// iShirt.go: Абстрактный продукт
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) { s.logo = logo }
func (s *Shirt) setSize(size int)    { s.size = size }
func (s *Shirt) getLogo() string     { return s.logo }
func (s *Shirt) getSize() int        { return s.size }

// adidasShirt.go: Конкретный продукт
type AdidasShirt struct {
	Shirt
}

// nikeShirt.go: Конкретный продукт
type NikeShirt struct {
	Shirt
}

// main.go: Клиентский код

// func main() {
// 	adidasFactory, _ := GetSportsFactory("adidas")
// 	nikeFactory, _ := GetSportsFactory("nike")

// 	nikeShoe := nikeFactory.makeShoe()
// 	nikeShirt := nikeFactory.makeShirt()

// 	adidasShoe := adidasFactory.makeShoe()
// 	adidasShirt := adidasFactory.makeShirt()

// 	printShoeDetails(nikeShoe)
// 	printShirtDetails(nikeShirt)

// 	printShoeDetails(adidasShoe)
// 	printShirtDetails(adidasShirt)
// }

// func printShoeDetails(s IShoe) {
// 	fmt.Printf("Logo: %s", s.getLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.getSize())
// 	fmt.Println()
// }

// func printShirtDetails(s IShirt) {
// 	fmt.Printf("Logo: %s", s.getLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.getSize())
// 	fmt.Println()
// }
