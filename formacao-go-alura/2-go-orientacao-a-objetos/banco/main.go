package main

import "fmt"

type ContaCorrente struct {
	titular        string
	numertoAgencia int
	numeroConta    int
	saldo          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaSilvia := ContaCorrente{}
	contaSilvia.titular = "Silvia"
	contaSilvia.saldo = 500

	fmt.Println(contaSilvia.saldo)

	fmt.Println(contaSilvia.Sacar(200))
	fmt.Println(contaSilvia.saldo)

}
