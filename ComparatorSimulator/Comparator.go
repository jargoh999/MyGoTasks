package main

import "fmt"

func main() {
	fmt.Println("enter first number")
	var firstNumber int
	fmt.Scanln(&firstNumber)
	fmt.Println("enter second number")
	var secondNumber int
	fmt.Scanln(&secondNumber)
	fmt.Println(compare(firstNumber, secondNumber))
}

func compare(firstNumber int, secondNumber int) int {
	if firstNumber == secondNumber {
		return 0
	}
	if secondNumber > firstNumber {
		return -1
	}
	if firstNumber > secondNumber {
		return 1
	}
	return compare(firstNumber, secondNumber)
}
