// Manages game state (players, cards, decks, etc.)
package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Players  []Player
	DrawDeck []Card
	Discard  []Card
}

// Initializes a new game
func MakeGame(p []Player) Game {
	g := Game{
		Players: p,
	}

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

	return g
}

type Player struct {
	Name string
	Hand []Card
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
