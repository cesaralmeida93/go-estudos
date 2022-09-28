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
- pode-se instanciar uma struct usando a palavra `new`, porém a variável que recebe precisa ser um ponteiro

```go
var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.saldo = 500

fmt.Println(*contaDaCris)
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

## 3-Retornos, pacotes e visibilidade

- pode-se definir mais de um retorno para uma função

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

status, valor := contaSilvia.Depositar(2000)
fmt.Print(status, valor)
```


- pode-se definir pacotes no Go para melhorar a legibilidade do código
- no exemplo abaixo, defini tudo o que se refere a contas no pacote `contas`

```go
package contas

type ContaCorrente struct {
	Titular        string
	NumertoAgencia int
	NumeroConta    int
	Saldo          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Valor do depósito menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
```

- para importar o pacote `contas` no arquivo `main.go`, deve-se primeiro executar o comando `go mod init nome-do-projeto`
- isso vai cirar o arquivo `go.mod` com os módulos utilizados(nesse caso, apenas o módulo de `banco`)
- no arquivo `main.go`ficam apenas as chamadas para `contas`


```go
package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaSilvia)

	fmt.Println(status)
	fmt.Println(contaSilvia)
	fmt.Println(contaDoGustavo)
}
```

