package network

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nilock/c4/c4"
)

func InitServer() {
	fmt.Println("server mode...")
	server, err := net.Listen("tcp", "127.0.0.1:4444")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := server.Accept()
	gs := initModel(true, conn)

	game := tea.NewProgram(gs)

	go func() {
		for {
			bf := make([]byte, 5120)

			n, err := conn.Read(bf)
			if err != nil {
				panic(err)
			}
			var gs c4.GameState
			json.Unmarshal(bf[:n], &gs)

			game.Send(gs)
		}
	}()

	if err := game.Start(); err != nil {
		fmt.Printf("error starting server game: %v", err)
		os.Exit(1)
	}

}
