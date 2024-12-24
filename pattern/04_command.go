package main

import "fmt"

// Команда интерфейс
type Command interface {
	Execute()
}

// Конкретные команды
type LightOnCommand struct{}
type LightOffCommand struct{}

func (c *LightOnCommand) Execute() {
	fmt.Println("Light is ON")
}

func (c *LightOffCommand) Execute() {
	fmt.Println("Light is OFF")
}

// Инициатор
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	remote := &RemoteControl{}

	lightOn := &LightOnCommand{}
	lightOff := &LightOffCommand{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
