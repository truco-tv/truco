package truco

type Player struct {
	Cards       []*Card
	Envido      bool
	EnvidoValue int
	Flor        bool
	FlorValue   int
	HasPerico   bool
	HasPerica   bool
}

func NewPlayer() *Player {
	player := &Player{}
	player.EnvidoValue = 0
	player.FlorValue = 0

	return player
}

func (player *Player) AppendCard(card *Card) {
	if card.IsPerica() {
		player.HasPerica = true
		player.AddEnvido(PericaEnvidoValue)
	}

	if card.IsPerico() {
		player.HasPerico = true
		player.AddEnvido(PericoEnvidoValue)
	}

	newCards := append(player.Cards, card)
	player.Cards = newCards
}

func (player *Player) HasEnvido() bool {
	player.EnvidoValue = 0
	typesOfSuits := make(map[int]int)
	envidoBySuits := make(map[int]int)

	for _, card := range player.Cards {
		typesOfSuits[card.Suit] += 1
		envidoBySuits[card.Suit] += card.EnvidoValue + 10
	}

	var hasEnvido bool

	for suit, value := range typesOfSuits {
		hasEnvido = false
		if value > 1 || (value > 0 && player.HasPerica || player.HasPerico) {
			hasEnvido = true
			player.AddEnvido(envidoBySuits[suit])
			break
		}
	}

	player.Envido = hasEnvido

	return hasEnvido
}

func (player *Player) HasFlor() {

}

func (player *Player) HasReservada() {

}

func (player *Player) AddEnvido(value int) {
	player.EnvidoValue += value
}

func (player *Player) AddFlor(value int) {
	player.FlorValue += value
}

func (player *Player) TransformPerica(vira, card *Card) *Card {
	if vira.Suit != card.Suit {
		return card
	}

	if card.Value != 10 && card.Value != 11 && card.Value != 12 {
		return card
	}

	if vira.Value != 10 && card.Value == 10 {
		card.SetIsPerica(true)
		return card
	}

	if vira.Value == 10 && card.Value == 12 {
		card.SetIsPerica(true)
		return card
	}

	return card
}

func (player *Player) TransformPerico(vira, card *Card) *Card {
	if vira.Suit != card.Suit {
		return card
	}

	if card.Value != 10 && card.Value != 11 && card.Value != 12 {
		return card
	}

	if vira.Value != 11 && card.Value == 11 {
		card.SetIsPerico(true)
		return card
	}

	if vira.Value == 11 && card.Value == 12 {
		card.SetIsPerico(true)
		return card
	}

	return card
}
