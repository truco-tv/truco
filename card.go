package truco

const (
	Cup   int = 1
	Coin  int = 2
	Stick int = 3
	Sword int = 4
)

type Card struct {
	Suit        int
	Value       int
	GameValue   int
	Placeholder bool
	IsVira      bool
	IsPerica    bool
	IsPerico    bool
}

func NewCard(value, suit int) *Card {
	card := &Card{Suit: suit, Value: value, GameValue: value}
	card.Placeholder = false
	card.IsVira = false
	card.IsPerica = false
	card.IsPerico = false

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

func (card *Card) SetPlaceHolder(value bool) *Card {
	card.Placeholder = value
	return card
}

func (card *Card) SetIsVira(value bool) *Card {
	card.IsVira = value
	return card
}

func (card *Card) SetIsPerica(value bool) *Card {
	card.IsPerica = value
	return card
}

func (card *Card) SetIsPerico(value bool) *Card {
	card.IsPerico = value
	return card
}
