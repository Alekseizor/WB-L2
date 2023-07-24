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
