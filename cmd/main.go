package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/michael4d45/chess"
)

func main() {
	board := chess.FilledBoard()
	fmt.Println(board.String())

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to chess")
	rankReg := `([a-h]|A-H])`
	fileReg := `[1-8]`
	posReg := rankReg + fileReg
	movePosReg := posReg + ` ` + posReg
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1) // for windows

		if strings.Compare("exit", text) == 0 {
			break
		}

		if matched, _ := regexp.MatchString(`^`+movePosReg+`$`, text); matched {
			pos1 := text[:2]
			pos2 := text[3:5]
			board.MovePiece(pos1, pos2)
		}

		if matched, _ := regexp.MatchString(`^player[1|2]$`, text); matched {
			num := text[6] - '1'
			player := board.Players[num]
			fmt.Println(player)
		}

		fmt.Println(board.String())
	}
}
