package model

import "sort"

// Sorted by tag number
var PileType_values []PileType

func init() {
	PileType_values = make([]PileType, len(PileType_value))
	i := 0
	for _, v := range PileType_value {
		PileType_values[i] = PileType(v)
		i++
	}
	sort.Slice(PileType_values, func(i, j int) bool { return int32(PileType_values[i]) < int32(PileType_values[j]) })
}

func (s Suit) IsBlack() bool {
	switch s {
	case Suit_CLUBS, Suit_SPADES:
		return true
	default:
		return false
	}
}

func (s Suit) IsRed() bool {
	switch s {
	case Suit_DIAMONDS, Suit_HEARTS:
		return true
	default:
		return false
	}
}

func (s1 Suit) IsAlternateColor(s2 Suit) bool {
	// Note that this entire block is only using true conditions to work properly with, for example, NO_SUIT
	if s1.IsRed() {
		return s2.IsBlack()
	}

	if s1.IsBlack() {
		return s2.IsRed()
	}

	return false
}

func (gs *GameState) GetPile(pt PileType) *Pile {
	if gs == nil || pt == PileType_NO_PILE {
		return nil
	}

	for _, v := range gs.GetPiles() {
		if v.GetPileType() == pt {
			return v
		}
	}

	return nil
}

func (pt PileType) IsFoundation() bool {
	return pt >= PileType_FOUNDATION_0 && pt <= PileType_FOUNDATION_3
}

func (pt PileType) IsTableau() bool {
	return pt >= PileType_TABLEAU_0 && pt <= PileType_TABLEAU_6
}
