package main

import (
	"log"
	"net/http"

	"rest_api_prueba/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
