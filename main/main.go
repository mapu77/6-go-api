package main

import (
	"log"
	"net/http"
	"github.com/mapu77/AD-Labs/6-go-api/routes"
	"os"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
