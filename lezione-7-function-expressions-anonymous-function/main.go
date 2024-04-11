package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Function expression")
	}

	f()

	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(add(10, 20))

	func() {
		fmt.Println("Anonymous function")
	}()
}
