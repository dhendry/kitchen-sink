package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dhendry/kitchen-sink/solitaire/webui/server"
)

func main() {
	server.RegisterApiHandlers()
	//server.RegisterUiHandlers()

	frontendUrl, e := url.Parse("http://localhost:8080/")
	if e != nil {
		panic(e)
	}

	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)
	http.Handle("/", proxy)

	// TODO: Configurable port
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))


	log.Println("Exiting")
}
