package game

import (
	cr "crypto/rand"
	"errors"
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

func NewShuffledDeck(seed int64) (result []*model.Card) {
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
	gs = &model.GameState{
		GameId: &model.GameState_Id{
			Seed:   rand.Int63(),
			Nonce: rand.Int63(),
		},
		StateToken: &model.GameState_StateToken{},
	}

	// Initialize all the relevant piles
	gs.Piles = make([]*model.Pile, 0, len(model.PileType_values)-1)
	for _, pileType := range model.PileType_values {
		if pileType == model.PileType_NO_PILE {
			continue
		}
		gs.Piles = append(gs.Piles, &model.Pile{PileType: pileType})
	}

	srcDeck := NewShuffledDeck(gs.GetGameId().GetSeed())
	srcDeckIdx := 0

	// Deal into the tableau piles:
	for tpIdx := 0; tpIdx < 7; tpIdx++ {
		// Note that this is taking advantage of the continuous range property of pile indexes
		// It could be a bit more efficient but we really dont care here
		pile := gs.GetPile(model.PileType(int(model.PileType_TABLEAU_0) + tpIdx))

		// Grab a range of the main deck for the specific tableau pile
		endDeckIdx := srcDeckIdx + 1 + tpIdx
		pile.Cards = append(pile.Cards, srcDeck[srcDeckIdx:endDeckIdx]...) // NOTE that this NOT a slice from the main deck but a copy (no shared array)
		pile.Cards[len(pile.Cards)-1].FaceUp = true

		srcDeckIdx = endDeckIdx
	}

	// Remainder of the cards go into the deck
	// Note that the append here is to force a copy to a new array. Its not technically necessary but is done proactively
	// to avoid any sort of shared mutable backing arrays
	gs.GetPile(model.PileType_DECK).Cards = append([]*model.Card(nil), srcDeck[srcDeckIdx:]...)

	return
}

// This function assumes that the provided game state is initially valid. It does NOT validate the game state.
func ApplyMove(gs *model.GameState, move model.Move) (error) {
	if gs == nil || gs.StateToken == nil || gs.Piles == nil {
		// Super basic checks on the game state - note that the gs is assumed to be valid when passed to this function
		return &ValidationError{msg: "GameState is invalid"}
	}

	if move.GetSrcPile() == model.PileType_NO_PILE || move.GetDestPile() == model.PileType_NO_PILE {
		return errors.New("invalid move pile")
	}

	if move.GetSrcPile() == move.GetDestPile() {
		return errors.New("src and dest piles are the same")
	}

	// TODO: Have a move limit somewhere
	// TODO: Consider how to communicate an error state where the game state is just totally broken (ValidationError?)

	destPile := gs.GetPile(move.GetDestPile())
	srcPile := gs.GetPile(move.GetSrcPile())

	if move.GetNumCards() <= 0 || int(move.GetNumCards()) > len(srcPile.GetCards()) {
		return errors.New("invalid number of cards to move")
	}

	//
	// General validation over, lets start processing the move itself
	//

	// This if-else-if is setup based on destination
	if model.PileType_WASTE  == move.GetDestPile() {
		if move.GetSrcPile() != model.PileType_DECK {
			// Source must be the deck
			return errors.New("invalid move")
		}
		if move.GetNumCards() != 1 {
			// Can only move one card at a time from the deck to the waste.
			// Note that this is therefore hardcoded for "Klondike deal 1"
			return errors.New("invalid move")
		}

		// Make the move
		destPile.Cards = append(destPile.Cards, srcPile.Cards[len(srcPile.Cards) - 1])
		srcPile.Cards = srcPile.Cards[:len(srcPile.Cards) - 1]

		// Update face states:
		destPile.Cards[len(destPile.Cards) - 1].FaceUp = true
		// TODO: Consider if its worth changing face states for other buried cards in the dest pile
	} else if model.PileType_DECK == move.GetDestPile() {
		if move.GetSrcPile() != model.PileType_WASTE {
			// Source must be the waste pile
			return errors.New("invalid move")
		}
		if int(move.GetNumCards()) != len(srcPile.GetCards()) {
			// We must be moving all the waste cards back to the deck
			return errors.New("invalid move")
		}
		if len(destPile.GetCards()) != 0 {
			// The deck must also be empty
			return errors.New("invalid move")
		}

		// Make the move:
		destPile.Cards = make([]*model.Card, move.GetNumCards())
		for i, v := range srcPile.GetCards() {
			v.FaceUp = false
			destPile.Cards[len(destPile.Cards)-i-1] = v;
		}
		srcPile.Cards = nil // TODO: Confirm that setting to nil is the right thing to do here - I THINK so
	} else {
		return errors.New("not implemented")
	}

	gs.StateToken.MoveNum++

	return nil
}
