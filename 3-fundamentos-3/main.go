package main

import (
	"fmt"
)

func main() {

	testParametro := func() {
		fmt.Println("test")
	}
	testParametro2 := func() {
		fmt.Println("test2")
	}

	test(testParametro, testParametro2)
}
 
func test(valoresString...func()) {

	for _, x := range valoresString { 
		x()
	}
}
