package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge" //asterisco per puntare a name
}

func main() {
	name := "tifa"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name)

	m := &name //pointer ... una variabile che punta alla variabile name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m) //un asterisco davanti al pointer m dà come output la variabile .. in questo caso "tifa"
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}


