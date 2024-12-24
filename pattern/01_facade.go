package main

import "fmt"

// Подсистемы
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU started")
}

func (c *CPU) Execute() {
	fmt.Println("CPU executing instructions")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory loaded")
}

type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("HardDrive reading data")
}

// Фасад
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) StartComputer() {
	fmt.Println("Starting computer...")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.Read()
	cf.cpu.Execute()
}

func main() {
	computer := NewComputerFacade()
	computer.StartComputer()
}
