package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 20
}

func main() {

	ch := make(chan int)

	go sendData(ch)

	value := <-ch

	fmt.Println("Received: ", value)

}
