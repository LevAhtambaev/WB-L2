package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние — это шаблон проектирования поведения, который позволяет объекту изменять свое поведение при изменении его внутреннего состояния.
Шаблон выделяет поведение, связанное с состоянием, в отдельные классы состояний и заставляет исходный объект делегировать работу экземпляру этих классов, а не действовать самостоятельно.
*/

// Рассмотрим пример - система управления активностью пользователя в приложении.
// В зависимости от текущего состояния пользователя ("онлайн", "оффлайн", "занят"), приложение должно предоставлять различные функции и возможности.

import "fmt"

// UserState определяет интерфейс состояния пользователя.
type UserState interface {
	PerformAction()
}

// OnlineState реализует состояние "онлайн".
type OnlineState struct{}

func (s *OnlineState) PerformAction() {
	fmt.Println("Пользователь находится в сети. Он может просматривать контент и общаться с другими пользователями.")
}

// OfflineState реализует состояние "оффлайн".
type OfflineState struct{}

func (s *OfflineState) PerformAction() {
	fmt.Println("Пользователь не в сети. Он не может просматривать контент или общаться с другими пользователями.")
}

// BusyState реализует состояние "занят".
type BusyState struct{}

func (s *BusyState) PerformAction() {
	fmt.Println("Пользователь занят. Он может только реагировать на уведомления")
}

// User представляет пользователя.
type User struct {
	state UserState
}

func (u *User) SetState(state UserState) {
	u.state = state
}

func (u *User) PerformAction() {
	u.state.PerformAction()
}

func main() {
	// Создание объекта пользователя
	user := &User{}

	// Установка состояния "онлайн" и выполнение действия
	user.SetState(&OnlineState{})
	user.PerformAction()

	// Установка состояния "оффлайн" и выполнение действия
	user.SetState(&OfflineState{})
	user.PerformAction()

	// Установка состояния "занят" и выполнение действия
	user.SetState(&BusyState{})
	user.PerformAction()
}
