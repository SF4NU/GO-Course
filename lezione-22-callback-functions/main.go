package main

import (
	"fmt"
)

func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	addName("Marius", func(s string) {
		fmt.Printf("Hi my name is %v", s)
	})
}
