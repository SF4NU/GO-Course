package main

import "fmt"

func main() {

	var nameOne string = "Mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20 //non si possono avere decimali con int
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8321831283128381238.7 //meglio questo

	scoreThree := 1.5

}
