package main

import (
	"fmt"
)

type Animal interface {
	GetInfo() string
}

type Cat struct {
	Name  string
	Color string
}

type Dog struct {
	Name  string
	Breed string
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("Cat: %s, Color: %s", c.Name, c.Color)
}

func (d Dog) GetInfo() string {
	return fmt.Sprintf("Dog: %s, Breed: %s", d.Name, d.Breed)
}

func printAnimalInfo(animal Animal) {
	fmt.Println(animal.GetInfo())
}

func main() {
	cat := Cat{Name: "Whiskers", Color: "Gray"}
	dog := Dog{Name: "Bubbles", Breed: "Labrador"}
	printAnimalInfo(cat)
	printAnimalInfo(dog)
}
