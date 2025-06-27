package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=**** dbname=**** password=**** host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
