package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(500)
	PagarBoleto(&contaLuisa, 400)

	fmt.Println(contaLuisa.ObterSaldo())

}
