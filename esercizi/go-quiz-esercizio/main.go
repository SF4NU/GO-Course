package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func main() {
	problems, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(problems)

	count := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Printf("Risposte corrette: %v \n", count)
			fmt.Println("Press Enter to exit...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Quanto fa %v ? \n", record[0])
		reader := bufio.NewReader(os.Stdin)
		answer, _ := getInput("Scrivi la risposta corretta: ", reader)

		if answer == record[1] {
			count++
			fmt.Printf("%v è uguale a %v \n \n", record[0], answer)
		} else {
			fmt.Printf("Sbagliato! %v non fa %v \n \n", record[0], answer)
		}
	}
}
