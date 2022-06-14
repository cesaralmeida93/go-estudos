# 2: Fundamentos da Linguagem III(funções, for, while e do-while)

## For

```go
package main

import "fmt"

func main() {
	
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	
}
```

### While?
- não exite "while" nativo no go

```go
package main

import "fmt"

func main() {

	test := 0

	for test <= 10 {
		fmt.Println("VALOR: ", test)
		test++
	}
}
```

### Do-While?
- executa 1 vez, depois vê se a condição é true

```go
package main

import "fmt"

func main() {

	anExpression := false

	for ok := true; ok; ok = anExpression {

		fmt.Println("PASSOU AQUI")

	}
}

```

### For com Range

```go
package main

import "fmt"

func main() {

	test := []string{"test1", "test2", "test3"}

	for _, value := range test {
		fmt.Println(value)
	}
}

```

## Funções

### Retorno nomeado e retorno múltiplo

```go
package main

import (
	"fmt"
)

func main() {
	value, err := test()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)

}
 
func test() (retornoString string, retornoErro error) {
	retornoErro = nil
	retornoString = "test"

	return
}
```

### Passando funções por parametro e recebendo funções como retorno
- Passando funções por parâmetro:

```go
package main

import (
	"fmt"
)

func main() {

	funcaoTest := func(test string, testInt int) {
		fmt.Println(test, testInt)
	}

	test(funcaoTest)
}
 
func test(value func(string, int)) {
	value("otavio", 20)
}
```

- Recebendo funções como retorno:

```go
package main

import (
	"fmt"
)

func main() {


	funcao := test()

	funcao("otavio", 20)
}
 
func test() func(string, int) {
	
	funcaoTest := func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}

	return funcaoTest

}
```

### Funções Anônimas
- faz a instância da função e já executa, ao invés de atribuir a uma variável ee executar essa variável

```go
package main

import (
	"fmt"
)

func main() {

	test()

}
 
func test() {

	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("otávio", 20)

}
```

### Recebendo N parâmetros em uma função

```go
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
```