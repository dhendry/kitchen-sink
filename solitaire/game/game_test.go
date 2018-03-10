package game

import (
	"fmt"
	"testing"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/stretchr/testify/assert"
)

func TestNewShuffledDeck_AllCardsPresent(t *testing.T) {
	//assert.Fail(t, fmt.Sprintf("%+v",NewShuffledDeck()))

	deck := NewShuffledDeck()
	cardsFound := make([]bool, 52)

	for _, v := range deck {
		cardsFound[(int(v.GetSuit())-1)*13+int(v.GetRank())-1] = true
	}

	for _, v := range cardsFound {
		assert.True(t, v, fmt.Sprintf("%+v", NewShuffledDeck()))
	}
}

func TestNewShuffledDeck_DifferentShuffles(t *testing.T) {
	d1 := NewShuffledDeck()
	d2 := NewShuffledDeck()

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
