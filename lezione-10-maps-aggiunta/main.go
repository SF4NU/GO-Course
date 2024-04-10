package main

import (
	"fmt"
)

func main() {
	userData := map[string]int{
		"john":   20,
		"luigi":  18,
		"bowser": 22,
	}

	fmt.Println(userData)

	userData["john"] = 25

	fmt.Println(userData)

	for prop, value := range userData {
		fmt.Println(prop, value)
	}

	newStudentData := map[int]string{
		20: "john",
		18: "luigi",
		22: "bowser",
	}

	fmt.Println(newStudentData)
}
