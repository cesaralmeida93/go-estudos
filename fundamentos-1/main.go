package main

import "fmt"

func main() {
	// test := map[string]string{
	// 	"test": "test",
	// }

	// var test2 interface{}

	// test2 = 20.5

	// testJson := map[string]interface{}{
	// 	"test21323": "Test",
	// 	"testst":    20,
	// }

	// var test []string = []string{"test", "test", "test", "test"}

	// fmt.Println(len(test))
	// fmt.Println(cap(test))

	// test = append(test, "otavio")

	// fmt.Println(len(test))
	// fmt.Println(cap(test))

	var user User = User{
		name:  "otavio",
		age:   20,
		test2: "test2",
	}

	fmt.Print(user)

}

type User struct {
	name  string
	age   int
	test2 string
}
