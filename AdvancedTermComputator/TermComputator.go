package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var number int
	fmt.Scanln(&number)
	fmt.Println("the equivalent e constant of this number is", termsComputator(number))
}

func termsComputator(number int) float64 {
	constant := 1.0
	for index := 1; index < number; index++ {
		value := 1.0 / float64(computeFactorial(index))
		constant += value
	}
	return constant
}

func computeFactorial(number int) int {
	factorial := 1
	for factors := 1; factors <= number; factors++ {
		factorial *= factors
	}
	return factorial
}
