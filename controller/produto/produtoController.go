package produto_controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	produto_dao "github.org/gabrielga-dev/raw-go-products-api/dao/produto"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := produto_dao.BuscaTodosOsProdutos()
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
			produto_dao.CriarNovoProduto(nome, descricao, preco, quantidade)
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
		produto_dao.DeletarProduto(idInt)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	produto := produto_dao.BuscaProdutoPorID(idInt)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
		id, idErr := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, precoErr := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, quantErr := strconv.Atoi(r.FormValue("quantidade"))
		if idErr == nil && precoErr == nil && quantErr == nil {
			produto_dao.AtualizarProduto(id, nome, descricao, preco, quantidade)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			fmt.Printf("Error parsing form values: %v, %v, %v\n", idErr, precoErr, quantErr)
			return
		}
	}
}
