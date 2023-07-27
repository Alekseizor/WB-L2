package pattern

import "fmt"

// Интерфейс состояния
type State interface {
	Process() State
}

// Состояние 1
type WelcomeState struct{}

func (w *WelcomeState) Process() State {
	fmt.Println("Hello!")
	return &OrderCompletionState{}
}

// Состояние 2
type OrderCompletionState struct{}

func (o *OrderCompletionState) Process() State {
	fmt.Println("Thank you for placing an order!")
	return &WelcomeState{}
}

// Контекст
type Client struct {
	state State
}

func (c *Client) request() {
	newState := c.state.Process() // Обработка запроса текущего состояния
	c.state = newState
}

func main8() {
	client := &Client{state: &WelcomeState{}} // Устанавливаем начальное состояние

	// Выполняем несколько запросов для демонстрации изменения состояний
	client.request()
	client.request()
}

//Применимость паттерна стратегия:
//1. Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется
//2. Когда код объекта должен легко расширяться с добавлением новых состояний и/или изменением существующих.
//Плюсы паттерна:
//1. Избавляет от множества больших условных операторов машины состояний.
//2. Концентрирует в одном месте код, связанный с определённым состоянием.
//Минусы паттерна:
//1. Может неоправданно усложнить код, если состояний мало и они редко меняются.
//Реальные примеры использования:
//1. Чат боты
//2. Корзина в интернет-магазине
//3. Различные автоматы по продаже чего-либо,постаматы
