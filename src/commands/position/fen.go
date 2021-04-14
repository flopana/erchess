package position

import (
	"fmt"
)

func Fen(arguments []string)  {
	board := makeBasicBoard()

	fmt.Println(board)
}
func makeBasicBoard() Board{
	var board Board
	board.Field = make([][]Piece, 8)
	for i := range board.Field{
		board.Field[i] = make([]Piece, 8)
	}

	return board
}