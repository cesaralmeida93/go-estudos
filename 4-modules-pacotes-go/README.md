# 4: Como funcionam os modules e pacotes no Go

## Estruturas de pastas e GOPATH

- o go define o GOPATH para a raiz do usuário/go
- /pkg(lugar onde o go guarda todas as dependências já instaladas em seus projetos)
- go literalmete dá um git clone na dependência quando você instala alguma coisa

- /src(Lugar onde o código deve ser guardado)

- /bin(Lugar onde o go guarda os binários que "go install" compila)

## GOPATH vs GOROOT
- $GOROOT: Onde vai estar o código da instalação do go(compilador)
- $GOPATH: Onde vai ficar a seu código, instalações e dependências

## O que são pacotes?
- São subpastas dentro do seu projeto(cada pasta é um pacote)

## O que são sa modules?
- Coleção de pacotes Go guardados em um arquivo com o nome de **go.mod**

## Estrutura do go.mod
- module<nome>
- go<versao-do-go>
- require<caminho-do-module><versão-do-pacote>(Pacotes que o seu projeto precisa para executar)
- replace<caminho-do-module> => /home/test/test-module(Pacotes que você "substitui" por alguma versão que você tenha local daquele mesmo pacote)
- exlude<caminho-do-module><versão-do-pacote>(Pacotes/versões que você não quer que executem no seu projeto)


```go
module github.com/HunnTeRUS/bookstore_users-api

go 1.16

require (
		github.com/Bookstore-GolangMS/bookstore_oauth-go v0.0.0-20210823130442-bd7a30307979
		github.com/Bookstore-GolangMS/bookstore_oauth-go v0.0.0-20210823130442-ddd480307979
)
```

## Estrutura do go.sum
- go mod init gera esse arquivo
- aponta os hashes dos commits, versões, e pacotes instalados no go.mod
- não é recomendado alterar esse arquivo

