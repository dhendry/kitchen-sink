package game

import (
	"fmt"
	"testing"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestNewShuffledDeck_AllCardsPresent(t *testing.T) {
	deck := NewShuffledDeck(rand.Int63())
	cardsFound := make([]bool, 52)

	for _, v := range deck {
		cardsFound[(int(v.GetSuit())-1)*13+int(v.GetRank())-1] = true
	}

	for _, v := range cardsFound {
		assert.True(t, v, fmt.Sprintf("%+v", NewShuffledDeck(rand.Int63())))
	}
}

func TestNewShuffledDeck_DifferentShuffles(t *testing.T) {
	d1 := NewShuffledDeck(rand.Int63())
	d2 := NewShuffledDeck(rand.Int63())

	assert.NotEqual(t, d1, d2)

	// Double check the data structures are isolated properly - paranoia
	for _, v := range d1 {
		assert.False(t, v.GetFaceUp())
		v.FaceUp = true
		assert.True(t, v.GetFaceUp()) // Kinda silly but just making sure
	}
	for _, v := range d2 {
		assert.False(t, v.GetFaceUp())
	}
}

func TestNewGameState_Scratch(t *testing.T) {
	// gs := NewGameState()
	// bytes, e := proto.Marshal(gs)
	// assert.Nil(t, e)
	//
	// assert.Fail(t, strconv.Itoa(len(bytes)))
}

func TestApplyMove_CycleDeckToWaste(t *testing.T) {
	gs := NewGameState()
	initialState := proto.Clone(gs).(*model.GameState)

	initialDeckCount := len(gs.GetPile(model.PileType_DECK).Cards)
	for i := 0; i < initialDeckCount; i++ {
		// Invalid:
		{
			ApplyMoveForTest(t, false, gs, model.Move{
				SrcPile:  model.PileType_WASTE,
				DestPile: model.PileType_DECK,
				NumCards: 1,
			})
		}

		ApplyMoveForTest(t, true, gs, model.Move{
			SrcPile:  model.PileType_DECK,
			DestPile: model.PileType_WASTE,
			NumCards: 1,
		})
	}

	// Stuff which is invalid:
	{
		ApplyMoveForTest(t, false, gs, model.Move{
			SrcPile:  model.PileType_DECK,
			DestPile: model.PileType_WASTE,
			NumCards: 1,
		})
		ApplyMoveForTest(t, false, gs, model.Move{
			SrcPile:  model.PileType_WASTE,
			DestPile: model.PileType_DECK,
			NumCards: 1, // INVALID
		})
	}

	ApplyMoveForTest(t, true, gs, model.Move{
		SrcPile:  model.PileType_WASTE,
		DestPile: model.PileType_DECK,
		NumCards: int32(len(gs.GetPile(model.PileType_WASTE).Cards)),
	})

	_ = initialState
	// TODO: Check final state
}

func ApplyMoveForTest(t *testing.T, shouldSucceed bool, gs *model.GameState, move model.Move) {

	initialState := proto.Clone(gs).(*model.GameState)

	err := ApplyMove(gs, move)
	assert.NoError(t, ValidateGameState(gs)) // No matter what, the game state should still be valid

	assert.Equal(t, shouldSucceed, err == nil)

	if err == nil {
		assert.NotEqual(t, initialState, gs)
		assert.NotEqual(t, initialState.GetPile(move.GetSrcPile()), gs.GetPile(move.GetSrcPile()))
		assert.NotEqual(t, initialState.GetPile(move.GetDestPile()), gs.GetPile(move.GetDestPile()))
		assert.NotEqual(t, initialState.GetStateToken(), gs.GetStateToken())

		gsNilledPiles := proto.Clone(gs).(*model.GameState)
		gsNilledPiles.GetPile(move.GetSrcPile()).Cards = nil
		gsNilledPiles.GetPile(move.GetDestPile()).Cards = nil
		gsNilledPiles.StateToken = nil
		initialStateNilledPiles := proto.Clone(initialState).(*model.GameState)
		initialStateNilledPiles.GetPile(move.GetSrcPile()).Cards = nil
		initialStateNilledPiles.GetPile(move.GetDestPile()).Cards = nil
		initialStateNilledPiles.StateToken = nil

		assert.Equal(t, initialStateNilledPiles, gsNilledPiles)
	} else {
		// there was an error, nothing should have changed
		assert.Equal(t, initialState, gs)
	}
}
