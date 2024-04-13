package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func startTimer() {
	time.Sleep(1 * time.Second)
	fmt.Println("Il quiz comincia tra 3 secondi")
	fmt.Println("3")
	time.Sleep(1 * time.Second)
	fmt.Println("2")
	time.Sleep(1 * time.Second)
	fmt.Println("1")
	time.Sleep(1 * time.Second)
	fmt.Println("Via!")
}

func main() {
	problems, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(problems)

	count := 0

	go func() {
		// reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ci saranno 20 domande di matematica, avrai 30 secondi di tempo per rispondere a tutto, se scade il tempo il quiz finisce indifferentemente da quante risposte hai dato.")

		reader := bufio.NewReader(os.Stdin)

		choice, _ := getInput("Se vuoi cambiare il tempo schiaccia 1 altrimenti schiaccia 2 per iniziare il quiz con il tempo predefinito di 30 secondi", reader)

		var chosenTime int = 30

		if choice == "1" {
			func() {
				userTime, _ := getInput("Imposta quanti secondi di tempo vuoi avere:", reader)

				convertedTime, err := strconv.Atoi(userTime)
				if err != nil {
					log.Fatal(err)
				}
				chosenTime = convertedTime
				startTimer()
			}()
		} else if choice == "2" {
			fmt.Println("Hai scelto di continuare con il tempo predefinito di 30 secondi!")
			startTimer()

		} else {
			fmt.Printf("La risposta: %v, non è accettata \n", choice)
			fmt.Println("Schiaccia invio per uscire...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		timer := time.NewTimer(time.Duration(chosenTime) * time.Second)

		for {
			select {
			case <-timer.C:
				fmt.Println("Tempo finito")
				fmt.Printf("Risposte corrette: %v su 20 \n", count)
				fmt.Println("Schiaccia invio per uscire...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				return
			default:
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Risposte corrette: %v su 20 \n", count)
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
	}()

	select {}
}
