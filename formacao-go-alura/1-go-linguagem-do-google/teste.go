// package main

// import "fmt"

// func main() {
// 	var frutas [4]string
// 	frutas[0] = "Abacaxi"
// 	frutas[1] = "Laranja"
// 	frutas[2] = "Morango"

// 	for k, v := range frutas {
// 		fmt.Println(frutas[k])
// 		fmt.Println(v)
// 	}

// 	fmt.Println(frutas[3])
// }

package main

import (
	"fmt"
)

func main() {
	estados := devolveEstadosDoSudeste()
	fmt.Println(estados)
}

func devolveEstadosDoSudeste() [4]string {
	var estados [4]string
	estados[0] = "RJ"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"
	return estados
}
