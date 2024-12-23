package uci_test

import (
	"strings"
	"testing"
	"time"

	"github.com/neet-007/chess_engine_go/pkg/uci"
)

var GUI = []string{}

func testTell(text ...string) {
	builder := strings.Builder{}

	for _, t := range text {
		builder.WriteString(t)
	}

	GUI = append(GUI, builder.String())
}

func TestUci(t *testing.T) {
	uci_ := uci.NewUCI(testTell)
	input := make(chan string)
	go uci_.Main(input)

	tests := []struct {
		name   string
		cmd    string
		wanted []string
	}{
		{"uci", "uci", []string{"id name chessEngine", "id author moayed", "option name Hash type spin default 32 min 1 max 1024", "option name Threads type spin default 1 min 1 max 16", "uciok"}},
		{"isready", "isready", []string{"readyok"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GUI = []string{}
			input <- tt.cmd
			time.Sleep(10 * time.Millisecond)
			for ix, want := range tt.wanted {
				if len(GUI) <= ix {
					t.Errorf("%v: wanted %#v in ix=%v but got nothing", tt.name, want, ix)
					continue
				}
				if len(want) > len(GUI[ix]) {
					t.Errorf("%v: wanted %#v (in index %v) but we got %#v", tt.name, want, ix, GUI[ix])
					continue
				}
				if GUI[ix][:len(want)] != want {
					t.Errorf("%v: Error. Should be %#v but we got %#v", tt.name, want, GUI[ix])
				}
			}

		})
	}

}
