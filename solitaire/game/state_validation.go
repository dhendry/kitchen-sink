package game

import (
	"github.com/dhendry/kitchen-sink/solitaire/model"
)

// This is considered a fatal error within the context of a particular game and is used to indicate the game is no longer
// playable
type ValidationError struct {
	msg string
}

func (ve *ValidationError) Error() string {
	return ve.msg
}

//func ValidateGameState(gs *model.GameState) (isValid bool, msg string) {
func ValidateGameState(gs *model.GameState) error {
	if gs == nil {
		return &ValidationError{msg: "GameState is nil"}
	}

	if e := correctPilesAndPileOrder(gs); e != nil {
		return e
	}

	if e := containsExactlyOneDeck(gs); e != nil {
		return e
	}

	return nil
}

func correctPilesAndPileOrder(gs *model.GameState) error {
	if len(gs.Piles) != len(model.PileType_values)-1 {
		return &ValidationError{msg: "wrong number of piles"}
	}

	for i, v := range model.PileType_values[1:] {
		if gs.Piles[i].GetPileType() != v {
			return &ValidationError{msg: "pile out of order, expected " + v.String()}
		}
	}

	return nil
}

func containsExactlyOneDeck(gs *model.GameState) error {
	cardsFound := make([]bool, 52)

	for _, pile := range gs.GetPiles() {
		for _, card := range pile.GetCards() {
			if card.GetRank() == model.Rank_NO_RANK || card.GetSuit() == model.Suit_NO_SUIT {
				return &ValidationError{msg: "card has invalid rank or suit"}
			}

			cardIdx := (int(card.GetSuit())-1)*13 + int(card.GetRank()) - 1
			if cardsFound[cardIdx] {
				// Already seem this card
				return &ValidationError{msg: "duplicate card " + card.String()}
			}
			cardsFound[cardIdx] = true
		}
	}

	// Check that all have been found
	for _, v := range cardsFound {
		if !v {
			return &ValidationError{msg: "missing card"}
		}
	}

	return nil
}
