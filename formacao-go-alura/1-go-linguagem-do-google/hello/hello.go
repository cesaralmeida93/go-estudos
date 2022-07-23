package main

import (
	"fmt"
)

func main() {
	var nome string

	fmt.Println("Digite o nome do Usuário")
	fmt.Scan(&nome)

	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	// fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é: ", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
