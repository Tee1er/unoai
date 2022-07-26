// Entry point for the program
package main

import (
	"embed"
	"unoai/cli"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// app := NewApp()

	// // Create application with options
	// err := wails.Run(&options.App{
	// 	Title:     "unoai",
	// 	Width:     1024,
	// 	Height:    768,
	// 	Assets:    assets,
	// 	OnStartup: app.startup,
	// 	Bind: []interface{}{
	// 		app,
	// 	},
	// })

	// if err != nil {
	// 	println("Error:", err)
	// }

	// Run the CLI interface

	cli.StartGame()
}
