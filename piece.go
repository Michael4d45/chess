package chess

// Piece is an object used on the board
type Piece struct {
	player *Player
	pieceType string
}

var pieceNames = map[string]string{
	"R": "Rook",
	"B": "Bishop",
	"N": "Knight",
	"K": "King",
	"Q": "Queen",
	"P": "Pawn",
}

// NewPiece converts a byte to a piece
func NewPiece(pieceType string, p *Player) *Piece {
	_, has := pieceNames[pieceType]
	if !has {
		return nil
	}
	piece := new(Piece)
	piece.player = p
	piece.pieceType = pieceType
	return piece
}

func (p Piece) checkCanMove(x1 byte, y1 byte, x2 byte, y2 byte, board Board) (bool, *Piece) {
	piece2 := board.Spaces[x2][y2]
	switch(p.pieceType){
	case "R":
		return true, piece2
	case "B":
		return true, piece2
	case "N":
		return true, piece2
	case "K":
		return true, piece2
	case "Q":
		return true, piece2
	case "P":
		if p.player.direction == "N" {

		}
		if p.player.direction == "S" {

		}
		return true, piece2
	}
	return false, nil
}

func (p Piece) String() string{
	return p.player.name + p.pieceType
}