package network

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	// main "github.com/nilock/c4"
	"github.com/nilock/c4/c4"
)

type NetworkedGameState struct {
	gs     c4.GameState
	myTurn bool
	conn   net.Conn
	log    []string
}

func initModel(server bool, conn net.Conn) NetworkedGameState {
	return NetworkedGameState{
		gs:     c4.NewGameState(),
		myTurn: !server,
		conn:   conn,
		log:    []string{},
	}
}

func (n NetworkedGameState) Init() tea.Cmd {
	return nil
}

func (n NetworkedGameState) View() string {
	return fmt.Sprintf(n.gs.View())
}

func (n NetworkedGameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	initTurnNum := n.gs.Game.MoveCount
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return n, tea.Quit
		}

		if n.myTurn {
			// forward keystrokes to n.gs
			ngs, _ := n.gs.Update(msg)
			n.gs = ngs.(c4.GameState)

			n.passGameState()

			// set to opponent's turn if we have placed
			// a piece
			if initTurnNum != n.gs.Game.MoveCount {
				n.myTurn = false
			}

		}
	case c4.GameState:
		if !n.myTurn {
			n.gs = msg

			if n.gs.Game.MoveCount != initTurnNum {
				n.myTurn = true
			}
		}

	}

	return n, nil
}

// passGameState serializses the current gamestate and
// writes it to the network connection
func (n NetworkedGameState) passGameState() {
	gsJSON, err := json.Marshal(n.gs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	n.conn.Write(gsJSON)
}
