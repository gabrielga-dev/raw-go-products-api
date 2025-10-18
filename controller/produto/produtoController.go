package produto_controller

import (
	"net/http"
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
