package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2() int {
	return 5 + 7
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2())
}