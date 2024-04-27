package main

import (
	"fmt"
	"math"
)

func main() {
	for customers := 0; customers < math.MaxInt; customers++ {
		fmt.Println("enter your account number")
		var customerAza int
		fmt.Scanln(&customerAza)
		fmt.Println("enter the customer balance at the beginning of the month")
		var customerBalance int
		fmt.Scanln(&customerBalance)
		fmt.Println("enter the total item charged by the customer this month")
		var totalItem int
		fmt.Scanln(&totalItem)
		fmt.Println("enter the total credit applied to the customer this month")
		var totalCredit int
		fmt.Scanln(&totalCredit)
		fmt.Println("enter the allowed credit limit")
		var allowedCredit int
		fmt.Scanln(&allowedCredit)
		newBalance := customerBalance + totalItem - totalCredit
		fmt.Println("your Chapter4 balance is", newBalance)
		if newBalance > totalCredit {
			fmt.Println("credit limit exceeded")
		}
		fmt.Println("do you want to calculate for another customer")
		var promptAgain string
		fmt.Scanln(&promptAgain)
		if promptAgain == "no" {
			break
		}
	}

}
