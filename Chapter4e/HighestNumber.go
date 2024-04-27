package main

import (
	"fmt"
)

func main() {
	var myArray [10]int
	for counter := 0; counter < 10; counter++ {
		fmt.Println("enter a number")
		var number int
		fmt.Scanln(&number)
		myArray[counter] = number
	}
	newArray := sortArrayAgain(myArray)
	fmt.Println("the two largest numbers", newArray[8], "and", newArray[9])
}
func sortArrayAgain(array [10]int) [10]int {
	temp := 0
	for index := 0; index < len(array); index++ {
		for index1 := index + 1; index1 < len(array); index1++ {
			if array[index] > array[index1] {
				temp = array[index]
				array[index] = array[index1]
				array[index1] = temp
			}
		}
	}
	return array
}
