package main

import (
	"fmt"

	"github.com/neet-007/chess_engine_go/pkg/board"
)

func main() {
	board_ := board.NewBoardFromInitial(board.GetDefaultBoard())

	/*
		ret := board_.BitBoardBoardToByte()
		for _, row := range ret {
			for _, col := range row {
				fmt.Printf("%s ", string(col))
			}
			fmt.Println()
		}
	*/

	a1 := board.Bitboard(1) << board.A1
	a8 := board.Bitboard(1) << board.A8
	fmt.Printf("a1 has white rook %t\n", (board_.WhiteRooks&a1) != 0)
	fmt.Printf("a1 has black rook %t\n", (board_.BlackRooks&a1) != 0)
	fmt.Printf("a8 has white rook %t\n", (board_.WhiteRooks&a8) != 0)
	fmt.Printf("a8 has black rook %t\n", (board_.BlackRooks&a8) != 0)

}
