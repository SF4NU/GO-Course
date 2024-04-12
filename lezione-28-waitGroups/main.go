package main

import (
	"fmt"
	"sync"
)

func printMessage(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) //Add the number of goroutines you're waiting for

	go printMessage(&wg, "Hello")
	go printMessage(&wg, "Go")
	go printMessage(&wg, "WaitGroup")

	wg.Wait() //Wait for all the goroutines to finish
	fmt.Println("All messages printed.")
}
