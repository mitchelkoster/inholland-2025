package main

import "fmt"

// Generic Engine
type Engine interface {
	Start()
}

// Engine Component
type DieselEngine struct{}

func (de DieselEngine) Start() {
	fmt.Println("Engine starting...")
}

// Electric Engine Component
type ElectricEngine struct{}

func (ee ElectricEngine) Start() {
	fmt.Println("Electric engine starting...")
}

// The can has an engine (composit)
type Car struct {
	engine    Engine // Interface
	buildYear int
}

func (c Car) Honk() {
	fmt.Println("Beep!")
}

func main() {
	// Observer
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()

	// Composition
	car1 := Car{engine: DieselEngine{}, buildYear: 2010}
	car2 := Car{engine: ElectricEngine{}, buildYear: 2021}

	cars := []Car{car1, car2}

	for _, c := range cars {
		c.engine.Start()
		c.Honk()
	}
}
