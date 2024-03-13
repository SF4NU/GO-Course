package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	
	switch opt {
	case "a":
		name, _ := getInput("Insert item name: ", reader)
		price, _ := getInput("Insert item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		fmt.Println(tip+"$")
	case "s":
		fmt.Println("You chose s")
	default:
		choseOpt := strings.ToUpper(opt)
		fmt.Printf("%v was not a valid option... \n", choseOpt)
		promptOptions(b) //fa ripartire la funzione se non si sceglie una delle 3 opzioni
	}
}

func main() {
	myBill := createBill()
	promptOptions((myBill))

}
