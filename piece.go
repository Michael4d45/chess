package chess

// Piece is an object used on the board
type Piece struct {
	player    *Player
	pieceType string
	Moves     []string
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
	piece2 := b.GetPiece(pos2)
	switch p.pieceType {
	case "R":
		switch pos1.OrientatedTo(pos2) {
		case "rank":
			if b.CheckEmptySpacesRank(pos1, pos2) {
				b.Swap(pos1, pos2, "rank")
				return true
			}
		case "file":
			if b.CheckEmptySpacesFile(pos1, pos2) {
				b.Swap(pos1, pos2, "file")
				return true
			}
		}
	case "B":
		if pos1.OrientatedTo(pos2) == "diagonal" {
			if b.CheckEmptySpacesDiagonal(pos1, pos2) {
				b.Swap(pos1, pos2, "move")
				return true
			}
		}
	case "N":
		xDistance, yDistance := pos1.GetDistances(pos2)
		if xDistance == 1 && yDistance == 2 || yDistance == 1 && xDistance == 2 {
			b.Swap(pos1, pos2, "move")
			return true
		}
	case "K":
		if pos1.NextTo(pos2) {
			b.Swap(pos1, pos2, "move")
			return true
		} else if len(p.Moves) == 0 && !p.InCheck(pos1, b) { // Castling
			xDistance, yDistance := pos1.GetDistances(pos2)
			if yDistance == 2 && xDistance == 0 {
				castlePos1 := Position{}
				castlePos2 := Position{}
				switch pos1.RankDirection(pos2) {
				case "W":
					castlePos1 = Position{x: pos1.x, y: 0}
					castlePos2 = Position{x: pos1.x, y: pos2.y + 1}
				case "E":
					castlePos1 = Position{x: pos1.x, y: 7}
					castlePos2 = Position{x: pos1.x, y: pos2.y - 1}
				}
				castle := b.GetPiece(castlePos1)
				posDir := pos1.RankDirection(pos2)
				castleDir := pos1.RankDirection(castlePos1)
				if castle != nil &&
					castle.pieceType == "R" &&
					len(castle.Moves) == 0 &&
					pos1.OrientatedTo(castlePos1) == "rank" &&
					posDir == castleDir &&
					b.CheckEmptySpacesRank(pos1, castlePos1) {
					b.Swap(pos1, pos2, "castle")
					b.Swap(castlePos1, castlePos2, "castle")
					return true
				}
			}
		}
	case "Q":
		switch pos1.OrientatedTo(pos2) {
		case "rank":
			if b.CheckEmptySpacesRank(pos1, pos2) {
				b.Swap(pos1, pos2, "")
				return true
			}
		case "file":
			if b.CheckEmptySpacesFile(pos1, pos2) {
				b.Swap(pos1, pos2, "")
				return true
			}
		case "diagonal":
			if b.CheckEmptySpacesDiagonal(pos1, pos2) {
				b.Swap(pos1, pos2, "")
				return true
			}
		}
	case "P":
		if p.player.direction == pos1.FileDirection(pos2) {
			if pos1.NextTo(pos2) {
				switch pos1.OrientatedTo(pos2) {
				case "file":
					if piece2 == nil {
						b.Swap(pos1, pos2, "move")
						return true
					}
				case "diagonal":
					if piece2 != nil {
						b.Swap(pos1, pos2, "take")
						return true
					}
					// En passant
					pos3 := Position{}
					switch pos1.RankDirection(pos2) {
					case "W":
						pos3 = Position{pos1.x, pos1.y - 1}
					case "E":
						pos3 = Position{pos1.x, pos1.y + 1}
					}
					piece3 := b.GetPiece(pos3)
					if piece3 != nil {
						if piece3.pieceType == "P" && piece3.LastMove() == "double" {
							b.Swap(pos1, pos2, "En passant")
							return true
						}
					}
				}
			} else if len(p.Moves) == 0 {
				if pos1.OrientatedTo(pos2) == "file" {
					xDistance, _ := pos1.GetDistances(pos2)
					if xDistance == 2 && b.CheckEmptySpacesFile(pos1, pos2) {
						b.Swap(pos1, pos2, "double")
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

// LastMove returns the last move of the move list.
func (p *Piece) LastMove() string {
	if len(p.Moves) == 0 {
		return "empty"
	}
	return p.Moves[len(p.Moves)-1]
}

// InCheck returns true if the piece is currently in check.
func (p *Piece) InCheck(pos Position, b *Board) bool {
	return false
}
