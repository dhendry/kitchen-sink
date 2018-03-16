package main

import (
	"log"
	"net/http"

	"github.com/dhendry/kitchen-sink/solitaire/webui/server"
)

func main() {
	server.RegisterApiHandlers()
	server.RegisterUiHandlers()

	// TODO: Configurable port
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("Exiting")
}
