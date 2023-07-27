package pattern

type TypeToy string

const (
	BearType  TypeToy = "bear"
	RobotType TypeToy = "robot"
)

type Toy interface {
	GetName() string
}

type Bear struct {
}

func NewBear() Toy {
	return &Bear{}
}

func (b Bear) GetName() string {
	return "teddy bear white"
}

type Robot struct {
}

func NewRobot() Toy {
	return &Robot{}
}

func (b Robot) GetName() string {
	return "a battery-powered police robot"
}

func New(variety TypeToy) Toy {
	switch variety {
	case BearType:
		return NewBear()
	case RobotType:
		return NewRobot()
	default:
		return nil
	}
}

func main6() {
	_ = New(BearType)
	_ = New(RobotType)
}

//Применимость паттерна фабричный метод:
//1. Когда у вас есть общий интерфейс для создания объектов, но конкретный класс, который следует использовать, определяется только на этапе выполнения.
//2. Когда необходимо делегировать ответственность за инстанцирование объектов.
//Плюсы паттерна:
//1. Вы избегаете тесной связи между классом создателя и конкретными классами продуктов.
//2. Принцип единственной ответственности. Вы можете переместить код создания продукта в одно место в программе, что упростит поддержку кода.
//Минусы паттерна:
//1. Код может стать более сложным
//Реальные примеры использования:
//1. Создание похожих по методам между собой объектов
