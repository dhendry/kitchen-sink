package game

import (
	cr "crypto/rand"
	"math"
	"math/big"
	stupidRand "math/rand"

	"github.com/dhendry/kitchen-sink/solitaire/model"
)

var allCards [52]model.Card

// A non-deterministically initialized random - one of the silliest things so far about go
var rand stupidRand.Rand

func init() {
	// Setup the random number - why the hell does go use a deterministic random number generator by default???
	secureSeed, err := cr.Int(cr.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	rand = *stupidRand.New(stupidRand.NewSource(secureSeed.Int64()))

	// Create the pre-allocated set of all cards
	allCards = [52]model.Card{}
	idx := 0
	for suitIdx := range model.Suit_name {
		if suitIdx == int32(model.Suit_NO_SUIT) {
			continue
		}

		for rankIdx := range model.Rank_name {
			if rankIdx == int32(model.Rank_NO_RANK) {
				continue
			}

			allCards[idx] = model.Card{
				Suit: model.Suit(suitIdx),
				Rank: model.Rank(rankIdx),
			}
			idx++
		}
	}

	if idx != 52 {
		panic("Bad final index: " + string(idx))
	}
}

func NewShuffledDeck() (result []*model.Card) {
	result = make([]*model.Card, 52)
	for i, v := range rand.Perm(52) {
		// Copy the value of the card so the returned deck is mutable
		cardCopy := allCards[v]
		result[i] = &cardCopy
	}
	return
}

func NewGameState() (gs *model.GameState) {
	// Create the game state
	gs = &model.GameState{}

	// Initialize all the relevant piles
	gs.Piles = make([]*model.Pile, 0, len(model.PileType_values)-1)
	for _, pileType := range model.PileType_values {
		if pileType == model.PileType_NO_PILE {
			continue
		}
		gs.Piles = append(gs.Piles, &model.Pile{PileType: pileType})
	}

	srcDeck := NewShuffledDeck()
	srcDeckIdx := 0

	// Deal into the tableau piles:
	for tpIdx := 0; tpIdx < 7; tpIdx++ {
		// Note that this is taking advantage of the continuous range property of pile indexes
		// It could be a bit more efficient but we really dont care here
		pile := gs.GetPile(model.PileType(int(model.PileType_TABLEAU_0) + tpIdx))

		// Grab a range of the main deck for the specific tableau pile
		endDeckIdx := srcDeckIdx + 1 + tpIdx
		pile.Cards = append(pile.Cards, srcDeck[srcDeckIdx:endDeckIdx]...) // NOTE that this NOT a slice from the main deck but a copy (no shared array)
		pile.Cards[len(pile.Cards) - 1].FaceUp = true

		srcDeckIdx = endDeckIdx
	}

	// Remainder of the cards go into the deck
	// Note that the append here is to force a copy to a new array. Its not technically necessary but is done proactively
	// to avoid any sort of shared mutable backing arrays
	gs.GetPile(model.PileType_DECK).Cards = append([]*model.Card(nil), srcDeck[srcDeckIdx:]...)

	return
}
