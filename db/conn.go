package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // Quando usado em localhost
	// host     = "go_db"  	  // Quando usado em imagem docker
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + 
							"password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to: " + dbname)

	return db, nil
}