package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)


const (

)

func Connection() *sql.DB{
	database, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=tweeter sslmode=disable")


	if err != nil{
		panic(err)
	}

	err = database.Ping()

	if err != nil {
		panic(err)
	}

	return database
}