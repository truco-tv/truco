package truco

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewCheatSheet(t *testing.T) {
	vira := &Card{Value: 7, Suit: Cup}
	deck := NewDeck()

	cards := NewCheatSheet(vira, deck)

	spew.Dump(cards)

	if !cards[SevenCoinRankedDeck].IsSevenCoin() {
		t.Errorf("Ranked Deck is not being generated correctly 7 Coins")
	}

	if !cards[SevenSwordRankedDeck].IsSevenSword() {
		t.Errorf("Ranked Deck is not being generated correctly 7 Swords")
	}

	if !cards[BastilloRankedDeck].IsBastillo() {
		t.Errorf("Ranked Deck is not being generated correctly bastillo")
	}

	if !cards[EspadillaRankedDeck].IsEspadilla() {
		t.Errorf("Ranked Deck is not being generated correctly espadilla")
	}

	if !cards[PericaRankedDeck].IsPerica() {
		t.Errorf("Ranked Deck is not being generated correctly perica")
	}

	if !cards[PericoRankedDeck].IsPerico() {
		t.Errorf("Ranked Deck is not being generated correctly perico")
	}
}
