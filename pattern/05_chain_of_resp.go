package pattern

import (
	"fmt"
)

// Обработчик запроса
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *BaseHandler) Handle(request string) {
	if request == "" {
		fmt.Println("Обработка запроса завершена")
		return
	}
	h.nextHandler.Handle(request)
}

// Конкретные обработчики

type FirstHandler struct {
	nextHandler Handler
}

func (h *FirstHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *FirstHandler) Handle(request string) {
	if request == "first" {
		fmt.Println("Первый обработчик обрабатывает запрос")
		return
	}
	h.nextHandler.Handle(request)

}

type SecondHandler struct {
	nextHandler Handler
}

func (h *SecondHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *SecondHandler) Handle(request string) {
	if request == "second" {
		fmt.Println("Второй обработчик обрабатывает запрос")
		return
	}
	h.nextHandler.Handle(request)
}

type ThirdHandler struct {
	nextHandler Handler
}

func (h *ThirdHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ThirdHandler) Handle(request string) {
	if request == "third" {
		fmt.Println("Третий обработчик обрабатывает запрос")
		return
	}
	fmt.Println("Ошибка запроса")
}

func main5() {
	// Создание цепочки обработчиков
	baseHandler := &BaseHandler{}
	firstHandler := &FirstHandler{}
	secondHandler := &SecondHandler{}
	thirdHandler := &ThirdHandler{}

	baseHandler.SetNext(firstHandler)
	firstHandler.SetNext(secondHandler)
	secondHandler.SetNext(thirdHandler)

	// Обработка запросов
	baseHandler.Handle("first")
}

//Применимость паттерна цепочка вызовов:
//1. Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
//2. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
//Плюсы:
//1. Гибкость и расширяемость
//2. Реализует принцип единственной обязанности
//Недостатки:
//1. Запрос может остаться никем не обработанным
//2. Возможность зацикливания
//Реальные примеры использования:
//1. Обработка HTTP-запросов в веб-приложении
//2. Последовательная проверка условий
