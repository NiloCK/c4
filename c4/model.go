package c4

import tea "github.com/charmbracelet/bubbletea"

type GameState struct {
	NextMove int
	Game     Game
}

func NewGameState() GameState {
	return GameState{
		0,
		NewGame(),
	}
}

func (g GameState) Init() tea.Cmd {
	return nil
}

func (g GameState) headerLine() string {
	ret := " "
	for i := 0; i < g.NextMove; i++ {
		ret += "  "
	}
	if g.Game.CurrentMover() == Red {
		ret += RedStr
	} else if g.Game.CurrentMover() == Yellow {
		ret += YelStr
	}
	return ret
}

// View returns a view of the current game state, including
//  - the board, with its placed pieces
//  - the in-play piece above the board
//  - numeric labels for each column
func (g GameState) View() string {
	// todo: add prompt for winner

	return g.headerLine() + "\n" + g.Game.View() + " 1 2 3 4 5 6 7\n"
}

func (g GameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return g, tea.Quit
		case "left":
			g.NextMove = (g.NextMove + 7 - 1) % 7
		case "right":
			g.NextMove = (g.NextMove + 1) % 7
		case "down":
			g.Game = g.Game.Move(g.NextMove)
		case "1":
			g.Game = g.Game.Move(0)
		case "2":
			g.Game = g.Game.Move(1)
		case "3":
			g.Game = g.Game.Move(2)
		case "4":
			g.Game = g.Game.Move(3)
		case "5":
			g.Game = g.Game.Move(4)
		case "6":
			g.Game = g.Game.Move(5)
		case "7":
			g.Game = g.Game.Move(6)
		}
	}

	return g, nil
}
