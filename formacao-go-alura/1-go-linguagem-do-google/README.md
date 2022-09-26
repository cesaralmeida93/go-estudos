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


## 5-As principais coleções do GO

- **array**: Tamanho fixo

```go
var sites [4]string

sites[0] = "https://random-status-code.herokuapp.com/"
sites[1] = "https://www.alura.com.br"
sites[2] = "https://www.caelum.com.br"

fmt.Println(sites)
```

função `len(nome_array)` informa quantos items tem no array/slice, não necessariamente o tamanho do array/slice
para ver a capacidade do array/slice, deve-se utilizar a função `cap(nome_array)`

- **slice**: Tamanho dinâmico
- sempre que se dá um `append` em um slice, a capacidade do mesmo dobra de tamanho(se já estiver na última posição)

```go
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "items")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "ïtems")

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "items")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "items")
}
```

- ao se fazer um `for`, pode-se utilizar a palavra reservada `range`

```go
func main() {
	var frutas [4]string
	frutas[0] = "Abacaxi"
	frutas[1] = "Laranja"
	frutas[2] = "Morango"

	for k, v := range frutas {
		fmt.Println(frutas[k])
		fmt.Println(v)
	}

	fmt.Println(frutas[3])
}
```

## 6-Lendo arquivos

- **os.Open()**: retorna ponteiro

```go
	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(arquivo)
```

- **ioutil.ReadFile()**: retorna array de bytes(que pode ser convertido para string)

```go
arquivo, err := ioutil.ReadFile("sites.txt")

if err != nil {
	fmt.Println("Ocorreu um erro:", err)
}

fmt.Println(string(arquivo))
return sites
```

- **bufio**: permite ler a linha inteira do arquivo

```go
func lerSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err == nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	fmt.Println(sites)
	return sites

}
```

## 7-Escrevendo Arquivos

- **os.OpenFile**: permite passar flags e permissões para abrir o arquivo

```go
arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

if err != nil {
	fmt.Println("Ocorreu um erro:", err)
}

fmt.Println(arquivo)
```

- **os.File.WriteString()**: permite escrever no arquivo

```go
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
```