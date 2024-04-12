package main

import (
	"fmt"
)

func printNumber(item, defaultValue int) (int, int) {
	return item, defaultValue
}

func printString(item, defaultValue string) (string, string) {
	return item, defaultValue
}

func printBool(item, defaultValue bool) (bool, bool) {
	return item, defaultValue
}

//generic function

func printItem[T any](item, defaultValue T) (T, T) {
	return item, defaultValue
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	num1, num2 := printNumber(10, 20)
	num3, num4 := printString("one", "two")
	num5, num6 := printBool(false, true)

	num7, num8 := printItem(10, 20)
	num9, num10 := printItem("one", "two")
	num11, num12 := printItem(false, true)

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)
	fmt.Println(num11)
	fmt.Println(num12)

	intSlice := []int{1, 2, 3, 4, 5}
	PrintSlice(intSlice)

	stringSlice := []string{"apple", "banana", "cherry"}
	PrintSlice(stringSlice)
}
