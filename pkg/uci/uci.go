package uci

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/neet-007/chess_engine_go/pkg/engine"
)

func Uci(frGUI chan string, tell func(text ...string)) {
	tell("hello from uci")
	frEng, toEng := engine.Engine()
	cmd := ""
	bestMove := ""
	quit := false

	for !quit {
		select {
		case cmd = <-frGUI:
		case bestMove = <-frEng:
			{
				handleBestMove(bestMove)
				continue
			}
		}
		switch cmd {
		case "uci":
		case "stop":
			{
				handleStop(toEng)
			}
		case "quit", "q":
			quit = true
			continue
		}
	}
}

func handleBestMove(bestMove string) {

}

func handleStop(toEng chan string) {
	toEng <- "stop"
}

func mainTell(text ...string) {
	builder := strings.Builder{}

	for _, t := range text {
		builder.WriteString(t)
	}

	fmt.Println(builder.String())
}

func input() chan string {
	line := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()

	return line
}
