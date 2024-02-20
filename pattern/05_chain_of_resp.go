package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Цепочка вызовов — это шаблон поведенческого проектирования, который позволяет передавать запросы по цепочке обработчиков.
Получив запрос, каждый обработчик решает либо обработать запрос, либо передать его следующему обработчику в цепочке.
Шаблон позволяет нескольким объектам обрабатывать запрос без привязки класса отправителя к конкретным классам получателей.
Цепочку можно составлять динамически во время выполнения с помощью любого обработчика, который соответствует стандартному интерфейсу обработчика.
*/

// Рассмотрим пример - обработка запросов на обслуживание клиентов в отделе технической поддержки.

import "fmt"

// Handler представляет интерфейс обработчика запроса.
type Handler interface {
	HandleRequest(request string)
	SetNext(handler Handler)
}

// BaseHandler реализует базовую функциональность обработчика.
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) HandleRequest(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// TechnicalSupportHandler представляет обработчик запросов технической поддержки.
type TechnicalSupportHandler struct {
	next Handler
}

func (h *TechnicalSupportHandler) HandleRequest(request string) {
	fmt.Println("Запрос на техническую поддержку:", request)
	fmt.Println("Техническая поддержка обрабатывает запрос.")
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *TechnicalSupportHandler) SetNext(handler Handler) {
	h.next = handler
}

// SalesHandler представляет обработчик запросов отдела продаж.
type SalesHandler struct {
	next Handler
}

func (h *SalesHandler) HandleRequest(request string) {
	fmt.Println("Запрос от клиента отдела продаж:", request)
	fmt.Println("Отдел продаж обрабатывает запрос.")
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *SalesHandler) SetNext(handler Handler) {
	h.next = handler
}

// MainHandler представляет главный обработчик, который будет передавать запросы по цепочке.
type MainHandler struct {
	BaseHandler
}

func main() {
	// Создаем обработчиков запросов
	technicalSupportHandler := &TechnicalSupportHandler{}
	salesHandler := &SalesHandler{}

	// Устанавливаем следующего обработчика в цепочке
	mainHandler := &MainHandler{}
	mainHandler.SetNext(technicalSupportHandler)
	technicalSupportHandler.SetNext(salesHandler)

	// Обработка запросов
	mainHandler.HandleRequest("Проблема с подключением к интернету")
	mainHandler.HandleRequest("Запрос на покупку нового продукта")
}
