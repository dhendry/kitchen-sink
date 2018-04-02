package db

import "github.com/dhendry/kitchen-sink/solitaire/model"


func GetSolitaireDataAccess() SolitaireDataAccess {

	return nil
}


type SolitaireDataAccess interface {
	SaveNewGameState(gs *model.GameState) error
}
