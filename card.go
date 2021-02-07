package truco

const (
	Cup               int = 1
	Coin              int = 2
	Stick             int = 3
	Sword             int = 4
	PericaEnvidoValue int = 9
	PericoEnvidoValue int = 10
)

type Card struct {
	Value       int
	Suit        int
	GameValue   int
	EnvidoValue int
	Placeholder bool
	IsVira      bool
	Perica      bool
	Perico      bool
}

func NewCard(value, suit int) *Card {
	card := &Card{Suit: suit, Value: value, GameValue: value}
	card.Placeholder = false
	card.IsVira = false
	card.Perica = false
	card.Perico = false
	card.EnvidoValue = value

	if value == 10 || value == 11 || value == 12 {
		card.EnvidoValue = 0
	}

	return card
}

func (card *Card) IsSevenCoin() bool {
	return card.Suit == Coin && card.Value == 7
}

func (card *Card) IsSevenSword() bool {
	return card.Suit == Sword && card.Value == 7
}

func (card *Card) IsBastillo() bool {
	return card.Suit == Stick && card.Value == 1
}

func (card *Card) IsEspadilla() bool {
	return card.Suit == Sword && card.Value == 1
}

func (card *Card) IsPerica() bool {
	return card.Perica == true
}

func (card *Card) IsPerico() bool {
	return card.Perico == true
}

func (card *Card) SetPlaceHolder(value bool) *Card {
	card.Placeholder = value
	return card
}

func (card *Card) SetIsVira(value bool) *Card {
	card.IsVira = value
	return card
}

func (card *Card) SetIsPerica(value bool) *Card {
	card.Perica = value
	return card
}

func (card *Card) SetIsPerico(value bool) *Card {
	card.Perico = value
	return card
}
