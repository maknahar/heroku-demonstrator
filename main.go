package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//DATABASE_URL := os.Getenv("DATABASE_URL")
	//EstablishDBConnection(DATABASE_URL)

	router := GetRouter()
	log.Println("listening in port", port)
	http.ListenAndServe(":"+port, router)
}

var conn *pgx.ConnPool

// Establish a connection to DB
func EstablishDBConnection(url string) error {
	if conn == nil {
		connConfig, err := pgx.ParseURI(url)
		if err != nil {
			log.Fatalf("Invalid Database URL, Err: %+v", err)
			return err
		}
		poolConfig := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 5}
		conn, err = pgx.NewConnPool(poolConfig)
		if err != nil {
			log.Fatalf("Unable to connect to Database, Err: %+v", err)
			return err
		}
		log.Println("Successfully established database connection")
	}
	return nil
}
