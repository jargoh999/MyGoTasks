package main

import "fmt"

func main() {
	fmt.Println(noPlusAdder(-7, -3))
}
func noPlusAdder(firstNumber int, secondNumber int) int {
	if isNegative(firstNumber, secondNumber) {
		return negativeNumberAdder(firstNumber, secondNumber)
	}
	if eitherOfTheNumberIsNegative(firstNumber, secondNumber) {
		return eitherNegOrPos(firstNumber, secondNumber)
	}
	adder := 1
	for adder <= secondNumber {
		firstNumber++
		adder++
	}
	return firstNumber
}
func isNegative(firstNumber int, secondNumber int) bool {
	return firstNumber < 0 && secondNumber < 0
}
func negativeNumberAdder(firstNumber int, secondNumber int) int {
	adder := 0
	for adder > secondNumber {
		firstNumber--
		adder--
	}
	return firstNumber

}
func eitherOfTheNumberIsNegative(firstNumber int, secondNumber int) bool {
	return firstNumber < 0 || secondNumber < 0
}
func eitherNegOrPos(firstNumber int, secondNumber int) int {
	if firstNumber < 0 {
		secondNumber = -firstNumber - secondNumber
		return -secondNumber
	}
	if secondNumber < 0 {
		firstNumber = -secondNumber - firstNumber
		return -firstNumber
	}
	return 0
}
