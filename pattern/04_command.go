package pattern

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

type Receiver struct {
}

func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

/*
Паттерн "Команда" применяется, когда требуется отделить вызывающий объект от объекта, который выполнит конкретную операцию.
Это позволяет создать гибкую и расширяемую систему команд, которые можно объединять, отменять или повторять.

Плюсы использования паттерна "Команда" включают:

1. Расширяемость: Паттерн позволяет легко добавлять новые команды, не изменяя код вызывающих объектов.
Это позволяет динамически добавлять и изменять функциональность системы.
2. Отделение вызывающего объекта от получателя: Паттерн помогает отделить логику вызывающего объекта от самой операции, которую нужно выполнить.
Это позволяет легко конфигурировать и комбинировать различные команды.
3. Управление отменой операций: Паттерн "Команда" предоставляет механизм отмены операций.
Каждая команда может реализовывать методы отмены и повтора операции, что улучшает управление выполнением команд в системе.

Несколько возможных минусов паттерна "Команда":

1. Усложнение структуры системы: Использование паттерна "Команда" может привести к повышенной сложности архитектуры системы из-за наличия дополнительных
классов и интерфейсов.
2. Возможное повышение затрат памяти: Если система содержит большое количество команд, это может привести к увеличению затрат оперативной памяти.

Необходимость использования паттерна "Команда" зависит от специфики задачи. Он полезен, когда требуется отделить вызов операции от ее непосредственного исполнения,
позволяя создавать гибкую и расширяемую систему команд.
*/
