# 1: Fundamentos da Linguagem I(variáveis, arrays/slices e structs)

## Variáveis

### Tipos Básicos:
- Numbers
- Booleans
- Strings

### Como instanciar variáveis:

```go
package main

import "fmt"

func main() {

    var test string = "ewewewewew"
    test1 := "ewewewewewe"
}
```

```go
package main

import "fmt"

func main() {

    test1 := map[string]string{
		"test": "test",
	}
    fmt.Printf("%T", test1)

}
```

### Variáveis públicas e privadas:
- primeira letra minúscula: privada
- primeira letra maiúscula: pública

```go
package main

import "fmt"

func main() {

    test := 30
    fmt.Printf("%T", test)
}

func test() {

}

func Test() {
	
}
```

### Tipos e tipo "any" (interface{}):
- varável que aceita qualquer tipo de valor
- útil para quando não se sabe o tipo do valor que será recebido

```go
package main

import "fmt"

func main() {

	var test2 interface{}

	test2 = 20

	testJson := map[string]interface{}{
		"test21323": "Test",
		"testst": 20,
	}

	fmt.Printf("%T", test2)
}
```

## Arrays e Slices
- Array: Lista com tamanho definido
- Slice: Lista com tamanho indefinido

```go
package main

import "fmt"

func main() {

	var test [4]string = [4]string{"test", "test", "test", "test"}
	var test1 []string = []string{"test", "test", "test", "test"}

	test1 = append(test1, "cesar")
	fmt.Println(len(test))
	fmt.Println(cap(test))
	fmt.Println(len(test1))
	fmt.Println(cap(test1))
}
```

## Structs
- Idêntico ao que é definido como "objeto" nas outras linguagens

```go
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
```