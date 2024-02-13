package c4

type Game struct {
	Board     Board
	MoveCount int
	Winner    Piece
}

// check returns the winner of the game
func (g Game) check() Piece {
	if vertical := g.checkVertical(); vertical != None {
		return vertical
	}

	if horizontal := g.checkHorizontal(); horizontal != None {
		return horizontal
	}

	if diagonal := g.checkDiagonals(); diagonal != None {
		return diagonal
	}

	return None
}

func (g Game) checkVertical() Piece {
	for _, p := range []Piece{Red, Yellow} {

		for _, col := range g.Board.Columns {
			for start := range []int{0, 1, 2} {
				if col[start] == p &&
					col[start+1] == p &&
					col[start+2] == p &&
					col[start+3] == p {
					return p
				}
			}
		}
	}
	return None
}

func (g Game) checkHorizontal() Piece {
	for _, p := range []Piece{Red, Yellow} {
		for i := 0; i+3 < len(g.Board.Columns); i++ {
			for j := 0; j+3 < len(g.Board.Columns[i]); j++ {
				if g.Board.Columns[i][j] == p &&
					g.Board.Columns[i+1][j] == p &&
					g.Board.Columns[i+2][j] == p &&
					g.Board.Columns[i+3][j] == p {
					return p
				}
			}
		}
	}
	return None
}

func (g Game) checkDiagonals() Piece {
	var p Piece
	for r := 0; r+3 < len(g.Board.Columns); r++ {
		for c := 0; c+3 < len(g.Board.Columns[0]); c++ {
			p = g.diagonalFrom(r, c)
			if p != None {
				return p
			}
		}
	}
	return None
}

func (g Game) diagonalFrom(r int, c int) (winner Piece) {
	// bottom-left to top-right diagonals
	if g.Board.Columns[r][c] == g.Board.Columns[r+1][c+1] &&
		g.Board.Columns[r][c] == g.Board.Columns[r+2][c+2] &&
		g.Board.Columns[r][c] == g.Board.Columns[r+3][c+3] {
		return g.Board.Columns[r][c]
	}

	h := len(g.Board.Columns[0]) - 1 - c
	// top-left to bottom-right diagonals
	if g.Board.Columns[r][h] == g.Board.Columns[r+1][h-1] &&
		g.Board.Columns[r][h] == g.Board.Columns[r+2][h-2] &&
		g.Board.Columns[r][h] == g.Board.Columns[r+3][h-3] {
		return g.Board.Columns[r][h]
	}
	return None
}

func NewGame() Game {
	return Game{
		Board:     NewBoard(),
		MoveCount: 0,
		Winner:    None,
	}
}

func (g Game) View() string {
	return g.Board.Print()
}

func (g Game) CurrentMover() Piece {
	if g.Winner != None {
		return None
	}

	if g.MoveCount%2 == 0 {
		return Yellow
	} else {
		return Red
	}
}

func (g Game) Move(column int) Game {
	if column > 6 || column < 0 {
		return g
	}
	if g.Winner != None {
		return g
	}

	g.MoveCount++

	if g.MoveCount%2 == 0 {
		updated, _ := g.Board.Add(Red, column)
		g.Board = updated
	} else {
		updated, _ := g.Board.Add(Yellow, column)
		g.Board = updated
	}

	g.Winner = g.check()

	return g
}
