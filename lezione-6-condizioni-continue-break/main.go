package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 { //anche le if / else sono uguali a js e php
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //la parola chiave keyword obbliga la if a continuare il ciclo e salta quello che c'è fuori dall'if quindi il printf (in questo caso salta il printf solo una volta dato che la condizione è index == 1)
		}
		if index > 2 {
			fmt.Println("breaking at position", index)
			break //uguale a js
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
