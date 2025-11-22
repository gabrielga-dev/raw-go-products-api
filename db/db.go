package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		dbPort = 5432 // Default port
	}
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	conexao := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		dbUsername,
		dbName,
		dbPassword,
		dbHost,
		dbPort,
	)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return db
}
