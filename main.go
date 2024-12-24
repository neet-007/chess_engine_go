package main

import (
	"github.com/neet-007/chess_engine_go/pkg/board"
	"github.com/neet-007/chess_engine_go/pkg/shared"
)

func main() {
	b := board.Board{}
	shared.Tell = shared.MainTell
	b.NewGame()
	b.Print()
	b.PrintAllBB()
}
