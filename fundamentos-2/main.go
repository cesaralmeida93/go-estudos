package main

import (
	"fmt"
)

func main() {

	// if retorno, err := test(); err != nil {
	// 	fmt.Println("true")
	// 	fmt.Println(retorno)
	// 	fmt.Print(err)
	// }

	// if 1 > 1 {
	// 	fmt.Print("True")
	// } else if 3 > 1 {
	// 	fmt.Print("test")
	// } else {
	// 	fmt.Print("false")
	// }

	// test := "test3"

	// switch test {

	// case "test", "test3":
	// 	fmt.Print("CAIU NA PRIMEIRA CONDIÇÃO")
	// 	fallthrough

	// case "test2":
	// 	fmt.Print("CAIU NA SEGUNDA CONDIÇÃO")
	// }


	test := "test9"

	switch test {

	case "test", "test2", "test434", "test1323":
		fmt.Print("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test3":
		fmt.Print("CAIU NA SEGUNDA CONDIÇÃO")

	case "test4":
		fmt.Print("CAIU NA TERCEIRA CONDIÇÃO")	

	default:
		fmt.Print("CAIU NO DEFAULT")
	}

}

// func test() (string, error) {
// 	return "test", errors.New("test")
// }
