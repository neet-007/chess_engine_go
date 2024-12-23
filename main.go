package main

import (
	"github.com/neet-007/chess_engine_go/pkg/shared"
	"github.com/neet-007/chess_engine_go/pkg/uci"
)

func main() {
	shared.Tell("hello from main")
	uci_ := uci.NewUCI()
	uci_.Main(shared.Input(), shared.MainTell)
}
