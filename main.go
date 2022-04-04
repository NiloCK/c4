package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	c4 "github.com/nilock/c4/game"
)

type GameState struct {
	nextMove int
	game c4.Game
}

func initModel() GameState {
	return GameState{
		4,
		c4.NewGame(),
	}
}

func (g GameState) Init() tea.Cmd {
	return nil
}

func (g GameState) headerLine() string {
	ret := " "
	for i:=0; i<g.nextMove; i++ {
		ret += "  "
	}
	if g.game.CurrentMover() == c4.Red {
		ret += c4.RedStr
	} else if g.game.CurrentMover() == c4.Yellow {
		ret += c4.YelStr
	}
	return ret
}

func (g GameState) View() string {
	// todo: add prompt for winner
	
	return g.headerLine() + "\n" + g.game.View() + " 1 2 3 4 5 6 7\n"
}

func (g GameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return g, tea.Quit
		case "left":
			g.nextMove = (g.nextMove + 7 - 1) % 7
		case "right":
			g.nextMove = (g.nextMove + 1) % 7
		case "down":
			g.game = g.game.Move(g.nextMove)
		case "1":
			g.game = g.game.Move(0)
		case "2":
			g.game = g.game.Move(1)
		case "3":
			g.game = g.game.Move(2)
		case "4":
			g.game = g.game.Move(3)
		case "5":
			g.game = g.game.Move(4)
		case "6":
			g.game = g.game.Move(5)
		case "7":
			g.game = g.game.Move(6)
		}
	}

	return g, nil
}

func main() {
	game := tea.NewProgram(initModel())
	
	if err := game.Start(); err != nil {
		fmt.Printf("error starting game: %v", err)
		os.Exit(1)
	}
}