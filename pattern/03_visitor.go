package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Посетитель — это шаблон поведенческого проектирования, который позволяет добавлять новые варианты поведения в существующую иерархию классов без изменения существующего кода.
// Шаблон Посетитель позволяет добавлять поведение в структуру без фактического изменения структуры.

// Рассмотрим такой пример - приложение для управления задачами и у нас есть различные типы задач, такие как задачи на программирование, задачи на дизайн, задачи на тестирование.

import "fmt"

// Task представляет общий интерфейс для всех типов задач.
type Task interface {
	Accept(visitor TaskVisitor)
}

// ProgrammingTask - задача на программирование.
type ProgrammingTask struct {
	Name     string
	Priority int
}

func (t *ProgrammingTask) Accept(visitor TaskVisitor) {
	visitor.VisitProgrammingTask(t)
}

// DesignTask - задача на дизайн.
type DesignTask struct {
	Name     string
	Priority int
}

func (t *DesignTask) Accept(visitor TaskVisitor) {
	visitor.VisitDesignTask(t)
}

// TestingTask - задача на тестирование.
type TestingTask struct {
	Name     string
	Priority int
}

func (t *TestingTask) Accept(visitor TaskVisitor) {
	visitor.VisitTestingTask(t)
}

// TaskVisitor определяет интерфейс посетителя.
type TaskVisitor interface {
	VisitProgrammingTask(task *ProgrammingTask)
	VisitDesignTask(task *DesignTask)
	VisitTestingTask(task *TestingTask)
}

// TaskPriorityVisitor реализует посетителя, который выводит приоритет задачи.
type TaskPriorityVisitor struct{}

func (v *TaskPriorityVisitor) VisitProgrammingTask(task *ProgrammingTask) {
	fmt.Printf("Приоритет задачи на программирование %s: %d\n", task.Name, task.Priority)
}

func (v *TaskPriorityVisitor) VisitDesignTask(task *DesignTask) {
	fmt.Printf("Приоритет задачи на дизайн %s: %d\n", task.Name, task.Priority)
}

func (v *TaskPriorityVisitor) VisitTestingTask(task *TestingTask) {
	fmt.Printf("Приоритет задачи на тестирование %s: %d\n", task.Name, task.Priority)
}

// В main создаются задачи разных типов, а затем они посещаются посетителем для вывода их приоритетов.

func main() {
	tasks := []Task{
		&ProgrammingTask{Name: "Реализовать функцию", Priority: 1},
		&DesignTask{Name: "Сделать дизайн для сайта", Priority: 2},
		&TestingTask{Name: "Написать тесты", Priority: 3},
	}

	visitor := &TaskPriorityVisitor{}

	for _, task := range tasks {
		task.Accept(visitor)
	}
}
