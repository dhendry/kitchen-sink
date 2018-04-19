package server

import (
	"net/http"

	"github.com/dhendry/kitchen-sink/solitaire/db"
	"github.com/dhendry/kitchen-sink/solitaire/game"
	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/golang/protobuf/jsonpb"
	"golang.org/x/net/context"
)

func RegisterApiHandlers() {
	http.HandleFunc("/api/v1/new", newGame)


}

func newGame(resp http.ResponseWriter, req *http.Request) {
	gs := game.NewGameState()

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := (&jsonpb.Marshaler{Indent: "    "}).Marshal(resp, gs)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Served new game to", req.RemoteAddr, req.UserAgent())
}

type PlayServiceImpl struct {
}

func (ps *PlayServiceImpl) NewGame(ctx context.Context, ngr *NewGameRequest) (gs *model.GameState, err error){
	log.Info(ctx)

	gs = game.NewGameState()
	db.GetSolitaireDataAccess().SaveNewGameState(gs)
	return
}
