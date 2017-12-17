package main

import (
	"log"
	"net/http"
	"github.com/mapu77/AD-Labs/6-go-api/routes"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}