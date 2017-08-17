package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	port := ":8000"

	log.Fatal(http.ListenAndServe(port, router))
}

