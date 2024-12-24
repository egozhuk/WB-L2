package main

import "fmt"

// Продукт
type House struct {
	doors   int
	windows int
	roof    string
}

// Интерфейс строителя
type HouseBuilder interface {
	SetDoors() HouseBuilder
	SetWindows() HouseBuilder
	SetRoof() HouseBuilder
	GetHouse() House
}

// Реализация строителя
type ConcreteHouseBuilder struct {
	house House
}

func (b *ConcreteHouseBuilder) SetDoors() HouseBuilder {
	b.house.doors = 4
	return b
}

func (b *ConcreteHouseBuilder) SetWindows() HouseBuilder {
	b.house.windows = 6
	return b
}

func (b *ConcreteHouseBuilder) SetRoof() HouseBuilder {
	b.house.roof = "Flat Roof"
	return b
}

func (b *ConcreteHouseBuilder) GetHouse() House {
	return b.house
}

// Директор
type Director struct {
	builder HouseBuilder
}

func (d *Director) SetBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) Construct() House {
	return d.builder.SetDoors().SetWindows().SetRoof().GetHouse()
}

func main() {
	builder := &ConcreteHouseBuilder{}
	director := &Director{}
	director.SetBuilder(builder)
	house := director.Construct()
	fmt.Printf("House built with %d doors, %d windows, and a %s.\n", house.doors, house.windows, house.roof)
}
