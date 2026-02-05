package board

import (
	"fmt"
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
	case White_all:
		return "all white"
	case Black_all:
		return "all black"
	default:
		return "?"
	}
}

const (
	// White pieces (0-5)
	White_pawn Piece = iota
	White_knight
	White_bishop
	White_rook
	White_queen
	White_king

	// Black pieces (6-11)
	Black_pawn
	Black_knight
	Black_bishop
	Black_rook
	Black_queen
	Black_king

	// Occupancy/Utility (12-13)
	White_all
	Black_all

	// For Mailbox and general use only (Not for bitboards!)
	No_piece = 14
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

var RookRelevantMasks = [64]uint64{}
var BishopRelevantMasks = [64]uint64{}

var KnightAttacks = [64]uint64{}
var PawnAttacks = [2][64]uint64{}
var KingAttacks = [64]uint64{}

type Square uint8

func (s Square) String() string {
	switch s {
	case A1:
		return "a1"
	case B1:
		return "b1"
	case C1:
		return "c1"
	case D1:
		return "d1"
	case E1:
		return "e1"
	case F1:
		return "f1"
	case G1:
		return "g1"
	case H1:
		return "h1"
	case A2:
		return "a2"
	case B2:
		return "b2"
	case C2:
		return "c2"
	case D2:
		return "d2"
	case E2:
		return "e2"
	case F2:
		return "f2"
	case G2:
		return "g2"
	case H2:
		return "h2"
	case A3:
		return "a3"
	case B3:
		return "b3"
	case C3:
		return "c3"
	case D3:
		return "d3"
	case E3:
		return "e3"
	case F3:
		return "f3"
	case G3:
		return "g3"
	case H3:
		return "h3"
	case A4:
		return "a4"
	case B4:
		return "b4"
	case C4:
		return "c4"
	case D4:
		return "d4"
	case E4:
		return "e4"
	case F4:
		return "f4"
	case G4:
		return "g4"
	case H4:
		return "h4"
	case A5:
		return "a5"
	case B5:
		return "b5"
	case C5:
		return "c5"
	case D5:
		return "d5"
	case E5:
		return "e5"
	case F5:
		return "f5"
	case G5:
		return "g5"
	case H5:
		return "h5"
	case A6:
		return "a6"
	case B6:
		return "b6"
	case C6:
		return "c6"
	case D6:
		return "d6"
	case E6:
		return "e6"
	case F6:
		return "f6"
	case G6:
		return "g6"
	case H6:
		return "h6"
	case A7:
		return "a7"
	case B7:
		return "b7"
	case C7:
		return "c7"
	case D7:
		return "d7"
	case E7:
		return "e7"
	case F7:
		return "f7"
	case G7:
		return "g7"
	case H7:
		return "h7"
	case A8:
		return "a8"
	case B8:
		return "b8"
	case C8:
		return "c8"
	case D8:
		return "d8"
	case E8:
		return "e8"
	case F8:
		return "f8"
	case G8:
		return "g8"
	case H8:
		return "h8"
	default:
		return "?"
	}
}

const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8

	No_square = 65
)

type Move uint16

const (
	QuietMoveFlag   uint16 = 0 // 0000
	DoublePushFlag  uint16 = 1 // 0001
	KingCastleFlag  uint16 = 2 // 0010
	QueenCastleFlag uint16 = 3 // 0011
	CaptureFlag     uint16 = 4 // 0100
	EpCaptureFlag   uint16 = 5 // 0101

	// Promotions (Bit 3 of the 4 bits is the "Promotion" bit)
	KnightPromotionFlag uint16 = 8  // 1000
	BishopPromotionFlag uint16 = 9  // 1001
	RookPromotionFlag   uint16 = 10 // 1010
	QueenPromotionFlag  uint16 = 11 // 1011

	// Promotion-Captures (Bit 2 + Bit 3)
	KnightPromoCaptureFlag uint16 = 12 // 1100
	BishopPromoCaptureFlag uint16 = 13 // 1101
	RookPromoCaptureFlag   uint16 = 14 // 1110
	QueenPromoCaptureFlag  uint16 = 15 // 1111
)

func NewMove(from Square, to Square, flags uint16) Move {
	return Move((flags&0xF)<<12 | (uint16(from)&0x3F)<<6 | (uint16(to) & 0x3F))
}

func (m Move) To() int {
	return int(m & 0x3F)
}

func (m Move) From() int {
	return int((m >> 6) & 0x3F)
}

func (m Move) Flags() int {
	return int(m >> 12)
}

type CurrentTurn uint8

const (
	WhiteTurn CurrentTurn = iota
	BlackTurn
)

const (
	WhiteCastleKingside = iota
	WhiteCastleQueenside
	BlackCastleKingside
	BlackCastleQueenside
)

// NOTE: Using LERF (little endian rank file)
type Board struct {
	Bitboards   [14]uint64
	Empty       uint64
	Occupied    uint64
	Mailbox     [64]uint8
	CurrentTurn CurrentTurn
	Flags       uint8
	HalfMoves   int
	FullMoves   int
	EpSquare    uint8
}

func NewBoard() *Board {
	var mailbox [64]uint8
	for i := range mailbox {
		mailbox[i] = No_piece
	}

	return &Board{
		CurrentTurn: WhiteTurn,
		Empty:       ^uint64(0),
		Mailbox:     mailbox,
		EpSquare:    No_square,
	}
}

func (b *Board) GetFlag(flag uint8) bool {
	return (b.Flags & (1 << flag)) != 0
}

func (b *Board) SetFlag(flag uint8) {
	b.Flags |= 1 << flag
}

func GetSquareIndex(file uint8, rank uint8) uint8 {
	return uint8((rank << 3) + file)
}

func GetFileIndex(index uint8) uint8 {
	return uint8(index & 7)
}

func GetRankIndex(index uint8) uint8 {
	return uint8((index >> 3))
}

func (b *Board) PrintBoard() {
	for rank := range 8 {
		rank = 7 - rank
		for file := range 8 {
			piece := b.Mailbox[rank*8+file]
			if piece == No_piece {
				fmt.Printf("*")
			} else {
				fmt.Printf("%s", Piece(piece))
			}
		}
		fmt.Println()
	}
}

func (b *Board) PrintBitboard() {
	var mailbox [64]uint8
	for i := uint8(0); i < 8; i++ {
		for j := uint8(0); j < 8; j++ {
			mailbox[i*8+j] = No_piece
		}
	}

	for i := uint8(0); i < 12; i++ {
		bitboard := b.Bitboards[i]
		piece := Piece(i)

		for j := uint8(0); j < 64; j++ {
			if bitboard&(1<<j) != 0 {
				mailbox[j] = uint8(piece)
			}
		}
	}

	old := b.Mailbox
	b.Mailbox = mailbox
	b.PrintBoard()
	b.Mailbox = old
}

func (b *Board) PutInSquare(file uint8, rank uint8, piece Piece) {
	index := GetSquareIndex(file, rank)

	if piece < No_piece {
		b.Bitboards[piece] |= 1 << index
		b.Occupied |= 1 << index
		// NOTE: b.empty will be updated later as ^b.occupied

		if piece < Black_pawn {
			b.Bitboards[White_all] |= 1 << index
		} else {
			b.Bitboards[Black_all] |= 1 << index
		}

		b.Mailbox[rank*8+file] = uint8(piece)
	}
}

func (b *Board) UpdateEmpty() {
	b.Empty = ^b.Occupied
}

func RankUp(bitboard uint64, rank uint8) uint64 {
	return bitboard << (8 * rank)
}

func RankDown(bitboard uint64, rank uint8) uint64 {
	return bitboard >> (8 * rank)
}

func PushPawnOne(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		return (RankUp(pawns, 1)) & empty
	} else {
		return (RankDown(pawns, 1)) & empty
	}
}

func PushPawnDouble(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		rank4 := uint64(0x00000000FF000000)
		pushOne := PushPawnOne(pawns, empty, color)
		return RankUp(pushOne, 1) & empty & rank4
	} else {
		rank5 := uint64(0x000000FF00000000)
		pushOne := PushPawnOne(pawns, empty, color)
		return RankDown(pushOne, 1) & empty & rank5
	}
}

func CanPushPawnSquares(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		return RankDown(empty, 1) & pawns
	} else {
		return RankUp(empty, 1) & pawns
	}
}

func CanPushPawnDoubleSquares(pawns uint64, empty uint64, color uint8) uint64 {
	if color == 0 {
		rank4 := uint64(0x00000000FF000000)
		emptyRank3 := RankDown(empty&rank4, 1) & empty
		return CanPushPawnSquares(pawns, emptyRank3, color)
	} else {
		rank5 := uint64(0x000000FF00000000)
		emptyRank6 := RankUp(empty&rank5, 1) & empty
		return CanPushPawnSquares(pawns, emptyRank6, color)
	}
}

func generateSlowAttacks(sq int, blockers uint64, isRook bool) uint64 {
	attacks := uint64(0)

	var dr [4]int
	var df [4]int

	if isRook {
		dr = [4]int{1, -1, 0, 0}
		df = [4]int{0, 0, 1, -1}
	} else {
		dr = [4]int{1, 1, -1, -1}
		df = [4]int{1, -1, 1, -1}
	}

	startRank := sq / 8
	startFile := sq % 8

	for i := range 4 {
		r, f := startRank+dr[i], startFile+df[i]

		for r >= 0 && r < 8 && f >= 0 && f < 8 {
			targetSq := uint64(1 << (r*8 + f))
			attacks |= targetSq

			if (targetSq & blockers) != 0 {
				break
			}

			r, f = r+dr[i], f+df[i]
		}
	}

	return attacks
}

func init() {
	InitMagics()

	for position := range 64 {
		f := int(GetFileIndex(uint8(position)))
		r := int(GetRankIndex(uint8(position)))

		knOffsets := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
		for _, off := range knOffsets {
			targetF, targetR := f+off[0], r+off[1]
			if targetF >= 0 && targetF < 8 && targetR >= 0 && targetR < 8 {
				KnightAttacks[position] |= 1 << GetSquareIndex(uint8(targetF), uint8(targetR))
			}
		}

		for df := -1; df <= 1; df++ {
			for dr := -1; dr <= 1; dr++ {
				if df == 0 && dr == 0 {
					continue
				}

				targetF, targetR := f+df, r+dr
				if targetF >= 0 && targetF < 8 && targetR >= 0 && targetR < 8 {
					KingAttacks[position] |= 1 << GetSquareIndex(uint8(targetF), uint8(targetR))
				}
			}
		}

		for _, df := range []int{-1, 1} {
			targetF, targetR := f+df, r+1
			if targetF >= 0 && targetF < 8 && targetR >= 0 && targetR < 8 {
				PawnAttacks[0][position] |= 1 << GetSquareIndex(uint8(targetF), uint8(targetR))
			}
		}

		for _, df := range []int{-1, 1} {
			targetF, targetR := f+df, r-1
			if targetF >= 0 && targetF < 8 && targetR >= 0 && targetR < 8 {
				PawnAttacks[1][position] |= 1 << GetSquareIndex(uint8(targetF), uint8(targetR))
			}
		}
	}

	for position := range RookRelevantMasks {
		var mask uint64

		file := GetFileIndex(uint8(position))
		rank := GetRankIndex(uint8(position))

		for i := file + 1; i < 7; i++ {
			mask |= 1 << GetSquareIndex(uint8(i), rank)
		}
		for i := file - 1; i > 0; i-- {
			mask |= 1 << GetSquareIndex(uint8(i), rank)
		}
		for i := rank + 1; i < 7; i++ {
			mask |= 1 << GetSquareIndex(file, uint8(i))
		}
		for i := rank - 1; i > 0; i-- {
			mask |= 1 << GetSquareIndex(file, uint8(i))
		}

		RookRelevantMasks[position] = mask
	}

	for position := range BishopRelevantMasks {
		var mask uint64

		file := GetFileIndex(uint8(position))
		rank := GetRankIndex(uint8(position))

		for i, j := file+1, rank+1; i < 7 && j < 7; i, j = i+1, j+1 {
			mask |= 1 << GetSquareIndex(uint8(i), uint8(j))
		}
		for i, j := file-1, rank+1; i > 0 && j < 7; i, j = i-1, j+1 {
			mask |= 1 << GetSquareIndex(uint8(i), uint8(j))
		}
		for i, j := file+1, rank-1; i < 7 && j > 0; i, j = i+1, j-1 {
			mask |= 1 << GetSquareIndex(uint8(i), uint8(j))
		}
		for i, j := file-1, rank-1; i > 0 && j > 0; i, j = i-1, j-1 {
			mask |= 1 << GetSquareIndex(uint8(i), uint8(j))
		}

		BishopRelevantMasks[position] = mask
	}

	for sq := range 64 {
		rookMask := RookRelevantMasks[sq]
		occupancy := uint64(0)

		for {
			attacks := generateSlowAttacks(sq, occupancy, true)

			magic := RookMagics[sq]
			shift := RookShifts[sq]
			offset := RookTableOffsets[sq]
			index := (uint64(occupancy) * magic) >> shift

			RookTable[offset+int(index)] = attacks

			occupancy = (occupancy - rookMask) & rookMask
			if occupancy == 0 {
				break
			}
		}

		bishopMask := BishopRelevantMasks[sq]
		occupancy = uint64(0)

		for {
			attacks := generateSlowAttacks(sq, occupancy, false)

			magic := BishopMagics[sq]
			shift := BishopShifts[sq]
			offset := BishopTableOffsets[sq]
			index := (uint64(occupancy) * magic) >> shift

			BishopTable[offset+int(index)] = attacks

			occupancy = (occupancy - bishopMask) & bishopMask
			if occupancy == 0 {
				break
			}
		}
	}
}
