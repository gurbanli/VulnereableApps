package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Repository struct {
}

var Database *sql.DB

func (r *Repository) InitializeDatabase() {
	db, err := sql.Open("mysql", "root:Pass4DB@tcp(127.0.0.1:3306)/ingress")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	Database = db
}
