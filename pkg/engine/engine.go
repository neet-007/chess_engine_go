package engine

import (
	"github.com/neet-007/chess_engine_go/pkg/shared"
)

func Engine() (chan string, chan string) {
	shared.Tell("from engine")
	frEng := make(chan string)
	toEng := make(chan string)

	go func() {
		for cmd := range toEng {
			switch cmd {
			case "quit":
				{

				}
			case "stop":
				{

				}
			}
		}

	}()

	return frEng, toEng
}
