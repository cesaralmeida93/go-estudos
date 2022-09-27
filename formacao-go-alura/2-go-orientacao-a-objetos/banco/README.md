# GO: Orienttação a objetos

## 1-Primeira Struct
- **struct:** "estutura" onde se pode armazenar variáveis e funções

```go
type ContaCorrente struct {
	titular        string
	numertoAgencia int
	numeroConta    int
	saldo          float64
}
```

- pode-se instanciar uma struct não passando o nome dos campos, se todos forem preenchidos

```go
package main

import "fmt"

type ContaCorrente struct {
	titular        string
	numertoAgencia int
	numeroConta    int
	saldo          float64
}

func main() {
	contaGuilherme := ContaCorrente{
		titular:        "Guilherme",
		numertoAgencia: 589,
		numeroConta:    123456,
		saldo:          125.50,
	}
	contaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
}
```

## 2-Referência, ponteiro e métodos
- pode-se insntanciar uma struct usando a palavra `new`, porém a variável que recebe precisa ser um ponteiro

```go
var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.saldo = 500

fmt.Println(*contaDaCris)

```

- pode-se instanciar uma struct não passando o nome dos campos, se todos forem preenchidos

```go
package main

import "fmt"

type ContaCorrente struct {
	titular        string
	numertoAgencia int
	numeroConta    int
	saldo          float64
}

func main() {
	contaGuilherme := ContaCorrente{
		titular:        "Guilherme",
		numertoAgencia: 589,
		numeroConta:    123456,
		saldo:          125.50,
	}
	contaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
}
```

- no Go, pode-se criar funções para a struct, semelhante ao método de uma classe

```go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}
```

- pode-se criar funções com quantidade de parâmetros indeterminados, chamada de `função variática`

```go
package main

import (
    "fmt"
)

func Somando(numeros ...int) int {
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}

func main() {
    fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))
}
```