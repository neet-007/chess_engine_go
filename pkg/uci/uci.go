package uci

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/neet-007/chess_engine_go/pkg/engine"
	"github.com/neet-007/chess_engine_go/pkg/shared"
)

func InitTell() {
	shared.Tell = MainTell
}

var saveBm string = ""

func formatCmd(cmd string) string {
	return strings.TrimSpace(strings.ToLower(cmd))
}

func Uci(frGUI chan string, tell func(text ...string)) {
	tell("hello from uci")
	frEng, toEng := engine.Engine()
	var cmd string
	var bestMove string
	biInfinite := false
	quit := false
	for !quit {
		select {
		case cmd = <-frGUI:
			{
			}
		case bestMove = <-frEng:
			{
				handleBestMove(bestMove, &biInfinite)
				continue
			}
		}
		words := strings.Split(cmd, " ")
		words[0] = formatCmd(words[0])
		switch words[0] {
		case "uci":
			{
				handleUci()
			}
		case "setoption":
			{
				handleSetOption(words)
			}
		case "isready":
			{
				handleIsReady()
			}
		case "ucinewgame":
			{
				handleUciNewGame()
			}
		case "position":
			{
				handlePosition(words)
			}
		case "debug":
			{
				handleDebug(words)
			}
		case "register":
			{
				handleRegister(words)
			}
		case "go":
			{
				handleGo(words)
			}
		case "ponderhit":
			{
				handlePonderHit()
			}
		case "stop":
			{
				handleStop(toEng, &biInfinite)
			}
		case "quit", "q":
			handleQuit(toEng)
			quit = true
			continue
		}
	}
}

func handleUci() {
	shared.Tell("id name chessEngine")
	shared.Tell("id auther moayed")

	shared.Tell("option name Hash type spin default 32 min 1 max 1024")
	shared.Tell("option name Threads type spin default 1 min 1 max 16")
	shared.Tell("uciok")
}

func handleIsReady() {
	shared.Tell("readyok")
}

func handleSetOption(option []string) {
	shared.Tell("set option with option " + strings.Join(option, " "))
	shared.Tell("not impleatned")
}

func handleBestMove(bestMove string, biInfinite *bool) {
	if *biInfinite {
		saveBm = bestMove
		return
	}
	shared.Tell(bestMove)
}

func handleUciNewGame() {
	shared.Tell("ucinewgame not implemented")
}
func handlePosition(words []string) {
	if len(words) > 1 {
		words[1] = formatCmd(words[1])
		switch words[1] {
		case "startpos":
			{
				shared.Tell("position startpos not implemented")
			}
		case "fen":
			{
				shared.Tell("position fen not implemented")
			}
		default:
			{
				shared.Tell("position " + words[1] + " not implemmented")
			}
		}
	}
}
func handleDebug(words []string) {
	shared.Tell("debug " + strings.Join(words, " ") + " not implemented")
}
func handleRegister(words []string) {
	shared.Tell("register " + strings.Join(words, " ") + " not implemented")
}
func handleGo(words []string) {
	if len(words) > 1 {
		words[1] = formatCmd(words[1])
		switch words[1] {
		case "searchmoves":
			{
				shared.Tell("go searchmoves not implemetned")
			}
		case "ponder":
			{
				shared.Tell("go ponder not implemented")
			}
		case "wtime":
			{
				shared.Tell("go wtime not implemented")
			}
		case "btime":
			{
				shared.Tell("go btime not implemented")
			}
		case "winc":
			{
				shared.Tell("go winc not impleanted")
			}
		case "binc":
			{
				shared.Tell("go binc not implemnetd")
			}
		case "movestogo":
			{
				shared.Tell("go movestogo not implemented")
			}
		case "depth":
			{
				shared.Tell("go depth not implemented")
			}
		case "nodes":
			{
				shared.Tell("go nodes not implemented")
			}
		case "movetime":
			{
				shared.Tell("go movetime not implemented")
			}
		case "mate":
			{
				shared.Tell("go mate not implemented")
			}
		case "infinite":
			{
				shared.Tell("go infinite not implemnetd")
			}
		default:
			{
				shared.Tell("go " + words[1] + "not implemenetd")
			}
		}
	} else {
		shared.Tell("go string not implemnted")
	}
}
func handlePonderHit() {
	shared.Tell("ponder not implemented")
}

func handleStop(toEng chan string, biInfinite *bool) {
	if *biInfinite {
		if saveBm != "" {
			shared.Tell(saveBm)
			saveBm = ""
		}
		toEng <- "stop"
		*biInfinite = false
	}
}

func handleQuit(toEng chan string) {
	toEng <- "stop"
}

func MainTell(text ...string) {
	builder := strings.Builder{}

	for _, t := range text {
		builder.WriteString(t)
	}

	fmt.Println(builder.String())
}

func Input() chan string {
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
