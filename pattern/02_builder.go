package pattern

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Строитель — это порождающий паттерн проектирования, который позволяет создавать объекты пошагово.

	Целью шаблона проектирования Builder является отделение конструкции сложного объекта от его представления. 
	Таким образом, один и тот же процесс построения может создавать разные представления.
	
	+ Позволяет изменять внутреннее представление продукта.
	+ Изолирует код, реализующий конструирование и представление.
	+ Дает более тонкий контроль над процессом конструирования.

	- Для каждого типа продукта необходимо создать отдельный ConcreteBuilder.
	- Классы Builder должны быть изменяемыми.
	- Может затруднить/усложнить внедрение зависимостей.
*/

const (
	// BrendCopmuterHP - бренд компьютера HP
	BrendCopmuterHP = "hp"
	// BrendCopmuterAsus - бренд компьютера Asus
	BrendCopmuterAsus = "asus"
)

// IBuilder определяет интерфейс для создания объекта и его сборки
type IBuilder interface {
	SetCore()
	SetRAM()
	SetBrend()
	SetGraphicCard()
	SetMonitor()
	GetCopmuter() Computer
}

// Computer - объект-компьютер
type Computer struct {
	Core        int
	RAM         int
	Brend       string
	GraphicCard int
	Monitor     int
}

// NewBuilder - конструктор строителя в зависимости от бренда компьютера
func NewBuilder(brand string) IBuilder {
	switch brand {
	case BrendCopmuterHP:
		return &HpBuilder{}
	case BrendCopmuterAsus:
		return &AsusBuilder{}
	default:
		return nil
	}
}

// AsusBuilder - строитель для компьютеров Asus
type AsusBuilder struct {
	Core        int
	RAM         int
	Brend       string
	GraphicCard int
	Monitor     int
}

// HpBuilder - строитель для компьютеров HP
type HpBuilder struct {
	Core        int
	RAM         int
	Brend       string
	GraphicCard int
	Monitor     int
}

// SetCore устанавливает количество ядер процессора
func (ac *AsusBuilder) SetCore() { ac.Core = 4 }

// SetRAM устанавливает объем оперативной памяти
func (ac *AsusBuilder) SetRAM() { ac.RAM = 16 }

// SetBrend устанавливает бренд компьютера
func (ac *AsusBuilder) SetBrend() { ac.Brend = "Asus" }

// SetGraphicCard устанавливает видеокарту
func (ac *AsusBuilder) SetGraphicCard() { ac.GraphicCard = 1 }

// SetMonitor устанавливает количество мониторов
func (ac *AsusBuilder) SetMonitor() { ac.Monitor = 1 }

// GetCopmuter возвращает собранный компьютер
func (ac *AsusBuilder) GetCopmuter() Computer {
	return Computer{
		Core:        ac.Core,
		RAM:         ac.RAM,
		Brend:       ac.Brend,
		GraphicCard: ac.GraphicCard,
		Monitor:     ac.Monitor,
	}
}

// SetCore устанавливает количество ядер процессора
func (hc *HpBuilder) SetCore() { hc.Core = 1 }

// SetRAM устанавливает объем оперативной памяти
func (hc *HpBuilder) SetRAM() { hc.RAM = 8 }

// SetBrend устанавливает бренд компьютера
func (hc *HpBuilder) SetBrend() { hc.Brend = "HP" }

// SetGraphicCard устанавливает видеокарту
func (hc *HpBuilder) SetGraphicCard() { hc.GraphicCard = 2 }

// SetMonitor устанавливает количество мониторов
func (hc *HpBuilder) SetMonitor() { hc.Monitor = 2 }

// GetCopmuter возвращает собранный компьютер
func (hc *HpBuilder) GetCopmuter() Computer {
	return Computer{
		Core:        hc.Core,
		RAM:         hc.RAM,
		Brend:       hc.Brend,
		GraphicCard: hc.GraphicCard,
		Monitor:     hc.Monitor,
	}
}

// Director определяет интерфейс для сборки компьютера
type Director struct {
	builder IBuilder
}

// NewDirector - конструктор для Director
func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}

// SetCollecor менят строителя
func (d *Director) SetCollecor(builder IBuilder) {
	d.builder = builder
}

// CreateComputer собирает компьютер
func (d *Director) CreateComputer() Computer {
	d.builder.SetCore()
	d.builder.SetRAM()
	d.builder.SetBrend()
	d.builder.SetGraphicCard()
	d.builder.SetMonitor()
	return d.builder.GetCopmuter()
}

// func main() {
// 	hpBuilder := NewBuilder("hp")
// 	asusBuilder := NewBuilder("asus")

// 	// Создание директора с указанием конкретного строителя для HP компьютера
// 	direcor := NewDirector(hpBuilder)

// 	// Создание HP компьютера через директора
// 	computerHP := direcor.CreateComputer()
// 	fmt.Println(computerHP)

// 	// Изменение строителя на AsusBuilder
// 	direcor.SetCollecor(asusBuilder)

// 	// Создание компьютера Asus через директора с новым строителем
// 	computerAsus := direcor.CreateComputer()
// 	fmt.Println(computerAsus)

// }
