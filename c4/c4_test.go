package c4

import (
	"log"
	"os"
	"testing"
)

func TestYellowVerticleWin(t *testing.T) {
	log.SetOutput(os.Stdout)

	g := NewGame()

	for i := 0; i < 8; i++ {
		g = g.Move(i % 2)
		g.Board.Print()
		log.Printf("Winner: %s", g.Winner)
	}

	if g.Winner != Yellow {
		t.Errorf("Expected yellow to win")
	}
}

func TestRedVerticleWin(t *testing.T) {
	log.SetOutput(os.Stdout)

	g := NewGame()

	for i := 0; i < 8; i++ {
		if g.CurrentMover() == Yellow {
			g = g.Move(i % len(g.Board.Columns))
		} else {
			g = g.Move(1)
		}
		g.Board.Print()
	}

	if g.Winner != Red {
		t.Errorf("Expected red to win")
	}
}

func TestCurrentMover(t *testing.T) {
	g := NewGame()

	lastMover := g.CurrentMover()

	if lastMover != Yellow {
		t.Errorf("expected first-turn Yellow")
	}

	for i := 0; i < 6; i++ {
		g = g.Move(0)
		if g.CurrentMover() == lastMover {
			t.Errorf("expected the mover to alternate. lastMove: %s, currentMove: %s",
				lastMover, g.CurrentMover())
		}
		lastMover = g.CurrentMover()
	}

}
