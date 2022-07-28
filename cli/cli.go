package cli

import (
	"fmt"
	"os"
	"strconv"
	"unoai/game"

	"os/signal"
	"syscall"

	"github.com/logrusorgru/aurora"
)

func StartGame() {

	// Handle keyboard interrupt <Ctrl+C>
	CloseHandler()

	// Create players
	players := CreatePlayers()

	// Create a new game
	g := game.MakeGame(players)

	fmt.Println(aurora.Green("Starting game ..."))

	// Display game
	for {
		fmt.Println()
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

	// Repeat until valid input
	for {

		// Scan for input
		input := ""
		fmt.Scanln(&input)

		var t game.Turn

		if input == "" {
			c := game.Card{
				Value: game.Zero,
				Color: game.None,
			}

			t = game.MakeTurn(c, true)
		} else {
			cardIndex, err := strconv.Atoi(input)
			if cardIndex < 1 || cardIndex > len(currentPlayer.Hand) || err != nil {
				fmt.Println(aurora.Red("-- Invalid card / turn."))
				continue
			}
			c := currentPlayer.Hand[cardIndex-1]
			t = game.MakeTurn(c, false)
		}

		valid, errMsg := g.ValidateTurn(t)
		// If invalid turn, print error message and try again
		if !valid {
			fmt.Printf("-- Invalid turn: %s\n", aurora.Red(errMsg))
			continue
		}

		success := g.PlayTurn(g.TurnCtr%len(g.Players), t)

		// If play successful, break and go to next player. If not, try again.
		if success {
			fmt.Println(aurora.Green("-- Play successful."))
			break
		} else {
			fmt.Println(aurora.Red("-- Play unsuccessful."))
			continue
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
				str := c.ShortColorString()

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
	}
	// Display last card played in discard pile
	fmt.Printf("%s: %s\n", aurora.Red("Discard"), g.Discard[0].ShortColorString())
}

func GetPlayerInfo() []string {
	// Ask for # of players
	fmt.Println("How many players?")
	numPlayers := 0
	fmt.Scanln(&numPlayers)

	playerNames := make([]string, numPlayers)

	// Ask for player names
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %d, what is your name?\n", aurora.Magenta(i+1))
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

// Handle keyboard interrupt <Ctrl+C>
func CloseHandler() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigs
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()
}
