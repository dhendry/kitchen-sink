package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuit_IsBlack(t *testing.T) {
	assert.False(t, Suit_NO_SUIT.IsBlack())
	assert.True(t, Suit_CLUBS.IsBlack())
	assert.False(t, Suit_DIAMONDS.IsBlack())
	assert.False(t, Suit_HEARTS.IsBlack())
	assert.True(t, Suit_SPADES.IsBlack())
}

func TestSuit_IsRed(t *testing.T) {
	assert.False(t, Suit_NO_SUIT.IsRed())
	assert.False(t, Suit_CLUBS.IsRed())
	assert.True(t, Suit_DIAMONDS.IsRed())
	assert.True(t, Suit_HEARTS.IsRed())
	assert.False(t, Suit_SPADES.IsRed())
}

func TestSuit_IsAlternateColor(t *testing.T) {
	assert.False(t, Suit_NO_SUIT.IsAlternateColor(Suit_NO_SUIT))
	assert.False(t, Suit_NO_SUIT.IsAlternateColor(Suit_CLUBS))
	assert.False(t, Suit_NO_SUIT.IsAlternateColor(Suit_HEARTS))
	assert.False(t, Suit_DIAMONDS.IsAlternateColor(Suit_NO_SUIT))
	assert.False(t, Suit_SPADES.IsAlternateColor(Suit_NO_SUIT))

	assert.False(t, Suit_SPADES.IsAlternateColor(Suit_CLUBS))
	assert.False(t, Suit_DIAMONDS.IsAlternateColor(Suit_HEARTS))

	assert.True(t, Suit_DIAMONDS.IsAlternateColor(Suit_CLUBS))
	assert.True(t, Suit_HEARTS.IsAlternateColor(Suit_SPADES))
	assert.True(t, Suit_CLUBS.IsAlternateColor(Suit_HEARTS))
	assert.True(t, Suit_SPADES.IsAlternateColor(Suit_DIAMONDS))
}

func TestGetPileTypeValues_Sorted(t *testing.T) {
	fmt.Printf("%v\n", PileType_values)
	indexes := make(map[PileType]int, len(PileType_values))
	for i, v := range PileType_values {
		indexes[v] = i
	}

	// Break these into chunks that actually matter (kind of but not really).
	// Just do some random checks.

	// Deck and waste:
	assert.True(t, indexes[PileType_DECK] < indexes[PileType_WASTE])

	// Tableau:
	assert.True(t, indexes[PileType_TABLEAU_0] < indexes[PileType_TABLEAU_2])
	assert.True(t, indexes[PileType_TABLEAU_2] < indexes[PileType_TABLEAU_4])
	assert.True(t, indexes[PileType_TABLEAU_3] < indexes[PileType_TABLEAU_6])
	assert.True(t, indexes[PileType_TABLEAU_0] < indexes[PileType_TABLEAU_6])

	// Foundation
	assert.True(t, indexes[PileType_FOUNDATION_0] < indexes[PileType_FOUNDATION_1])
	assert.True(t, indexes[PileType_FOUNDATION_2] < indexes[PileType_FOUNDATION_3])
	assert.True(t, indexes[PileType_FOUNDATION_0] < indexes[PileType_FOUNDATION_3])
}
