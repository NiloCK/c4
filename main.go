package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	c4 "github.com/nilock/c4/c4"
	network "github.com/nilock/c4/network"
)

func initModel() c4.GameState {
	return c4.NewGameState()
}

var serve = flag.Bool("serve", false, "host a mnetworked multiplayer game")
var join = flag.String("join", "", "join a networked game hosted at supplied address")

func init() {
	flag.Parse()
}

func main() {

	game := tea.NewProgram(initModel())
	if *serve {
		// todo
		network.InitServer()
	}

	if *join != "" {
		network.InitClient(*join)
	}

	// game := tea.NewProgram(initModel())

	if err := game.Start(); err != nil {
		fmt.Printf("error starting game: %v", err)
		os.Exit(1)
	}
}
