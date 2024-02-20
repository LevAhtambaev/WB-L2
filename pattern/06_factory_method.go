package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод — это творческий шаблон проектирования, который предоставляет интерфейс для создания объектов в суперклассе, но позволяет подклассам изменять тип создаваемых объектов.
Фабричный метод определяет метод, который следует использовать для создания объектов вместо прямого вызова конструктора (новый оператор).
Подклассы могут переопределить этот метод, чтобы изменить класс создаваемых объектов.
*/

// Рассмотрим пример - фабрика мебели, которая производит различные виды мебели, такие как стулья и столы
// В этом примере интерфейс Furniture определяет методы, общие для всех видов мебели.
// Каждый конкретный тип мебели (Chair, Table) реализует этот интерфейс и предоставляет свою уникальную реализацию методов.
// Фабрики ChairFactory и TableFactory реализуют интерфейс FurnitureFactory и создают конкретные виды мебели соответственно.

import "fmt"

// FurnitureFactory - фабрика мебели
type FurnitureFactory interface {
	Create() Furniture
}

// Furniture - интерфейс для мебели
type Furniture interface {
	GetType() string
	Assemble()
}

// Chair - стул
type Chair struct{}

func (c *Chair) GetType() string {
	return "Стул"
}

func (c *Chair) Assemble() {
	fmt.Println("Сборка стула")
}

// Table - стол
type Table struct{}

func (t *Table) GetType() string {
	return "Стол"
}

func (t *Table) Assemble() {
	fmt.Println("Сборка стола")
}

// ChairFactory - фабрика стульев
type ChairFactory struct{}

func (f *ChairFactory) Create() Furniture {
	return &Chair{}
}

// TableFactory - фабрика столов
type TableFactory struct{}

func (f *TableFactory) Create() Furniture {
	return &Table{}
}

func main() {
	// Создание фабрик
	chairFactory := &ChairFactory{}
	tableFactory := &TableFactory{}

	// Создание мебели с использованием фабрик
	chair := chairFactory.Create()
	table := tableFactory.Create()

	// Сборка мебели
	chair.Assemble()
	table.Assemble()
}
