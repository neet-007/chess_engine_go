package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	White_pawn = iota
	Black_pawn
	White_knight
	Black_knight
	White_bishop
	Black_bishop
	White_rook
	Black_rook
	White_queen
	Black_queen
	White_king
	Black_king
)

const (
	a1 = iota
	b1
	c1
	d1
	e1
	f1
	g1
	h1

	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2

	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3

	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4

	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5

	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6

	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7

	a8
	b8
	c8
	d8
	e8
	f8
	g8
	h8
)

// NOTE: Using LERF (little endian rank file)
type Board struct {
	bitboards [14]uint64
	empty     uint64
	occupied  uint64
}

type Engine struct {
	id     string
	author string
}

func (b *Board) getSquareIndex(file uint8, rank uint8) uint8 {
	return uint8((rank << 3) + file)
}

func (b *Board) getFileIndex(index uint8) uint8 {
	return uint8(index & 7)
}

func (b *Board) getRankIndex(index uint8) uint8 {
	return uint8((index >> 3))
}

func parseFEN(fen string) error {
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return fmt.Errorf("Invalid FEN: %s expected 6 parts found %d", fen, len(parts))
	}

	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return fmt.Errorf("Invalid FEN: %s expected 8 ranks found %d", fen, len(ranks))
	}

	for i, rank := range ranks {
		total := 0
		for j, char := range rank {
			if unicode.IsDigit(char) {
				total += int(char - '0')
				fmt.Printf("%d files are empty in rank %d\n", int(char-'0'), i)
				continue
			}

			fmt.Printf("piece %s is in file %d in rank %d\n", string(char), j, i)
			total++
		}

		if total != 8 {
			return fmt.Errorf("Invalid FEN: %s expected 8 ranks found %d", fen, len(ranks))
		}
	}

	switch parts[1] {
	case "w":
		{
			fmt.Println("White to move")
		}
	case "b":
		{
			fmt.Println("Black to move")
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
					fmt.Println("White Kingside castling rights")
				}
			case 'Q':
				{
					fmt.Println("White Queenside castling rights")
				}
			case 'k':
				{
					fmt.Println("Black Kingside castling rights")
				}
			case 'q':
				{
					fmt.Println("Black Queenside castling rights")
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
	}

	if halfMoves, err := strconv.Atoi(parts[4]); err != nil {
		return fmt.Errorf("Invalid half moves: %s expected integer found %s", fen, parts[4])
	} else {
		fmt.Printf("Half moves is %d\n", halfMoves)
	}

	if fullMoves, err := strconv.Atoi(parts[5]); err != nil {
		return fmt.Errorf("Invalid full moves: %s expected integer found %s", fen, parts[5])
	} else {
		fmt.Printf("Full moves is %d\n", fullMoves)
	}

	return nil
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

	fmt.Printf("> ")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if err := parseFEN(line); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\n> ")
	}
}
