# GO: Crie uma aplicação web

## 1-Servidor e struct de produtos
- para criar um servidor deve-se utilizar o pacote `http`, é ele que comtém as funções de rota e servidor

```go
package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
```

- go permite criar código html de maneira dinâmica, no exemplo abaixo o html renderiza as informaçoes contidas no servidor go

```html
{{define "Index"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
        crossorigin="anonymous">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    <title>Alura loja</title>
</head>
<nav class="navbar navbar-light bg-light mb-4">
    <a class="navbar-brand" href="/">Alura Loja</a>
</nav>

<body>
    <div class="container">
        <section class="card">
            <div>
                <table class="table table-striped table-hover mb-0">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Descrição</th>
                            <th>Preço</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .}}
                        <tr>
                            <td>{{.Nome}}</td>
                            <td>{{.Descricao}}</td>
                            <td>{{.Preco}}</td>
                            <td>{{.Quantidade}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </section>

    </div>
</body>

</html>
{{end}}
```

## 2-Conectando com banco de dados

[como intanciar postgres com docker](https://cursos.alura.com.br/forum/topico-dica-postgres-com-docker-204423)

- no pgadmin, criar a tabela de produtos e inserir valores

```sql
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into produtos(nome, descricao, preco, quantidade) values 
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 99, 5);
```

- através do terminal, instalar a lib postgres do go

```bash
go get github.com/lib/pq
```

- criar uma função que conecta com o banco de dados

```go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"0.0.0.0", 5051, "postgres", "qwerty", "postgres")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
}
```

- reescrever a função `index` para coenctar com o banco, selecionar todas as linhas da tabela `produtos` e adicionar a lista de `produtos`

```go
func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
```
