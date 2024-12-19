package main

import (
	"fmt"

	"github.com/neet-007/chess_engine_go/pkg/board"
)

func main() {
	board_ := board.NewBoardFromInitial(board.GetDefaultBoard())

	ret := board_.BitBoardBoardToByte()
	for _, row := range ret {
		for _, col := range row {
			fmt.Printf("%s ", string(col))
		}
		fmt.Println()
	}
	/*
	 */
}
