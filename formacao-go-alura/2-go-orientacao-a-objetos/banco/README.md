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