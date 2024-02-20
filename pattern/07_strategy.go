package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия — это шаблон поведенческого проектирования, который позволяет вам определить семейство алгоритмов, поместить каждый из них в отдельный класс и сделать их объекты взаимозаменяемыми.
Исходный объект, называемый контекстом, содержит ссылку на объект стратегии.
Контекст делегирует выполнение поведения связанному объекту стратегии.
Чтобы изменить способ выполнения контекстом своей работы, другие объекты могут заменить текущий связанный объект стратегии другим.
*/

// Рассмотрим пример - система доставки еды, которая предлагает различные стратегии доставки в зависимости от времени заказа.
// Например, в пиковые часы или в праздничные дни может быть доступна стратегия "экспресс-доставка" с более быстрой доставкой, а в обычные дни - стандартная доставка.
// Таким образом, система доставки еды может легко переключаться между различными стратегиями доставки в зависимости от обстоятельств.

import "fmt"

// DeliveryStrategy определяет интерфейс стратегии доставки.
type DeliveryStrategy interface {
	Deliver()
}

// StandardDeliveryStrategy реализует стандартную стратегию доставки.
type StandardDeliveryStrategy struct{}

func (s *StandardDeliveryStrategy) Deliver() {
	fmt.Println("Стандартная доставка: 2-3 рабочих дня")
}

// ExpressDeliveryStrategy реализует стратегию экспресс-доставки.
type ExpressDeliveryStrategy struct{}

func (s *ExpressDeliveryStrategy) Deliver() {
	fmt.Println("Экспресс-доставка: в течение 24 часов")
}

// DeliveryService представляет сервис доставки.
type DeliveryService struct {
	strategy DeliveryStrategy
}

func (d *DeliveryService) SetStrategy(strategy DeliveryStrategy) {
	d.strategy = strategy
}

func (d *DeliveryService) Deliver() {
	d.strategy.Deliver()
}

func main() {
	// Создание объекта сервиса доставки
	deliveryService := &DeliveryService{}

	// Установка стратегии стандартной доставки
	deliveryService.SetStrategy(&StandardDeliveryStrategy{})
	deliveryService.Deliver()

	// Установка стратегии экспресс-доставки
	deliveryService.SetStrategy(&ExpressDeliveryStrategy{})
	deliveryService.Deliver()
}
