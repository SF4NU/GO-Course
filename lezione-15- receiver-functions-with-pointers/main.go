package main

import "fmt"

func main() {
	myBill := newBill("mario's bill")

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("potato soup", 5.50)
	myBill.addItem("tomato soup", 10.50)
	myBill.updateTip(10)

	fmt.Println(myBill.format())

}