package db

import (
	"database/sql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=loja password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
