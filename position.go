package chess

// Position represents a spot on the board.
type Position struct {
	x byte
	y byte
}

// NextTo retruns true if two positions are once square away.
func (pos1 Position) NextTo(pos2 Position) bool {
	xDistance := pos1.x - pos2.x
	if pos1.x < pos2.x {
		xDistance = pos2.x - pos1.x
	}
	yDistance := pos1.y - pos2.y
	if pos1.y < pos2.y {
		yDistance = pos2.y - pos1.y
	}
	if xDistance <= 1 && yDistance <= 1 {
		return true
	}
	return false
}

// OrientatedTo returns a description of how two positions relate to each other.
func (pos1 Position) OrientatedTo(pos2 Position) string {
	xDistance := pos1.x - pos2.x
	if pos1.x < pos2.x {
		xDistance = pos2.x - pos1.x
	}
	yDistance := pos1.y - pos2.y
	if pos1.y < pos2.y {
		yDistance = pos2.y - pos1.y
	}
	if xDistance == 0 {
		return "rank"
	}
	if yDistance == 0 {
		return "file"
	}
	if xDistance == yDistance {
		return "diagonal"
	}
	return "none"
}