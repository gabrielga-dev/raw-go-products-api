package produto_controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	produto_model "github.org/gabrielga-dev/raw-go-products-api/model/produtoModel"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := produto_model.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func Novo(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NovoProduto", nil)
}

func Cadastrar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, precoErr := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, quantErr := strconv.Atoi(r.FormValue("quantidade"))
		if precoErr == nil && quantErr == nil {
			produto_model.CriarNovoProduto(nome, descricao, preco, quantidade)
			http.Redirect(w, r, "/", 301)
		} else {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			fmt.Printf("Error parsing form values: %v, %v\n", precoErr, quantErr)
			return
		}
	}
}

func Deleta(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		produto_model.DeletarProduto(idInt)
		http.Redirect(w, r, "/", 301)
	} else {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
}
