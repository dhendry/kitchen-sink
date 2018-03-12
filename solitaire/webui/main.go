package main

import (
	"log"
	"net/http"
)

func main() {
	RegisterApiHandlers()

	// TODO: Configurable port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
