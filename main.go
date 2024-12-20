package main

import (
	"github.com/neet-007/chess_engine_go/pkg/shared"
	"github.com/neet-007/chess_engine_go/pkg/uci"
)

func main() {
	uci.InitTell()
	shared.Tell("hello from main")
	uci.Uci(uci.Input(), uci.MainTell)
}
