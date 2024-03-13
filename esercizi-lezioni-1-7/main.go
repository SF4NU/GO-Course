package main

import (
	"fmt"
	// "sort"
)

//1

//Calculate the year given the date of birth and age

// func calculateYear(y float64, a float64) float64{
// 	return y + a;
// }

// func main() {
// 	saveNewYear := calculateYear(2001, 22)
// 	fmt.Println(saveNewYear)

// }

//2

//Create a program that calculates the average weight of 5 people.

// func calculateAvgWeight(a float64, b float64, c float64, d float64, e float64) float64{
// 	return (a + b + c + d + e) / 5
// }

// func main() {
// 	saveWeight := calculateAvgWeight(55, 88, 102, 48, 75)

// 	fmt.Printf("%v Kg", saveWeight)
// }

//3

// Create an array with the number 0 to 10 and loop through it, then change the numbers inside the loop and sort it then loop through it

// var numbers = []int{2, 4, 12, 23, 11, 55, 34, 8, 21, 39, 78}

// func main() {
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Printf("Number %v \n", numbers[i])
// 	// }

// 	sort.Ints(numbers)

// 	for i, v := range numbers {
// 		fmt.Printf("number: %v is at index: %v \n", v, i)
// 	}

// }

//4 Make a for that counts through 1 to 10

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

//5

//Make a program that divides x by 2 if it’s greater than 0

func divide(number float64) float64{
	if number > 0 {
		return (number / 2)
	} else {
		return (number)
	}
}

func main() {
	fmt.Println(divide(55))
	fmt.Println(divide(-28))
	fmt.Println(divide(200))
	fmt.Println(divide(10))
	fmt.Println(divide(0))
}