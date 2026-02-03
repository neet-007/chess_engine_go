package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Piece uint8

func (p Piece) String() string {
	switch p {
	case White_pawn:
		return "P"
	case Black_pawn:
		return "p"
	case White_knight:
		return "N"
	case Black_knight:
		return "n"
	case White_bishop:
		return "B"
	case Black_bishop:
		return "b"
	case White_rook:
		return "R"
	case Black_rook:
		return "r"
	case White_queen:
		return "Q"
	case Black_queen:
		return "q"
	case White_king:
		return "K"
	case Black_king:
		return "k"
	default:
		return "?"
	}
}

const (
	White_pawn Piece = iota
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
	No_piece
)

var CharToPiece = map[rune]Piece{
	'P': White_pawn,
	'p': Black_pawn,
	'N': White_knight,
	'n': Black_knight,
	'B': White_bishop,
	'b': Black_bishop,
	'R': White_rook,
	'r': Black_rook,
	'Q': White_queen,
	'q': Black_queen,
	'K': White_king,
	'k': Black_king,
}

var KnigthAttacks = [64]uint64{}
var PawnAttacks = [2][64]uint64{}
var KingAttacks = [64]uint64{}

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

type currentTurn uint8

const (
	WhiteTurn currentTurn = iota
	BlackTurn
)

const (
	WhiteCastleKingside = iota
	WhiteCastleQueenside
	BlackCastleKingside
	BlackCastleQueenside
	EnPassant
)

// NOTE: Using LERF (little endian rank file)
type Board struct {
	bitboards   [14]uint64
	empty       uint64
	occupied    uint64
	mailbox     [64]Piece
	currentTurn currentTurn
	flags       uint8
}

type Engine struct {
	id     string
	author string
	board  Board
}

func (b *Board) setFlag(flag uint8) {
	b.flags |= 1 << flag
}

func getSquareIndex(file uint8, rank uint8) uint8 {
	return uint8((rank << 3) + file)
}

func getFileIndex(index uint8) uint8 {
	return uint8(index & 7)
}

func getRankIndex(index uint8) uint8 {
	return uint8((index >> 3))
}

func (b *Board) printBoard() {
	for rank := range 8 {
		for file := range 8 {
			piece := b.mailbox[rank*8+file]
			if piece == No_piece {
				fmt.Printf("*")
			} else {
				fmt.Printf("%s", piece)
			}
		}
		fmt.Println()
	}
}

func (b *Board) printBitboard() {
	var mailbox [64]Piece
	for i := uint8(0); i < 8; i++ {
		for j := uint8(0); j < 8; j++ {
			mailbox[i*8+j] = No_piece
		}
	}

	for i := uint8(0); i < 12; i++ {
		bitboard := b.bitboards[i]
		piece := Piece(i)

		for j := uint8(0); j < 64; j++ {
			if bitboard&(1<<j) != 0 {
				rank := getRankIndex(j)
				file := getFileIndex(j)

				mailbox[rank*8+file] = piece
			}
		}
	}

	old := b.mailbox
	b.mailbox = mailbox
	b.printBoard()
	b.mailbox = old
}

func (b *Board) putInSquare(file uint8, rank uint8, count uint8, piece Piece) {
	index := getSquareIndex(file, rank)
	b.bitboards[piece] |= 1 << index

	b.occupied |= 1 << index
	b.empty &= ^(1 << index)

	for i := uint8(0); i < count; i++ {
		b.mailbox[rank*8+file+i] = piece
	}
}

func init() {
	for position := range KnigthAttacks {
		var calculatedKnight uint64
		var calculatedPawnWhite uint64
		var calculatedPawnBlack uint64
		var calculatedKing uint64

		if position+1 < 64 {
			calculatedKnight |= 1 << (position + 1)
		}
		if position+9 < 64 {
			calculatedPawnWhite |= 1 << (position + 9)
			calculatedKing |= 1 << (position + 9)
		}
		if position+8 < 64 {
			calculatedKing |= 1 << (position + 8)
		}
		if position+7 < 64 {
			calculatedPawnWhite |= 1 << (position + 7)
			calculatedKing |= 1 << (position + 7)
		}
		if 63-position-9 > -1 {
			calculatedPawnBlack |= 1 << (63 - position - 9)
			calculatedKing |= 1 << (63 - position - 9)
		}
		if 63-position-8 > -1 {
			calculatedKing |= 1 << (63 - position - 8)
		}
		if 63-position-7 > -1 {
			calculatedPawnBlack |= 1 << (63 - position - 7)
			calculatedKing |= 1 << (63 - position - 7)
		}
		if position-1 > -1 {
			calculatedKnight |= 1 << (position - 1)
		}

		if position+17 < 64 {
			calculatedKnight |= 1 << (position + 17)
		}
		if position+15 < 64 {
			calculatedKnight |= 1 << (position + 15)
		}
		if position+10 < 64 {
			calculatedKnight |= 1 << (position + 10)
		}
		if position+6 < 64 {
			calculatedKnight |= 1 << (position + 6)
		}
		if position-17 > -1 {
			calculatedKnight |= 1 << (position - 17)
		}
		if position-15 > -1 {
			calculatedKnight |= 1 << (position - 15)
		}
		if position-10 > -1 {
			calculatedKnight |= 1 << (position - 10)
		}
		if position-6 > -1 {
			calculatedKnight |= 1 << (position - 6)
		}

		KnigthAttacks[position] = calculatedKnight
		PawnAttacks[0][position] = calculatedPawnWhite
		PawnAttacks[1][63-position] = calculatedPawnBlack
		KingAttacks[position] = calculatedKing
	}
}

func rankUp(bitboard uint64, rank uint8) uint64 {
	return bitboard << (8 * rank)
}

func rankDown(bitboard uint64, rank uint8) uint64 {
	return bitboard >> (8 * rank)
}

func pushPawnOne(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		return (rankUp(pawns, 1)) & empty
	} else {
		return (rankDown(pawns, 1)) & empty
	}
}

func pushPawnDouble(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		rank4 := uint64(0x00000000FF000000)
		pushOne := pushPawnOne(pawns, empty, color)
		return rankUp(pushOne, 1) & empty & rank4
	} else {
		rank5 := uint64(0x000000FF00000000)
		pushOne := pushPawnOne(pawns, empty, color)
		return rankDown(pushOne, 1) & empty & rank5
	}
}

func canPushPawnSquares(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		return rankDown(empty, 1) & pawns
	} else {
		return rankUp(empty, 1) & pawns
	}
}

func canPushPawnDoubleSquares(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		rank4 := uint64(0x00000000FF000000)
		emptyRank3 := rankDown(empty&rank4, 1) & empty
		return canPushPawnSquares(pawns, emptyRank3, color)
	} else {
		rank5 := uint64(0x000000FF00000000)
		emptyRank6 := rankUp(empty&rank5, 1) & empty
		return canPushPawnSquares(pawns, emptyRank6, color)
	}
}

func parseFEN(board *Board, fen string) (err error) {
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
		for _, char := range rank {
			if unicode.IsDigit(char) {
				count := int(char - '0')
				board.putInSquare(uint8(total), uint8(i), uint8(count), No_piece)
				total += count
				continue
			}

			board.putInSquare(uint8(total), uint8(i), 1, CharToPiece[char])
			total++
		}

		if total != 8 {
			return fmt.Errorf("Invalid FEN: %s expected 8 ranks found %d", fen, len(ranks))
		}
	}

	switch parts[1] {
	case "w":
		{
			board.currentTurn = WhiteTurn
		}
	case "b":
		{
			board.currentTurn = BlackTurn
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
					board.setFlag(WhiteCastleKingside)
				}
			case 'Q':
				{
					board.setFlag(WhiteCastleQueenside)
				}
			case 'k':
				{
					board.setFlag(BlackCastleKingside)
				}
			case 'q':
				{
					board.setFlag(BlackCastleQueenside)
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
		board.setFlag(EnPassant)
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

	engine := Engine{
		id:     "chess_engine",
		author: "Moayed",
		board: Board{
			currentTurn: WhiteTurn,
		},
	}
	_ = engine

	fmt.Printf("> ")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		index, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		if index < 0 || index > 63 {
			fmt.Println("Invalid index")
			continue
		}

		knightPosition := KnigthAttacks[index]
		pawnPositionWhite := PawnAttacks[0][index]
		pawnPositionBlack := PawnAttacks[1][index]
		for i := uint8(0); i < 64; i++ {
			if (knightPosition & (1 << i)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
		for i := uint8(0); i < 64; i++ {
			if (pawnPositionWhite & (1 << i)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
		for i := uint8(0); i < 64; i++ {
			if (pawnPositionBlack & (1 << i)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
		fmt.Printf("> ")
	}
}
