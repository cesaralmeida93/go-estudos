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