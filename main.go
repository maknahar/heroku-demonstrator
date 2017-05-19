package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := GetRouter()
	log.Println("listening in port", port)
	http.ListenAndServe(":"+port, router)
}
