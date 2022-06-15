package main

import (
	"fmt"
)

func main() {

	x := 100
	y := &x
	x++

	fmt.Println(x, *y)

	testValue := "Otávio"
	copyStringVALUE(testValue)
	fmt.Println(testValue)

	originalStringValue(&testValue)
	fmt.Println(testValue)

}

func copyStringVALUE(stringValue string) {
	stringValue = "TEST"
	fmt.Println(stringValue)
}

func originalStringValue(stringValue *string) {
	*stringValue = "TEST"
	fmt.Println(*stringValue)
}