package main

import "fmt"

func main() {
	fmt.Println("enter a number to get its factorial")
	var number int
	fmt.Scanln(&number)
	fmt.Println("the factorial of ", number, " is ", computeFactorial(number))
}

func computeFactorial(number int) int {
	factorial := 1
	for factors := 1; factors <= number; factors++ {
		factorial *= factors
	}
	return factorial
}
