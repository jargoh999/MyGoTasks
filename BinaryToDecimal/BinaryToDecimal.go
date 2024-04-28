package main

import (
	"fmt"
	"math"
)

func main() {
	digit := 1101
	val := convertBinaryToDecimal(digit)
	fmt.Println(val)

}

func convertBinaryToDecimal(binaryValueToConvert int) int {
	position := 0
	decimalToReturn := 0
	for binaryValueToConvert > 0 {
		justAVariable := binaryValueToConvert % 10
		anotherVariable := justAVariable * int(math.Pow(2, float64(position)))
		decimalToReturn += anotherVariable
		binaryValueToConvert /= 10
		position++
	}
	return decimalToReturn
}
