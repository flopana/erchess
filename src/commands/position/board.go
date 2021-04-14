package position

type Piece struct {
	Color rune
	PieceType rune
}

type Board struct {
	Field      [][]Piece
	EnpassantX int
	EnpassantY int
	ToMove     rune

}
