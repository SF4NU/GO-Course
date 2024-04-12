package main

import (
	"fmt"
)

func main() {
	//unbuffered channel
	unbufferedCh := make(chan int)

	//buffered channel with a capacity of 2
	bufferedCh := make(chan string, 2)

	go func() {
		unbufferedCh <- 20
	}()

	go func() {
		bufferedCh <- "hello"
		bufferedCh <- "world"
		close(bufferedCh)
	}()

	value := <-unbufferedCh
	fmt.Println("Received from unbuffered channel: ", value)

	for msg := range bufferedCh {
		fmt.Println("Received from buffered channel: ", msg)
	}
}
