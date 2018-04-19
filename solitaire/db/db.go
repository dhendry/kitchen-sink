package db

import (
	"github.com/dhendry/kitchen-sink/solitaire/model"
)

var Db SolitaireDataAccess //= &mySqlSolitaireDataAccess{}

type SolitaireDataAccess interface {
	SaveNewGameState(gs *model.GameState) error
}
