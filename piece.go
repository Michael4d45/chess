package chess

// Piece is an object used on the board
type Piece struct {
	player    *Player
	pieceType string
	moveNum   int
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

func (p *Piece) move(pos1 Position, pos2 Position, b *Board) bool {
	//piece2 := b.Spaces[x2][y2]
	switch p.pieceType {
	case "R":
		switch pos1.OrientatedTo(pos2) {
		case "rank":
			if b.CheckEmptySpacesRank(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		case "file":
			if b.CheckEmptySpacesFile(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		}
	case "B":
		if pos1.OrientatedTo(pos2) == "diagonal" {
			if b.CheckEmptySpacesDiagonal(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		}
	case "N":
		xDistance, yDistance := pos1.GetDistances(pos2)
		if xDistance == 1 && yDistance == 2 || yDistance == 1 && xDistance == 2 {
			b.Swap(pos1, pos2)
			return true
		}
	case "K":
		if pos1.NextTo(pos2) {
			b.Swap(pos1, pos2)
			return true
		}
	case "Q":
		switch pos1.OrientatedTo(pos2) {
		case "rank":
			if b.CheckEmptySpacesRank(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		case "file":
			if b.CheckEmptySpacesFile(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		case "diagonal":
			if b.CheckEmptySpacesDiagonal(pos1, pos2) {
				b.Swap(pos1, pos2)
				return true
			}
		}
	case "P":
		if p.player.direction == pos1.FileDirection(pos2) {
			if pos1.NextTo(pos2) {
				switch pos1.OrientatedTo(pos2) {
				case "file":
					b.Swap(pos1, pos2)
					return true
				case "diagonal":
					b.Swap(pos1, pos2)
					return true
				}
			} else if (p.moveNum == 0) {
				if pos1.OrientatedTo(pos2) == "file" {
					xDistance, _ := pos1.GetDistances(pos2)
					if xDistance == 2 && b.CheckEmptySpacesFile(pos1, pos2) {
						b.Swap(pos1, pos2)
						return true
					}
				}
			}
		}
	}
	return false
}

func (p Piece) String() string {
	return p.player.name + p.pieceType
}
