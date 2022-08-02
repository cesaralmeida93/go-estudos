# GO: A linguagem do Google

## 1-Introdução ao GO
Go é uma linguage compilada, então é necessário compular e executar

```
go build main.go
./main
```

- Dá pra concatenar os dois comandos, utilizando o go run(compila e executa)

```
go run main.go
```

- o `;` é opcional em go
- a `{` sempre fica ao lado da função, não em baixo.
- go tem várias convenções para deixar o desenvolvimento mais rápido.

## 2-Trabalhando com Variáveis

- o pacote `fmt` contém funções como `Println`, `Scanf` e `Scan`
- a função `Scanf` recebe um modificador `%s` para string, `%f` para float, `%d` para int, e um ponteiro para a variável que guarará a entrada do usuário. Ex:

```go
var comando int
fmt.Scan(&comando)
fmt.Println("O endereço da minha variável comando é: ", &comando)
```

- Pode-se verificar o tipo de uma variável com a função `TypeOf` do pacote `reflect`.

```go
versao := 1.1
fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
```

## 3-Controlando o fluxo de input

- `if` em go tem 2 características principais:
    1. Não tem parênteses
    2. Sempre tem que ser uma expressão que retorna um booleano
    
```go
if comando == 1 {
fmt.Println("Monitorando")
} else if comando == 2 {
    fmt.Println("Exibindo Logs...")

} else if comando == 0 {
    fmt.Println("Saindo do Programa")
} else {
    fmt.Println("Não conheço esse comando")
}
```

- pode-se usar o `switch` no lugar do if
- `switch` no go não utiliza a palavra `break` para cada caso.

```go
switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não conheço este comando")
	}
```

## 4-Fazendo requisições para a web

- É comum utiliar o pacote `net/http` para fazer requisições web GO

- Pode-se utilizar a propriedade `StatusCode` da função GET do pacote `net/http` para verificar o status de retorno da chamada para a URL.

```go
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, " está com problemas!, Status Code: ", resp.StatusCode)
	}
}
```


## 4-As principais coleções do GO

- É comum utiliar o pacote `net/http` para fazer requisições web GO




