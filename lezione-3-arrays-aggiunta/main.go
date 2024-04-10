package main

import "fmt"

func main() {
	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	var people [3]string

	people[0] = "Luigi"
	people[1] = "John"
	people[2] = "Bowser"

	fmt.Println(people)

	var arr = [3]string{"John", "Bowser", "Luigi"}
	fmt.Println(arr)
	fmt.Println(cap(arr))
	for item := 0; item < len(arr); item++ {
		fmt.Println(arr[item])
	}

	//slices

	num := []int{10, 20, 30, 40, 50}
	fmt.Println(num[:])
	fmt.Println(num[0:4])
	fmt.Println(num[3:])
	num = append(num, 30, 40, 50)
	fmt.Println(num)

	num1 := make([]int, 5, 20)

	num1[0] = 10
	num1[1] = 20
	num1[2] = 30
	num1[3] = 40
	num1[4] = 50

	fmt.Println(num1)
	fmt.Println(cap(num1))

	peoples := []string{"John", "Bowser", "Luigi"}

	for i := 0; i < len(peoples); i++ {
		fmt.Println(peoples[i])
	}

	for index, value := range peoples {
		fmt.Println(index, value)
	}
}
