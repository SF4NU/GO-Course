package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20, 25, 30} //in go un array ha una lunghezza fissa e bisogna dichiararla

	var ages = [3]int{20, 25, 30} //modo più veloce per scrivere un array e funziona allo stesso modo

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices usano array ma si possono manipolare meglio
	var scores = []int{100, 50, 60} //slice
	scores[2] = 25
	scores = append(scores, 85) //siccome append restituisce un nuovo slice deve essere riassegnato a scores

	fmt.Println(scores, len(scores))

	//slice ranges 

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
