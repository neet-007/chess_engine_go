package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neet-007/chess_engine_go/pkg/castlings"
	"github.com/neet-007/chess_engine_go/pkg/shared"
)

type Color int

const (
	NP12     = 12
	NP       = 6
	WHITE    = Color(0)
	BLACK    = Color(1)
	StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
)

const (
	Pawn int = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	WP = iota
	BP
	WN
	BN
	WB
	BB
	WR
	BR
	WQ
	BQ
	WK
	BK
	Empty = 15
)

type Board struct {
	Sq       [64]int
	Count    [12]int
	PiecesBB [6]Bitboard
	WBBB     [2]Bitboard
	King     [2]int

	Ep int
	castlings.Castlings
	Smt Color
}

func (b *Board) AllBB() Bitboard {
	return b.WBBB[0] | b.WBBB[1]
}

func (b *Board) Clear() {
	b.Smt = WHITE
	b.Castlings = 0
	b.Ep = 0
	b.WBBB[0], b.WBBB[1] = 0, 0

	for i := 0; i < NP; i++ {
		b.PiecesBB[i] = 0
	}

	for i := A1; i < H8; i++ {
		b.Sq[i] = Empty
	}

	for i := 0; i < NP12; i++ {
		b.Count[i] = 0
	}

}

func (b *Board) NewGame() {
	b.Clear()
	ParseFEN(StartPos)
}

func (b *Board) SetSq(p12, s int) {
	b.Sq[s] = p12

	if p12 == Empty {
		b.WBBB[WHITE].Clr(uint(s))
		b.WBBB[BLACK].Clr(uint(s))

		for p := 0; p < NP; p++ {
			b.PiecesBB[p].Clr(uint(s))
		}

		return
	}

	p := piece(p12)
	sd := p12Color(p12)

	if p12 == King {
		b.King[p] = s
	}

	b.WBBB[sd].Set(uint(s))
	b.PiecesBB[p].Set(uint(s))
}

func ParseFEN(fen string) {
	fenIx := 0
	sq := 0

	for row := 7; row >= 0; row-- {
		for sq = row * 8; sq < row*8+8; {

			char := string(fen[fenIx])
			fenIx++

			if char == "/" {
				continue
			}

			if i, err := strconv.Atoi(char); err == nil {
				fmt.Println(i, "empty from sq", sq)
				sq += i
				continue
			}
			fmt.Println(char, " at sq ", sq)

			// TODO: set the square on the board
			//b.SetSq(Fen2Sq[char], sq)

			sq++
		}
	}

	remaining := strings.Split(strings.TrimSpace(fen[fenIx:]), " ")

	if len(remaining) > 0 {
		if remaining[0] == "w" {
			// TODO: make to move white
		} else if remaining[0] == "b" {
			// TODO: make to move black
		} else {
			r := fmt.Sprintf("%s sq=%d fenIx=%d\n", strings.Join(remaining, " "), sq, fenIx)

			shared.Tell("parse fen remaingin ", r, "\n")
			shared.Tell("parse fen ", remaining[0], "invalid smt")
			// TODO: reset to move to white
		}

	}

	// TODO: reset board castlings
	if len(remaining) > 1 {
		_ = castlings.ParseCastlings(remaining[1])
	}

	// TODO: reset es passent
	if len(remaining) > 2 {
		if remaining[2] != "-" {
			// TODO: set es passent on board
			_ = Fen2Sq[remaining[2]]
		}
	}

	// TODO: reset the 50 rule
	if len(remaining) > 3 {
		_ = parse50(remaining[3])
	}
}

func parse50(fen50 string) int {
	r50, err := strconv.Atoi(fen50)
	if err != nil || r50 < 0 {
		shared.Tell("parse 50 50 move rule in fenstring ", fen50, " is not a valid number >= 0 ")
		return 0
	}
	return r50
}

func ParseMvs(moves string) {
	movesList := strings.Split(moves, " ")

	for _, move := range movesList {
		fmt.Println("make move ", move)
	}
}

func piece(p12 int) int {
	return p12 >> 1
}

func p12Color(p12 int) Color {
	return Color(p12 & 0x1)
}
