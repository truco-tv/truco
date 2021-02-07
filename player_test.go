package truco

import (
	"testing"
)

func TestPlayerEnvido(t *testing.T) {
	player := NewPlayer()
	//vira := &Card{Value: 3, Suit: Coin}

	card1 := NewCard(5, Sword)
	card2 := NewCard(6, Sword)
	card3 := NewCard(3, Stick)

	player.AppendCard(card1)
	player.AppendCard(card2)
	player.AppendCard(card3)

	if player.HasEnvido() != true {
		t.Errorf("HasEnvido is wrong")
	}

	if player.EnvidoValue != 31 {
		t.Errorf("Envido Value is not right %d", player.EnvidoValue)
	}

}

func TestPlayerEnvido2(t *testing.T) {
	player := NewPlayer()
	vira := &Card{Value: 1, Suit: Stick}

	card1 := NewCard(1, Cup)
	card1 = player.TransformPerico(vira, card1)
	card1 = player.TransformPerica(vira, card1)

	card2 := NewCard(5, Stick)
	card2 = player.TransformPerico(vira, card2)
	card2 = player.TransformPerica(vira, card2)

	card3 := NewCard(6, Coin)
	card3 = player.TransformPerico(vira, card3)
	card3 = player.TransformPerica(vira, card3)

	player.AppendCard(card1)
	player.AppendCard(card2)
	player.AppendCard(card3)

	if player.HasEnvido() != false {
		t.Errorf("HasEnvido is wrong")
	}
}

func TestPlayerEnvido3(t *testing.T) {
	player := NewPlayer()
	vira := &Card{Value: 1, Suit: Stick}

	card1 := NewCard(11, Stick)
	card1 = player.TransformPerico(vira, card1)
	card1 = player.TransformPerica(vira, card1)

	card2 := NewCard(5, Stick)
	card2 = player.TransformPerico(vira, card2)
	card2 = player.TransformPerica(vira, card2)

	card3 := NewCard(6, Coin)
	card3 = player.TransformPerico(vira, card3)
	card3 = player.TransformPerica(vira, card3)

	player.AppendCard(card1)
	player.AppendCard(card2)
	player.AppendCard(card3)

	if player.HasEnvido() != true {
		t.Errorf("HasEnvido is wrong")
	}

	if player.EnvidoValue != 35 {
		t.Errorf("Envido Value is not right %d", player.EnvidoValue)
	}
}

func TestTransformPerico(t *testing.T) {
	player := NewPlayer()
	vira := &Card{Value: 1, Suit: Stick}
	card := &Card{Value: 11, Suit: Stick}

	card = player.TransformPerico(vira, card)
	if card.IsPerico() != true {
		t.Error("Wrong perico")
	}
}

func TestTransformPerico2(t *testing.T) {
	player := NewPlayer()
	vira := &Card{Value: 11, Suit: Stick}
	card := &Card{Value: 12, Suit: Stick}

	card = player.TransformPerico(vira, card)
	if card.IsPerico() != true {
		t.Error("Wrong perico")
	}
}

func TestTransformPerico3(t *testing.T) {
	player := NewPlayer()
	vira := &Card{Value: 7, Suit: Stick}
	card := &Card{Value: 3, Suit: Stick}

	card = player.TransformPerico(vira, card)
	if card.IsPerico() == true {
		t.Error("Wrong perico")
	}
}
