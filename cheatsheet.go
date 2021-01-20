package truco

func NewCheatSheet(vira *Card, deck *Deck) []*Card {
	viraFound, _ := deck.FindCard(vira)

	sevenCoin := &Card{Suit: Coin, Value: 7}
	sevenSword := &Card{Suit: Sword, Value: 7}
	bastillo := &Card{Suit: Stick, Value: 1}
	espadilla := &Card{Suit: Sword, Value: 1}

	viraFound.SetPlaceHolder(true).SetIsVira(true)

	sevenCoin.SetPlaceHolder(true)
	sevenSword.SetPlaceHolder(true)
	bastillo.SetPlaceHolder(true)
	espadilla.SetPlaceHolder(true)

	sevenCoinRanked, _ := deck.FindCard(sevenCoin)
	deck.Cards[SevenCoinDefaultDeck] = sevenCoin

	sevenSwordRanked, _ := deck.FindCard(sevenSword)
	deck.Cards[SevenSwordDefaultDeck] = sevenSword

	bastilloRanked, _ := deck.FindCard(bastillo)
	deck.Cards[BastilloDefaultDeck] = bastillo

	espadillaRanked, _ := deck.FindCard(espadilla)
	deck.Cards[EspadillaDefaultDeck] = espadilla

	pericaRanked, perica, pericaPositionInDeck := Perica(vira, deck)
	pericaRanked.SetIsPerica(true)
	perica.SetPlaceHolder(true)
	deck.Cards[pericaPositionInDeck] = perica

	pericoRanked, perico, pericoPositionInDeck := Perico(vira, deck)
	pericoRanked.SetIsPerico(true)
	perico.SetPlaceHolder(true)
	deck.Cards[pericoPositionInDeck] = perico

	newCards := deck.Cards[0:]
	newCards = append(newCards, sevenCoinRanked, sevenSwordRanked, bastilloRanked, espadillaRanked, pericaRanked, pericoRanked)

	return newCards
}

func Perica(vira *Card, deck *Deck) (*Card, *Card, int) {
	if vira.Value == 10 {
		perica := NewCard(12, vira.Suit)
		pericaOnDeck, positionInDeck := deck.FindCard(perica)
		return pericaOnDeck, perica, positionInDeck
	}

	perica := NewCard(10, vira.Suit)
	pericaOnDeck, positionInDeck := deck.FindCard(perica)
	return pericaOnDeck, perica, positionInDeck
}

func Perico(vira *Card, deck *Deck) (*Card, *Card, int) {
	if vira.Value == 11 {
		perico := NewCard(12, vira.Suit)
		pericoOnDeck, positionInDeck := deck.FindCard(perico)
		return pericoOnDeck, perico, positionInDeck
	}

	perico := NewCard(11, vira.Suit)
	pericoOnDeck, positionInDeck := deck.FindCard(perico)
	return pericoOnDeck, perico, positionInDeck
}
