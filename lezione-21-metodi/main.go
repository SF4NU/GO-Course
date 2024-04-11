package main

import (
	"fmt"
)

type Person struct {
	Gender     string
	FirstName  string
	LastName   string
	Age        int
	Profession string
	IsMarried  bool
}

func (p Person) printInfo() {
	fmt.Println("I'm", p.FirstName, p.LastName, "Gender: ", p.Gender, "Age: ", p.Age, "Profession: ", p.Profession, "Married: ", p.IsMarried)
}

func main() {
	Marius := Person{
		Gender:     "Maschio",
		FirstName:  "Marius",
		LastName:   "Marius",
		Age:        22,
		Profession: "Software Developer",
		IsMarried:  false,
	}

	Marius.printInfo()
}
