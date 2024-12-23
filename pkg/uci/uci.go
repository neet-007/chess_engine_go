package uci

import (
	"fmt"
	"strings"

	"github.com/neet-007/chess_engine_go/pkg/board"
	"github.com/neet-007/chess_engine_go/pkg/engine"
	"github.com/neet-007/chess_engine_go/pkg/shared"
)

type UCI struct {
	Options map[string]string
	Debug   bool
}

func initTell(tell func(text ...string)) {
	shared.Tell = tell
}

func NewUCI(tell func(test ...string)) *UCI {
	initTell(tell)
	return &UCI{Debug: false}
}

var saveBm string = ""

func formatCmd(cmd string) string {
	return strings.TrimSpace(strings.ToLower(cmd))
}

func (u *UCI) Main(frGUI chan string) {
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
				u.handleBestMove(bestMove, &biInfinite)
				continue
			}
		}
		words := strings.Split(cmd, " ")
		words[0] = formatCmd(words[0])
		switch words[0] {
		case "uci":
			{
				u.handleUci()
			}
		case "setoption":
			{
				u.handleSetOption(words)
			}
		case "isready":
			{
				u.handleIsReady()
			}
		case "ucinewgame":
			{
				u.handleUciNewGame()
			}
		case "position":
			{
				u.handlePosition(cmd)
			}
		case "debug":
			{
				u.handleDebug(words)
			}
		case "register":
			{
				u.handleRegister(words)
			}
		case "go":
			{
				u.handleGo(words)
			}
		case "ponderhit":
			{
				u.handlePonderHit()
			}
		case "stop":
			{
				u.handleStop(toEng, &biInfinite)
			}
		case "quit", "q":
			u.handleQuit(toEng)
			quit = true
			continue
		}
	}
}

func (u *UCI) handleUci() {
	shared.Tell("id name chessEngine")
	shared.Tell("id author moayed")

	shared.Tell("option name Hash type spin default 32 min 1 max 1024")
	shared.Tell("option name Threads type spin default 1 min 1 max 16")
	shared.Tell("uciok")
}

func (u *UCI) handleIsReady() {
	shared.Tell("readyok")
}

func (u *UCI) handleSetOption(option []string) {
	shared.Tell("set option with option " + strings.Join(option, " "))
	shared.Tell("info not impleatned")
}

func (u *UCI) handleBestMove(bestMove string, biInfinite *bool) {
	if *biInfinite {
		saveBm = bestMove
		return
	}
	shared.Tell(bestMove)
}

func (u *UCI) handleUciNewGame() {
	shared.Tell("info ucinewgame not implemented")
}

func (u *UCI) handlePosition(cmd string) {
	cmd = strings.TrimSpace(strings.TrimPrefix(cmd, "position"))
	parts := strings.Split(cmd, "moves")
	if len(parts) == 0 || len(parts) > 2 {
		err := fmt.Errorf("%v wrong length=%v", parts, len(parts))
		shared.Tell("info position error ", err.Error())
		return
	}

	alt := strings.Split(parts[0], " ")

	alt[0] = strings.TrimSpace(alt[0])
	if alt[0] == "startpos" {
		parts[0] = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	} else if alt[0] == "fen" {
		parts[0] = strings.TrimSpace(strings.TrimPrefix(parts[0], "fen"))
	} else {
		err := fmt.Errorf("%#v must be %#v or %#v", alt[0], "startpos", "fen")
		shared.Tell("info position error ", err.Error())
		return
	}

	board.ParseFEN(parts[0])

	if len(parts) == 2 {
		parts[1] = formatCmd(parts[1])
		board.ParseMvs(parts[1])
	}
}

func (u *UCI) handleDebug(words []string) {
	shared.Tell("info debug " + strings.Join(words, " ") + " not implemented")
}

func (u *UCI) handleRegister(words []string) {
	shared.Tell("info register " + strings.Join(words, " ") + " not implemented")
}

func (u *UCI) handleGo(words []string) {
	if len(words) > 1 {
		words[1] = formatCmd(words[1])
		switch words[1] {
		case "searchmoves":
			{
				shared.Tell("info go searchmoves not implemetned")
			}
		case "ponder":
			{
				shared.Tell("info go ponder not implemented")
			}
		case "wtime":
			{
				shared.Tell("info go wtime not implemented")
			}
		case "btime":
			{
				shared.Tell("info go btime not implemented")
			}
		case "winc":
			{
				shared.Tell("info go winc not impleanted")
			}
		case "binc":
			{
				shared.Tell("info go binc not implemnetd")
			}
		case "movestogo":
			{
				shared.Tell("info go movestogo not implemented")
			}
		case "depth":
			{
				shared.Tell("info go depth not implemented")
			}
		case "nodes":
			{
				shared.Tell("info go nodes not implemented")
			}
		case "movetime":
			{
				shared.Tell("info go movetime not implemented")
			}
		case "mate":
			{
				shared.Tell("info go mate not implemented")
			}
		case "infinite":
			{
				shared.Tell("info go infinite not implemnetd")
			}
		default:
			{
				shared.Tell("info go " + words[1] + "not implemenetd")
			}
		}
	} else {
		shared.Tell("info go string not implemnted")
	}
}

func (u *UCI) handlePonderHit() {
	shared.Tell("info ponder not implemented")
}

func (u *UCI) handleStop(toEng chan string, biInfinite *bool) {
	if *biInfinite {
		if saveBm != "" {
			shared.Tell(saveBm)
			saveBm = ""
		}
		toEng <- "stop"
		*biInfinite = false
	}
}

func (u *UCI) handleQuit(toEng chan string) {
	toEng <- "stop"
}
