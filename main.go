package main

import (
	"fmt"
	"net/http"

	dto "github.com/gabrielga-dev/migratto/dto"
	migration_service "github.com/gabrielga-dev/migratto/service/migration"
	"github.org/gabrielga-dev/raw-go-products-api/routes"
)

func main() {
	err := inicializaBancoDeDados()
	if err != nil {
		fmt.Println("Error during database initialization:", err)
		return
	}
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

func inicializaBancoDeDados() error {
	config := dto.ConfigDTO{
		DatabaseDriver:   "postgres",
		DatabaseHost:     "localhost",
		DatabasePort:     5432,
		DatabaseName:     "go-produtos",
		DatabaseUsername: "admin",
		DatabasePassword: "admin123",
		Sslmode:          "disable",
		MigrationsDir:    "./migrations",
		Log:              true,
	}

	fmt.Println("Starting migrations with Migratto...")
	return migration_service.Migrate(config)
}
