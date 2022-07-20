package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}
