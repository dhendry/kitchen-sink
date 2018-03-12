package game

import (
	"testing"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateGameState_Unmodified(t *testing.T) {
	gs := NewGameState()
	assert.NoError(t, ValidateGameState(gs))
}

func TestValidateGameState_OutOfOrderPile(t *testing.T) {
	gs := NewGameState()
	gs.Piles[3], gs.Piles[5] = gs.Piles[5], gs.Piles[3]

	assert.EqualError(t, ValidateGameState(gs), "pile out of order, expected TABLEAU_1")
}

func TestValidateGameState_NilPile(t *testing.T) {
	gs := NewGameState()
	gs.Piles[3] = nil

	assert.EqualError(t, ValidateGameState(gs), "pile out of order, expected TABLEAU_1")
}

func TestValidateGameState_NewBadCard(t *testing.T) {
	gs := NewGameState()

	pile := gs.GetPile(model.PileType_FOUNDATION_1)
	pile.Cards = append(pile.Cards, &model.Card{})

	assert.EqualError(t, ValidateGameState(gs), "card has invalid rank or suit")
}

func TestValidateGameState_ExtraCard(t *testing.T) {
	gs := NewGameState()

	pile := gs.GetPile(model.PileType_FOUNDATION_3) // Arbitrary
	pile.Cards = append(pile.Cards, &model.Card{Rank: model.Rank_ACE, Suit: model.Suit_HEARTS})

	assert.EqualError(t, ValidateGameState(gs), "duplicate card rank:ACE suit:HEARTS ")
}

func TestValidateGameState_MissingCard(t *testing.T) {
	gs := NewGameState()

	pile := gs.GetPile(model.PileType_TABLEAU_6) // Arbitrary, but one that has cards to choose from
	pile.Cards = pile.Cards[1:]

	assert.EqualError(t, ValidateGameState(gs), "missing card")
}
