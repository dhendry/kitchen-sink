package db

import (
	"github.com/dhendry/kitchen-sink/solitaire/model"
)

var SolitaireDataAccess SolitaireDataAccess

func GetSolitaireDataAccess() SolitaireDataAccess {
	return mysql.Db
}


type SolitaireDataAccess interface {
	SaveNewGameState(gs *model.GameState) error
}
