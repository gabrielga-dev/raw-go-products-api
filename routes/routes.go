package routes

import (
	"net/http"

	produto_controller "github.org/gabrielga-dev/raw-go-products-api/controller/produto"
)

func CarregaRotas() {
	http.HandleFunc("/", produto_controller.Index)
	http.HandleFunc("/produto/novo", produto_controller.Novo)
	http.HandleFunc("/produto/insert", produto_controller.Cadastrar)
	http.HandleFunc("/produto/delete", produto_controller.Deleta)
	http.HandleFunc("/produto/edit", produto_controller.Edit)
	http.HandleFunc("/produto/update", produto_controller.Update)
}
