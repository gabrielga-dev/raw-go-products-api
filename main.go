package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.org/gabrielga-dev/raw-go-products-api/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
