package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/michael4d45/chess"
)

const rankReg = `([a-h]|A-H])`
const fileReg = `[1-8]`
const posReg = rankReg + fileReg
const movePosReg = posReg + ` ` + posReg

func main() {
	board := chess.FilledBoard()

	actions := []string{}
	if len(os.Args) > 1 {
		actions = strings.Split(os.Args[1], ",")
	}
	i := 0

	fmt.Println(board.String())

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to chess")
	for {
		fmt.Print("-> ")
		text := ""

		if i < len(actions) {
			text = actions[i]
			fmt.Print(text)
			i++
		} else {
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\r", "", -1) // for windows
		}

		if !doAction(text, &board) {
			break
		}

		fmt.Println(board.String())
	}
}

func doAction(action string, b *chess.Board) bool {
	if strings.Compare("exit", action) == 0 {
		return false
	}

	if matched, _ := regexp.MatchString(`^`+movePosReg+`$`, action); matched {
		pos1 := action[:2]
		pos2 := action[3:5]
		b.MovePiece(pos1, pos2)
	}

	if matched, _ := regexp.MatchString(`^player[1|2]$`, action); matched {
		num := action[6] - '1'
		player := b.Players[num]
		fmt.Println(player)
	}

	if matched, _ := regexp.MatchString(`^piece ` + posReg + `$`, action); matched {
		pos := action[6:8]
		piece := b.GetPieceByString(pos)
		fmt.Println(piece)
		fmt.Println(piece.Moves)
	}
	return true
}
