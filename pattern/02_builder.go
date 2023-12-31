package pattern

type Car struct {
	Brand string
	Model string
	Color string
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) SetBrand(brand string) {
	c.Brand = brand
}

func (c *Car) SetModel(model string) {
	c.Model = model
}

func (c *Car) SetColor(color string) {
	c.Color = color
}

//Применимость паттерна Строитель:
//1. Когда нужно создавать объекты, состоящие из множества частей с разными вариациями или конфигурациями.
//2. Когда необходимо создавать сложные объекты, но не хочется загромождать конструктор класса большим количеством параметров.
//3. Когда нужно обеспечить пошаговое создание объекта, чтобы иметь возможность контролировать каждый этап процесса.
//Плюсы паттерна Строитель:
//1. Упрощает процесс создания сложных объектов, скрывая детали конструирования.
//2. Позволяет создавать различные вариации объектов, используя один тот же код строительства.
//3. Изолирует клиентский код от классов конкретных продуктов.
//Минусы паттерна Строитель:
//1. В некоторых случаях увеличивает сложность кода из-за введения дополнительных методов.
//2. Может быть избыточным, если объекты имеют простую структуру.
//Реальные примеры использования паттерна Строитель:
//1. Оформление заказа
//2. Переводы в банке
//3. Игры
