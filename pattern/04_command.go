package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Комманда - это паттерн поведенческого проектирования, в котором объект используется для инкапсуляции всей информации,
	необходимой для выполнения действий или запуска события в более позднее время.

	Шаблон команды полезен:
		1. Когда вам нужно выполнить задачи, но вы хотите отделить управление задачами от выполнения самой задачи.
		2. Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
		3. Когда вам нужна операция отмены.

	+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	+ Позволяет реализовать простую отмену и повторение операций.
	+ Позволяетреализовать отложенный запуск операций.
	+ Позволяет собрать сложные команды из простых.

	- Усложняет кот программы из-за введения множества дополнительных класснов.
*/

// UiCommand - командный интерфейс
type UiCommand interface {
	Print()
	Remove()
}

// ButtonUi - является конкретной командой, структура реализующая интерфейс UiCommand
type ButtonUi struct {
	name string
}

// NewButtonUi - конструктор для новой команды
func NewButtonUi(name string) *ButtonUi {
	return &ButtonUi{name}
}

// Print - метод для вывода названия команды
func (buttonUi *ButtonUi) Print() {
	fmt.Printf("Printing %s Button\n", buttonUi.name)
}

// Remove - метод для вывода сообщения об удалении команды
func (buttonUi *ButtonUi) Remove() {
	fmt.Printf("Removing %s Button\n", buttonUi.name)
}

// InputUi - является конкретной командой, структура реализующая интерфейс UiCommand
type InputUi struct{}

// NewInputUi - конструктор для новой команды
func NewInputUi() *InputUi {
	return &InputUi{}
}

// Print - метод для вывода названия команды
func (inputUi *InputUi) Print() {
	fmt.Println("Printing Input")
}

// Remove - метод для вывода сообщения об удалении команды
func (inputUi *InputUi) Remove() {
	fmt.Println("Removing Input")
}

// TableUi - является конкретной командой, структура реализующая интерфейс UiCommand
type TableUi struct {
}

// NewTableUi - конструктор для новой команды
func NewTableUi() *TableUi {
	return &TableUi{}
}

// Print - метод для вывода названия команды
func (tableUi *TableUi) Print() {
	fmt.Println("Printing Table")
}

// Remove - метод для вывода сообщения об удалении команды
func (tableUi *TableUi) Remove() {
	fmt.Println("Removing Table")
}

// UiControl - структура выполнителя (invoker). Invoker ничего не знает о конкретной команде, он знает только об интерфейсе.
type UiControl struct {
	commandList []UiCommand
}

// NewUiControl - конструктор выполнителя
func NewUiControl() *UiControl {
	return &UiControl{}
}

// AddElement - добавляет команду в очередь
func (uiControl *UiControl) AddElement(uiCommand UiCommand) {
	uiCommand.Print()
	uiControl.commandList = append(uiControl.commandList, uiCommand)
}

// RemoveElement - удаляет команду из очереди
func (uiControl *UiControl) RemoveElement(uiCommand UiCommand) {
	uiCommand.Remove()

	newList := []UiCommand{}

	for _, elem := range uiControl.commandList {
		if elem != uiCommand {
			newList = append(newList, elem)
		}
	}

	uiControl.commandList = newList
}

// Undo - отменяет последнюю команду
func (uiControl *UiControl) Undo() {
	uiCommand := uiControl.commandList[len(uiControl.commandList)-1]
	uiControl.RemoveElement(uiCommand)
}

// func main() {
// 	uiControl := NewUiControl()
// 	inputUi := NewInputUi()
// 	tableUi := NewTableUi()
// 	buttonUi := NewButtonUi("Submit")

// 	uiControl.AddElement(inputUi)
// 	uiControl.AddElement(tableUi)
// 	uiControl.AddElement(buttonUi)

// 	uiControl.RemoveElement(tableUi)

// 	uiControl.AddElement(NewButtonUi("Cancel"))
// 	uiControl.AddElement(NewTableUi())
// 	uiControl.AddElement(NewInputUi())
// 	uiControl.AddElement(NewButtonUi("Wrong button"))

// 	uiControl.Undo()
// 	uiControl.Undo()
// 	fmt.Println(uiControl.commandList)
// }
