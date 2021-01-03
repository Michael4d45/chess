package chess

// Position represents a spot on the board.
type Position struct {
	x byte
	y byte
}

// GetDistances returns x and y distances
func (pos1 Position) GetDistances(pos2 Position) (byte, byte) {
	xDistance := pos1.x - pos2.x
	if pos1.x < pos2.x {
		xDistance = pos2.x - pos1.x
	}
	yDistance := pos1.y - pos2.y
	if pos1.y < pos2.y {
		yDistance = pos2.y - pos1.y
	}
	return xDistance, yDistance
}

// NextTo retruns true if two positions are once square away.
func (pos1 Position) NextTo(pos2 Position) bool {
	xDistance, yDistance := pos1.GetDistances(pos2)
	if xDistance <= 1 && yDistance <= 1 {
		return true
	}
	return false
}

// OrientatedTo returns a description of how two positions relate to each other.
func (pos1 Position) OrientatedTo(pos2 Position) string {
	xDistance, yDistance := pos1.GetDistances(pos2)
	switch {
	case xDistance == 0:
		return "rank"
	case yDistance == 0:
		return "file"
	case xDistance == yDistance:
		return "diagonal"
	default:
		return "none"
	}
}

// RankDirection returns cardinal directions of positions
func (pos1 Position) RankDirection(pos2 Position) string {
	switch {
	case pos1.y > pos2.y:
		return "W"
	case pos1.y < pos2.y:
		return "E"
	default:
		return ""
	}
}

// FileDirection returns cardinal directions of positions
func (pos1 Position) FileDirection(pos2 Position) string {
	switch {
	case pos1.x > pos2.x:
		return "N"
	case pos1.x < pos2.x:
		return "S"
	default:
		return ""
	}
}
