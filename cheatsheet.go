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

	newCards := deck.Cards[0:]
	newCards = append(newCards, sevenCoinRanked, sevenSwordRanked, bastilloRanked, espadillaRanked)

	return newCards
}
