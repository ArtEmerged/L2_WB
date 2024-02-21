package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Комманда — это паттерн поведенческого проектирования, в котором объект используется для инкапсуляции всей информации,
	необходимой для выполнения действий или запуска события в более позднее время.

	Шаблон "команды" полезен:
		1. Когда вам нужно выполнить задачи, но вы хотите отделить управление задачами от выполнения самой задачи.
		2. Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
		3. Когда вам нужна операция отмены.

	+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	+ Позволяет реализовать простую отмену и повторение операций.
	+ Позволяетреализовать отложенный запуск операций.
	+ Позволяет собрать сложные команды из простых.

	- Усложняет кот программы из-за введения множества дополнительных класснов.
*/

// UICommand - командный интерфейс
type UICommand interface {
	Print()
	Remove()
}

// Button - является конкретной командой, структура реализующая интерфейс Command
type Button struct {
	name string
}

// NewButton - конструктор для новой команды
func NewButton(name string) *Button {
	return &Button{name}
}

// Print - метод для вывода названия команды
func (button *Button) Print() {
	fmt.Printf("Printing %s Button\n", button.name)
}

// Remove - метод для вывода сообщения об удалении команды
func (button *Button) Remove() {
	fmt.Printf("Removing %s Button\n", button.name)
}

// Input - является конкретной командой, структура реализующая интерфейс Command
type Input struct{}

// NewInput - конструктор для новой команды
func NewInput() *Input {
	return &Input{}
}

// Print - метод для вывода названия команды
func (input *Input) Print() {
	fmt.Println("Printing Input")
}

// Remove - метод для вывода сообщения об удалении команды
func (input *Input) Remove() {
	fmt.Println("Removing Input")
}

// Table - является конкретной командой, структура реализующая интерфейс Command
type Table struct {
}

// NewTable - конструктор для новой команды
func NewTable() *Table {
	return &Table{}
}

// Print - метод для вывода названия команды
func (table *Table) Print() {
	fmt.Println("Printing Table")
}

// Remove - метод для вывода сообщения об удалении команды
func (table *Table) Remove() {
	fmt.Println("Removing Table")
}

// UIControl - структура выполнителя (invoker). Invoker ничего не знает о конкретной команде, он знает только об интерфейсе.
type UIControl struct {
	commandList []UICommand
}

// NewControl - конструктор выполнителя
func NewControl() *UIControl {
	return &UIControl{}
}

// AddElement - добавляет команду в очередь
func (uiControl *UIControl) AddElement(uiCommand UICommand) {
	uiCommand.Print()
	uiControl.commandList = append(uiControl.commandList, uiCommand)
}

// RemoveElement - удаляет команду из очереди
func (uiControl *UIControl) RemoveElement(uiCommand UICommand) {
	uiCommand.Remove()

	newList := []UICommand{}

	for _, elem := range uiControl.commandList {
		if elem != uiCommand {
			newList = append(newList, elem)
		}
	}

	uiControl.commandList = newList
}

// Undo - отменяет последнюю команду
func (uiControl *UIControl) Undo() {
	uiCommand := uiControl.commandList[len(uiControl.commandList)-1]
	uiControl.RemoveElement(uiCommand)
}

// func main() {
// 	uiControl := NewUIControl()
// 	inputUI := NewInputUI()
// 	tableUI := NewTableUI()
// 	buttonUI := NewButtonUI("Submit")

// 	uiControl.AddElement(inputUI)
// 	uiControl.AddElement(tableUI)
// 	uiControl.AddElement(buttonUI)

// 	uiControl.RemoveElement(tableUI)

// 	uiControl.AddElement(NewButtonUi("Cancel"))
// 	uiControl.AddElement(NewTableUI())
// 	uiControl.AddElement(NewInputUI())
// 	uiControl.AddElement(NewButtonUI("Wrong button"))

// 	uiControl.Undo()
// 	uiControl.Undo()
// 	fmt.Println(uiControl.commandList)
// }
