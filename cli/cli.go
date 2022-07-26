package cli

import (
	"fmt"
	"unoai/game"

	"github.com/fatih/color"
	"github.com/logrusorgru/aurora"

	"strconv"
)

func StartGame() {

	// Create players
	players := CreatePlayers()

	// Create a new game
	g := game.MakeGame(players)

	fmt.Println(aurora.Green("Starting game ..."))

	// Display game
	for {
		DisplayGame(g)
		InputTurn(&g)
	}

}

// Get user input for a turn
func InputTurn(g *game.Game) {
	currentPlayer := g.Players[g.TurnCtr%len(g.Players)]
	// Ask for turn
	fmt.Printf("%s, input your turn:\n", aurora.Magenta(currentPlayer.Name))
	// Get card the user wishes to play
	fmt.Println("Enter a card number to play, or press <Enter> to draw a card.")
	input := ""
	fmt.Scanln(&input)

	// If user pressed enter, draw a card
	if input == "" {
		c := game.Card{
			Value: game.Zero,
			Color: game.None,
		}
		turn := game.MakeTurn(c, true)
		if g.PlayTurn(g.TurnCtr%len(g.Players), turn) {
			fmt.Println(aurora.Green("-- Draw successful."))
		} else {
			fmt.Println(aurora.Red("-- Draw unsuccessful."))
		}
	} else {
		cardIndex, err := strconv.Atoi(input)
		if cardIndex < 1 || cardIndex > len(currentPlayer.Hand) || err != nil {
			fmt.Println(aurora.Red("-- Invalid card / turn."))
			return
		}
		c := currentPlayer.Hand[cardIndex-1]
		turn := game.MakeTurn(c, false)
		if g.PlayTurn(g.TurnCtr%len(g.Players), turn) {
			fmt.Println(aurora.Green("-- Play successful."))
		} else {
			fmt.Println(aurora.Red("-- Play unsuccessful."))
		}

	}
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

	color.Unset()
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
