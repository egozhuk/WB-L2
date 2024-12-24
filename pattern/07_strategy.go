package main

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

// Конкретные стратегии
type StrategyA struct{}

func (s *StrategyA) Execute() {
	fmt.Println("Executing Strategy A")
}

type StrategyB struct{}

func (s *StrategyB) Execute() {
	fmt.Println("Executing Strategy B")
}

// Контекст
type Context07 struct {
	strategy Strategy
}

func (c *Context07) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context07) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := &Context07{}

	strategyA := &StrategyA{}
	strategyB := &StrategyB{}

	context.SetStrategy(strategyA)
	context.ExecuteStrategy()

	context.SetStrategy(strategyB)
	context.ExecuteStrategy()
}
