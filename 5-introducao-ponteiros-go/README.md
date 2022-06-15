# 5: Introdução à ponteiros em Go

- passar valores entre funções sem ponteiros, cria uma cópia
- todos os "cantos" do código acessar o mesmo valor da variável, se utilizar ponteiro

```go
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
```

