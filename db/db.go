package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarComBancoDados() *sql.DB {
	conexao := "user=postgres dbname=presentes password=1609 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
