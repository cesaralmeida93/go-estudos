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
