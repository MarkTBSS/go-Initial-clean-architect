package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func DatabaseConnect() *sqlx.DB {
	connString := "postgres://postgres:Pass1234@localhost:5432/postgres"
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	db.DB.SetMaxOpenConns(10)
	return db
}
