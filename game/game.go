// Manages game state (players, cards, decks, etc.)
package game

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Game struct {
	Players []Player
	TurnCtr int
	// Regular dir. is 1, reversed is -1.
	TurnIncr int
	DrawDeck []Card
	Discard  []Card
	GameOver bool
}

// Initializes a new game
func MakeGame(p []Player) Game {
	g := Game{
		Players: p,
	}

	g.TurnCtr = 0
	// Regular direction. Reversed is -1.
	g.TurnIncr = 1

	// Populate draw deck

	// Per color:

	// Two of each number card 1-9
	// One of each 0
	// Two of each +2
	// Two of each reverse
	// Two of each skip

	// Four +4
	// Four Wilds

	// Populate for cards of each color (+ the 8 wild cards)
	for i := 0; i < 4; i++ {
		color := CardColor(i)
		// Zeroes
		g.DrawDeck = append(g.DrawDeck, Card{color, Zero})

		for j := 0; j < 2; j++ {
			//Two of each number card 1-9
			for k := 1; k <= 9; k++ {
				g.DrawDeck = append(g.DrawDeck, Card{color, CardType(j)})
			}
			// Non-number cards
			g.DrawDeck = append(g.DrawDeck, Card{color, DrawTwo})
			g.DrawDeck = append(g.DrawDeck, Card{color, Reverse})
			g.DrawDeck = append(g.DrawDeck, Card{color, Skip})
		}
		//Wilds & +4s
		g.DrawDeck = append(g.DrawDeck, Card{None, Wild})
		g.DrawDeck = append(g.DrawDeck, Card{None, WildDrawFour})
	}

	// Setup rng
	rand.Seed(time.Now().UnixNano())

	// Shuffle draw deck
	rand.Shuffle(len(g.DrawDeck), func(a, b int) {
		g.DrawDeck[a], g.DrawDeck[b] = g.DrawDeck[b], g.DrawDeck[a]
	})

	// Deal cards
	g.Deal()

	// First card is drawn from the draw deck
	g.Discard = append(g.Discard, g.Draw(1)...)

	return g
}

// Deals 7 cards to each player, drawn from the draw deck.
func (g *Game) Deal() {
	for i := 0; i < len(g.Players); i++ {
		for j := 0; j < 7; j++ {
			g.Players[i].Hand = append(g.Players[i].Hand, g.DrawDeck[0])
			g.DrawDeck = g.DrawDeck[1:]
		}
	}
}

// Draws `n` cards from the top of the draw deck and returns them.
// TODO finish this
func (g *Game) Draw(n int) []Card {
	cards := make([]Card, 0, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Drawing card: %v", g.DrawDeck[0])
		cards = append(cards, g.DrawDeck[0])
		g.DrawDeck = g.DrawDeck[1:]
	}
	return cards
}

// Plays a turn created with MakeTurn.
func (g *Game) PlayTurn(playerIndex int, t Turn) {
	p := g.Players[playerIndex]

	// If the player is drawing, give them a card and end their turn.
	if t.Draw {
		p.Hand = append(p.Hand, g.Draw(1)...)
		g.TurnCtr += g.TurnIncr
		return
	}

	// Find & remove the card from the player's hand
	for i := 0; i < len(p.Hand); i++ {
		if p.Hand[i] == t.Card {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			break
		}
	}

	// Add the card to the discard pile
	g.Discard = append([]Card{t.Card}, g.Discard...)

	// If the card is a reverse, reverse the direction of the game
	if t.Card.Value == Reverse {
		g.TurnIncr = -1
	}

	// If the card is a skip, skip the next player
	if t.Card.Value == Skip {
		g.TurnCtr += g.TurnIncr * 2
	}

	// If the card is a draw two, draw two cards for the next player
	if t.Card.Value == DrawTwo {
		// nextPlayerHand := &g.Players[(playerIndex+g.TurnIncr)%len(g.Players)].Hand
		// // nextPlayerHand = append(nextPlayerHand, g.Draw(2)...)
	}
}

type Player struct {
	Name string
	Hand []Card
}

type Turn struct {
	// If the card is a wild, change the color of the card itself.
	Card Card
	// If the player draws a card instead of playing one.
	Draw bool
}

func MakeTurn(card Card, draw bool) Turn {
	t := Turn{
		Card: card,
		Draw: draw,
	}
	if !t.IsValid() {
		log.Fatal("Invalid turn")
	}
	return t
}

// Checks if the turn is valid
func (t Turn) IsValid() bool {
	// if the card is a wild & has no color set, it's invalid
	if (t.Card.Value == WildDrawFour || t.Card.Value == Wild) && t.Card.Color == None {
		return false
	}
	return true
}

type Card struct {
	Color CardColor
	Value CardType
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", []string{"Red", "Green", "Blue", "Yellow", "None"}[c.Color], []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "DrawTwo", "Reverse", "Skip", "Wild", "WildDrawFour"}[c.Value])
}

// enum for card types
type CardType int

const (
	Zero CardType = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Skip
	Reverse
	DrawTwo
	Wild
	WildDrawFour
)

// enum for card colors
type CardColor int

const (
	Red CardColor = iota
	Green
	Blue
	Yellow
	None
)
