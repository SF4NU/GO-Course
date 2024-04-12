package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go printNumbers()

	fmt.Println("Main function call")

	time.Sleep(1 * time.Second)
}
