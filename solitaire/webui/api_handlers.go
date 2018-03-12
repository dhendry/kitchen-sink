package main

import (
	"net/http"

	"github.com/dhendry/kitchen-sink/solitaire/game"
	"github.com/golang/protobuf/jsonpb"
)

func RegisterApiHandlers() {
	http.HandleFunc("/new", newGame)
}

func newGame(resp http.ResponseWriter, req *http.Request) {
	gs := game.NewGameState()

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := (&jsonpb.Marshaler{Indent: "    "}).Marshal(resp, gs)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
}