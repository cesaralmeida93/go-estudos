package main

import "fmt"

func main() {

	var user User = User{
		name: "cesar",
		age: 28,
		test2: "test2",
	}

	fmt.Print(user)

}

type User struct {
	name string
	age int
	test2 string
}