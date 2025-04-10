package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	database_url := os.Getenv("DATABASE_URL")
	println(database_url)
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
