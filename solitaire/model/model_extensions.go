package model

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
