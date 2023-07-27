package pattern

import "log"

type Animal interface {
	Drink() error
}

func main7() {
	cat := NewCat()
	err := Live("cat", cat)
	if err != nil {
		log.Println(err)
	}
}

func Live(name string, animal Animal) error {
	err := animal.Drink()
	if err != nil {
		return err
	}
	log.Println(name)
	return nil
}

type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (c *Cat) Drink() error {
	return nil
}

type Dog struct {
}

func NewDog() *Dog {
	return &Dog{}
}

func (c *Dog) Drink() error {
	return nil
}

//Применимость паттерна стратегия:
//1. Когда у вас есть множество похожих классов, отличающихся только некоторым поведением
//2. Когда вы не хотите обнажать детали реализации алгоритмов для других классов
//Плюсы паттерна:
//1. Горячая замена алгоритмов на лету.
//2. Изолирует код и данные алгоритмов от остальных классов.
//3. Уход от наследования к делегированию.
//Минусы паттерна:
//1. Увеличение сложности кода
//Реальные примеры использования:
//1. С моей точки зрения, данный паттерн используется повсеместно в разработке на Go: прокидывание реализации БД для сервиса, различных алгоритмов и так далее.
