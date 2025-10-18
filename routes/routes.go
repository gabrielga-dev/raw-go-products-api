package routes

import (
	"net/http"

	produto_controller "github.org/gabrielga-dev/raw-go-products-api/controller/produto"
)

func CarregaRotas() {
	http.HandleFunc("/", produto_controller.Index)
	http.HandleFunc("/produto/novo", produto_controller.Novo)
}
