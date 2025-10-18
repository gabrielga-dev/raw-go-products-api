package produto_model

import (
	"fmt"

	"github.org/gabrielga-dev/raw-go-products-api/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectProdutos, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM produto")
	produtos := []Produto{}
	if err != nil {
		fmt.Println(err)
	} else {
		p := Produto{}
		for selectProdutos.Next() {
			var id, quantidade int
			var nome, descricao string
			var preco float64
			err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err != nil {
				fmt.Println(err)
			}
			p.Id = id
			p.Nome = nome
			p.Descricao = descricao
			p.Preco = preco
			p.Quantidade = quantidade
			produtos = append(produtos, p)
		}
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insertProduto, err := db.Prepare("INSERT INTO produto(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		fmt.Println(err)
	}
	insertProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id int) {
	db := db.ConectaComBancoDeDados()
	deleteProduto, err := db.Prepare("DELETE FROM produto WHERE id=$1")
	if err != nil {
		fmt.Println(err)
	}
	deleteProduto.Exec(id)
	defer db.Close()
}
