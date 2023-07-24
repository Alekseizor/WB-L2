package pattern

import "fmt"

// Интерфейс посетителя
type BodyVisitor interface {
	VisitSedan(sedan Sedan)
	VisitSUV(suv SUV)
}

// Конкретный посетитель - расчет стоимости автомобиля
type CostCalculator struct {
	totalCost int
}

func (c *CostCalculator) VisitSedan(sedan Sedan) {
	fmt.Println("Посещение легкового автомобиля")
}

func (c *CostCalculator) VisitSUV(suv SUV) {
	fmt.Println("Посещение внедорожника")
}

// Базовый класс кузова
type Body interface {
	Accept(visitor BodyVisitor)
}

// Класс легкового автомобиля
type Sedan struct{}

func (s Sedan) Accept(visitor BodyVisitor) {
	visitor.VisitSedan(s)
}

// Класс внедорожника
type SUV struct{}

func (s SUV) Accept(visitor BodyVisitor) {
	visitor.VisitSUV(s)
}

func main3() {
	// Создаем экземпляр посетителя
	calculator := &CostCalculator{}

	// Создаем экземпляры автомобилей
	sedan := Sedan{}
	suv := SUV{}

	// Применяем посетителя к автомобилям
	sedan.Accept(calculator)
	suv.Accept(calculator)
}

//Применимость паттерна Посетитель:
//1. Добавление методов классам, без изменения самих классов и их интерфейса.
//2. Если необходимо добавить методы для классов, но при этом данные методы не вписываются в абстракцию, которая существует
//3. Необходимость реализовать функционал, который будет еще изменяться.
//4. Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
//Плюсы паттерна Посетитель:
//1. Позволяет выполнять операции над группой объектов, не изменяя их классы.
//2. Гибкость, так как интерфейс "Посититель" может быть имплементирован самыми различными способами
//Минусы паттерна Посетитель:
//1. В некоторых случаях увеличивает сложность кода.
//Реальные примеры использования паттерна Посетитель:
//1. Оформление отчетов и каких-либо документов на основе имеющихся данных
