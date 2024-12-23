package main

import (
	"github.com/neet-007/chess_engine_go/pkg/shared"
	"github.com/neet-007/chess_engine_go/pkg/uci"
)

func main() {
	uci_ := uci.NewUCI(shared.Tell)
	uci_.Main(shared.Input())
}
