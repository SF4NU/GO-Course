package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
	// "math/rand"
)

//1

//Make a program that rolls a dice (1 to 6)

// func main() {

// 	randomNum := rand.Intn(6)

// 	fmt.Println(randomNum)

// }

//2

// Take the string ‘hello world’ and slice it in two. Can you take a slice of a slice?

// func main() {
// 	hello := "hello world"

// 	sliceHelloWorld := hello[0:5]

// 	sliceAgain := hello[0:4]

// 	fmt.Println(sliceHelloWorld)
// 	fmt.Println(sliceAgain)
// }

//3

//Create a method that sums two numbers. 

func sumTwo(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func userInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	number1, _ := userInput("Inserisci il primo numero: ", reader) 
	number2, _ := userInput("Inserisci il secondo numero: ", reader)
	
	n1, err := strconv.ParseFloat(number1, 64)
	if err != nil {
		println("Non inserire caratteri diversi dai numeri!")
		main()
	}
	n2, err := strconv.ParseFloat(number2, 64)
	if err != nil {
		println("Non inserire caratteri diversi dai numeri!")
		main()
	}

	sum := sumTwo(n1, n2)

	fmt.Println("La tua somma è: ", sum)

}
