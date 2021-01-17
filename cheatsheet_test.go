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
}
