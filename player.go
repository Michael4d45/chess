package chess

// Player moves pieces around
type Player struct {
	name string
	direction string
	graveyard []Piece
}