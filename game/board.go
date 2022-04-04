package c4

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Piece string

const (
    Red = "red"
    Yellow = "yellow"
	None = "none"
)

var RedStr = lipgloss.NewStyle().SetString("⊚").Foreground(
	lipgloss.Color("#ff0000")).String()
var YelStr = lipgloss.NewStyle().SetString("❉").Foreground(
	lipgloss.Color("#efff00")).String()


type Board struct {
	columns [7][6]Piece
}

func (b Board) Add(p Piece, column int) (Board, error) {
	if p == None {
		return b, fmt.Errorf("added pieces must be Red or Yellow")
	}

	for i := range b.columns[column] {
		if b.columns[column][i] == None {
			b.columns[column][i] = p
			return b, nil
		}
	}
	
	return b, fmt.Errorf("column %d is full", column)
}

func (b Board) Print() string {
	s := strings.Builder{}
	s.WriteString("===============\n")

	for i:=len(b.columns[0]) - 1; i>=0; i-- {
		s.WriteRune('|')
		for j := 0; j<len(b.columns); j++ {
			// str := ""
			color := b.columns[j][i]
			switch color {
			case Red:
				// str := lipgloss.NewStyle().SetString("  ").Background(lipgloss.Color("#fefe00"))
				// s.WriteString(str.String())
				s.WriteString(RedStr)
			case Yellow:
				// str += yelStr.Render("y")
				s.WriteString(YelStr)
			case None:
				// str += " "
				s.WriteString(" ")
			}
			s.WriteRune('|')
			// str += fmt.Sprint(b.columns[j][i])
		}
		// s.WriteRune('|')
		s.WriteRune('\n')
		// rows = append(rows, str)
	}

	s.WriteString("===============\n")
	
	// fmt.Println(s.String())
	return s.String()
}

func NewBoard() Board {
	b := Board{
		columns: [7][6]Piece{
			{},{},{},{},{},{},{},
		},
	}
	for c, col := range b.columns {
		for r := range col {
			b.columns[c][r] = None
		}
	}
	return b
}
 