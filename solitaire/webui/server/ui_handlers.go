package server

import (
	"html/template"
	"net/http"

)

func RegisterUiHandlers() {
	http.HandleFunc("/index.html", serveIndex)

	fs := http.FileServer(http.Dir("server/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func serveIndex(resp http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("server/index.html")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(resp, nil)

	log.Println("Served index to", req.RemoteAddr, req.UserAgent())
}
