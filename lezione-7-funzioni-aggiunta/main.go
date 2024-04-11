package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world!")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func showNumbers(anotherParam string, numbers ...int) {
	fmt.Println(anotherParam, numbers)
}

func showUsers(s ...string) {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func add(x int, y int) int {
	return x + y
}

func main() {
	hello()
	greet("Marius")
	showNumbers("Hello", 1, 2, 3, 4, 5, 6, 7)
	showUsers("alex", "john", "luigi")

	sum := add(10, 20)

	fmt.Println(sum)
}
