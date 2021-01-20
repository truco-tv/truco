package truco

type Player struct {
	Cards []*Card
}

func NewPlayer() *Player {
	player := &Player{}

	return player
}

func (player *Player) appendCard(card *Card) {
	newCards := append(player.Cards, card)
	player.Cards = newCards
}

func (player *Player) hasEnvido() {
	for _, card := range player.Cards {

	}
}

func (player *Player) hasFlor() {

}

func (player *Player) hasReservada() {

}

func (player *Player) envidoValue() {

}

func (player *Player) florValue() {

}
