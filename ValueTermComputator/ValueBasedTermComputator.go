package main

import "fmt"

func main() {
	fmt.Println("enter numbers to check")
	var input int
	fmt.Scanln(&input)
	fmt.Println("enter your number of terms")
	var numberOfTerms int
	fmt.Scanln(&numberOfTerms)
	fmt.Println("the number of terms to check for", numberOfTerms, "with the factorial of", input, "is", newTermsComputator(input, numberOfTerms))
}

func newTermsComputator(number int, value int) float64 {
	constant := 1.0
	for index := 1; index < number; index++ {
		value := raiseToPower(index, value) / float64(computeFactorial(index))
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

func raiseToPower(exponent int, baseNumber int) float64 {
	result := 1.0
	for timesToMultiply := 1; timesToMultiply <= exponent; timesToMultiply++ {
		result *= float64(baseNumber)
	}
	return result
}
