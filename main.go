package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Servidor ouvindo em http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", novaApi()))

}
