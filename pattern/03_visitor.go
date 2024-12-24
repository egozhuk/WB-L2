package main

import "fmt"

// Элемент интерфейса
type Element interface {
	Accept(v Visitor)
}

// Конкретные элементы
type ConcreteElementA struct{}
type ConcreteElementB struct{}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

// Посетитель интерфейса
type Visitor interface {
	VisitConcreteElementA(e *ConcreteElementA)
	VisitConcreteElementB(e *ConcreteElementB)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("Visiting Element A")
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("Visiting Element B")
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
