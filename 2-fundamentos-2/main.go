package main

import (
	"fmt"
)

func main() {

	test := "test8"

	switch test {

	case "test", "test99":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "tchutchutchu":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")

	default:
		fmt.Println("CAIU NO DEFAULT")
		
	}
	 
}


