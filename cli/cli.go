package cli

import (
	"fmt"
	"unoai/game"

	"github.com/fatih/color"
)

func StartGame() {

	// Create players
	players := CreatePlayers()

	// Create a new game
	g := game.MakeGame(players)

	// Display game
	DisplayGame(g)

}

func DisplayGame(g game.Game) {
	for i := 0; i < len(g.Players); i++ {
		// If it's the current player's turn, display their cards.
		// Player names
		fmt.Printf("%s: ", g.Players[i].Name)

		if i == g.TurnCtr%len(g.Players) {
			// Iterate thru deck
			for j := 0; j < len(g.Players[i].Hand); j++ {

				c := g.Players[i].Hand[j]
				str := ""

				// Color card
				switch c.Color {
				case game.Red:
					color.Set(color.FgRed)
				case game.Blue:
					color.Set(color.FgBlue)
				case game.Green:
					color.Set(color.FgGreen)
				case game.Yellow:
					color.Set(color.FgYellow)
				case game.None:
					color.Set(color.FgBlack)
					color.Set(color.BgHiWhite)
				}

				if c.Value <= 9 {
					// # cards
					str = fmt.Sprintf("%d", c.Value)
				} else if c.Value == game.Skip {
					// Skip card
					str = "S"
				} else if c.Value == game.Reverse {
					str = "R"
				} else if c.Value == game.DrawTwo {
					str = "+2"
				} else if c.Value == game.Wild {
					str = "W"
				} else if c.Value == game.WildDrawFour {
					str = "+4"
				}

				fmt.Printf("%s ", str)
			}

			fmt.Printf("\n")

		} else {
			handStr := ""
			for j := 0; j < len(g.Players[i].Hand); j++ {
				handStr += "🃵 "
			}
			fmt.Printf("%s ", handStr)
			fmt.Printf("\n")
		}

		color.Unset()
	}
}

func GetPlayerInfo() []string {
	// Ask for # of players
	color.Blue("How many players?")
	numPlayers := 0
	fmt.Scanln(&numPlayers)

	playerNames := make([]string, numPlayers)

	// Ask for player names
	for i := 0; i < numPlayers; i++ {
		color.Blue("Player %d, what is your name?", i+1)
		name := ""
		fmt.Scanln(&name)
		playerNames[i] = name
	}

	return playerNames

}

func CreatePlayers() []game.Player {
	// Ask user for players
	playerNames := GetPlayerInfo()

	// Create players
	players := make([]game.Player, len(playerNames))

	for i := 0; i < len(players); i++ {
		players[i] = game.Player{
			Name: playerNames[i],
		}
	}

	return players
}
