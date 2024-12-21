package board

import "fmt"

type Board_ struct {
	WhitePawns   Bitboard
	WhiteKnights Bitboard
	WhiteBishops Bitboard
	WhiteRooks   Bitboard
	WhiteQueens  Bitboard
	WhiteKing    Bitboard

	BlackPawns   Bitboard
	BlackKnights Bitboard
	BlackBishops Bitboard
	BlackRooks   Bitboard
	BlackQueens  Bitboard
	BlackKing    Bitboard
}

type Castling uint

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
	Castling
	Smt Color
}

func (b *Board) AllBB() Bitboard {
	return b.WBBB[0] | b.WBBB[1]
}

func (b *Board) Clear() {
	b.Smt = WHITE
	b.Castling = 0
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
	//parseFEN(StartPos)
}

var defaultBoard = [][]byte{
	{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{'.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.'},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'},
}

func GetDefaultBoard() [][]byte {
	copyBoard := make([][]byte, len(defaultBoard))
	for i := range defaultBoard {
		copyBoard[i] = make([]byte, len(defaultBoard[i]))
		copy(copyBoard[i], defaultBoard[i])
	}
	return copyBoard
}

var CharsToPosInitail = map[string]byte{
	"a1": 'R', "b1": 'N', "c1": 'B', "d1": 'Q', "e1": 'K', "f1": 'B', "g1": 'N', "h1": 'R',
	"a2": 'P', "b2": 'P', "c2": 'P', "d2": 'P', "e2": 'P', "f2": 'P', "g2": 'P', "h2": 'P',
	"a3": '.', "b3": '.', "c3": '.', "d3": '.', "e3": '.', "f3": '.', "g3": '.', "h3": '.',
	"a4": '.', "b4": '.', "c4": '.', "d4": '.', "e4": '.', "f4": '.', "g4": '.', "h4": '.',
	"a5": '.', "b5": '.', "c5": '.', "d5": '.', "e5": '.', "f5": '.', "g5": '.', "h5": '.',
	"a6": '.', "b6": '.', "c6": '.', "d6": '.', "e6": '.', "f6": '.', "g6": '.', "h6": '.',
	"a7": 'p', "b7": 'p', "c7": 'p', "d7": 'p', "e7": 'p', "f7": 'p', "g7": 'p', "h7": 'p',
	"a8": 'r', "b8": 'n', "c8": 'b', "d8": 'q', "e8": 'k', "f8": 'b', "g8": 'n', "h8": 'r',
}

func NewBoardFromInitial(initBoard [][]byte) Board_ {
	if len(initBoard) != 8 || len(initBoard[0]) != 8 {
		panic("invalid board")
	}

	board := Board_{}
	for i := 0; i < 64; i++ {
		rank := 7 - (i / 8)
		file := 7 - (i % 8)
		squareIndex := 8*rank + file
		var bi Bitboard = 1 << squareIndex
		switch initBoard[rank][file] {
		case 'r':
			{
				board.BlackRooks |= bi
			}
		case 'n':
			{
				board.BlackKnights |= bi
			}
		case 'b':
			{
				board.BlackBishops |= bi
			}
		case 'q':
			{
				board.BlackQueens |= bi
			}
		case 'k':
			{
				board.BlackKing |= bi
			}
		case 'p':
			{
				board.BlackPawns |= bi
			}
		case 'R':
			{
				board.WhiteRooks |= bi
			}
		case 'N':
			{
				board.WhiteKnights |= bi
			}
		case 'B':
			{
				board.WhiteBishops |= bi
			}
		case 'Q':
			{
				board.WhiteQueens |= bi
			}
		case 'K':
			{
				board.WhiteKing |= bi
			}
		case 'P':
			{
				board.WhitePawns |= bi
			}
		}
	}
	return board
}

func NewBoardFromInitialLERF(initBoard [][]byte) Board_ {
	if len(initBoard) != 8 || len(initBoard[0]) != 8 {
		panic("invalid board")
	}

	board := Board_{}
	for i := 0; i < 64; i++ {
		rank := 7 - (i / 8) // Flip the rank
		file := 7 - (i % 8) // Flip the file
		squareIndex := 8*rank + file
		var bi Bitboard = 1 << squareIndex

		switch initBoard[rank][file] {
		case 'r':
			board.BlackRooks |= bi
		case 'n':
			board.BlackKnights |= bi
		case 'b':
			board.BlackBishops |= bi
		case 'q':
			board.BlackQueens |= bi
		case 'k':
			board.BlackKing |= bi
		case 'p':
			board.BlackPawns |= bi
		case 'R':
			board.WhiteRooks |= bi
		case 'N':
			board.WhiteKnights |= bi
		case 'B':
			board.WhiteBishops |= bi
		case 'Q':
			board.WhiteQueens |= bi
		case 'K':
			board.WhiteKing |= bi
		case 'P':
			board.WhitePawns |= bi
		}
	}

	return board
}

func (b Board_) BitBoardBoardToByte() [][]byte {
	ret := make([][]byte, 8)
	for i := range 8 {
		ret[i] = make([]byte, 8)
	}
	for i := 0; i < 64; i++ {
		ret[i/8][i%8] = '.'
		if (b.BlackRooks>>i)&1 == 1 {
			ret[i/8][i%8] = 'r'
		}
		if (b.BlackKnights>>i)&1 == 1 {
			ret[i/8][i%8] = 'n'
		}
		if (b.BlackBishops>>i)&1 == 1 {
			ret[i/8][i%8] = 'b'
		}
		if (b.BlackQueens>>i)&1 == 1 {
			ret[i/8][i%8] = 'q'
		}
		if (b.BlackKing>>i)&1 == 1 {
			ret[i/8][i%8] = 'k'
		}
		if (b.BlackPawns>>i)&1 == 1 {
			ret[i/8][i%8] = 'p'
		}
		if (b.WhiteRooks>>i)&1 == 1 {
			ret[i/8][i%8] = 'R'
		}
		if (b.WhiteKnights>>i)&1 == 1 {
			ret[i/8][i%8] = 'N'
		}
		if (b.WhiteBishops>>i)&1 == 1 {
			ret[i/8][i%8] = 'B'
		}
		if (b.WhiteQueens>>i)&1 == 1 {
			ret[i/8][i%8] = 'Q'
		}
		if (b.WhiteKing>>i)&1 == 1 {
			ret[i/8][i%8] = 'K'
		}
		if (b.WhitePawns>>i)&1 == 1 {
			ret[i/8][i%8] = 'P'
		}
	}
	return ret
}

func (b Board_) PrintBoardBitBoards() {
	fmt.Printf("br %064b\n", b.BlackRooks)
	fmt.Printf("bn %064b\n", b.BlackKnights)
	fmt.Printf("bb %064b\n", b.BlackBishops)
	fmt.Printf("bq %064b\n", b.BlackQueens)
	fmt.Printf("bk %064b\n", b.BlackKing)
	fmt.Printf("bp %064b\n", b.BlackPawns)
	fmt.Printf("wr %064b\n", b.WhiteRooks)
	fmt.Printf("wn %064b\n", b.WhiteKnights)
	fmt.Printf("wb %064b\n", b.WhiteBishops)
	fmt.Printf("wq %064b\n", b.WhiteQueens)
	fmt.Printf("wk %064b\n", b.WhiteKing)
	fmt.Printf("wp %064b\n", b.WhitePawns)
}
