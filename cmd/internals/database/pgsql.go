package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


func Connection() *sql.DB{
	database, err := sql.Open("sqlite3","../../../database.db")


	if err != nil{
		panic(err)
	}

	return database
}