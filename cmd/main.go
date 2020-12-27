package main

import (
	"fmt"
	"github.com/michael4d45/chess"
)

func main() {
	board := chess.FilledBoard()
	fmt.Println("board:", board.String())
}
