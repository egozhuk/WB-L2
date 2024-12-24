package main

import "fmt"

// Продукт интерфейса
type Product interface {
	Use()
}

// Конкретные продукты
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using Product A")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using Product B")
}

// Фабрика интерфейса
type Factory interface {
	CreateProduct() Product
}

// Конкретные фабрики
type FactoryA struct{}

func (f *FactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type FactoryB struct{}

func (f *FactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factoryA := &FactoryA{}
	productA := factoryA.CreateProduct()
	productA.Use()

	factoryB := &FactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()
}
