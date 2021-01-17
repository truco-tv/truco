package truco

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if !deck.Cards[SevenCoinDefaultDeck].IsSevenCoin() {
		t.Errorf("Seven of Coins is not in the right position %d %d", deck.Cards[SevenCoinDefaultDeck].Value, deck.Cards[SevenCoinDefaultDeck].Suit)
	}

	if !deck.Cards[SevenSwordDefaultDeck].IsSevenSword() {
		t.Errorf("Seven of Swords is not in the right position %d %d", deck.Cards[SevenSwordDefaultDeck].Value, deck.Cards[SevenSwordDefaultDeck].Suit)
	}

	if !deck.Cards[BastilloDefaultDeck].IsBastillo() {
		t.Errorf("Bastillo is not in the right position %d %d", deck.Cards[BastilloDefaultDeck].Value, deck.Cards[BastilloDefaultDeck].Suit)
	}

	if !deck.Cards[EspadillaDefaultDeck].IsEspadilla() {
		t.Errorf("Espadilla is not in the right position %d %d", deck.Cards[EspadillaDefaultDeck].Value, deck.Cards[EspadillaDefaultDeck].Suit)
	}
}

func TestFindCard(t *testing.T) {
	deck := NewDeck()
	card := &Card{Suit: Sword, Value: 1, GameValue: 1}

	cardFound, position := deck.FindCard(card)
	if cardFound == nil {
		t.Errorf("Espadilla not found in deck")
	}

	if position != EspadillaDefaultDeck {
		t.Errorf("Espadilla in wrong position default deck")
	}

}
