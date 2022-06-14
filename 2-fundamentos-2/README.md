# 2: Fundamentos da Linguagem II(estruturas de controle com if e switch)

## If/Else

```go
package main

import "fmt"

func main() {

	if 1 > 1 {
		fmt.Print("true")
	} else if 3 > 1 {
		fmt.Print("teste")
	} else {
		fmt.Print("false")
	}
}
```

## If com Init
- usado nos casos onde a variável é criadas apenas para verificar uma condição
- importante: o valor dela não é acessível fora do escopo do If

```go
package main

import (
	"errors"
	"fmt"
)

func main() {

	if retorno1, err := test1(); err != nil {
		fmt.Println("true")
		fmt.Println(retorno1)
		fmt.Println(err)
	}

	if retornoTest := test(); retornoTest == "test" {
		fmt.Println("true")
	}
	 
}

func test() string {
	return "test"
}

func test1() (string, error) {
	return "test1", errors.New("test1")
}
```


## Switch

```go
package main

import (
	"fmt"
)

func main() {

	test := "test"

	switch test {
	case "test":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test2":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
	}
	 
}
```

### Aceitando vários valores nos cases

```go
package main

import (
	"fmt"
)

func main() {

	test := "test2"

	switch test {
	case "test", "test3":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test2":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
	}
	 
}
```


### Fallthrough e Break
- faz o programa "cair" nos outros cases do switch, mesmo se o programa já "entrou" no case certo.
- o break faz o programam parar(já para sozinho sem o fallthrough)

```go
package main

import (
	"fmt"
)

func main() {

	test := "test"

	switch test {
	case "test", "test3":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")
		fallthrough

	case "test2":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
		fallthrough

	case "test5":
		fmt.Println("CAIU NA TERCEIRA CONDIÇÃO")
		fallthrough

	case "test6":
		fmt.Println("CAIU NA QUARTA CONDIÇÃO")
		break
	
	case "test7":
		fmt.Println("CAIU NA QUINTA CONDIÇÃO")
	}
	 
}
```

### Default
- case que serã o selecionado se o valor passado não corresponder a nenhum dos outros cases

```go
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
```