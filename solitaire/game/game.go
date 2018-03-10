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

func NewShuffledDeck() (result []model.Card) {
	result = make([]model.Card, 52, 52)
	for i, v := range rand.Perm(52) {
		result[i] = allCards[v]
	}
	return
}
