package chess

import (
	"errors"
)

// Board has a 2D array representing piece placements
type Board struct {
	Spaces  [8][8]*Piece
	Players [2]*Player
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
				if j % 2 == i % 2 {
					s += "  "
				} else {
					s += "**"
				}
			}
		}
		s += "|\n"
		for j := 0; j < len(b.Spaces[i]); j++ {
			if j % 2 == i % 2 {
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
	board := Board{Players: [2]*Player{p1, p2}}

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

// MovePiece will try to move one piece from one space to another.
func (b *Board) MovePiece(pos1 string, pos2 string) error {
	x1, y1, err1 := toXY(pos1)
	if err1 != nil {
		return err1
	}

	x2, y2, err2 := toXY(pos2)
	if err2 != nil {
		return err2
	}

	piece1 := b.Spaces[x1][y1]
	canTake, piece2 := piece1.checkCanMove(x1, y1, x2, y2, *b)
	if canTake {
		if piece2 != nil {
			for i := 0; i < len(b.Players); i++ {
				if piece2.player.name == b.Players[i].name {
					_ = append(b.Players[i].graveyard, *piece2)
					break
				}
			}
		}
	}

	return nil
}
