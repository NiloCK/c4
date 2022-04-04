package c4

type Game struct {
    board Board
    moveCount int
	winner Piece
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

		for _, col := range g.board.columns {
			for start := range []int{0,1,2} {
				if col[start] == p &&
				col[start + 1] == p &&
				col[start + 2] == p &&
				col[start + 3] == p {
					return p
				}
			}
		}
	}
	return None
}

func (g Game) checkHorizontal() Piece {
	for _, p := range []Piece{Red, Yellow} {
		for i := 0; i+3< len(g.board.columns); i++ {
			for j := 0; j+3< len(g.board.columns[i]); j++ {
				if g.board.columns[i][j] == p &&
				g.board.columns[i+1][j] == p &&
				g.board.columns[i+2][j] == p &&
				g.board.columns[i+3][j] == p {
					return p
				}
			}
		}
	}
	return None
}

func (g Game) checkDiagonals() Piece {
	var p Piece
	for r := 0; r+3<len(g.board.columns); r++ {
		for c :=0; c+3<len(g.board.columns[0]); c++ {
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
	if g.board.columns[r][c] == g.board.columns[r+1][c+1] &&
		g.board.columns[r][c] == g.board.columns[r+2][c+2] &&
		g.board.columns[r][c] == g.board.columns[r+3][c+3] {
		return g.board.columns[r][c]
	}
	
	h := len(g.board.columns[0]) - 1 - c
	// top-left to bottom-right diagonals
	if g.board.columns[r][h] == g.board.columns[r+1][h-1] &&
		g.board.columns[r][h] == g.board.columns[r+2][h-2] &&
		g.board.columns[r][h] == g.board.columns[r+3][h-3] {
		return g.board.columns[r][h]
	}
	return None
}

func NewGame() Game {
    return Game{
		board: NewBoard(),
		moveCount: 0,
		winner: None,
	}
}

func (g Game) View() string {
	return g.board.Print()
}

func (g Game) CurrentMover() Piece {
	if g.winner != None {
		return None
	}

	if g.moveCount % 2 == 0 {
		return Yellow
	} else {
		return Red
	}
}

func (g Game) Move(column int) Game {
	if column > 6 || column < 0 {
		return g
	}
	if g.winner != None {
		return g
	}

	g.moveCount++

	if g.moveCount % 2 == 0 {
		updated, _ := g.board.Add(Red, column)
		g.board = updated
	} else {
		updated, _ := g.board.Add(Yellow, column)
		g.board = updated
	}
	
	g.winner = g.check()

	return g
}

