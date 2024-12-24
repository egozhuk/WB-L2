package main

import "fmt"

// Состояние интерфейса
type State interface {
	Handle()
}

// Конкретные состояния
type StateA struct{}

func (s *StateA) Handle() {
	fmt.Println("Handling State A")
}

type StateB struct{}

func (s *StateB) Handle() {
	fmt.Println("Handling State B")
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &StateA{}
	stateB := &StateB{}

	context.SetState(stateA)
	context.Request()

	context.SetState(stateB)
	context.Request()
}
