package game

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMakeGame(t *testing.T) {
	p := []Player{RandomPlayer(), RandomPlayer()}
	g := MakeGame(p)

	LogGameState(&g, t)
	if len(g.Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(g.Players))
	}
	if len(g.DrawDeck) != 108 {
		t.Errorf("Expected 108 cards in draw deck, got %d", len(g.DrawDeck))
	}
	if len(g.Discard) != 0 {
		t.Errorf("Expected 0 cards in discard, got %d", len(g.Discard))
	}
}

// Generates a player with a random deck (only for testing)
// Player name is "Test Player" & hand will consist only of cards from 0-9 (no wilds, skips, etc)
func RandomPlayer() Player {
	hand := []Card{}
	for i := 0; i < 7; i++ {
		hand = append(hand, Card{Color: CardColor(rand.Intn(4)), Value: CardType(rand.Intn(9))})
	}
	p := Player{
		Name: "Test Player",
		Hand: hand,
	}
	return p
}

// Pretty-prints the game state, for debugging.
// Requires the testing.T struct in order to log to the test output.
func LogGameState(g *Game, t *testing.T) {
	s := fmt.Sprintf("Game:\n")
	for i := 0; i < len(g.Players); i++ {
		t.Log(g.Players[i])
	}
	t.Log(s)
	draw := fmt.Sprintf("Draw Deck: %d cards \n", len(g.DrawDeck))
	for j := 0; j < len(g.DrawDeck); j++ {
		draw += fmt.Sprintf("%v, ", g.DrawDeck[j])
	}
	t.Log(draw)
	discard := fmt.Sprintf("Discard Deck: %d cards \n", len(g.Discard))
	for k := 0; k < len(g.Discard); k++ {
		discard += fmt.Sprintf("%v, ", g.Discard[k])
	}
	t.Log(discard)
}
