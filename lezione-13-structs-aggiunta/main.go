package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Programmer struct {
	Person
	isProgrammer bool
}

func main() {
	// h := Person {
	// 	FirstName: "John",
	// 	LastName:  "Smith",
	// }

	// var h Person

	// h.FirstName = "john"
	// h.LastName = "smith"

	// fmt.Println(h)

	var h Programmer

	h.Person.FirstName = "John"
	h.Person.LastName = "Smith"
	h.isProgrammer = true
	fmt.Println(h.FirstName)
	fmt.Println(h.isProgrammer)

	g := struct {
		FirstName string
		LastName  string
	}{
		FirstName: "John",
		LastName:  "Smith",
	}

	fmt.Println(g.FirstName, g.LastName)
}
