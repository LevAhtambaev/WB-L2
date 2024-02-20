package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Builder — это творческий шаблон проектирования, который позволяет шаг за шагом создавать сложные объекты.
// Шаблон позволяет создавать разные типы и представления объекта, используя один и тот же код построения.

// В отличие от других творческих шаблонов, Builder не требует, чтобы продукты имели общий интерфейс.
// Это позволяет производить различные продукты, используя один и тот же процесс строительства.

// Рассмотрим пример использования паттерна Builder для создания автомобилей с различными характеристиками.

import "fmt"

// Car представляет автомобиль.
type Car struct {
	model      string
	numDoors   int
	engineType string
	extras     []string
}

// CarBuilder - интерфейс для создания автомобиля.
type CarBuilder interface {
	SetModel()
	SetNumDoors()
	SetEngineType()
	GetCar() *Car
}

// SedanCarBuilder создает седан.
type SedanCarBuilder struct {
	car *Car
}

func NewSedanCarBuilder() *SedanCarBuilder {
	return &SedanCarBuilder{
		car: &Car{},
	}
}

func (b *SedanCarBuilder) SetModel() {
	b.car.model = "Седан"
}

func (b *SedanCarBuilder) SetNumDoors() {
	b.car.numDoors = 4
}

func (b *SedanCarBuilder) SetEngineType() {
	b.car.engineType = "V4"
}

func (b *SedanCarBuilder) GetCar() *Car {
	return b.car
}

// SportsCarBuilder создает спортивный автомобиль.
type SportsCarBuilder struct {
	car *Car
}

func NewSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{
		car: &Car{},
	}
}

func (b *SportsCarBuilder) SetModel() {
	b.car.model = "Спортивный автомобиль"
}

func (b *SportsCarBuilder) SetNumDoors() {
	b.car.numDoors = 2
}

func (b *SportsCarBuilder) SetEngineType() {
	b.car.engineType = "V8"
}

func (b *SportsCarBuilder) GetCar() *Car {
	return b.car
}

// CarDirector cоздает автомобиль с использованием строителя.
type CarDirector struct {
	builder CarBuilder
}

func NewCarDirector(builder CarBuilder) *CarDirector {
	return &CarDirector{
		builder: builder,
	}
}

func (d *CarDirector) ConstructCar() *Car {
	d.builder.SetModel()
	d.builder.SetNumDoors()
	d.builder.SetEngineType()
	return d.builder.GetCar()
}

func main() {
	sedanBuilder := NewSedanCarBuilder()
	sedanDirector := NewCarDirector(sedanBuilder)
	sedan := sedanDirector.ConstructCar()

	fmt.Println("Седан:")
	fmt.Println("Модель:", sedan.model)
	fmt.Println("Кол-во дверей:", sedan.numDoors)
	fmt.Println("Тип двигателя:", sedan.engineType)

	sportsCarBuilder := NewSportsCarBuilder()
	sportsCarDirector := NewCarDirector(sportsCarBuilder)
	sportsCar := sportsCarDirector.ConstructCar()

	fmt.Println("Спортивная машина:")
	fmt.Println("Модель:", sportsCar.model)
	fmt.Println("Кол-во дверей:", sportsCar.numDoors)
	fmt.Println("Тип двигателя:", sportsCar.engineType)
}
