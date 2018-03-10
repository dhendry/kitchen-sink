package model

import (
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
