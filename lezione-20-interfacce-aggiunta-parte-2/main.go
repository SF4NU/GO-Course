package main

import "fmt"

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
}

func (p Person) Describe() string {
	return "Hi, I'm " + p.Name
}

type Car struct {
	Model string
}

func (c Car) Describe() string {
	return "This is a " + c.Model + " car."
}

func displayDescription(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	person := Person{Name: "Alice"}
	car := Car{Model: "Toyota"}

	displayDescription(person)
	displayDescription(car)
}
