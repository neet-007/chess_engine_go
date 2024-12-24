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
	StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - "
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
	Smt    Color
	Rule50 int
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
	ParseFEN(b, StartPos)
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

	if p == King {
		b.King[sd] = s
	}

	b.WBBB[sd].Set(uint(s))
	b.PiecesBB[p].Set(uint(s))
}

func (b *Board) Move(to, fr, pr int) bool {
	p12 := b.Sq[fr]

	newEp := 0
	switch {
	case p12 == WK && b.Castlings != 0:
		{
			b.Castlings.Off(castlings.ShortW | castlings.LongW)
			if fr == E1 {
				if to == G1 {
					b.SetSq(WR, F1)
					b.SetSq(Empty, G1)
				} else {
					b.SetSq(WR, D1)
					b.SetSq(Empty, A1)
				}
			}
		}
	case p12 == BK && b.Castlings != 0:
		{
			b.Castlings.Off(castlings.ShortB | castlings.LongB)
			if fr == E8 {
				if to == G8 {
					b.SetSq(WR, F8)
					b.SetSq(Empty, G8)
				} else {
					b.SetSq(WR, D8)
					b.SetSq(Empty, A8)
				}
			}
		}
	case p12 == WR && b.Castlings != 0:
		{
			if fr == F1 {
				b.Castlings.Off(castlings.LongW)
			} else if fr == H1 {
				b.Castlings.Off(castlings.ShortW)
			}
		}
	case p12 == BR && b.Castlings != 0:
		{
			if fr == F8 {
				b.Castlings.Off(castlings.LongW)
			} else if fr == H8 {
				b.Castlings.Off(castlings.ShortW)
			}
		}
	case p12 == WP && b.Sq[to] == Empty:
		{
			if to-fr == 16 {
				newEp = fr + 8
			} else if to-fr == 7 {
				b.SetSq(Empty, to+8)
			} else if to-fr == 9 {
				b.SetSq(Empty, to-8)
			}
		}
	case p12 == BP && b.Sq[to] == Empty:
		{
			if to-fr == 16 {
				newEp = fr + 8
			} else if to-fr == 7 {
				b.SetSq(Empty, to+8)
			} else if to-fr == 9 {
				b.SetSq(Empty, to-8)
			}
		}
	}
	b.Ep = newEp
	b.Sq[fr] = Empty

	if pr != Empty {
		b.SetSq(pr, to)
	} else {
		b.SetSq(p12, to)
	}

	b.Smt = b.Smt ^ 0x1
	// TODO: check if in check to return false

	return true
}

func (b *Board) Print() {
	fmt.Println()
	txtStm := "BLACK"
	if b.Smt == WHITE {
		txtStm = "WHITE"
	}
	txtEp := "-"
	if b.Ep != 0 {
		txtEp = Sq2Fen[b.Ep]
	}

	fmt.Printf("%v to move; ep: %v  castling:%v  \n", txtStm, txtEp, b.Castlings.String())

	fmt.Println("  +------+------+------+------+------+------+------+------+")
	for lines := 8; lines > 0; lines-- {
		fmt.Println("  |      |      |      |      |      |      |      |      |")
		fmt.Printf("%v |", lines)
		for ix := (lines - 1) * 8; ix < lines*8; ix++ {
			fmt.Printf("   %v  |", Int2fen(b.Sq[ix]))
		}
		fmt.Println()
		fmt.Println("  |      |      |      |      |      |      |      |      |")
		fmt.Println("  +------+------+------+------+------+------+------+------+")
	}

	fmt.Printf("       A      B      C      D      E      F      G      H\n")
}

func (b *Board) PrintAllBB() {
	txtStm := "BLACK"
	if b.Smt == WHITE {
		txtStm = "WHITE"
	}
	txtEp := "-"
	if b.Ep != 0 {
		txtEp = Sq2Fen[b.Ep]
	}
	fmt.Printf("%v to move; ep: %v   castling:%v\n", txtStm, txtEp, b.Castlings.String())

	fmt.Println("white pieces")
	fmt.Println(b.WBBB[WHITE].Stringln())
	fmt.Println("black pieces")
	fmt.Println(b.WBBB[BLACK].Stringln())

	fmt.Println("wP")
	fmt.Println((b.PiecesBB[Pawn] & b.WBBB[WHITE]).Stringln())
	fmt.Println("wN")
	fmt.Println((b.PiecesBB[Knight] & b.WBBB[WHITE]).Stringln())
	fmt.Println("wB")
	fmt.Println((b.PiecesBB[Bishop] & b.WBBB[WHITE]).Stringln())
	fmt.Println("wR")
	fmt.Println((b.PiecesBB[Rook] & b.WBBB[WHITE]).Stringln())
	fmt.Println("wQ")
	fmt.Println((b.PiecesBB[Queen] & b.WBBB[WHITE]).Stringln())
	fmt.Println("wK")
	fmt.Println((b.PiecesBB[King] & b.WBBB[WHITE]).Stringln())

	fmt.Println("bP")
	fmt.Println((b.PiecesBB[Pawn] & b.WBBB[BLACK]).Stringln())
	fmt.Println("bN")
	fmt.Println((b.PiecesBB[Knight] & b.WBBB[BLACK]).Stringln())
	fmt.Println("bB")
	fmt.Println((b.PiecesBB[Bishop] & b.WBBB[BLACK]).Stringln())
	fmt.Println("bR")
	fmt.Println((b.PiecesBB[Rook] & b.WBBB[BLACK]).Stringln())
	fmt.Println("bQ")
	fmt.Println((b.PiecesBB[Queen] & b.WBBB[BLACK]).Stringln())
	fmt.Println("bK")
	fmt.Println((b.PiecesBB[King] & b.WBBB[BLACK]).Stringln())
}

func Int2fen(fenInt int) string {
	switch fenInt {
	case Empty:
		return ""
	case WP:
		return "P"
	case BP:
		return "p"
	case WK:
		return "K"
	case BK:
		return "k"
	case WB:
		return "B"
	case BB:
		return "b"
	case WR:
		return "R"
	case BR:
		return "r"
	case WN:
		return "N"
	case BN:
		return "n"
	case WQ:
		return "Q"
	case BQ:
		return "q"
	default:
		panic("invalid fen int")
	}
}

func ParseFEN(b *Board, fen string) {
	if b == nil {
		fmt.Println("board is nil")
		return
	}
	fenIx := 0
	sq := 0

	fen = strings.TrimSpace(fen)
	for row := 7; row >= 0; row-- {
		for sq = row * 8; sq < row*8+8; {

			char := string(fen[fenIx])
			fenIx++

			if char == "/" {
				continue
			}

			if i, err := strconv.Atoi(char); err == nil {
				for j := 0; j < i; j++ {
					b.SetSq(Empty, sq)
					sq++
				}
				continue
			}
			if strings.IndexAny(P12ToFen, char) == -1 {
				shared.Tell("info string invalid piece ", char, " try next one")
				continue
			}

			// TODO: set the square on the board
			b.SetSq(fen2pc(char), sq)

			sq++
		}
	}

	remaining := strings.Split(strings.TrimSpace(fen[fenIx:]), " ")

	if len(remaining) > 0 {
		if remaining[0] == "w" {
			// TODO: make to move white
			b.Smt = WHITE
		} else if remaining[0] == "b" {
			// TODO: make to move black
			b.Smt = BLACK
		} else {
			r := fmt.Sprintf("%s sq=%d fenIx=%d\n", strings.Join(remaining, " "), sq, fenIx)

			shared.Tell("parse fen remaingin ", r, "\n")
			shared.Tell("parse fen ", remaining[0], "invalid smt")
			// TODO: reset to move to white
			b.Smt = WHITE
		}

	}

	// TODO: reset board castlings
	b.Castlings = 0
	if len(remaining) > 1 {
		b.Castlings = castlings.ParseCastlings(remaining[1])
	}

	// TODO: reset es passent
	b.Ep = 0
	if len(remaining) > 2 {
		if remaining[2] != "-" {
			// TODO: set es passent on board
			b.Ep = Fen2Sq[remaining[2]]
		}
	}

	// TODO: reset the 50 rule
	if len(remaining) > 3 {
		b.Rule50 = parse50(remaining[3])
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
	movesList := strings.Fields(strings.ToLower(moves))

	for _, move := range movesList {
		move := strings.TrimSpace(move)

		if len(move) != 4 && len(move) != 5 {
			shared.Tell("parse move ", move, " move length must be 4 or 5")
			return
		}

		_, ok := Fen2Sq[move[:2]]
		if !ok {
			shared.Tell("parse move ", move, " from square is invalid")
			return
		}

		// TODO: get the piece on the board
		//p12 := board.sq[fr]

		// TODO: get the color of the piece
		//pCol = p12Color(p12)

		// TODO: check if the color is the one to move

		_, ok = Fen2Sq[move[2:4]]
		if !ok {
			shared.Tell("parse move ", move, " to square is invalid")
			return
		}

		promotion := 0
		if len(move) == 5 {
			if !strings.ContainsAny(move[4:5], "qnbr") {
				shared.Tell("parse move ", move, " invalid promotion")
				return
			}
		}
		promotion = promotion
	}
	// TODO: make board move method
}

func fen2pc(c string) int {
	for p, x := range P12ToFen {
		if string(x) == c {
			return p
		}
	}
	return Empty
}
func piece(p12 int) int {
	return p12 >> 1
}

func p12Color(p12 int) Color {
	return Color(p12 & 0x1)
}
