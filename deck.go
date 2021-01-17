package truco

const (
	SevenCoinDefaultDeck  int = 13
	SevenSwordDefaultDeck int = 15
	BastilloDefaultDeck   int = 30
	EspadillaDefaultDeck  int = 31

	SevenCoinRankedDeck  int = 40
	SevenSwordRankedDeck int = 41
	BastilloRankedDeck   int = 42
	EspadillaRankedDeck  int = 43
)

type Deck struct {
	Cards [40]*Card
}

func NewDeck() *Deck {
	deck := &Deck{}

	firstRange := []int{4, 5, 6, 7}
	secondRange := []int{10, 11, 12}
	thirdRange := []int{1, 2, 3}
	suitRange := []int{Cup, Coin, Stick, Sword}

	deckOrderRange := [][]int{firstRange, secondRange, thirdRange}
	cardIndex := 0
	for _, deckRange := range deckOrderRange { //ranges 4-7 10-12 1-3 | i
		for _, value := range deckRange { //ranges 4,5,6,7 | 10,11,12 | 1,2,3 |
			for _, suit := range suitRange { // 1,2,3,4 (suits)
				card := NewCard(value, suit)
				deck.Cards[cardIndex] = card
				cardIndex++
			}
		}
	}

	return deck
}

func (deck *Deck) FindCard(card *Card) (*Card, int) {
	for position, deckCard := range deck.Cards {
		if deckCard.Value == card.Value && deckCard.Suit == card.Suit {
			return deckCard, position
		}
	}

	return nil, 0
}
