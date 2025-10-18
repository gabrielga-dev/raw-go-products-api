package main

import (
	"net/http"

	"github.org/gabrielga-dev/raw-go-products-api/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
