package main

import (
	"fmt"
	"github.com/google/uuid"
)

type allowedValuesHere interface {
	Admin | Cart | Customer | Item | Product
}

func main() {

	firstAdmin := Admin{}
	generateId(firstAdmin.getId())

	myPrint(firstAdmin)

}

func myPrint(value interface{}) {
	fmt.Println(value)
}

func generateId(value interface{}) int {
	if value.getId() == 0 {
		return int(uuid.New().ID())
	}
	return value.getId()
}
