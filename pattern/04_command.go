package pattern

import "fmt"

var overdraftLimit = -500

const (
	Deposit int = iota
	Withdraw
)

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited:", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		return true
	}
	return false
}

type Command interface {
	Call()
}

type BankAccountCommand struct {
	account *BankAccount
	action  int
	amount  int
}

func NewBankAccountCommand(account *BankAccount, action int, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

type Button interface {
	Press()
}

type DepositPress struct {
	cmd *BankAccountCommand
}

func (dp *DepositPress) Press() {
	dp.cmd.Call()
}

type WithdrawPress struct {
	cmd *BankAccountCommand
}

func (dp *WithdrawPress) Press() {
	dp.cmd.Call()
}

func main4() {
	ba := BankAccount{}                             //есть какой-то пользовательский аккаунт
	cmd := NewBankAccountCommand(&ba, Deposit, 100) //создаем новый экземляр для выполнения команд
	dp := DepositPress{cmd: cmd}
	dp.Press()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	dp2 := WithdrawPress{cmd: cmd2}
	dp2.Press()
	fmt.Println(ba)
}

//Применимость паттерна команда:
//1. Когда вы хотите параметризовать клиентские запросы и отделить их от получателей операций.
//2. Когда вы хотите поддерживать выполнение операций в разных временах, запускать их с задержкой, создавать очереди запросов или отменять операции.
//3. Когда вы хотите обеспечить отделение клиентов выполнивших запросы от объектов, которые выполняют эти запросы.
//
//Плюсы использования паттерна команда:
//1. Уменьшение связанности: Команда помогает изолировать клиентов от получателей операций, что уменьшает связанность и делает код более гибким и расширяемым.
//2. Поддержка отмены и повтора операций: Паттерн Command предоставляет удобный способ реализации отмены и повтора операций.
//3. Управление очередью и выполнением запросов: Команды могут использоваться для управления очередью запросов и определения порядка их выполнения.
//
//Минусы использования паттерна команда:
//
//1. Увеличение сложности: Использование паттерна Command может привести к увеличению сложности кода из-за необходимости создания дополнительных классов и интерфейсов.
//2. Увеличение объема памяти: Каждая команда, хранящаяся в истории или очереди, может занимать дополнительное место в памяти.
//3. Дополнительное время выполнения: Введение объектов команд может создавать дополнительные слои абстракции и влиять на производительность системы.
//
//Реальные примеры использования паттерна команда:
//
//1. Интерактивные программы с обработкой пользовательских команд и возможностью отмены и повтора операций.
//2. Управление домом, где команды могут использоваться для управления устройствами и автоматизации действий.
