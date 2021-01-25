package chess

import (
	"errors"
)

// Board has a 2D array representing piece placements
type Board struct {
	Size    byte
	Spaces  [][]*Piece
	Players []*Player
}

// GetMaxPos returns the maximum position for a piece
func (b *Board) GetMaxPos() byte {
	return b.Size - 1
}

// CheckEmptySpacesFile will return true if all the spaces are empty between the two positions.
func (b *Board) CheckEmptySpacesFile(pos1 Position, pos2 Position) bool {
	y := pos1.y
	xStart := pos1.x
	xEnd := pos2.x
	if xStart > xEnd {
		xStart, xEnd = xEnd, xStart
	}
	for x := xStart + 1; x < xEnd; x++ {
		if b.Spaces[x][y] != nil {
			return false
		}
	}
	return true
}

// CheckEmptySpacesRank will return true if all the spaces are empty between the two positions.
func (b *Board) CheckEmptySpacesRank(pos1 Position, pos2 Position) bool {
	x := pos1.x
	yStart := pos1.y
	yEnd := pos2.y
	if yStart > yEnd {
		yStart, yEnd = yEnd, yStart
	}
	for y := yStart + 1; y < yEnd; y++ {
		if b.Spaces[x][y] != nil {
			return false
		}
	}
	return true
}

// CheckEmptySpacesDiagonal true if all the spaces are empty between the two positions.
func (b *Board) CheckEmptySpacesDiagonal(pos1 Position, pos2 Position) bool {
	xStart := pos1.x
	xEnd := pos2.x
	yStart := pos1.y
	yEnd := pos2.y
	if yStart > yEnd {
		if xStart > xEnd {
			x := xStart - 1
			for y := yStart - 1; y > yEnd; y-- {
				if b.Spaces[x][y] != nil {
					return false
				}
				x--
			}
		} else {
			x := xStart + 1
			for y := yStart - 1; y > yEnd; y-- {
				if b.Spaces[x][y] != nil {
					return false
				}
				x++
			}
		}
	} else {
		if xStart > xEnd {
			x := xStart - 1
			for y := yStart + 1; y < yEnd; y++ {
				if b.Spaces[x][y] != nil {
					return false
				}
				x--
			}
		} else {
			x := xStart + 1
			for y := yStart + 1; y < yEnd; y++ {
				if b.Spaces[x][y] != nil {
					return false
				}
				x++
			}
		}
	}
	return true
}

func (b *Board) String() string {
	s := "\n"
	for i := 0; i < len(b.Spaces); i++ {
		for j := 0; j < len(b.Spaces[i]); j++ {
			space := b.Spaces[i][j]
			s += "|"
			if space != nil {
				s += space.String()
			} else {
				if j%2 == i%2 {
					s += "  "
				} else {
					s += "**"
				}
			}
		}
		s += "|\n"
		for j := 0; j < len(b.Spaces[i]); j++ {
			if j%2 == i%2 {
				s += "|__"
			} else {
				s += "|**"
			}
		}
		s += "|\n"
	}
	return s
}

func toXY(s string) (byte, byte, error) {
	ErrBadPosition := errors.New("bad position: " + s)
	if len(s) != 2 {
		return 0, 0, ErrBadPosition
	}
	rank := s[0] - 'a'
	file := 7 - (s[1] - '1')
	if rank < 0 || rank >= 8 {
		return 0, 0, ErrBadPosition
	}
	if file < 0 || file >= 8 {
		return 0, 0, ErrBadPosition
	}
	return file, rank, nil
}

// FilledBoard generates a board with set pieces.
func FilledBoard() Board {
	p1 := new(Player)
	p2 := new(Player)
	p1.name = "W"
	p2.name = "B"
	p1.direction = "N"
	p2.direction = "S"
	size := byte(8)
	spaces := make([][]*Piece, size, size)
	for i := range spaces {
        spaces[i] = make([]*Piece, size)
    }
	board := Board{Size: size, Spaces: spaces, Players: []*Player{p1, p2}}

	board.AssignSpace("a1", "R", p1)
	board.AssignSpace("b1", "N", p1)
	board.AssignSpace("c1", "B", p1)
	board.AssignSpace("d1", "Q", p1)
	board.AssignSpace("e1", "K", p1)
	board.AssignSpace("f1", "B", p1)
	board.AssignSpace("g1", "N", p1)
	board.AssignSpace("h1", "R", p1)
	board.AssignSpace("a2", "P", p1)
	board.AssignSpace("b2", "P", p1)
	board.AssignSpace("c2", "P", p1)
	board.AssignSpace("d2", "P", p1)
	board.AssignSpace("e2", "P", p1)
	board.AssignSpace("f2", "P", p1)
	board.AssignSpace("g2", "P", p1)
	board.AssignSpace("h2", "P", p1)

	board.AssignSpace("a8", "R", p2)
	board.AssignSpace("b8", "N", p2)
	board.AssignSpace("c8", "B", p2)
	board.AssignSpace("d8", "Q", p2)
	board.AssignSpace("e8", "K", p2)
	board.AssignSpace("f8", "B", p2)
	board.AssignSpace("g8", "N", p2)
	board.AssignSpace("h8", "R", p2)
	board.AssignSpace("a7", "P", p2)
	board.AssignSpace("b7", "P", p2)
	board.AssignSpace("c7", "P", p2)
	board.AssignSpace("d7", "P", p2)
	board.AssignSpace("e7", "P", p2)
	board.AssignSpace("f7", "P", p2)
	board.AssignSpace("g7", "P", p2)
	board.AssignSpace("h7", "P", p2)

	return board
}

// AssignSpace fills the board with pieces.
func (b *Board) AssignSpace(pos string, piece string, player *Player) error {
	x, y, err := toXY(pos)
	if err != nil {
		return err
	}
	b.Spaces[x][y] = NewPiece(piece, player)
	return nil
}

// GetPiece returns pice at this position on the board.
func (b *Board) GetPiece(pos Position) *Piece {
	return b.Spaces[pos.x][pos.y]
}

// GetPieceByString get peice by string.
func (b *Board) GetPieceByString(pos string) *Piece {
	x, y, _ := toXY(pos)
	return b.GetPiece(Position{x, y})
}

// MovePiece will try to move one piece from one space to another.
func (b *Board) MovePiece(posString1 string, posString2 string) error {
	x1, y1, err1 := toXY(posString1)
	if err1 != nil {
		return err1
	}
	pos1 := Position{x1, y1}

	x2, y2, err2 := toXY(posString2)
	if err2 != nil {
		return err2
	}
	pos2 := Position{x2, y2}

	piece1 := b.GetPiece(pos1)
	if piece1 == nil {
		return nil
	}
	moved := piece1.move(pos1, pos2, b)
	if moved {
	}

	return nil
}

// Swap will swap two piece positions.
func (b *Board) Swap(pos1 Position, pos2 Position, move string) {
	piece := b.GetPiece(pos1)
	if piece != nil && move != "" {
		piece.Moves = append(piece.Moves, move)
	}
	b.Spaces[pos2.x][pos2.y], b.Spaces[pos1.x][pos1.y] = b.Spaces[pos1.x][pos1.y], b.Spaces[pos2.x][pos2.y]
}
