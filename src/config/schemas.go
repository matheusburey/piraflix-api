package config

import (
	"database/sql"
	"log"
)

func InitSchema(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS movies (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        image TEXT,
        type VARCHAR(50),
        code VARCHAR(50) UNIQUE,
        about TEXT,
        origin VARCHAR(50),
		duration INTEGER,
		release_date VARCHAR(10),
		trailer TEXT,
		description TEXT,
		language VARCHAR(50),
		views INTEGER,
		active BOOLEAN DEFAULT TRUE,
		rating FLOAT DEFAULT 0,
		categories TEXT[]
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}
}
