package helpers

import (
	_ "github.com/go-sql-driver/mysql" //exporting to use as DB config
	"log"
	"github.com/jmoiron/sqlx"
)
var host = "db4free.net"
var user = "awaaz_admin"
var password = "qwerty12345"
var port = "3306"
var database = "awaaz_case_store"
// GetDBInstance return the connection Instance of DB
func GetDBInstance () *sqlx.DB {
	db, err := sqlx.Open("mysql", user+ ":" + password +"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}