package network

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nilock/c4/c4"
)

func InitClient(addr string) {
	fmt.Println("client mode...")

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Client connection: %s", conn.RemoteAddr().String())

	ngs := initModel(false, conn)
	prog := tea.NewProgram(ngs)

	go func() {
		for {
			bf := make([]byte, 5120)

			n, err := conn.Read(bf)
			if err != nil {
				panic(err)
			}
			var gs c4.GameState
			json.Unmarshal(bf[:n], &gs)

			prog.Send(gs)
		}
	}()

	if err := prog.Start(); err != nil {
		fmt.Printf("error starting game: %v", err)
		os.Exit(1)
	} else {
		fmt.Printf("game started")
	}

}
