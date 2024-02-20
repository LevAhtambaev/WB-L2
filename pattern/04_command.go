package pattern

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Команда — это шаблон поведенческого проектирования, который превращает запрос в автономный объект, содержащий всю информацию о запросе.
Это преобразование позволяет передавать запросы в качестве аргументов метода, задерживать или ставить в очередь выполнение запроса, а также поддерживать отменяемые операции.
Преобразование позволяет отложенное или удаленное выполнение команд, сохранение истории команд и т. д.
*/

// Рассмотрим пример с системой управления заказами в ресторане.
// Мы можем использовать паттерн Команда для реализации команд - создание заказа, добавление блюда к заказу, удаление блюда из заказа.

import "fmt"

// OrderReceiver представляет получателя команды.
type OrderReceiver struct{}

// CreateOrder выполняет команду создания заказа.
func (r *OrderReceiver) CreateOrder(orderID int) {
	fmt.Printf("Создание заказа с ID %d\n", orderID)
}

// MenuItem представляет блюдо в меню ресторана.
type MenuItem struct {
	Name  string
	Price float64
}

// Command определяет интерфейс команды.
type Command interface {
	Execute()
}

// CreateOrderCommand представляет команду создания заказа.
type CreateOrderCommand struct {
	orderReceiver *OrderReceiver
	orderID       int
}

func NewCreateOrderCommand(receiver *OrderReceiver, orderID int) *CreateOrderCommand {
	return &CreateOrderCommand{
		orderReceiver: receiver,
		orderID:       orderID,
	}
}

func (cmd *CreateOrderCommand) Execute() {
	cmd.orderReceiver.CreateOrder(cmd.orderID)
}

// AddItemToOrderCommand представляет команду добавления блюда к заказу.
type AddItemToOrderCommand struct {
	orderReceiver *OrderReceiver
	orderID       int
	item          MenuItem
}

func NewAddItemToOrderCommand(receiver *OrderReceiver, orderID int, item MenuItem) *AddItemToOrderCommand {
	return &AddItemToOrderCommand{
		orderReceiver: receiver,
		orderID:       orderID,
		item:          item,
	}
}

func (cmd *AddItemToOrderCommand) Execute() {
	fmt.Printf("Добавление блюда %s к заказу с ID %d\n", cmd.item.Name, cmd.orderID)
}

// Invoker представляет инициатор команды - управляет выполнением команды.
type Invoker struct {
	command Command
}

func (inv *Invoker) SetCommand(command Command) {
	inv.command = command
}

func (inv *Invoker) ExecuteCommand() {
	inv.command.Execute()
}

func main() {
	orderReceiver := &OrderReceiver{}

	createOrderCommand := NewCreateOrderCommand(orderReceiver, 1001)
	addItemToOrderCommand := NewAddItemToOrderCommand(orderReceiver, 1001, MenuItem{Name: "Блюдо", Price: 5})

	invoker := &Invoker{}

	// Выполнение команды создания заказа
	invoker.SetCommand(createOrderCommand)
	invoker.ExecuteCommand()

	// Выполнение команды добавления блюда к заказу
	invoker.SetCommand(addItemToOrderCommand)
	invoker.ExecuteCommand()
}
