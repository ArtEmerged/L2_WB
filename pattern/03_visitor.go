package pattern

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет создавать новые операции,
	не меняя классы объектов на которыми могут выполняться.

	Применяется когда над объектами сложной структуры объектов нужно выполнить некоторые, не связанные между собой операции,
	но вы не хотите "засорять" классы такими операциями.

	+ Новая функциональность в несколько классов добавляется сразу, не изменяя код этих классов.
	+ Объединяет родственные операции в одном классе.

	- Паттерн не оправдан, если иерархия компонентов часто меняется.
	- Может привести к нарушению инкапсуляции компонентов.
*/

// Service - интерфейс, представляющий услугу
type Service interface {
	Accept(hostingCalculatorVisitor HostingCalculatorVisitor) float64
}

// ComputeService - структура, представляющая вычислительную услугу
type ComputeService struct {
	price    float64
	quantity int
}

// NewComputeService - конструктор вычислительной услуги
func NewComputeService(quantity int) *ComputeService {
	return &ComputeService{
		price:    10.50,
		quantity: quantity,
	}
}

// GetPrice - полчает цены вычислительной услуги
func (compute *ComputeService) GetPrice() float64 {
	return compute.price
}

// GetQuantity - полчает количества вычислительных услуг
func (compute *ComputeService) GetQuantity() int {
	return compute.quantity
}

// Accept - принимает посетителя вычислительной услуги
func (compute *ComputeService) Accept(hostingCalculatorVisitor HostingCalculatorVisitor) float64 {
	return hostingCalculatorVisitor.ComputeVisit(compute)
}

// DatabaseService - структура, представляющая услугу базы данных
type DatabaseService struct {
	price         float64
	backPrice     float64
	quantity      int
	backupEnabled bool
}

// NewDatabaseService1 - конструктор базы данных без резервного копирования
func NewDatabaseService1(quantity int) *DatabaseService {
	return &DatabaseService{
		price:         100.00,
		backPrice:     30.00,
		backupEnabled: false,
		quantity:      quantity,
	}
}

// NewDatabaseService2 - конструктор базы данных с возможностью резервного копирования
func NewDatabaseService2(quantity int, backupEnabled bool) *DatabaseService {
	return &DatabaseService{
		price:         100.00,
		backPrice:     30.00,
		backupEnabled: backupEnabled,
		quantity:      quantity,
	}
}

// GetBackPrice - получет цены резервного копирования базы данных
func (databaseService *DatabaseService) GetBackPrice() float64 {
	return databaseService.backPrice
}

// GetPrice - получет цены услуги базы данных
func (databaseService *DatabaseService) GetPrice() float64 {
	return databaseService.price
}

// GetQuantity - получает количества услуг базы данных.
func (databaseService *DatabaseService) GetQuantity() int {
	return databaseService.quantity
}

// IsBackupEnabled - проверяет возможности резервного копирования базы данных
func (databaseService *DatabaseService) IsBackupEnabled() bool {
	return databaseService.backupEnabled
}

// Accept - принимает посетителя услуги базы данных
func (databaseService *DatabaseService) Accept(hostingCalculatorVisitor HostingCalculatorVisitor) float64 {
	return hostingCalculatorVisitor.DatabaseVisit(databaseService)
}

// HostingCalculatorVisitor - интерфейс посетителя
type HostingCalculatorVisitor interface {
	ComputeVisit(computeService *ComputeService) float64
	DatabaseVisit(databaseService *DatabaseService) float64
}

// HostingCalculatorVisitorImpl - реализация посетителя
type HostingCalculatorVisitorImpl struct{}

// NewHostingCalculatorVisitorImpl - конструктор посетителя
func NewHostingCalculatorVisitorImpl() *HostingCalculatorVisitorImpl {
	return &HostingCalculatorVisitorImpl{}
}

// ComputeVisit - метод для посещения вычислительной услуги
func (hcvi *HostingCalculatorVisitorImpl) ComputeVisit(computeService *ComputeService) float64 {
	return computeService.GetPrice() * float64(computeService.GetQuantity())
}

// DatabaseVisit - метод для посещения услуги базы данных
func (hcvi *HostingCalculatorVisitorImpl) DatabaseVisit(databaseService *DatabaseService) float64 {
	serviceCost := databaseService.GetPrice() * float64(databaseService.GetQuantity())

	backupCost := 0.00
	if databaseService.IsBackupEnabled() {
		backupCost = databaseService.GetBackPrice() * float64(databaseService.GetQuantity())
	}
	return serviceCost + backupCost
}

// func main() {
// 	usedServices := []Service{
// 		NewComputeService(3),
// 		NewDatabaseService1(3),
// 		NewDatabaseService2(2, true),
// 	}
// 	totalCost := calculateHostingCost(usedServices)

// 	fmt.Printf("Total cost of hosting is: %f\n", totalCost)
// }

// // calculateHostingCost - вычисляет общей стоимости хостинга
// func calculateHostingCost(services []Service) float64 {
// 	hostingCalculatorVisitorImpl := NewHostingCalculatorVisitorImpl()
// 	totalCost := 0.00

// 	for _, service := range services {
// 		totalCost += service.Accept(hostingCalculatorVisitorImpl)
// 	}
// 	return totalCost
// }
