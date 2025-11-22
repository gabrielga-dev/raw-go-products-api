package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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
	time.Sleep(5 * time.Second)
	config := getMigrattoConfig()
	fmt.Println("Starting migrations with Migratto...")
	return migration_service.Migrate(config)
}

func getMigrattoConfig() dto.ConfigDTO {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		dbPort = 5432 // Default port
	}
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	return dto.ConfigDTO{
		DatabaseDriver:   "postgres",
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseName:     dbName,
		DatabaseUsername: dbUsername,
		DatabasePassword: dbPassword,
		Sslmode:          "disable",
		MigrationsDir:    "./migrations",
		Log:              true,
	}
}
