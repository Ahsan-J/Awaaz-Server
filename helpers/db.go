package helpers

import (
	_ "github.com/go-sql-driver/mysql" //exporting to use as DB config
	"log"
	"github.com/jmoiron/sqlx"
)

// GetDBInstance return the connection Instance of DB
func GetDBInstance () *sqlx.DB {
	db, err := sqlx.Open("mysql", "root@tcp(localhost:3306)/awaaz_case_store")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}