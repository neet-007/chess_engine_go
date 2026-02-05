package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/neet-007/chess_engine_go/internal/board"
)

type Engine struct {
	id     string
	author string
	board  *board.Board
}

func NewEngine(id string, author string) *Engine {
	return &Engine{
		id:     id,
		author: author,
		board:  board.NewBoard(),
	}
}

func parseFEN(b *board.Board, fen string) (err error) {
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return fmt.Errorf("Invalid FEN: %s expected 6 parts found %d", fen, len(parts))
	}

	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return fmt.Errorf("Invalid FEN: %s expected 8 ranks found %d", fen, len(ranks))
	}

	for i, rank := range ranks {
		i = 7 - i
		total := 0
		for _, char := range rank {
			if unicode.IsDigit(char) {
				count := int(char - '0')
				total += count
				continue
			}

			b.PutInSquare(uint8(total), uint8(i), board.CharToPiece[char])
			total++
		}

		if total != 8 {
			return fmt.Errorf("Invalid FEN: %s expected 8 ranks found %d", fen, len(ranks))
		}
	}

	b.UpdateEmpty()

	switch parts[1] {
	case "w":
		{
			b.CurrentTurn = board.WhiteTurn
		}
	case "b":
		{
			b.CurrentTurn = board.BlackTurn
		}
	default:
		{
			return fmt.Errorf("Invalid turn: %s expected w or b found %s", fen, parts[1])
		}
	}

	if parts[2] == "-" {
		fmt.Println("No castling rights")
	} else {
		if len(parts[2]) > 4 {
			return fmt.Errorf("Invalid castling rights: %s expected 4 letters found %d", fen, len(parts[2]))
		}

		for _, char := range parts[2] {
			switch char {
			case 'K':
				{
					b.SetFlag(board.WhiteCastleKingside)
				}
			case 'Q':
				{
					b.SetFlag(board.WhiteCastleQueenside)
				}
			case 'k':
				{
					b.SetFlag(board.BlackCastleKingside)
				}
			case 'q':
				{
					b.SetFlag(board.BlackCastleQueenside)
				}
			default:
				{
					return fmt.Errorf("Invalid castling rights: %s expected K, Q, k or q found %s", fen, string(char))
				}
			}
		}
	}

	if parts[3] == "-" {
		fmt.Println("No en passant square")
	} else {
		// TODO: mabye validate the square?
		fmt.Printf("En passant square is %s\n", parts[3])
		file := parts[3][0] - 'a'
		rank := parts[3][1] - '1'

		if file > 7 || rank > 7 {
			return fmt.Errorf("Invalid en passant square: %s expected a1-h8 found %s", fen, parts[3])
		}

		b.EpSquare = (board.GetSquareIndex(uint8(file), uint8(rank)))
	}

	if halfMoves, err := strconv.Atoi(parts[4]); err != nil {
		return fmt.Errorf("Invalid half moves: %s expected integer found %s", fen, parts[4])
	} else {
		b.HalfMoves = halfMoves
		fmt.Printf("Half moves is %d\n", halfMoves)
	}

	if fullMoves, err := strconv.Atoi(parts[5]); err != nil {
		return fmt.Errorf("Invalid full moves: %s expected integer found %s", fen, parts[5])
	} else {
		b.FullMoves = fullMoves
		fmt.Printf("Full moves is %d\n", fullMoves)
	}

	return nil
}

func serializeFEN(b *board.Board) string {
	var builder strings.Builder

	for rank := range uint8(8) {
		rank = 7 - rank

		noPieceCount := 0
		for file := range uint8(8) {
			index := board.GetSquareIndex(file, rank)
			found := false
			var foundPiece board.Piece

			for i := range 12 {
				if b.Bitboards[i]&(1<<index) != 0 {
					if noPieceCount > 0 {
						builder.WriteString(strconv.Itoa(noPieceCount))
						noPieceCount = 0
					}

					found = true
					foundPiece = board.Piece(i)

					builder.WriteString(foundPiece.String())
					break
				}
			}

			if !found {
				noPieceCount++
			}
		}
		if noPieceCount > 0 {
			builder.WriteString(strconv.Itoa(noPieceCount))
			noPieceCount = 0
		}

		if rank != 0 {
			builder.WriteString("/")
		}
	}

	if b.CurrentTurn == board.WhiteTurn {
		builder.WriteString(" w")
	} else {
		builder.WriteString(" b")
	}

	builder.WriteString(" ")

	hasCastlingRights := false

	if b.GetFlag(board.WhiteCastleKingside) {
		builder.WriteString("K")
		hasCastlingRights = true
	}
	if b.GetFlag(board.WhiteCastleQueenside) {
		builder.WriteString("Q")
		hasCastlingRights = true
	}
	if b.GetFlag(board.BlackCastleKingside) {
		builder.WriteString("k")
		hasCastlingRights = true
	}
	if b.GetFlag(board.BlackCastleQueenside) {
		builder.WriteString("q")
		hasCastlingRights = true
	}

	if !hasCastlingRights {
		builder.WriteString("-")
	}

	builder.WriteString(" ")
	if b.EpSquare != board.No_square {
		builder.WriteString(board.Square(b.EpSquare).String())
	} else {
		builder.WriteString("-")
	}

	builder.WriteString(" ")
	builder.WriteString(strconv.Itoa(b.HalfMoves))
	builder.WriteString(" ")
	builder.WriteString(strconv.Itoa(b.FullMoves))

	return builder.String()
}

func readUCI(engine *Engine) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		command := scanner.Text()
		fmt.Println(command)

		parts := strings.Split(command, " ")

		switch parts[0] {
		case "uci":
			{
				fmt.Fprintf(os.Stdout, "id name %s\n", engine.id)
				fmt.Fprintf(os.Stdout, "id author %s\n", engine.author)
			}
		case "debug":
			{
				fmt.Fprintf(os.Stdout, "Uninmplemented command: %s\n", command)
			}
		case "isready":
			{
				fmt.Fprintf(os.Stdout, "readyok\n")
			}
		case "setoption":
			{
				name := parts[1]
				value := parts[2]

				fmt.Fprintf(os.Stdout, "option name %s value %s\n", name, value)
			}
		case "ucinewgame":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "go":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "ponderhit":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "position":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "quit":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "register":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		case "stop":
			{
				fmt.Fprintf(os.Stdout, "Unimplemented command: %s\n", command)
			}
		default:
			{
				fmt.Fprintf(os.Stdout, "Unknown command: %s\n", command)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	engine := Engine{
		id:     "chess_engine",
		author: "Moayed",
		board: &board.Board{
			CurrentTurn: board.WhiteTurn,
		},
	}
	_ = engine

	fmt.Printf("> ")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if err := parseFEN(engine.board, line); err != nil {
			fmt.Println(err)
			continue
		}

		serialized := serializeFEN(engine.board)
		if serialized == "" {
			fmt.Println("Invalid FEN")
		} else {
			fmt.Println(serialized)
			if serialized != line {
				fmt.Println("FEN is not equal to serialized FEN")
			}
		}
	}
}
