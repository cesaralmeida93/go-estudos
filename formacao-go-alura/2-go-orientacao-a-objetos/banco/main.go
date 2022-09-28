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
