package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{ //simili agli oggetti
		267272317: "mario",
		241423411: "luigi",
		124514251: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267272317])

	phonebook[124514251] = "bowser"
	fmt.Println(phonebook)

	phonebook[241423411] = "yoshi"
	fmt.Println(phonebook)

}
